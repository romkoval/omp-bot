package group

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroupCommander) Default(ctx context.Context, inputMessage *tgbotapi.Message) error {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"You wrote: "+inputMessage.Text+" type /help__logistic__group for command list",
	)

	_, err := c.bot.Send(msg)
	return err
}
