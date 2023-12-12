package domain

import (
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   CPF    `json:"cpf"`
}
