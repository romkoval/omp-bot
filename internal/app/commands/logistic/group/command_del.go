package group

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroupCommander) Del(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	var answ string
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		answ = "Please use format: /delete__logistic__group groupId"
	} else {
		ok, err := c.groupService.Remove(uint64(idx))
		if ok {
			answ = fmt.Sprintf("Group with id: %d successfuly removed", idx)
		} else {
			answ = fmt.Sprintf("Unable to remove group with id: %d, err: %s", idx, err)
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		answ,
	)

	_, err = c.bot.Send(msg)
	return err
}
