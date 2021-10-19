package group

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func (c *GroupCommander) Add(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var answ string

	if len(args) == 0 {
		answ = "Please use group title as an argument, e.g. /new__logistic__group New Group Title"
	} else {
		newid, err := c.groupService.Create(logistic.Group{Title: args})

		if err == nil {
			answ = fmt.Sprintf("Group successfuly created with id: %d", newid)
		} else {
			answ = fmt.Sprintf("Unable to create group, err: %s", err)
		}
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		answ,
	)

	c.bot.Send(msg)
}
