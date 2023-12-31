package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
)

type productRepository struct {
	databaseAdapter repository.IDatabaseAdapter
}

func NewProductRepository(databaseAdapter repository.IDatabaseAdapter) *productRepository {
	return &productRepository{databaseAdapter: databaseAdapter}
}

func (pr productRepository) Find(id string) (*domain.Product, error) {
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
		product.ToMongo(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr productRepository) Update(product *domain.Product) error {
	_, err := pr.databaseAdapter.Update(
		product.ID,
		product.ToMongo(),
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
