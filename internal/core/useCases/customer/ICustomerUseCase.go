package customer

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type CustomerCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Cpf   string `json:"cpf" binding:"required,IsCpfValid" `
}

type ICustomerUseCase interface {
	FindAll(context.Context) ([]domain.Customer, error)
	CreateCustomer(context.Context, CustomerCreateRequest) (domain.Customer, error)
	GetCustomer(ctx context.Context, params map[string]string) (domain.Customer, error)
}
