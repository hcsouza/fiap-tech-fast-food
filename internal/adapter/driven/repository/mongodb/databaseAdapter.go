package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoAdapter[T any] struct {
	client         mongo.Client
	collectionName string
	collection     mongo.Collection
	domain         T
}

func NewMongoAdapter[T any](client mongo.Client, collectionName string) repository.IDatabaseAdapter {
	collection := client.Database("db").Collection(collectionName)
	return &mongoAdapter[T]{
		collectionName: collectionName,
		client:         client,
		collection:     *collection,
	}
}

func (ad *mongoAdapter[T]) FindOne(key, value string) (interface{}, error) {
	ctx := context.TODO()
	var result T

	err := ad.collection.FindOne(
		ctx,
		bson.D{{Key: key, Value: value}},
	).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("record not found")
		} else {
			return nil, err
		}
	}
	return &result, err
}

func (ad *mongoAdapter[T]) Save(identifier string, data interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ad.collection.InsertOne(ctx, data)
	return res.InsertedID, err
}
