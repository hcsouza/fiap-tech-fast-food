package customer

import (
	"context"
	"errors"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/email"
)

var ErrCustomerInvalid = errors.New("customer has invalid attributes")
var ErrCustomerSearchParams = errors.New("invalid params to search customer")

type customerUseCase struct {
	repository repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) ICustomerUseCase {
	return &customerUseCase{
		repository: repo,
	}
}

func (interactor *customerUseCase) CreateCustomer(ctx context.Context, customerRequest CustomerCreateRequest) (*domain.Customer, error) {

	customerToCreate := domain.Customer{
		Name:  customerRequest.Name,
		Email: email.Email(customerRequest.Email),
		CPF:   cpf.CPF(customerRequest.Cpf),
	}

	if !customerToCreate.IsValid() {
		return nil, ErrCustomerInvalid
	}

	err := interactor.repository.Save(&customerToCreate)
	return &customerToCreate, err
}

func (interactor *customerUseCase) GetCustomer(ctx context.Context, params map[string]string) (*domain.Customer, error) {
	param, exists := params["cpf"]
	if !exists {
		return &domain.Customer{}, ErrCustomerSearchParams
	}
	return interactor.repository.Find(cpf.CPF(param))
}
