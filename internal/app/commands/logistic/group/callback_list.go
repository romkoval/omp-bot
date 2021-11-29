package group

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *GroupCommander) CallbackList(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		logger.ErrorKV(ctx, "failed to handle callback event", "event", callbackPath.CallbackData)
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			"failed to handle callback message",
		)
		c.bot.Send(msg)
		return err
	}
	return c.ListOffset(ctx, callback.Message, parsedData.Offset)
}
