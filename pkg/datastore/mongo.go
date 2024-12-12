package datastore

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"kotakemail.id/shared/interfaces"
)

type MongoDB struct {
	*BaseDatastore
	client *mongo.Client
	URI    string
}

func NewMongoDB(uri string) (interfaces.Datastore[*mongo.Client], error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		BaseDatastore: &BaseDatastore{},
		URI:           uri,
		client:        client,
	}, nil
}

func (m *MongoDB) GetConnection() *mongo.Client {
	return m.client
}

func (m *MongoDB) Shutdown() error {
	return m.client.Disconnect(context.Background())
}
