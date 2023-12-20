package repositories

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type ICustomerRepository interface {
	Create(ctx context.Context, customerToCreate domain.Customer) (domain.Customer, error)
	Find(ctx context.Context, params map[string]string) (domain.Customer, error)
}
