package main

import (
	"context"
	"fmt"
	stdLog "log"

	"github.com/nightlord189/example-hasher/internal/config"
	"github.com/nightlord189/example-hasher/internal/delivery/grpc"
	"github.com/nightlord189/example-hasher/internal/usecase"
	"github.com/nightlord189/example-hasher/pkg/log"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("start #1")

	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		stdLog.Fatalf("load config error: %v", err)
	}

	if err := log.InitLogger(cfg.LogLevel, "example-hasher-grpc-server"); err != nil {
		stdLog.Fatalf("error on init logger: %v", err)
	}
	ctx := context.Background()

	zerolog.Ctx(ctx).Info().Msg("start #2")

	usecaseInst := usecase.New()

	handler := grpc.New(cfg.GRPCPort, usecaseInst)

	zerolog.Ctx(ctx).Info().Msgf("running grpc handler on port %d", cfg.GRPCPort)

	if err := handler.Run(); err != nil {
		zerolog.Ctx(ctx).Error().Msgf("run handler error: %v", err)
	}
}
