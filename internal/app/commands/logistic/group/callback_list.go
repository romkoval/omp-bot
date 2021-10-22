package group

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *GroupCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("failed to handle callback event:%+v", callbackPath.CallbackData)
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			"failed to handle callback message",
		)
		c.bot.Send(msg)
		return err
	}
	return c.ListOffset(callback.Message, parsedData.Offset)
}
