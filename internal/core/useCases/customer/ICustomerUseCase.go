package customer

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type ICustomerUseCase interface {
	GetAll(context.Context) ([]domain.Customer, error)
}
