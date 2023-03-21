package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Collection[T interface{}] struct {
	collection *mongo.Collection
}

// collection: client.Database(config.DatabaseName).Collection(collectionName),
func (col *Collection[T]) Get(skip, limit int64, sort bson.M) ([]T, error) {
	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(sort)

	cursor, err := col.collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	return getDocumentsFromCursor[T](cursor), nil
}

func (col *Collection[T]) Add(newDocs []T) (int, error) {
	var addingDocs []interface{}
	for _, newDoc := range newDocs {
		addingDocs = append(addingDocs, newDoc)
	}

	result, err := col.collection.InsertMany(context.TODO(), addingDocs, nil)
	if err != nil {
		return 0, err
	}

	return len(result.InsertedIDs), nil
}

// TODO index based on ytid
func (col *Collection[T]) FindByKey(key string, vals []string) ([]T, error) {
	filter := bson.M{
		key: bson.M{
			"$in": vals,
		},
	}

	cursor, err := col.collection.Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}

	return getDocumentsFromCursor[T](cursor), nil
}

func getDocumentsFromCursor[T interface{}](cursor *mongo.Cursor) []T {
	var result []T
	for cursor.Next(context.TODO()) {
		var currentResult T
		err := bson.Unmarshal(cursor.Current, &currentResult)
		if err != nil {
			log.Fatal("error while marshalling into T", cursor.Current.String())
		}

		result = append(result, currentResult)
	}

	return result
}
