package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"videosservice/env"
)

func GetDatabase(config env.MongoConfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database(config.DatabaseName), client.Ping(context.TODO(), nil)
}
