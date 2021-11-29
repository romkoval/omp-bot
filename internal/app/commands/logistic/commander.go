package logistic

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/logistic/group"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot            *tgbotapi.BotAPI
	groupCommander Commander
}

func NewLogisticCommander(bot *tgbotapi.BotAPI) Commander {
	return &LogisticCommander{
		bot:            bot,
		groupCommander: group.NewGroupCommander(bot),
	}
}

func (c *LogisticCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "group":
		c.groupCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown group - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "group":
		c.groupCommander.HandleCommand(ctx, msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown group - %s", commandPath.Subdomain)
	}
}
