package checkout

import (
	os "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type CreateCheckout struct {
	CheckoutURL string `json:"checkout_url"`
	Message     string `json:"message"`
}

type UpdateCheckoutDTO struct {
	Status string `json:"status"`
}

type ICheckoutUseCase interface {
	CreateCheckout(orderId string) (*CreateCheckout, error)
	UpdateCheckout(orderId string, status os.OrderStatus) error
}
