package main

import (
	"log"
	"smtpcore/routes"

	"kotakemail.id/config"
	"kotakemail.id/pkg/cmd"
	"kotakemail.id/pkg/container"
	appcontext "kotakemail.id/pkg/context"
	"kotakemail.id/pkg/database"
	"kotakemail.id/pkg/logger"
	"kotakemail.id/pkg/storage"
)

func main() {
	ctx := appcontext.NewAppContext()
	cfg, err := config.NewConfig(ctx, "../../config", "rest-api")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	appLogger := logger.NewLogger(ctx, &cfg.Logging)

	container := container.NewContainer(ctx, appLogger)
	container.AddCommand(cmd.NewRestCommand(cfg, appLogger, routes.GetRoutes()))

	for _, dbCfg := range cfg.Databases {
		var db database.Database
		switch dbCfg.Type {
		case config.DB_MONGO:
			db, err = database.NewMongoDB(&dbCfg, appLogger)
		}
		if err != nil {
			appLogger.Fatal().Err(err).Msg("can't connect to database")
		}
		container.AddDatabase(db)
	}

	for _, storageCfg := range cfg.Storages {
		var stor storage.Storage
		switch storageCfg.Type {
		case config.STORAGE_LOCAL:
			stor, err = storage.NewLocalStorage(&storageCfg, appLogger)
		}
		if err != nil {
			appLogger.Fatal().Err(err).Msg("can't connect to database")
		}
		container.AddStorage(stor)
	}

	container.Run()
}
