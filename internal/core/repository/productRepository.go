package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	Find(id string) (*domain.Product, error)
	FindAllByCategory(category Category) ([]domain.Product, error)
	Save(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id string) error
}
