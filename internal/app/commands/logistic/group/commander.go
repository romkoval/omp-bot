package group

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
)

type CommandHandler interface {
	Help(ctx context.Context, inputMsg *tgbotapi.Message) error
	Get(ctx context.Context, inputMsg *tgbotapi.Message) error
	List(ctx context.Context, inputMsg *tgbotapi.Message) error
	Del(ctx context.Context, inputMsg *tgbotapi.Message) error

	Add(ctx context.Context, inputMsg *tgbotapi.Message) error
	Edit(ctx context.Context, inputMsg *tgbotapi.Message) error
	Default(ctx context.Context, inputMsg *tgbotapi.Message) error
}

type GroupCallbackHandler interface {
	CallbackList(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) error
}

type GroupCommander struct {
	bot          *tgbotapi.BotAPI
	groupService group.Service
}

func NewGroupCommander(bot *tgbotapi.BotAPI) *GroupCommander {
	groupService := group.NewDummyGroupService()

	return &GroupCommander{
		bot:          bot,
		groupService: groupService,
	}
}

func (c *GroupCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var err error
	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(ctx, callback, callbackPath)
	default:
		log.Printf("GroupCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
	if err != nil {
		log.Printf("failed to handle command with error:%s\n", err)
	}
}

func (c *GroupCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(ctx, msg)
	case "list":
		err = c.List(ctx, msg)
	case "get":
		err = c.Get(ctx, msg)
	case "delete":
		err = c.Del(ctx, msg)
	case "new":
		err = c.Add(ctx, msg)
	case "edit":
		err = c.Edit(ctx, msg)
	default:
		err = c.Default(ctx, msg)
	}
	if err != nil {
		log.Printf("failed to handle command with error:%s\n", err)
	}
}
