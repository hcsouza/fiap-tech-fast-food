package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/handlers"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup) {
	groupServer := gServer.Group("/v1")
	registerCustomerHandler(groupServer)
}

func registerCustomerHandler(groupServer *gin.RouterGroup) {
	customerInteractor := customer.NewCustomerUseCase()
	handlers.NewCustomerHandler(groupServer, customerInteractor)
}
