package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(gRouter *gin.Engine) {
	gRouter.GET("/health/liveness", getLivenessHandler)
	gRouter.GET("/health/readiness", getReadinessHandler)
}

// Liveness godoc
// @Summary Liveness probe
// @Description Liveness probe
// @Tags Health Routes
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health/liveness [get]
func getLivenessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

// Readiness godoc
// @Summary Readiness probe
// @Description Readiness probe
// @Tags Health Routes
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health/readiness [get]
func getReadinessHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
