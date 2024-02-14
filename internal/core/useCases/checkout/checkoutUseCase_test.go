package checkout_test

import (
	"errors"
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/checkout"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCheckoutUseCase(t *testing.T) {
	t.Parallel()

	productRepositoryMock := mocks.NewMockProductRepository(t)
	productUseCase := product.NewProductUseCase(productRepositoryMock)

	customerRepositoryMock := mocks.NewMockCustomerRepository(t)
	customerUseCase := customer.NewCustomerUseCase(customerRepositoryMock)
	orderRepositoryMock := mocks.NewMockOrderRepository(t)
	orderUseCase := order.NewOrderUseCase(orderRepositoryMock, productUseCase, customerUseCase)

	t.Run("should create checkout with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_STARTED}
		updatedOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		createdCheckout, err := useCase.CreateCheckout(orderId)

		assert.Nil(t, err)
		assert.Equal(t, createdCheckout.Message, "checkout created")
		assert.Equal(t, createdCheckout.CheckoutURL, "https://fake-checkout-fb94eb803a7a.herokuapp.com/payment/"+orderId)
	})

	t.Run("should not create checkout When order status is diff of STARTED", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		createdCheckout, err := useCase.CreateCheckout(orderId)

		assert.Nil(t, err)
		assert.Equal(t, createdCheckout.Message, "Order already has a checkout completed")
		assert.Equal(t, createdCheckout.CheckoutURL, "")
	})

	t.Run("should not create checkout When some error occurs during update operation", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_STARTED}
		updatedOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(errors.New("error updating order status"))

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		_, err := useCase.CreateCheckout(orderId)

		assert.NotNil(t, err)
	})

	t.Run("should update order status to PAYMENT_APPROVED with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.Nil(t, err)
	})

	t.Run("should update order status to PAYMENT_REFUSED with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_REFUSED}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_REFUSED)

		assert.Nil(t, err)
	})

	t.Run("should not update order status When new status is equal to the current", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "order already has a checkout completed")
	})

	t.Run("should not update order status When new status is not valid next status", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_COMPLETED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "order already has a checkout completed")
	})

	t.Run("should not update order status When some error occurrs during update operation", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &domain.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderRepositoryMock.On("FindById", orderId).Return(existentOrder, nil)
		orderRepositoryMock.On("Update", updatedOrder).Return(errors.New("error updating order status"))

		useCase := checkout.NewCheckoutUseCase(orderUseCase)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error updating order status PAYMENT_PENDING to PAYMENT_APPROVED")
	})
}
