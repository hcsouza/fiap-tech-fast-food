package checkout

import (
	"fmt"

	coreErrors "github.com/hcsouza/fiap-tech-fast-food/internal/core/errors"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type checkoutUseCase struct {
	order order.IOrderUseCase
}

func NewCheckoutUseCase(orderUseCase order.IOrderUseCase) ICheckoutUseCase {
	return &checkoutUseCase{
		order: orderUseCase,
	}
}

func (uc *checkoutUseCase) CreateCheckout(orderId string) (*CreateCheckout, error) {
	order, err := uc.order.FindById(orderId)
	nextStatus := orderStatus.ORDER_PAYMENT_PENDING

	if err != nil {
		return nil, err
	}

	if !order.OrderStatus.IsValidNextStatus(nextStatus.String()) {
		return &CreateCheckout{
			CheckoutURL: "",
			Message:     coreErrors.ErrCheckoutOrderAlreadyCompleted.Error(),
		}, nil
	}

	err = uc.order.UpdateOrderStatus(orderId, nextStatus)

	if err != nil {
		return nil, fmt.Errorf("error updating order status %s to %s", order.OrderStatus.String(), nextStatus.String())
	}

	return &CreateCheckout{
		CheckoutURL: fmt.Sprintf("https://fake-checkout-fb94eb803a7a.herokuapp.com/payment/%s", orderId),
		Message:     "checkout created",
	}, nil
}

func (uc *checkoutUseCase) UpdateCheckout(orderId string, status orderStatus.OrderStatus) error {
	order, err := uc.order.FindById(orderId)

	if err != nil {
		return err
	}

	if !order.OrderStatus.IsValidNextStatus(status.String()) {
		return coreErrors.ErrCheckoutOrderAlreadyCompleted
	}

	err = uc.order.UpdateOrderStatus(orderId, status)

	if err != nil {
		return fmt.Errorf("error updating order status %s to %s", order.OrderStatus.String(), status.String())
	}

	return nil
}
