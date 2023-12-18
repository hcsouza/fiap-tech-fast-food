package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hcsouza/fiap-tech-fast-food/docs"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api"
)

// @title Fast Food API
// @version 0.1.0
// @description Fast Food API for FIAP Tech course

// @host localhost:8080
// @BasePath /
func main() {
	gServer := gin.New()
	api.Run(gServer)
}
