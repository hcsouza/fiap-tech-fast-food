package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/routes"
)

func Run(gServer *gin.Engine) {
	gServer.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/liveness", "/health/readiness"),
		gin.Recovery(),
	)

	RegisterHealthRoutes(gServer)
	RegisterSwaggerRoutes(gServer)

	api := gServer.Group("/api")
	routes.RegisterBusinessRoutes(api)

	err := gServer.Run(fmt.Sprintf(":%s", "8080"))

	if err != nil {
		panic(err)
	}
}
