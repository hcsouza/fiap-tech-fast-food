package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hcsouza/fiap-tech-fast-food/cmd/configuration"
	_ "github.com/hcsouza/fiap-tech-fast-food/docs"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api"
)

// @title Fast Food API
// @version 0.1.0
// @description Fast Food API for FIAP Tech course

// @host localhost:8080
// @BasePath /
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
