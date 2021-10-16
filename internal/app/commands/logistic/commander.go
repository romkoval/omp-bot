package logistic

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/logistic/group"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot            *tgbotapi.BotAPI
	groupCommander Commander
}

func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot: bot,
		// groupCommander
		groupCommander: group.NewLogisticGroupCommander(bot),
	}
}

func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "group":
		c.groupCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown group - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "group":
		c.groupCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown group - %s", commandPath.Subdomain)
	}
}
