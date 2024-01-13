package telereporter

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type TeleReporter struct {
	reportChannel int64
	bot           *telebot.Bot
	logger        *zap.Logger
}

func New(bot *telebot.Bot, logger *zap.Logger) *TeleReporter {
	return &TeleReporter{
		reportChannel: -1002031798487, // for debugging bot
		bot:           bot,
		logger:        logger,
	}
}

func (c *TeleReporter) Started() {
	c.report("i have started!")
}

func (c *TeleReporter) report(msg string) {
	if _, err := c.bot.Send(&telebot.Chat{ID: c.reportChannel}, fmt.Sprintf(msg)); err != nil {
		c.logger.Error("failed to report message",
			zap.Int64("channel id", c.reportChannel),
			zap.String("message", msg),
			zap.Error(err))
	}
}
