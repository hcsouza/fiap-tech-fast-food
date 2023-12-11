package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api"
)

func main() {
	gServer := gin.New()
	api.Run(gServer)
}
