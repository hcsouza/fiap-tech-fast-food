package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type CustomerRespository interface {
	Find(cpf CPF) (*domain.Customer, error)
	Save(customer *domain.Customer) error
}
