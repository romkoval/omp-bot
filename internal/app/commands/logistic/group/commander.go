package group

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
)

type GroupCommander struct {
	bot          *tgbotapi.BotAPI
	groupService group.GroupService
}

func NewGroupCommander(bot *tgbotapi.BotAPI) *GroupCommander {
	groupService := group.NewDummyGroupService()

	return &GroupCommander{
		bot:          bot,
		groupService: groupService,
	}
}

func (c *GroupCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("GroupCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *GroupCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(msg)
	case "list":
		err = c.List(msg)
	case "get":
		err = c.Get(msg)
	case "delete":
		err = c.Del(msg)
	case "new":
		err = c.Add(msg)
	case "edit":
		err = c.Edit(msg)
	default:
		err = c.Default(msg)
	}
	if err != nil {
		log.Printf("failed to handle command with error:%s\n", err)
	}
}
