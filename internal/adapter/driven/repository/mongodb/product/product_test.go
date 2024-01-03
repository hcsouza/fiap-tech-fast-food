package repository_test

import (
	"errors"
	"testing"

	repository "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/product"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

var databaseAdapter *mocks.MockIDatabaseAdapter

func TestMain(m *testing.M) {
	databaseAdapter = &mocks.MockIDatabaseAdapter{}
}

func TestProductRepository(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		productsInterface := make([]interface{}, 0)

		productsInterface = append(productsInterface, map[string]string{"_id": "found"})
		productsInterface = append(productsInterface, map[string]string{"_id": "foundx"})

		databaseAdapter.On("FindAll").Return(productsInterface, nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		products, err := productRepository.FindAll()

		assert.Nil(t, err)
		assert.Len(t, len(products), 2)
		assert.Equal(t, products[0].ID, "found")
		assert.Equal(t, products[1].ID, "foundx")
	})
	t.Run("Should return error when something went wrong on get all products", func(t *testing.T) {
		databaseAdapter.On("FindAll").Return(nil, errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		products, err := productRepository.FindAll()

		assert.NotNil(t, err)
		assert.Nil(t, products)
	})
	t.Run("Should return all products by category 'lanche'", func(t *testing.T) {
		category := "lanche"
		productsInterface := make([]interface{}, 0)

		productsInterface = append(productsInterface, map[string]string{"_id": "found", "name": "x-salada", "category": "lanche"})
		productsInterface = append(productsInterface, map[string]string{"_id": "foundx", "name": "x-tudo", "category": "lanche"})
		productsInterface = append(productsInterface, map[string]string{"_id": "foundx", "name": "x-tudo", "category": "bebida"})

		databaseAdapter.On("FindAll", "category", "Lanche").Return(productsInterface, nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		products, err := productRepository.FindAllByCategory(Category(string(category)))

		assert.Nil(t, err)
		assert.Len(t, len(products), 2)
	})
	t.Run("Should return all products by category 'lanche' but not has anyone", func(t *testing.T) {
		category := "lanche"
		productsInterface := make([]interface{}, 0)

		productsInterface = append(productsInterface, map[string]string{"_id": "foundx", "name": "x-tudo", "category": "bebida"})

		databaseAdapter.On("FindAll", "category", "Lanche").Return(productsInterface, nil)

		productRepository := repository.NewProductRepository(databaseAdapter)

		products, err := productRepository.FindAllByCategory(Category(string(category)))

		assert.Nil(t, err)
		assert.Len(t, len(products), 0)
	})
	t.Run("Should return error when something went wrong on get all products by category", func(t *testing.T) {
		category := "lanche"

		databaseAdapter.On("FindAll", "category", "Lanche").Return(nil, errors.New("something went wrong"))

		productRepository := repository.NewProductRepository(databaseAdapter)

		products, err := productRepository.FindAllByCategory(Category(string(category)))

		assert.NotNil(t, err)
		assert.Nil(t, products)
	})
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
