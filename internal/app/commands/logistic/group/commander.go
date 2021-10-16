package group

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
)

type LogisticGroupCommander struct {
	bot          *tgbotapi.BotAPI
	groupService *group.Service
}

func NewLogisticGroupCommander(
	bot *tgbotapi.BotAPI,
) *LogisticGroupCommander {
	groupService := group.NewService()

	return &LogisticGroupCommander{
		bot:          bot,
		groupService: groupService,
	}
}

func (c *LogisticGroupCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LogisticGroupCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LogisticGroupCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
