package group

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *GroupCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("failed to handle callback event:%+v", callbackPath.CallbackData)
		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		c.bot.Send(msg)
	}
	c.ListOffset(callback.Message, parsedData.Offset)
}
