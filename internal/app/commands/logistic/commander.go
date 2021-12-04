package logistic

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	group_app "github.com/ozonmp/omp-bot/internal/app/commands/logistic/group"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot            *tgbotapi.BotAPI
	groupCommander Commander
}

func NewLogisticCommander(bot *tgbotapi.BotAPI, service group.Service) Commander {
	return &LogisticCommander{
		bot:            bot,
		groupCommander: group_app.NewGroupCommander(bot, service),
	}
}

func (c *LogisticCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "group":
		c.groupCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		logger.ErrorKV(ctx, "LogisticCommander.HandleCallback: unknown group", "group", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "group":
		c.groupCommander.HandleCommand(ctx, msg, commandPath)
	default:
		logger.ErrorKV(ctx, "LogisticCommander.HandleCommand: unknown group", "group", commandPath.Subdomain)
	}
}
