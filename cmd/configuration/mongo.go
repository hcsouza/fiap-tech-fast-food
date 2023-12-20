package configuration

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDbConfiguration(ctx context.Context) (mongo.Client, error) {

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://172.24.0.2:27017"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return *client, err
}
