package domain

import (
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/email"
)

type Customer struct {
	Id    string `json:"_id"`
	Name  string `json:"name"`
	Email Email  `json:"email"`
	CPF   CPF    `json:"cpf"`
}
