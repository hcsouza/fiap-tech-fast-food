package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hcsouza/fiap-tech-fast-food/cmd/configuration"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api"
)

func main() {

	mongoClient, err := configuration.InitMongoDbConfiguration(context.TODO())
	if err != nil {
		log.Fatal("error on create mongoConnection")
	}

	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	gServer := gin.New()
	api.Run(gServer, mongoClient)
}
