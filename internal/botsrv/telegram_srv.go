package botsrv

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/router"
	"github.com/ozonmp/omp-bot/internal/config"
)

type tlgbotSrv struct {
	bot *tgbotapi.BotAPI
}

type TlgbotSrv interface {
	Start(context.Context) error
}

func NewTlgbotSrv(token string, cfg *config.Config) (TlgbotSrv, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = cfg.Project.Debug

	return &tlgbotSrv{
		bot: bot,
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

	router := router.NewRouter(b.bot)

	go func() {
		select {
		case update := <-updates:
			router.HandleUpdate(update)
		case <-ctx.Done():
			return
		}
	}()
	return nil
}
