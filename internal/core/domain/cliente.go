package domain

import (
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type Cliente struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	CPF   CPF
}
