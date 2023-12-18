package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hcsouza/fiap-tech-fast-food/docs"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api"
)

// @title Fast Food API
func main() {
	gServer := gin.New()
	api.Run(gServer)
}
