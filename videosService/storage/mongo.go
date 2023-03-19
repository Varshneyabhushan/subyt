package storage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"videosservice/env"
)

func GetDatabase(config env.MongoConfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, errors.New("error while connecting to storage : " + err.Error())
	}

	return client.Database(config.DatabaseName), nil
}
