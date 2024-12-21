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
	config, err := config.NewConfig(ctx, "../../config", "rest-api")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.NewLogger(ctx, &config.Logging)

	container := container.NewContainer(ctx, logger)
	container.AddCommand(cmd.NewRestCommand(config, logger, routes.GetRoutes()))

	container.Run()
}
