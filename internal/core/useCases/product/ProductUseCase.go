package product

import (
	"errors"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type productUseCase struct {
	productsCollection []domain.Product
}

func NewProductUseCase() IProductUseCase {
	return &productUseCase{}
}

func (interactor *productUseCase) GetByCategory(category string) ([]domain.Product, error) {
	if !Category(category).IsValid() {
		return nil, errors.New("invalid category")
	}

	productsFromSearchCategory := []domain.Product{}
	products := []domain.Product{
		{Name: "X-Salada", Category: "Lanche", Price: 10.00},
		{Name: "Refrigerante", Category: "Bebida", Price: 5.00},
	}

	for _, product := range products {
		if product.Category == Category(category) {
			productsFromSearchCategory = append(productsFromSearchCategory, product)
		}
	}

	return productsFromSearchCategory, nil
}

func (interactor *productUseCase) Create(product domain.Product) error {
	if product.Name == "" {
		return errors.New("invalid product name")
	}

	if product.Category == "" || !Category(product.Category).IsValid() {
		return errors.New("invalid product category")
	}

	if product.Price <= 0 {
		return errors.New("invalid product price")
	}

	interactor.productsCollection = append(interactor.productsCollection, product)

	return nil
}

func (interactor *productUseCase) Update(productId string, product domain.Product) error {
	if product.Name == "" {
		return errors.New("invalid product name")
	}

	if product.Category == "" || !Category(product.Category).IsValid() {
		return errors.New("invalid product category")
	}

	if product.Price <= 0 {
		return errors.New("invalid product price")
	}

	for index, product := range interactor.productsCollection {
		if product.Name == productId { // Gerar UUID
			interactor.productsCollection[index] = product
			return nil
		}
	}

	return errors.New("product not found")
}

func (interactor *productUseCase) Delete(productId string) error {
	for index, product := range interactor.productsCollection {
		if product.Name == productId { // Gerar UUID
			interactor.productsCollection = append(interactor.productsCollection[:index], interactor.productsCollection[index+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}
