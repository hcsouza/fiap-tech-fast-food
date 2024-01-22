package order_test

import (
	"errors"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var orderRepositoryMock *mocks.MockOrderRepository

func TestOrderUseCase(t *testing.T) {
	t.Parallel()

	productRepositoryMock := mocks.NewMockProductRepository(t)
	productUseCase := product.NewProductUseCase(productRepositoryMock)

	customerRepositoryMock := mocks.NewMockCustomerRepository(t)
	customerUseCase := customer.NewCustomerUseCase(customerRepositoryMock)

	t.Run("should return order by given id", func(t *testing.T) {
		expectedOrder := &domain.Order{ID: "123"}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", "123").Return(expectedOrder, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrder, err := useCase.FindById("123")

		assert.Nil(t, err)
		assert.NotNil(t, resultOrder)
	})

	t.Run("should return empty result when not found order by id", func(t *testing.T) {
		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", "123").Return(nil, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrder, err := useCase.FindById("123")

		assert.NoError(t, err)
		assert.Nil(t, resultOrder)
	})

	t.Run("should return error in Repository when call FindById", func(t *testing.T) {
		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", "789").Return(nil, errors.New("repository error"))

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		result, err := useCase.FindById("789")

		assert.Error(t, err)
		assert.Nil(t, result)
		orderRepositoryMock.AssertExpectations(t)
	})

	t.Run("should return orders by status", func(t *testing.T) {
		expectedOrders := []domain.Order{
			{ID: "1", OrderStatus: ORDER_STARTED},
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindAllByStatus", ORDER_STARTED).Return(expectedOrders, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrders, err := useCase.GetAllByStatus(ORDER_STARTED)

		assert.NoError(t, err)
		assert.Len(t, resultOrders, len(expectedOrders))
	})

	t.Run("should return empty list when not found orders by status", func(t *testing.T) {
		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindAllByStatus", ORDER_COMPLETED).Return([]domain.Order{}, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrders, err := useCase.GetAllByStatus(ORDER_COMPLETED)

		assert.NoError(t, err)
		assert.Empty(t, resultOrders)
	})

	t.Run("should handle repository error", func(t *testing.T) {
		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindAllByStatus", ORDER_WAITING_PAYMENT).Return(nil, errors.New("repository error"))

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrders, err := useCase.GetAllByStatus(ORDER_WAITING_PAYMENT)

		assert.Error(t, err)
		assert.Nil(t, resultOrders)
	})

}
