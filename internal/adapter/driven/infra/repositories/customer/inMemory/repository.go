package inmemory

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repositories"
)

type customerRepository struct {
	customers map[string]domain.Customer
}

func NewCustomerRepository() repositories.ICustomerRepository {
	return &customerRepository{
		customers: make(map[string]domain.Customer),
	}
}

func (repo *customerRepository) Create(ctx context.Context, customerToCreate domain.Customer) (domain.Customer, error) {
	repo.customers[string(customerToCreate.CPF)] = customerToCreate
	return customerToCreate, nil
}

func (repo *customerRepository) Find(ctx context.Context, params map[string]string) (domain.Customer, error) {
	cpf, exists := params["cpf"]
	if !exists {
		return domain.Customer{}, nil
	}

	customer, ok := repo.customers[cpf]
	if ok {
		return customer, nil
	}
	return domain.Customer{}, nil

}

func (repo *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	v := make([]domain.Customer, 0, len(repo.customers))
	for _, value := range repo.customers {
		v = append(v, value)
	}
	return v, nil
}
