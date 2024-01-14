package cmd

import (
	"TwitterMediaBot/utils"
	"fmt"
	tele "gopkg.in/telebot.v3"
)

type Command interface {
	Run(tele.Context, utils.Input) ([]string, error)
	Config() Configuration
}

type Configuration struct {
	Aliases      []string
	Description  string
	HideFromHelp bool
}

func (c *Configuration) ToString() string {
	aliasesStr := func() string {
		aliases := ""
		for i := 0; i < len(c.Aliases); i++ {
			if i == 0 {
				aliases = fmt.Sprintf("/%s", c.Aliases[i])
			} else {
				aliases += fmt.Sprintf(", /%s", c.Aliases[i])
			}
		}
		return aliases
	}
	return fmt.Sprintf("%s: %s", aliasesStr(), c.Description)
}
