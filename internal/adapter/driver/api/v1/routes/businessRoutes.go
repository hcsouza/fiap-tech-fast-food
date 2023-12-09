package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/handlers"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/cliente"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup) {
	groupServer := gServer.Group("/v1")
	registerClienteHandler(groupServer)
}

func registerClienteHandler(groupServer *gin.RouterGroup) {
	clienteInteractor := cliente.NewClienteUseCase()
	handlers.NewClienteHandler(groupServer, clienteInteractor)
}
