package main

import (
	"log"

	"email-handler/grpc_handler"

	"kotakemail.id/config"
	"kotakemail.id/pkg/cmd"
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/logger"
)

func main() {
	ctx := appcontext.NewAppContext()
	cfg, err := config.NewConfig(ctx, "../../config", "email-handler")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	appLogger := logger.NewLogger(ctx, &cfg.Logging)

	container := container.NewContainer(ctx, appLogger)

	if err := container.InitDB(cfg); err != nil {
		appLogger.Fatal().Err(err).Msg("can't connect to database")
	}

	if err := container.InitStorage(cfg); err != nil {
		appLogger.Fatal().Err(err).Msg("can't connect to storage")
	}

	container.AddCommand(
		cmd.NewGrpcServer("grpc-email-handler", cfg, appLogger, cmd.GrpcServerOptions{
			RegisterService: grpc_handler.RegisterGrpcServices,
		}),
	)

	container.Run()
}
