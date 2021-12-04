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
	grpc_lgc_group_api "github.com/ozonmp/omp-bot/internal/pkg/lgc-group-api"
	"github.com/ozonmp/omp-bot/internal/service/logistic/group"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())

	token := os.Getenv("TOKEN")
	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.ErrorKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	lgcGroupApiConn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%d", cfg.GrpcLgcGroupApi.Host, cfg.GrpcLgcGroupApi.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.ErrorKV(ctx, "failed to create client", "err", err)
	}
	client_api := grpc_lgc_group_api.NewOmpGroupApiServiceClient(lgcGroupApiConn)
	service := group.NewGrpcGroupService(client_api, client_api)
	botsrv, err := botsrv.NewTlgbotSrv(token, service, &cfg)
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
