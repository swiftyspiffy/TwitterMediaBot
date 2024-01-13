package cmd

import (
	"TelegramTwitterMediaBot/utils"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

type Info struct {
	Logger *zap.Logger
}

func (c *Info) Config() Configuration {
	return Configuration{
		Aliases:     []string{"info"},
		Description: "Get information about this bot",
	}
}

func (c *Info) Run(ctx tele.Context, payload utils.Input) ([]string, error) {
	info := `Twitter Media Fetcher Bot

Some shit broken? Want a feature? ping me on twitter twitter.com/notnotnotclippy`

	return []string{info}, nil
}
