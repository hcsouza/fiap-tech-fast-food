package domain

import (
	"strings"

	"github.com/google/uuid"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type Product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

func (p *Product) IsValidCategory() bool {
	return p.Category.IsValid()
}

func (p *Product) IsValidPrice() bool {
	return p.Price > 0
}

func (p *Product) IsValidName() bool {
	return p.Name != ""
}

func (p *Product) Normalize() *Product {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}

	return &Product{
		ID:       p.ID,
		Name:     strings.ToLower(p.Name),
		Price:    p.Price,
		Category: Category(strings.ToLower(string(p.Category))),
	}
}

func (p *Product) ToMongo() map[string]interface{} {
	return map[string]interface{}{
		"_id":      uuid.New().String(),
		"name":     p.Name,
		"price":    p.Price,
		"category": p.Category,
	}
}

func (p Product) CollectionName() string {
	return "product"
}
