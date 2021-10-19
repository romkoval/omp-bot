package group

import (
	"encoding/json"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *GroupCommander) List(inputMessage *tgbotapi.Message) error {
	return c.ListOffset(inputMessage, uint64(0))
}

func (c *GroupCommander) ListOffset(inputMessage *tgbotapi.Message, cursor uint64) error {
	outputMsgText := "Here all the groups: \n\n"
	nextOffset := 5

	products, err := c.groupService.List(cursor, uint64(5))
	if err != nil {
		outputMsgText = "Nothing found"
		nextOffset = 0
	} else {
		for _, p := range products {
			outputMsgText += p.String()
			outputMsgText += "\n"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			*makeButton("Prev page", cursor, -5),
			*makeButton("Next page", cursor, nextOffset),
		),
	)

	_, err = c.bot.Send(msg)
	return err
}

func makeButton(title string, cursor uint64, offset int) *tgbotapi.InlineKeyboardButton {
	if offset < 0 {
		if math.Abs(float64(offset)) > float64(cursor) {
			cursor = 0
		} else {
			cursor = cursor - uint64(-1*offset)
		}
	} else {
		cursor = cursor + uint64(offset)
	}
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: cursor,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "group",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	button := tgbotapi.NewInlineKeyboardButtonData(title, callbackPath.String())
	return &button
}
