package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"kotakemail.id/config"
	"kotakemail.id/pkg/logger"
)

type MongoDB struct {
	*BaseDatabase
	client *mongo.Client
	db     *mongo.Database
	logger *logger.Logger
	cfg    *config.DatabaseConfig
	uri    string
}

func NewMongoDB(cfg *config.DatabaseConfig, appLogger *logger.Logger) (Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	appLogger.Info().Msgf("connecting to mongodb: %s", uri)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		appLogger.Error().Err(err).Msg("failed to connect to mongodb server")
		return nil, err
	}
	appLogger.Info().Msg("connected to mongodb server")
	db := &MongoDB{
		BaseDatabase: &BaseDatabase{},
		uri:          uri,
		cfg:          cfg,
		logger:       appLogger,
		client:       client,
		db:           client.Database(cfg.Database),
	}

	db.SetName(cfg.Name)
	return db, nil
}

func (m *MongoDB) GetConnection() interface{} {
	return m.db
}

func (m *MongoDB) Shutdown() error {
	return m.client.Disconnect(context.Background())
}
