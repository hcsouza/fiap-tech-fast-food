package repository_test

import (
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	databaseAdapter := &mocks.MockDatabaseAdapter{}
	t.Run("Should find a customer by CPF", func(t *testing.T) {
		cpf := CPF("12345678900")

		databaseAdapter.On("FindOne", string(cpf)).Return(&domain.Customer{}, nil)

		customerRepository := repository.NewCustomerRespository(databaseAdapter)

		customer, err := customerRepository.Find(cpf)

		assert.Nil(t, err)
		assert.NotNil(t, customer)
	})
	t.Run("Should save a customer", func(t *testing.T) {
		cpf := CPF("12345678900")
		customer := &domain.Customer{
			Name:  "John Doe",
			CPF:   CPF("12345678900"),
			Email: "john.doe@email.com",
		}

		databaseAdapter.
			On(
				"Save",
				string(cpf),
				map[string]interface{}{"name": customer.Name, "email": customer.Email},
			).
			Return(nil)

		customerRepository := repository.NewCustomerRespository(databaseAdapter)

		err := customerRepository.Save(customer)

		assert.Nil(t, err)
	})
}
