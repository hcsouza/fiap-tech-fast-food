package routes

import (
	"github.com/gin-gonic/gin"
	gw "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/httpClient"
	pg "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/httpClient/paymentGateway"
	adapterDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb"
	customerDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/customer"
	orderDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/order"
	productDB "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/product"
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/infra/config"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"net/http"

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
	registerOrderHandler(groupServer, dbClient)
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
	productRepo := productDB.NewProductRepository(
		adapterDB.NewMongoAdapter[domain.Product](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Product{}.CollectionName(),
		),
	)

	productInteractor := product.NewProductUseCase(productRepo)
	handlers.NewProductHandler(groupServer, productInteractor)
}

func registerOrderHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	customerRepo := customerDB.NewCustomerRepository(
		adapterDB.NewMongoAdapter[domain.Customer](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Customer{}.CollectionName(),
		),
	)

	productRepo := productDB.NewProductRepository(
		adapterDB.NewMongoAdapter[domain.Product](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Product{}.CollectionName(),
		),
	)

	orderRepo := orderDB.NewOrderRepository(
		adapterDB.NewMongoAdapter[domain.Order](
			dbClient,
			config.GetMongoCfg().Database,
			domain.Order{}.CollectionName(),
		),
	)

	client := http.Client{}
	gateway := gw.NewGateway(client)
	paymentGateway := pg.NewPaymentGateway(gateway)

	customerInteractor := customer.NewCustomerUseCase(customerRepo)
	productInteractor := product.NewProductUseCase(productRepo)
	orderInteractor := order.NewOrderUseCase(orderRepo, productInteractor, customerInteractor, paymentGateway)
	handlers.NewOrderHandler(groupServer, orderInteractor)
}
