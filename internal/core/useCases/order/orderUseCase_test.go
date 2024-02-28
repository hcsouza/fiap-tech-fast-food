package order_test

import (
	"errors"
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
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
		orderRepositoryMock.On("FindAllByStatus", ORDER_PAYMENT_PENDING).Return(nil, errors.New("repository error"))

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		resultOrders, err := useCase.GetAllByStatus(ORDER_PAYMENT_PENDING)

		assert.Error(t, err)
		assert.Nil(t, resultOrders)
	})

	// t.Run("should return all orders sorted by READY > PREPARING > RECEIVED", func(t *testing.T) {
	// 	expectedOrders := []domain.Order{
	// 		{ID: "1", OrderStatus: ORDER_BEING_PREPARED},
	// 		{ID: "2", OrderStatus: ORDER_COMPLETED},
	// 		{ID: "3", OrderStatus: ORDER_READY},
	// 		{ID: "4", OrderStatus: ORDER_READY},
	// 		{ID: "5", OrderStatus: ORDER_COMPLETED},
	// 		{ID: "6", OrderStatus: ORDER_BEING_PREPARED},
	// 	}

	// 	orderRepositoryMock = mocks.NewMockOrderRepository(t)
	// 	orderRepositoryMock.On("FindAll").Return(expectedOrders, nil)

	// 	useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

	// 	resultOrders, err := useCase.FindAll()

	// 	assert.NoError(t, err)
	// 	assert.Len(t, resultOrders, len(expectedOrders))
	// 	assert.Equal(t, ORDER_READY, resultOrders[0].OrderStatus)
	// 	assert.Equal(t, ORDER_READY, resultOrders[1].OrderStatus)
	// 	assert.Equal(t, ORDER_BEING_PREPARED, resultOrders[2].OrderStatus)
	// 	assert.Equal(t, ORDER_BEING_PREPARED, resultOrders[3].OrderStatus)
	// 	assert.Equal(t, ORDER_COMPLETED, resultOrders[4].OrderStatus)
	// })

	// t.Run("should return all orders sorted by createdAt", func(t *testing.T) {
	// 	currentTime := time.Now()

	// 	expectedOrders := []domain.Order{
	// 		{ID: "1", OrderStatus: ORDER_READY, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(2) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "2", OrderStatus: ORDER_READY, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(1) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "3", OrderStatus: ORDER_BEING_PREPARED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(4) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "4", OrderStatus: ORDER_BEING_PREPARED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(3) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "5", OrderStatus: ORDER_COMPLETED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(6) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "6", OrderStatus: ORDER_COMPLETED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(5) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 	}

	// 	orderRepositoryMock = mocks.NewMockOrderRepository(t)
	// 	orderRepositoryMock.On("FindAll").Return(expectedOrders, nil)

	// 	useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

	// 	resultOrders, err := useCase.FindAll()

	// 	assert.NoError(t, err)
	// 	assert.Len(t, resultOrders, len(expectedOrders))
	// 	assert.Equal(t, resultOrders[0].OrderStatus, ORDER_READY)
	// 	assert.Equal(t, resultOrders[1].OrderStatus, ORDER_READY)
	// 	assert.Equal(t, resultOrders[2].OrderStatus, ORDER_BEING_PREPARED)
	// 	assert.Equal(t, resultOrders[3].OrderStatus, ORDER_BEING_PREPARED)
	// 	assert.Equal(t, resultOrders[4].OrderStatus, ORDER_COMPLETED)
	// 	assert.Equal(t, resultOrders[5].OrderStatus, ORDER_COMPLETED)
	// 	assert.True(t, resultOrders[0].CreatedAt.Before(resultOrders[1].CreatedAt.Time))
	// 	assert.True(t, resultOrders[1].CreatedAt.Before(resultOrders[2].CreatedAt.Time))
	// 	assert.True(t, resultOrders[2].CreatedAt.Before(resultOrders[3].CreatedAt.Time))
	// 	assert.True(t, resultOrders[3].CreatedAt.Before(resultOrders[4].CreatedAt.Time))
	// 	assert.True(t, resultOrders[4].CreatedAt.Before(resultOrders[5].CreatedAt.Time))
	// })

	// t.Run("should return all orders without COMPLETED status", func(t *testing.T) {
	// 	currentTime := time.Now()

	// 	expectedOrders := []domain.Order{
	// 		{ID: "1", OrderStatus: ORDER_READY, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(2) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "2", OrderStatus: ORDER_READY, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(1) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "3", OrderStatus: ORDER_BEING_PREPARED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(4) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "4", OrderStatus: ORDER_COMPLETED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(4) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "5", OrderStatus: ORDER_BEING_PREPARED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(3) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "6", OrderStatus: ORDER_COMPLETED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(6) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 		{ID: "7", OrderStatus: ORDER_COMPLETED, CreatedAt: customTime.CustomTime{
	// 			Time: currentTime.Add(
	// 				time.Hour*time.Duration(5) +
	// 					time.Minute*time.Duration(0) +
	// 					time.Second*time.Duration(0),
	// 			),
	// 		}},
	// 	}

	// 	orderRepositoryMock = mocks.NewMockOrderRepository(t)
	// 	orderRepositoryMock.On("FindAll").Return(expectedOrders, nil)

	// 	useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

	// 	resultOrders, err := useCase.FindAll()

	// 	assert.NoError(t, err)
	// 	assert.Len(t, resultOrders, len(expectedOrders)-1)

	// 	for _, order := range resultOrders {
	// 		assert.NotEqual(t, ORDER_COMPLETED, order.OrderStatus)
	// 	}

	// 	assert.Equal(t, resultOrders[0].OrderStatus, ORDER_READY)
	// 	assert.Equal(t, resultOrders[1].OrderStatus, ORDER_READY)
	// 	assert.Equal(t, resultOrders[2].OrderStatus, ORDER_BEING_PREPARED)
	// 	assert.Equal(t, resultOrders[3].OrderStatus, ORDER_BEING_PREPARED)
	// 	assert.Equal(t, resultOrders[4].OrderStatus, ORDER_COMPLETED)
	// 	assert.Equal(t, resultOrders[5].OrderStatus, ORDER_COMPLETED)
	// 	assert.True(t, resultOrders[0].CreatedAt.Before(resultOrders[1].CreatedAt.Time))
	// 	assert.True(t, resultOrders[1].CreatedAt.Before(resultOrders[2].CreatedAt.Time))
	// 	assert.True(t, resultOrders[2].CreatedAt.Before(resultOrders[3].CreatedAt.Time))
	// 	assert.True(t, resultOrders[3].CreatedAt.Before(resultOrders[4].CreatedAt.Time))
	// 	assert.True(t, resultOrders[4].CreatedAt.Before(resultOrders[5].CreatedAt.Time))
	// })
}
