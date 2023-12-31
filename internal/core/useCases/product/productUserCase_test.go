package product_test

import (
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

var productRepositoryMock *mocks.MockProductRepository

func TestProductUseCase(t *testing.T) {
	t.Parallel()

	t.Run("should return all products", func(t *testing.T) {
		expected := []domain.Product{
			{
				ID:       "found",
				Name:     "x-salada",
				Price:    10.0,
				Category: Category("lanche"),
			},
			{
				ID:       "foundb",
				Name:     "coca-cola",
				Price:    5.0,
				Category: Category("bebida"),
			},
		}

		productRepositoryMock = mocks.NewMockProductRepository(t)
		productRepositoryMock.On("FindAll").Return(expected, nil)

		useCase := product.NewProductUseCase(productRepositoryMock)

		products, err := useCase.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Len(t, products, 2)
		assert.Equal(t, products[0].ID, expected[0].ID)
		assert.Equal(t, products[0].Name, expected[0].Name)
		assert.Equal(t, products[0].Price, expected[0].Price)
		assert.Equal(t, products[0].Category, expected[0].Category)
		assert.Equal(t, products[1].ID, expected[1].ID)
	})
	t.Run("should return products by category", func(t *testing.T) {
		category := Category("lanche")
		expected := []domain.Product{
			{
				ID:       "found",
				Name:     "x-salada",
				Price:    10.0,
				Category: Category("lanche"),
			},
		}

		productRepositoryMock = mocks.NewMockProductRepository(t)
		productRepositoryMock.On("FindAllByCategory", category).Return(expected, nil)

		useCase := product.NewProductUseCase(productRepositoryMock)

		products, err := useCase.GetByCategory(category)

		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Len(t, products, 1)
	})
	t.Run("should create a product", func(t *testing.T) {
		newProduct := &domain.Product{
			Name:     "x-salada",
			Price:    10.0,
			Category: Category("lanche"),
		}

		productRepositoryMock = mocks.NewMockProductRepository(t)
		productRepositoryMock.On("Save", newProduct.Normalize()).Return(nil)

		useCase := product.NewProductUseCase(productRepositoryMock)

		err := useCase.Create(newProduct)

		assert.Nil(t, err)
	})
	t.Run("should update a product", func(t *testing.T) {
		newProduct := &domain.Product{
			ID:       "found",
			Name:     "x-salada",
			Price:    10.0,
			Category: Category("lanche"),
		}

		productRepositoryMock = mocks.NewMockProductRepository(t)
		productRepositoryMock.On("Update", newProduct.Normalize()).Return(nil)

		useCase := product.NewProductUseCase(productRepositoryMock)

		err := useCase.Update(newProduct.ID, newProduct)

		assert.Nil(t, err)
	})
	t.Run("should delete a product", func(t *testing.T) {
		id := "found"

		productRepositoryMock = mocks.NewMockProductRepository(t)
		productRepositoryMock.On("Delete", id).Return(nil)

		useCase := product.NewProductUseCase(productRepositoryMock)

		err := useCase.Delete(id)

		assert.Nil(t, err)
	})
}
