package main

import (
	"TwitterMediaBot/cmd"
	"TwitterMediaBot/telereporter"
	"TwitterMediaBot/twitter"
	"TwitterMediaBot/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"os"
	"time"
)

type Handler struct {
	reporter *telereporter.TeleReporter
	bot      *tele.Bot
	twitter  *twitter.TwitterClient
	logger   *zap.Logger
}

func (h *Handler) Start() {
	cmds := h.getCommands()
	for _, command := range cmds {
		currentCommand := command
		if len(command.Config().Aliases) == 0 {
			panic("cmd has no aliases")
		}

		for _, alias := range currentCommand.Config().Aliases {
			h.bot.Handle(fmt.Sprintf("/%s", alias), func(c tele.Context) error {
				h.logger.Info("command received",
					zap.String("command", alias),
					zap.String("payload", c.Message().Payload),
					zap.String("author", c.Message().Sender.Username))
				result, err := currentCommand.Run(c, utils.NewInput(alias, c.Message().Payload))
				if err != nil {
					return c.Send(fmt.Sprintf("Error: %s", err.Error()))
				}
				// custom send logic is INSIDE the command file, command returned empty string, so we dont sent anything here
				if result == nil || len(result) == 0 {
					return nil
				}
				for _, res := range result {
					if len(res) < 1 {
						continue
					}
					_ = c.Send(res)
				}
				return nil
			})
		}
	}
	h.bot.Handle("/help", h.help)
	h.bot.Handle(tele.OnText, h.handleText)
	h.reporter.Started()
	h.bot.Start()
}

func (h *Handler) getCommands() []cmd.Command {
	return []cmd.Command{
		&cmd.Info{Logger: h.logger},
		&cmd.Start{Logger: h.logger},
	}
}

func (h *Handler) help(c tele.Context) error {
	cmds := h.getCommands()
	result := `You can just send me the link to a tweet containing media that you want and I'll do my best to get it for you!

Also, here are my commands:`
	for i := 0; i < len(cmds); i++ {
		// hide command from help output
		if cmds[i].Config().HideFromHelp {
			continue
		}
		config := cmds[i].Config()
		result += fmt.Sprintf(" - %s\n", config.ToString())
	}
	return c.Send(result)
}

func (h *Handler) handleText(c tele.Context) error {
	if c.Message().Text == "help" {
		return h.help(c)
	}

	tweetUrls := h.twitter.ParseMessageForTweetUrls(c.Message().Text)
	if len(tweetUrls) == 0 {
		return c.Send("I didn't recognize any Tweet urls in your message! If you're trying to run a command, try /help")
	}
	if err := c.Reply(fmt.Sprintf("I found %d tweet links in that message. Let me check for any unique media, one moment!", len(tweetUrls))); err != nil {
		return err
	}

	guestToken, err := h.twitter.GetGuestToken()
	if err != nil {
		return err
	}

	failedTweetRetrievals := []error{}
	mediaUrls := []string{}
	for _, url := range tweetUrls {
		tweet, err := h.twitter.GetTweet(url, guestToken)
		if err != nil {
			failedTweetRetrievals = append(failedTweetRetrievals, err)
			continue
		}
		if len(tweet.Data.TweetResult.Result.Legacy.Entities.Media) == 0 {
			failedTweetRetrievals = append(failedTweetRetrievals, errors.New(fmt.Sprintf("tweet (%s) doesn't appear to have any media attached to it!", url.ToString())))
		}
		for _, media := range tweet.Data.TweetResult.Result.Legacy.Entities.Media {
			mediaUrl, err := extractTweetMediaUrl(media)
			if err != nil {
				failedTweetRetrievals = append(failedTweetRetrievals, err)
				continue
			}
			mediaUrls = append(mediaUrls, mediaUrl)
		}
	}
	for _, url := range mediaUrls {
		sendDocument := tele.Document{
			File:                 tele.FromURL(url),
			Caption:              "",
			DisableTypeDetection: true,
		}
		// tele library handles rate limiting automatically
		if _, err := sendDocument.Send(c.Bot(), c.Chat(), &tele.SendOptions{DisableWebPagePreview: true}); err != nil {
			return err
		}
	}
	if len(failedTweetRetrievals) > 0 {
		retrievalErrMsg := "It looks like I failed to retrieve some tweets:\n\n"
		for _, failedRetrieval := range failedTweetRetrievals {
			retrievalErrMsg += fmt.Sprintf("- %s\n", failedRetrieval.Error())
		}
		if err := c.Send(retrievalErrMsg); err != nil {
			return err
		}
	}
	return nil
}

func extractTweetMediaUrl(media twitter.Media) (string, error) {
	switch media.Type {
	case "animated_gif":
		return media.VideoInfo.Variants[0].URL, nil
	case "photo":
		return media.MediaURLHTTPS, nil
	case "video":
		getVideoUrlAndBitrate := func(variants []twitter.Variants) (string, int) {
			largestBitrate := 0
			mp4VidUrl := ""
			for _, variant := range variants {
				if variant.Bitrate > largestBitrate {
					mp4VidUrl = variant.URL
					largestBitrate = variant.Bitrate
				}
			}
			return mp4VidUrl, largestBitrate
		}

		vidUrl, bitrate := getVideoUrlAndBitrate(media.VideoInfo.Variants)
		if bitrate == 0 {
			return "", errors.New("tweet has video with no valid variant links found")
		}
		return vidUrl, nil
	default:
		return "", errors.New(fmt.Sprintf("unknown/unsupported media type: %s", media.Type))
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	handler := Handler{
		reporter: telereporter.New(bot, logger),
		bot:      bot,
		twitter:  twitter.NewTwitterClient(logger),
		logger:   logger,
	}

	handler.Start()
}
