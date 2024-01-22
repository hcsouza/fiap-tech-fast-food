package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type productRepository struct {
	databaseAdapter repository.IDatabaseAdapter
}

func NewProductRepository(databaseAdapter repository.IDatabaseAdapter) *productRepository {
	return &productRepository{databaseAdapter: databaseAdapter}
}

func (pr productRepository) FindAll() ([]domain.Product, error) {
	products, err := pr.databaseAdapter.FindAll("", "")

	if err != nil {
		return nil, err
	}

	foundProducts := []domain.Product{}

	for _, product := range products {
		foundProducts = append(foundProducts, product.(domain.Product))
	}

	return foundProducts, nil
}

func (pr productRepository) FindAllByCategory(category Category) ([]domain.Product, error) {
	products, err := pr.databaseAdapter.FindAll("category", string(category))

	if err != nil {
		return nil, err
	}

	foundProducts := []domain.Product{}

	for _, product := range products {
		foundProducts = append(foundProducts, product.(domain.Product))
	}

	return foundProducts, nil
}

func (pr productRepository) FindById(id string) (*domain.Product, error) {
	product, err := pr.databaseAdapter.FindOne("_id", string(id))

	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	foundProduct := product.(*domain.Product)

	return foundProduct, nil
}

func (pr productRepository) Save(product *domain.Product) error {
	_, err := pr.databaseAdapter.Save(
		product.ToSaveMongo(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr productRepository) Update(product *domain.Product) error {
	_, err := pr.databaseAdapter.Update(
		product.ID,
		product.ToUpdateMongo(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr productRepository) Delete(id string) error {
	_, err := pr.databaseAdapter.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
