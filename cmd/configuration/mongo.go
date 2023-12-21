package configuration

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/infra/config"
)

func InitMongoDbConfiguration(ctx context.Context) (mongo.Client, error) {

	mongoCfg := config.GetMongoCfg()
	uri := fmt.Sprintf("mongodb://%s:%s", mongoCfg.Host, mongoCfg.Port)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		log.Fatal(err)
	}

	return *client, err
}
