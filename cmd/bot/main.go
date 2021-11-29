package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ozonmp/omp-bot/internal/botsrv"
	"github.com/ozonmp/omp-bot/internal/config"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())

	token := os.Getenv("TOKEN")
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Printf("Failed init configuration: %s", err)
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
	case <-sigs:
		// logger.InfoKV(ctx, "signal.Notify", "sig", fmt.Sprintf("%v", v))
		cancel()
	case <-ctx.Done():
		// logger.InfoKV(ctx, "ctx.Done", "done", fmt.Sprintf("%v", done))
	}
}
