package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/cmd/configuration"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func RegisterHealthRoutes(gRouter *gin.Engine) {
	gRouter.GET("/health/liveness", getLivenessHandler)
	gRouter.GET("/health/readiness", getReadinessHandler)
}

func getLivenessHandler(c *gin.Context) {
	c.JSON(healthCheck())
}

func getReadinessHandler(c *gin.Context) {
	c.JSON(healthCheck())
}

func healthCheck() (code int, obj any) {
	if databaseHealth() {
		return http.StatusOK, struct{ Status string }{Status: "OK"}
	} else {
		return http.StatusBadGateway, struct{ Status string }{Status: "FAILED"}
	}
}

func databaseHealth() (healthStatus bool) {
	healthStatus = false
	client, err := configuration.GetMongoClient()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	healthStatus = true
	return
}
