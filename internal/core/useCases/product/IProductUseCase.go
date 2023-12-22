package product

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type IProductUseCase interface {
	GetByCategory(category string) ([]domain.Product, error)
	Create(product domain.Product) error
	Update(productId string, product domain.Product) error
	Delete(productId string) error
}
