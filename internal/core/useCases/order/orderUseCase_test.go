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

	t.Run("should update order with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_STARTED,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}
		newOrder := order.OrderUpdateDTO{
			Cpf: "19119119100",
			OrderItemsDTO: []order.OrderItemDTO{{
				ProductId: "product-123",
				Quantity:  2,
			}},
		}
		updatedOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_STARTED,
			OrderItems:  []domain.OrderItem{{Product: domain.Product{ID: "product-123", Price: 5}, Quantity: 2}},
			Amount:      10,
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		productRepositoryMock.On("FindById", "product-123").Return(&domain.Product{ID: "product-123", Price: 5}, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		err := useCase.UpdateOrder(orderId, newOrder)

		assert.Nil(t, err)
	})

	t.Run("should not update order When status is diff of STARTED", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_PAYMENT_PENDING,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}
		newOrder := order.OrderUpdateDTO{
			Cpf: "19119119100",
			OrderItemsDTO: []order.OrderItemDTO{{
				ProductId: "product-123",
				Quantity:  2,
			}},
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		err := useCase.UpdateOrder(orderId, newOrder)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "order cannot be updated cause status is PAYMENT_PENDING")
	})

	t.Run("should update order status with sucess", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_STARTED,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}
		orderUpdated := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_PAYMENT_PENDING,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", orderUpdated).Return(nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		err := useCase.UpdateOrderStatus(orderId, ORDER_PAYMENT_PENDING)

		assert.Nil(t, err)
	})

	t.Run("should not update order status When new status is some previous status", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_PAYMENT_PENDING,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		err := useCase.UpdateOrderStatus(orderId, ORDER_STARTED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "order status PAYMENT_PENDING cannot updated to previous status STARTED")
	})

	t.Run("should not update order status When new status is not a valid next status", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{
			ID:          orderId,
			Customer:    domain.Customer{},
			OrderStatus: ORDER_PAYMENT_PENDING,
			OrderItems: []domain.OrderItem{{
				Product:  domain.Product{ID: "product-123"},
				Quantity: 1,
			}},
			Amount: 10,
		}

		orderRepositoryMock = mocks.NewMockOrderRepository(t)
		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

		err := useCase.UpdateOrderStatus(orderId, ORDER_READY)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "order status PAYMENT_PENDING cannot be updated to READY. Status available are: [PAYMENT_APPROVED PAYMENT_REFUSED]")
	})
}
