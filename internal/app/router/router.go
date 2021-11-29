package router

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/logistic"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot *tgbotapi.BotAPI

	// demoCommander
	// demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	// work
	// service
	// exchange
	// estate
	// rating
	// security
	// cinema
	// logistic
	logisticCommander Commander
	// product
	// education
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	return &Router{
		// bot
		bot: bot,
		// demoCommander
		// demoCommander: demo.NewDemoCommander(bot),
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		// work
		// service
		// exchange
		// estate
		// rating
		// security
		// cinema
		// logistic
		logisticCommander: logistic.NewLogisticCommander(bot),
		// product
		// education
	}
}

func (c *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			logger.ErrorKV(ctx, "recovered from panic", "value", panicValue)
		}
	}()

	logger.DebugKV(ctx, "handleUpdate", "update", update)

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(ctx, update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(ctx, update.Message)
	}
}

func (c *Router) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		logger.ErrorKV(ctx, "Router.handleCallback: error parsing callback data", "data", callback.Data, "err", err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		// c.demoCommander.HandleCallback(callback, callbackPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		c.logisticCommander.HandleCallback(ctx, callback, callbackPath)
	case "product":
		break
	case "education":
		break
	default:
		logger.WarnKV(ctx, "Router.handleCallback: unknown domain", "domain", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(ctx context.Context, msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(ctx, msg)
		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		logger.ErrorKV(ctx, "Router.handleCallback: error parsing callback data", "cmd", msg.Command(), "err", err)
		return
	}

	logger.DebugKV(ctx, "handleMessage", "Domain", commandPath.Domain)
	switch commandPath.Domain {
	case "demo":
		// c.demoCommander.HandleCommand(msg, commandPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		break
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		c.logisticCommander.HandleCommand(ctx, msg, commandPath)
	case "product":
		break
	case "education":
		break
	default:
		logger.WarnKV(ctx, "Router.handleCallback: unknown domain", "domain", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(ctx context.Context, inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	c.bot.Send(outputMsg)
}
