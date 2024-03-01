package controller

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/src/common/dto"
	"github.com/hcsouza/fiap-tech-fast-food/src/common/interfaces"
	"github.com/hcsouza/fiap-tech-fast-food/src/communication/gateway"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/entity"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/usecase"
)

type customerController struct {
	useCase interfaces.CustomerUseCase
}

func NewCustomerController(datasource interfaces.DatabaseSource) *customerController {
	gateway := gateway.NewCustomerGateway(datasource)
	return &customerController{
		useCase: usecase.NewCustomerUseCase(gateway),
	}
}

func (cc *customerController) CreateCustomer(ctx context.Context,
	customerRequest dto.CustomerCreateDTO) (*entity.Customer, error) {
	return cc.useCase.CreateCustomer(ctx, customerRequest)
}

func (cc *customerController) GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error) {
	return cc.useCase.GetCustomer(ctx, params)
}
