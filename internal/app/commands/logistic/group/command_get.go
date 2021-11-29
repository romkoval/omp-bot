package group

import (
	"context"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/logger"
)

func (c *GroupCommander) Get(ctx context.Context, inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	var answ string
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		answ = "Please use format: /get__logistic__group groupId"
	} else {
		group, err := c.groupService.Describe(uint64(idx))
		if err != nil {
			logger.ErrorKV(ctx, "fail to get group", "idx", idx, "err", err)
			answ = fmt.Sprintf("group not found by id: %d", idx)
		} else {
			answ = fmt.Sprintf("found group: %s", group.String())
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		answ,
	)

	_, err = c.bot.Send(msg)
	return err
}
