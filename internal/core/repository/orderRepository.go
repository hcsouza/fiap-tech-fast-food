package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type OrderRepository interface {
	FindById(id string) (*domain.Order, error)
	FindAllByStatus(status OrderStatus) ([]domain.Order, error)
	Save(order *domain.Order) (string, error)
	Update(order *domain.Order) error
}
