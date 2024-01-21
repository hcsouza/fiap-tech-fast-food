package product

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
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

func (interactor *productUseCase) GetByCategory(category Category) ([]domain.Product, error) {
	products, err := interactor.repository.FindAllByCategory(category)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (interactor *productUseCase) Create(product *domain.Product) error {
	err := interactor.repository.Save(product.Normalize())

	if err != nil {
		return err
	}

	return nil
}

func (interactor *productUseCase) Update(productId string, product *domain.Product) error {
	product.ID = productId

	err := interactor.repository.Update(product.Normalize())

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
