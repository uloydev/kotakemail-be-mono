package main

import (
	"log"
	"smtpcore/routes"

	"kotakemail.id/config"
	"kotakemail.id/pkg/cmd"
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/logger"
)

func main() {
	ctx := appcontext.NewAppContext()
	cfg, err := config.NewConfig(ctx, "../../config", "rest-api")
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

	container.AddCommand(cmd.NewRestCommand(cfg, appLogger, routes.GetRoutes()))

	container.Run()
}
