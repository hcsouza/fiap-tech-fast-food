package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/src/external/api/infra/config"
	"github.com/hcsouza/fiap-tech-fast-food/src/external/api/v1/middlewares"
	"github.com/hcsouza/fiap-tech-fast-food/src/external/api/v1/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(gServer *gin.Engine, dbClient mongo.Client) {
	gServer.Use(
		middlewares.CORSMiddleware(),
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/liveness", "/health/readiness"),
		gin.Recovery(),
	)

	RegisterHealthRoutes(gServer)
	RegisterSwaggerRoutes(gServer)

	api := gServer.Group("/api")
	routes.RegisterBusinessRoutes(api, dbClient)

	gServer.Run(fmt.Sprintf(":%s", config.GetApiCfg().Port))
}
