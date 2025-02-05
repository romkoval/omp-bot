package group

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func (c *GroupCommander) Edit(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	var answ string

	if groupId, title, err := parseEditGroupArgs(args); err != nil {
		answ = "Please use format: /edit__logistic__group groupId Updated Group Title"
	} else {
		err := c.groupService.Update(groupId, logistic.Group{Title: title})

		if err == nil {
			answ = fmt.Sprintf("Group successfuly updated by id: %d", groupId)
		} else {
			answ = fmt.Sprintf("Unable to update group, err: %s", err)
		}
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		answ,
	)

	_, err := c.bot.Send(msg)
	return err
}

func parseEditGroupArgs(args string) (groupId uint64, title string, err error) {
	splitted := strings.SplitN(args, " ", 2)
	if len(splitted) != 2 {
		err = errors.New("input format error, expected string with two values separated by space")
		return
	}
	groupId, err = strconv.ParseUint(splitted[0], 10, 64)
	if err != nil {
		return
	}
	title = splitted[1]
	return
}
