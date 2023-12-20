package routes

import (
	"github.com/gin-gonic/gin"
	adapterDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb"
	customerDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/handlers"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup, dbClient mongo.Client) {
	groupServer := gServer.Group("/v1")
	registerCustomerHandler(groupServer, dbClient)
}

func registerCustomerHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	repo := customerDB.NewCustomerRepository(
		adapterDB.NewMongoAdapter[domain.Customer](dbClient, domain.Customer{}.CollectionName()),
	)

	customerInteractor := customer.NewCustomerUseCase(repo)
	handlers.NewCustomerHandler(groupServer, customerInteractor)
}
