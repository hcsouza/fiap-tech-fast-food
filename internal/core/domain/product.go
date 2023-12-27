package domain

import (
	"github.com/google/uuid"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type Product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Quantity int      `json:"quantity"`
	Category Category `json:"category"`
}

func (p *Product) IsValid() bool {
	return p.Category.IsValid() // TODO: Melhorar validação
}

func (p *Product) ToMongo() map[string]interface{} {
	return map[string]interface{}{
		"_id":      uuid.New().String(),
		"name":     p.Name,
		"price":    p.Price,
		"quantity": p.Quantity,
		"category": p.Category,
	}
}

func (p *Product) CollectionName() string {
	return "product"
}
