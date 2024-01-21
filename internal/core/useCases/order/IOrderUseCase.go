package order

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	os "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type OrderCreateDTO struct {
	Cpf           string         `json:"customer"`
	OrderItemsDTO []OrderItemDTO `json:"orderItems"`
}

type OrderUpdateDTO struct {
	Cpf           string         `json:"customer"`
	OrderItemsDTO []OrderItemDTO `json:"orderItems"`
	OrderStatus   os.OrderStatus `json:"orderStatus"`
}

type OrderItemDTO struct {
	ProductId string `json:"product"`
	Amount    int    `json:"amount"`
}

type IOrderUseCase interface {
	FindById(id string) (*domain.Order, error)
	GetAllByStatus(status os.OrderStatus) ([]domain.Order, error)
	CreateOrder(order OrderCreateDTO) (string, error)
	UpdateOrder(orderId string, order OrderUpdateDTO) error
}
