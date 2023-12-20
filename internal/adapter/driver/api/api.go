package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(gServer *gin.Engine, dbClient mongo.Client) {
	gServer.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/liveness", "/health/readiness"),
		gin.Recovery(),
	)

	RegisterHealthRoutes(gServer)

	api := gServer.Group("/api")
	routes.RegisterBusinessRoutes(api)

	gServer.Run(fmt.Sprintf(":%s", "8080"))
}
