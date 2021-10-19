package group

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroupCommander) Del(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.groupService.Remove(uint64(idx))

	var answ string
	if ok {
		answ = fmt.Sprintf("Group with id: %d successfuly removed", idx)
	} else {
		answ = fmt.Sprintf("Unable to remove group with id: %d, err: %s", idx, err)
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		answ,
	)

	c.bot.Send(msg)
}
