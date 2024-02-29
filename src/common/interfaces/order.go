package interfaces

import (
	"github.com/hcsouza/fiap-tech-fast-food/src/common/dto"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/entity"
	vo "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type IOrderUseCase interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	GetAllByStatus(status vo.OrderStatus) ([]entity.Order, error)
	CreateOrder(order dto.OrderCreateDTO) (string, error)
	UpdateOrder(orderId string, order dto.OrderUpdateDTO) error
	UpdateOrderStatus(orderId string, status vo.OrderStatus) error
}
