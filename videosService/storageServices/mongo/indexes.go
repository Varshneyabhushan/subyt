package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitIndexes(service Service) error {
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.M{"publishedat": -1},
			Options: options.Index().SetName("publishedAtDSC"),
		},
		{
			Keys:    bson.M{"ytid": 1},
			Options: options.Index().SetUnique(true).SetName("youtubeIdASC"),
		},
	}

	_, err := service.database.Collection("videos").Indexes().CreateMany(context.Background(),
		indexes)

	return err
}
