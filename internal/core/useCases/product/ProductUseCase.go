package product

import (
	"fmt"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
)

type productUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) IProductUseCase {
	return &productUseCase{
		repository: repo,
	}
}

func (interactor *productUseCase) FindById(id string) (*domain.Product, error) {
	product, err := interactor.repository.FindById(id)

	if err != nil {
		mappedErrors := map[string]error{
			"record not found": fmt.Errorf("not found product id {%s}", id),
		}
		mappedError, ok := mappedErrors[err.Error()]

		if ok {
			return nil, mappedError
		}

		return nil, err
	}

	return product, nil
}

func (interactor *productUseCase) GetAll() ([]domain.Product, error) {
	products, err := interactor.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (interactor *productUseCase) GetByCategory(category string) ([]domain.Product, error) {
	products, err := interactor.repository.FindAllByCategory(category)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (interactor *productUseCase) Create(product *domain.Product) error {
	normalizedProduct, err := product.Normalize()

	if err != nil {
		return err
	}

	err = interactor.repository.Save(normalizedProduct)

	if err != nil {
		return err
	}

	return nil
}

func (interactor *productUseCase) Update(productId string, product *domain.Product) error {
	product.ID = productId
	normalizedProduct, err := product.Normalize()

	if err != nil {
		return err
	}

	err = interactor.repository.Update(normalizedProduct)

	if err != nil {
		return err
	}

	return nil
}

func (interactor *productUseCase) Delete(productId string) error {
	err := interactor.repository.Delete(productId)

	if err != nil {
		return err
	}

	return nil
}
