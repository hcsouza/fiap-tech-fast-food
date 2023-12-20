package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type customerRepository struct {
	databaseAdapter repository.IDatabaseAdapter
}

func NewCustomerRespository(databaseAdapter repository.IDatabaseAdapter) *customerRepository {
	return &customerRepository{databaseAdapter: databaseAdapter}
}

func (cr customerRepository) Find(cpf CPF) (*domain.Customer, error) {
	customer, err := cr.databaseAdapter.FindOne("cpf", string(cpf))

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	return customer.(*domain.Customer), nil
}

func (cr customerRepository) Save(customer *domain.Customer) error {
	_, err := cr.databaseAdapter.Save(
		string(customer.CPF),
		map[string]interface{}{
			"name":  customer.Name,
			"email": customer.Email,
			"cpf":   customer.CPF,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
