package domain

import (
	"time"

	"github.com/google/uuid"
	ct "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/customTime"
	os "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type Order struct {
	ID          string         `json:"_id" bson:"_id"`
	Customer    Customer       `json:"customer,omitempty"`
	OrderStatus os.OrderStatus `json:"orderStatus"`
	OrderItems  []OrderItem    `json:"orderItems"`
	CreatedAt   ct.CustomTime  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   ct.CustomTime  `json:"updatedAt" bson:"updatedAt"`
	Value       float64        `json:"value"`
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
		"value":       o.Value,
		"createdAt":   ct.CustomTime{Time: time.Now()},
	}
}

func (o *Order) ToUpdateMongo() map[string]interface{} {
	return map[string]interface{}{
		"orderStatus": o.OrderStatus,
		"orderItems":  o.OrderItems,
		"value":       o.Value,
		"updatedAt":   ct.CustomTime{Time: time.Now()},
	}
}

func (o Order) CollectionName() string {
	return "order"
}
