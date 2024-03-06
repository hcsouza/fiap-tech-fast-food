package dto

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
