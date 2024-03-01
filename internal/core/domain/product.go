package domain

import (
	"strings"

	"github.com/google/uuid"
	categoryValueObject "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"
)

type Product struct {
	ID       string  `json:"_id" bson:"_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type ProductDTO struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func (p *Product) IsValidPrice() bool {
	return p.Price > 0
}

func (p *Product) IsValidName() bool {
	return p.Name != ""
}

func (p *Product) Normalize() (*Product, error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}

	category, err := categoryValueObject.NewCategory(p.Category)

	if err != nil {
		return &Product{}, err
	}

	return &Product{
		ID:       p.ID,
		Name:     strings.ToLower(p.Name),
		Price:    p.Price,
		Category: category,
	}, nil
}

func (p *Product) ToSaveMongo() map[string]interface{} {
	return map[string]interface{}{
		"_id":      uuid.New().String(),
		"name":     p.Name,
		"price":    p.Price,
		"category": p.Category,
	}
}

func (p *Product) ToUpdateMongo() map[string]interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"price":    p.Price,
		"category": p.Category,
	}
}

func (p Product) CollectionName() string {
	return "product"
}
