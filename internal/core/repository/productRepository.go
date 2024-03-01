package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	FindById(id string) (*domain.Product, error)
	FindAllByCategory(category string) ([]domain.Product, error)
	Save(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id string) error
}
