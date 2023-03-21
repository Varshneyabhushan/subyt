package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"videosservice/env"
)

type Service struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewService(config env.MongoConfig) (Service, error) {
	clientOptions := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return Service{}, err
	}

	return Service{
		client:   client,
		database: client.Database(config.DatabaseName),
	}, client.Ping(context.TODO(), nil)
}

func GetCollection[T interface{}](service Service, collectionName string) Collection[T] {
	return Collection[T]{
		collection: service.database.Collection(collectionName),
	}
}
