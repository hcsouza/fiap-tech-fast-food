
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(gRouter *gin.Engine) {
	gRouter.GET("/health/liveness", getLivenessHandler)
	gRouter.GET("/health/readiness", getReadinessHandler)
}

func getLivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func getReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
