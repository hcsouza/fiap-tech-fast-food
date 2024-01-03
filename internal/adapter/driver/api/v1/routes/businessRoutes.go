package routes

import (
	"github.com/gin-gonic/gin"
	adapterDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb"
	customerDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/customer"
	productDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/product"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/infra/config"

	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driver/api/v1/handlers"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup, dbClient mongo.Client) {
	groupServer := gServer.Group("/v1")
	registerCustomerHandler(groupServer, dbClient)
	registerProductHandler(groupServer, dbClient)
}

func registerCustomerHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	repo := customerDB.NewCustomerRepository(
		adapterDB.NewMongoAdapter[domain.Customer](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Customer{}.CollectionName(),
		),
	)

	customerInteractor := customer.NewCustomerUseCase(repo)
	handlers.NewCustomerHandler(groupServer, customerInteractor)
}

func registerProductHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	repo := productDB.NewProductRepository(
		adapterDB.NewMongoAdapter[domain.Product](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Product{}.CollectionName(),
		),
	)

	productInteractor := product.NewProductUseCase(repo)
	handlers.NewProductHandler(groupServer, productInteractor)
}
