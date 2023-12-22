package domain

import (
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/email"
)

type Customer struct {
	Name  string `json:"name"`
	Email Email  `json:"email"`
	CPF   CPF    `json:"cpf"`
}

func (c *Customer) IsValid() bool {
	return len(c.Name) > 0 && c.Email.IsValid() && c.CPF.IsValid()
}

func (c *Customer) ToMongo() map[string]interface{} {
	return map[string]interface{}{
		"_id":   c.CPF,
		"name":  c.Name,
		"email": c.Email,
	}
}

func (c Customer) CollectionName() string {
	return "customer"
}
