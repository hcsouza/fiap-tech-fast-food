package entity

import (
	"time"

	"github.com/google/uuid"
	vo "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type Order struct {
	ID          string         `json:"_id" bson:"_id"`
	Customer    Customer       `json:"customer,omitempty"`
	OrderStatus vo.OrderStatus `json:"orderStatus"`
	OrderItems  []OrderItem    `json:"orderItems"`
	CreatedAt   vo.CustomTime  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   vo.CustomTime  `json:"updatedAt" bson:"updatedAt"`
	Amount      float64        `json:"amount"`
}

type OrderItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

func (o *Order) ToSaveMongo() map[string]interface{} {
	return map[string]interface{}{
		"_id":         uuid.New().String(),
		"customer":    o.Customer,
		"orderStatus": o.OrderStatus,
		"orderItems":  o.OrderItems,
		"amount":      o.Amount,
		"createdAt":   vo.CustomTime{Time: time.Now()},
	}
}

func (o *Order) ToUpdateMongo() map[string]interface{} {
	return map[string]interface{}{
		"orderStatus": o.OrderStatus,
		"orderItems":  o.OrderItems,
		"amount":      o.Amount,
		"updatedAt":   vo.CustomTime{Time: time.Now()},
	}
}

func (o Order) CollectionName() string {
	return "order"
}
