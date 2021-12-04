package botsrv

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/router"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
)

type tlgbotSrv struct {
	bot    *tgbotapi.BotAPI
	router *router.Router
}

type TlgbotSrv interface {
	Start(context.Context) error
}

func NewTlgbotSrv(token string, service group.Service, cfg *config.Config) (TlgbotSrv, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	logger.InfoKV(context.Background(), "Authorized on account", "user_name", bot.Self.UserName)

	bot.Debug = cfg.Project.Debug

	return &tlgbotSrv{
		bot:    bot,
		router: router.NewRouter(bot, service),
	}, nil
}

func (b *tlgbotSrv) Start(ctx context.Context) error {
	u := tgbotapi.UpdateConfig{
		Timeout: 10,
	}

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case update := <-updates:
				b.router.HandleUpdate(ctx, update)
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
