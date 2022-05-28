package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	DB *mongo.Database
}

func NewMongoConn(uri string, database string) *MongoClient {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db := client.Database(database)

	return &MongoClient{DB: db}
}

func upsert(ctx context.Context, collection *mongo.Collection, filter bson.D, item interface{}) (bool, error) {
	upsert := true
	result, err := collection.ReplaceOne(ctx, filter, item, &options.ReplaceOptions{
		Upsert: &upsert,
	})

	if err != nil {
		return false, err
	}

	if result.UpsertedCount > 0 {
		return true, nil
	}
	return false, nil
}
