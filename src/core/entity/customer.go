package entity

import (
	valueobject "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type Customer struct {
	Name  string            `json:"name"`
	Email valueobject.Email `json:"email"`
	CPF   valueobject.CPF   `json:"cpf"`
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
