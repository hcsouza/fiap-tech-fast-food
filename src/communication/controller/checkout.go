package controller

import (
	"github.com/hcsouza/fiap-tech-fast-food/src/common/dto"
	"github.com/hcsouza/fiap-tech-fast-food/src/common/interfaces"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/usecase"
	valueobject "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type checkoutController struct {
	useCase interfaces.CheckoutUseCase
}

func NewCheckoutController(datasource interfaces.DatabaseSource,
	orderUseCase interfaces.OrderUseCase) *checkoutController {
	return &checkoutController{
		useCase: usecase.NewCheckoutUseCase(orderUseCase),
	}
}

func (cc *checkoutController) CreateCheckout(orderId string) (*dto.CreateCheckout, error) {
	return cc.useCase.CreateCheckout(orderId)
}

func (cc *checkoutController) UpdateCheckout(orderId string, status valueobject.OrderStatus) error {
	return cc.useCase.UpdateCheckout(orderId, status)
}
