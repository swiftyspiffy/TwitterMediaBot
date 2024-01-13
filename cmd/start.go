package cmd

import (
	"TelegramTwitterMediaBot/utils"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Start struct {
	Logger *zap.Logger
}

func (c *Start) Config() Configuration {
	return Configuration{
		Aliases:     []string{"start"},
		Description: "How to use bot",
	}
}

func (c *Start) Run(ctx tele.Context, payload utils.Input) ([]string, error) {
	info := `Hello, I'm Twitter Media Fetcher Bot! Send me as many twitter links as you want, I'll try to get you the media in them! 

For more info, use /info`

	return []string{info}, nil
}
