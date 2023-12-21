package domain

import . "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/category"

type Product struct {
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Quantity int      `json:"quantity"`
	Category Category `json:"category"`
}
