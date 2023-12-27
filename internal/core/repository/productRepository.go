package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type ProductRepository interface {
	Find(id string) (*domain.Product, error)
	Save(product *domain.Product) error
}
