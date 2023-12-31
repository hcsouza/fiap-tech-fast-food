package repository_test

import (
	"errors"
	"testing"

	repository "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/product"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

var databaseAdapter *mocks.MockIDatabaseAdapter

func TestMain(m *testing.M) {
	databaseAdapter = &mocks.MockIDatabaseAdapter{}
}

func TestProductRepository(t *testing.T) {
	t.Run("Should find a product by ID", func(t *testing.T) {
		id := "found"

		databaseAdapter.On("FindOne", id).Return(&domain.Product{}, nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		product, err := productRepository.Find(id)

		assert.Nil(t, err)
		assert.NotNil(t, product)
	})
	t.Run("Should return nil when not found product by ID", func(t *testing.T) {
		id := "not-found"

		databaseAdapter.On("FindOne", id).Return(nil, nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		product, err := productRepository.Find(id)

		assert.Nil(t, err)
		assert.Nil(t, product)
	})
	t.Run("Should return error when something went wrong on get product by ID", func(t *testing.T) {
		id := "found"

		databaseAdapter.On("FindOne", id).Return(nil, errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		product, err := productRepository.Find(id)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})
	t.Run("Should save a product", func(t *testing.T) {
		product := &domain.Product{
			Name:     "Burguer",
			Price:    10.0,
			Quantity: 1,
			Category: "Lanche",
		}

		databaseAdapter.On("Save", product).Return(nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Save(product)

		assert.Nil(t, err)
	})
	t.Run("Should return error when something went wrong on save product", func(t *testing.T) {
		product := &domain.Product{
			Name:     "Burguer",
			Price:    10.0,
			Quantity: 1,
			Category: "Lanche",
		}

		databaseAdapter.On("Save", product).Return(errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Save(product)

		assert.NotNil(t, err)
	})
	t.Run("Should update a product", func(t *testing.T) {
		product := &domain.Product{
			ID:       "found",
			Name:     "Burguer",
			Price:    10.0,
			Quantity: 1,
			Category: "Lanche",
		}

		databaseAdapter.On("Update", product).Return(nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Update(product)

		assert.Nil(t, err)
	})
	t.Run("Should return error when something went wrong on update product", func(t *testing.T) {
		product := &domain.Product{
			ID:       "found",
			Name:     "Burguer",
			Price:    10.0,
			Quantity: 1,
			Category: "Lanche",
		}

		databaseAdapter.On("Update", product).Return(errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Update(product)

		assert.NotNil(t, err)
	})
	t.Run("Should delete a product", func(t *testing.T) {
		id := "found"

		databaseAdapter.On("Delete", id).Return(nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Delete(id)

		assert.Nil(t, err)
	})
	t.Run("Should return error when something went wrong on delete product", func(t *testing.T) {
		id := "found"

		databaseAdapter.On("Delete", id).Return(errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		err := productRepository.Delete(id)

		assert.NotNil(t, err)
	})
}
