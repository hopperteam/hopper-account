package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var database *mongo.Database
var userCollection *mongo.Collection

func ConnectDB(connectStr string, db string) error {
	dbOptions := options.Client().ApplyURI(connectStr)
	client, err := mongo.NewClient(dbOptions)
	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		return err
	}

	database = client.Database(db)
	userCollection = database.Collection("users")

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_, err = userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{ "email": 1 },
	})

	return err
}
