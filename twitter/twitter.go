package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

type TweetUrl struct {
	Username string
	Id       string
}

// ToString reconstructs the tweet url using the username and id
func (t *TweetUrl) ToString() string {
	return fmt.Sprintf("https://twitter.com/%s/status/%s", t.Username, t.Id)
}

type TwitterClient struct {
	authorization string
	logger        *zap.Logger
}

func NewTwitterClient(logger *zap.Logger) *TwitterClient {
	return &TwitterClient{
		// a shared bearer token for logged out users?
		authorization: "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		logger:        logger,
	}
}

// GetGuestToken retrives a token from Twitter's activate.json endpoint. Not only does this need to be a POST, but it NEEDS TO HAVE A LEGIT
// USER AGENT. The Guest Token is then included in gql calls, without which said calls will fail with "rate limit exceeded".
func (c *TwitterClient) GetGuestToken() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/guest/activate.json", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authorization", c.authorization)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var resp GetGuestTokenResponse
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(resBody, &resp); err != nil {
		return "", err
	}
	return resp.GuestToken, nil
}

// GetTweet makes call to Twitter's TweetResultByRestId gql endpoint. The authorization is static value used by Twitter's frontend for users that are *not* logged
// into the website. Logged in users will have their own authorization. Guest token is retrieved in a prior step. Returned is a large, detailed Tweet struct ref and error.
func (c *TwitterClient) GetTweet(tweetUrl *TweetUrl, guestToken string) (*Tweet, error) {
	url := "https://api.twitter.com/graphql/OGDXNj5PSaBRwg2MlpX0JQ/TweetResultByRestId?variables=%7B%22tweetId%22%3A%22" + tweetUrl.Id + "%22%2C%22withCommunity%22%3Afalse%2C%22includePromotedContent%22%3Afalse%2C%22withVoice%22%3Afalse%7D&features=%7B%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22tweetypie_unmention_optimization_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Afalse%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_media_download_video_enabled%22%3Afalse%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("authorization", c.authorization)
	req.Header.Set("x-guest-token", guestToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var resp Tweet
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resBody, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ParseMessageForTweetUrls uses two regular expressions to detect valid twitter.com* and x.com* tweet urls and then uses
// parseTweetUrl to extract tweet url components. Returned is a slice of TweetUrl refs.
func (c *TwitterClient) ParseMessageForTweetUrls(message string) []*TweetUrl {
	message = strings.ToLower(message)
	results := []*TweetUrl{}
	reTwitter := regexp.MustCompile(`(?:https?://)?(?:www\.)?twitter\.com/\S+/status/\d+`)
	reX := regexp.MustCompile(`(?:https?://)?(?:www\.)?x\.com/\S+/status/\d+`)

	twitterURLs := reTwitter.FindAllString(message, -1)
	xUrls := reX.FindAllString(message, -1)
	urls := append(twitterURLs, xUrls...)
	for i, url := range urls {
		urls[i] = strings.TrimPrefix(url, "http://")
		urls[i] = strings.TrimPrefix(urls[i], "https://")
		urls[i] = strings.TrimPrefix(urls[i], "www.")
		tweetUrl, err := c.parseTweetUrl(urls[i])
		if err != nil {
			c.logger.Error("failed parsing tweet url",
				zap.String("url", urls[i]),
				zap.Error(err))
		} else {
			results = append(results, tweetUrl)
		}
	}
	return results
}

// parseTweetUrl splits a tweet url into sections and then extracts the username and tweet id, returning it as a struct ref,
// as well as error for invalid tweet urls.
func (c *TwitterClient) parseTweetUrl(tweetUrl string) (*TweetUrl, error) {
	invalidTweetUrlErr := errors.New(fmt.Sprintf("tweet link (%s) doesn't appear valid. expected format: twitter.com/username/status/123123123123123", tweetUrl))
	if !strings.Contains(tweetUrl, "") {
		return nil, invalidTweetUrlErr
	}
	parts := strings.Split(tweetUrl, "/status/")
	if len(parts) != 2 {
		return nil, invalidTweetUrlErr
	}

	usernameSide := parts[0]
	if !strings.Contains(usernameSide, "/") {
		return nil, invalidTweetUrlErr
	}
	username := strings.ToLower(strings.Split(usernameSide, "/")[len(strings.Split(usernameSide, "/"))-1])

	tweetIdSide := parts[1]

	extractId := func(input string) string {
		tweetId := ""
		for _, chr := range []rune(tweetIdSide) {
			if !unicode.IsDigit(chr) {
				return tweetId
			}
			tweetId += fmt.Sprintf("%c", chr)
		}
		return tweetId
	}

	tweetId := extractId(tweetIdSide)
	if tweetId == "" {
		return nil, invalidTweetUrlErr
	}
	return &TweetUrl{
		Username: username,
		Id:       tweetId,
	}, nil
}
