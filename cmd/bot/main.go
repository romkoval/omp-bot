package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ozonmp/omp-bot/internal/botsrv"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/logger"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())

	token := os.Getenv("TOKEN")
	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.ErrorKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	botsrv, err := botsrv.NewTlgbotSrv(token, &cfg)
	if err != nil {
		panic(err)
	}
	botsrv.Start(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case v := <-sigs:
		logger.InfoKV(ctx, "cancel by signal.Notify", "sig", fmt.Sprintf("%v", v))
		cancel()
	case <-ctx.Done():
		logger.InfoKV(ctx, "ctx.Done")
	}
}
