package customer

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repositories"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type customerUseCase struct {
	repository repositories.ICustomerRepository
}

func NewCustomerUseCase(repo repositories.ICustomerRepository) ICustomerUseCase {
	return &customerUseCase{
		repository: repo,
	}
}

func (interactor *customerUseCase) GetAll(ctx context.Context) ([]domain.Customer, error) {
	return interactor.repository.GetAllCustomers(ctx)
}

func (interactor *customerUseCase) CreateCustomer(ctx context.Context, customerRequest CustomerCreateRequest) (domain.Customer, error) {

	customerToCreate := domain.Customer{
		Name:  customerRequest.Name,
		Email: customerRequest.Email,
		CPF:   cpf.CPF(customerRequest.Cpf),
	}

	created, err := interactor.repository.AddCustomer(ctx, customerToCreate)
	return created, err
}

func (interactor *customerUseCase) GetCustomer(ctx context.Context, params map[string]string) (domain.Customer, error) {
	return interactor.repository.GetCustomerWithParams(ctx, params)
}
