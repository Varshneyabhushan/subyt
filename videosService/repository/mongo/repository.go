package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{collection: collection}
}

func (repo *MongoRepository) Get(skip, limit int64, sort bson.M) ([]Video, error) {
	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(sort)

	cursor, err := repo.collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	return getDocumentsFromCursor(cursor), nil
}

func (repo *MongoRepository) Add(newDocs []Video) (int, error) {
	var addingDocs []interface{}
	for _, newDoc := range newDocs {
		addingDocs = append(addingDocs, newDoc)
	}

	result, err := repo.collection.InsertMany(context.TODO(), addingDocs, nil)
	if err != nil {
		return 0, err
	}

	return len(result.InsertedIDs), nil
}

func (repo *MongoRepository) FindByIds(ids []primitive.ObjectID) ([]Video, error) {
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}

	cursor, err := repo.collection.Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}

	return getDocumentsFromCursor(cursor), nil
}

func getDocumentsFromCursor(cursor *mongo.Cursor) []Video {
	var result []Video
	for cursor.Next(context.TODO()) {
		var currentResult Video
		err := bson.Unmarshal(cursor.Current, &currentResult)
		if err != nil {
			log.Fatal("error while marshalling into T", cursor.Current.String())
		}

		result = append(result, currentResult)
	}

	return result
}
