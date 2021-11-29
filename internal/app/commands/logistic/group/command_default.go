package group

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/logger"
)

func (c *GroupCommander) Default(ctx context.Context, inputMessage *tgbotapi.Message) error {
	logger.InfoKV(ctx, "userName", inputMessage.From.UserName, "text", inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You wrote: "+inputMessage.Text+" type /help__logistic__group for command list",
	)

	_, err := c.bot.Send(msg)
	return err
}
