package customer_test

import (
	"context"
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

var customerRepositoryMock *mocks.MockCustomerRepository

func TestCustomerUseCase(t *testing.T) {
	t.Parallel()

	t.Run("Should find a customer by CPF", func(t *testing.T) {
		cpf := CPF("12345678900")

		expected := domain.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   cpf,
		}

		customerRepositoryMock = mocks.NewMockCustomerRepository(t)
		customerRepositoryMock.On("Find", cpf).Return(&expected, nil)

		useCase := customer.NewCustomerUseCase(customerRepositoryMock)

		params := map[string]string{
			"cpf": string(cpf),
		}
		customer, err := useCase.GetCustomer(context.TODO(), params)

		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, customer.CPF, cpf)
	})

	t.Run("Should return error when search params was invalid", func(t *testing.T) {

		customerRepositoryMock = mocks.NewMockCustomerRepository(t)
		useCase := customer.NewCustomerUseCase(customerRepositoryMock)

		params := map[string]string{
			"nome": "john",
		}
		_, err := useCase.GetCustomer(context.TODO(), params)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, customer.ErrCustomerSearchParams)
	})

	t.Run("Should return error when a customer has invalid attributes", func(t *testing.T) {
		createRequest := customer.CustomerCreateDTO{
			Name:  "John Doe",
			Email: "email.com",
			Cpf:   "111",
		}

		customerRepositoryMock = mocks.NewMockCustomerRepository(t)

		useCase := customer.NewCustomerUseCase(customerRepositoryMock)
		_, err := useCase.CreateCustomer(context.TODO(), createRequest)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, customer.ErrCustomerInvalid)
	})

	t.Run("Should create customer successfully when has valid attributes", func(t *testing.T) {
		createRequest := customer.CustomerCreateDTO{
			Name:  "John Doe",
			Email: "john@email.com",
			Cpf:   "35679254077",
		}

		customerArg := domain.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   CPF("35679254077"),
		}
		customerRepositoryMock = mocks.NewMockCustomerRepository(t)
		customerRepositoryMock.On("Save", &customerArg).Return(nil)

		useCase := customer.NewCustomerUseCase(customerRepositoryMock)
		result, err := useCase.CreateCustomer(context.TODO(), createRequest)

		assert.Nil(t, err)
		assert.Equal(t, result, &domain.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   CPF("35679254077"),
		})
	})
}
