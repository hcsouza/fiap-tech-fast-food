package controller

import (
	"github.com/hcsouza/fiap-tech-fast-food/src/common/dto"
	"github.com/hcsouza/fiap-tech-fast-food/src/common/interfaces"
	"github.com/hcsouza/fiap-tech-fast-food/src/communication/gateway"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/entity"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/usecase"
	vo "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type orderController struct {
	useCase interfaces.OrderUseCase
}

func NewOrderController(datasource interfaces.DatabaseSource,
	productUseCase interfaces.ProductUseCase,
	customerUseCase interfaces.CustomerUseCase,
) *orderController {

	gateway := gateway.NewOrderGateway(datasource)
	return &orderController{
		useCase: usecase.NewOrderUseCase(gateway, productUseCase, customerUseCase),
	}
}

func (oc *orderController) FindAll() ([]entity.Order, error) {
	return oc.useCase.FindAll()
}

func (oc *orderController) FindById(id string) (*entity.Order, error) {
	return oc.useCase.FindById(id)
}

func (oc *orderController) GetAllByStatus(status vo.OrderStatus) ([]entity.Order, error) {
	return oc.useCase.GetAllByStatus(status)
}

func (oc *orderController) CreateOrder(order dto.OrderCreateDTO) (string, error) {
	return oc.useCase.CreateOrder(order)
}

func (oc *orderController) UpdateOrder(orderId string, order dto.OrderUpdateDTO) error {
	return oc.useCase.UpdateOrder(orderId, order)
}

func (oc *orderController) UpdateOrderStatus(orderId string, status vo.OrderStatus) error {
	return oc.useCase.UpdateOrderStatus(orderId, status)
}
