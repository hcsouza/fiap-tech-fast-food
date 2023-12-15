package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/interfaces"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type CustomerRespository interface {
	Find(cpf CPF) (*domain.Customer, error)
	Save(customer *domain.Customer) error
}

type customerRepository struct {
	databaseAdapter interfaces.IDatabaseAdapter
}

func NewCustomerRespository(databaseAdapter interfaces.IDatabaseAdapter) *customerRepository {
	return &customerRepository{databaseAdapter: databaseAdapter}
}

func (cr customerRepository) Find(cpf CPF) (*domain.Customer, error) {
	customer, err := cr.databaseAdapter.FindOne(string(cpf))

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	return customer.(*domain.Customer), nil
}

func (cr customerRepository) Save(customer *domain.Customer) error {
	err := cr.databaseAdapter.Save(
		string(customer.CPF),
		map[string]interface{}{
			"name":  customer.Name,
			"email": customer.Email,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
