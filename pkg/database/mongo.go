package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"kotakemail.id/config"
	"kotakemail.id/shared/interfaces"
)

type MongoDB struct {
	*BaseDatabase
	client *mongo.Client
	URI    string
}

func NewMongoDB(cfg config.DatabaseConfig) (interfaces.Datastore[*mongo.Client], error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		BaseDatabase: &BaseDatabase{},
		URI:          uri,
		client:       client,
	}, nil
}

func (m *MongoDB) GetConnection() *mongo.Client {
	return m.client
}

func (m *MongoDB) Shutdown() error {
	return m.client.Disconnect(context.Background())
}
