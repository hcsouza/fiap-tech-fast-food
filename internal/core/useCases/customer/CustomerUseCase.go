package customer

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type customerUseCase struct{}

func NewCustomerUseCase() ICustomerUseCase {
	return &customerUseCase{}
}

func (interactor *customerUseCase) GetAll(logContext context.Context) ([]domain.Customer, error) {
	return []domain.Customer{
		{Name: "Cliente 1", Email: "cliente@mail.com", CPF: "394.671.960-00"},
		{Name: "Cliente 2", Email: "cliente2@mail.com", CPF: "963.953.450-10"},
	}, nil
}
