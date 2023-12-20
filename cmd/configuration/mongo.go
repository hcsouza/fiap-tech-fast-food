package configuration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDbConfiguration(context.Context) (mongo.Client, error) {

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	return *client, err
}
