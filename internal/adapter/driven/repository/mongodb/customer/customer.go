package repository

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type customerRepository struct {
	databaseAdapter repository.IDatabaseAdapter
}

func NewCustomerRepository(databaseAdapter repository.IDatabaseAdapter) *customerRepository {
	return &customerRepository{databaseAdapter: databaseAdapter}
}

func (cr customerRepository) Find(cpf CPF) (*domain.Customer, error) {
	customer, err := cr.databaseAdapter.FindOne("_id", string(cpf))

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	found := customer.(*domain.Customer)
	composed := domain.Customer{
		CPF:   cpf,
		Name:  found.Name,
		Email: found.Email,
	}
	return &composed, nil
}

func (cr customerRepository) Save(customer *domain.Customer) error {
	_, err := cr.databaseAdapter.Save(
		string(customer.CPF),
		customer.ToMongo(),
	)

	if err != nil {
		return err
	}

	return nil
}
