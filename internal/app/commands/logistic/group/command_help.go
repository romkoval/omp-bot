package group

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroupCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__logistic__group - help\n"+
			"/get__logistic__group — get a entity\n"+
			"/list__logistic__group - list products\n"+
			"/delete__logistic__group - delete an existing entity\n"+
			"/new__logistic__group — create a new entity\n"+
			"/edit__logistic__group — edit an entity\n",
	)

	c.bot.Send(msg)
}
