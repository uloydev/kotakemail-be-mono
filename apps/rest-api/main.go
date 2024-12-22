package main

import (
	"log"
	"restapi/routes"

	"kotakemail.id/config"
	"kotakemail.id/pkg/cmd"
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/logger"
	"kotakemail.id/shared/base/rest/middleware"
)

// @title Kotak Email API
// @version 1.0
// @description Kotak Email Internal Rest API
// @contact.name Uloydev
// @contact.email wahyu@uloy.dev
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

	restCmd := cmd.NewRestCommand(
		cfg,
		appLogger,
		routes.GetRoutes(),
		middleware.RequestIDMiddleware(),
	)

	container.AddCommand(restCmd)
	container.Run()
}
