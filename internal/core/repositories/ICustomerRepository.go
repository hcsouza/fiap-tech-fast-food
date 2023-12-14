package repositories

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type ICustomerRepository interface {
	AddCustomer(ctx context.Context, customerToCreate domain.Customer) (domain.Customer, error)
	GetCustomerWithParams(ctx context.Context, params map[string]string) (domain.Customer, error)
	GetAllCustomers(ctx context.Context) ([]domain.Customer, error)
}
