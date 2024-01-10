package product

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type IProductUseCase interface {
	GetAll() ([]domain.Product, error)
	GetByCategory(category Category) ([]domain.Product, error)
	Create(product *domain.Product) error
	Update(productId string, product *domain.Product) error
	Delete(productId string) error
}
