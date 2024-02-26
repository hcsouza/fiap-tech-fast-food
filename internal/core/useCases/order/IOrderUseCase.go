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
}

type OrderItemDTO struct {
	ProductId string `json:"product"`
	Quantity  int    `json:"quantity"`
}

type IOrderUseCase interface {
	FindAll() ([]domain.Order, error)
	FindById(id string) (*domain.Order, error)
	GetAllByStatus(status os.OrderStatus) ([]domain.Order, error)
	CreateOrder(order OrderCreateDTO) (string, error)
	UpdateOrder(orderId string, order OrderUpdateDTO) error
	UpdateOrderStatus(orderId string, status os.OrderStatus) error
}
