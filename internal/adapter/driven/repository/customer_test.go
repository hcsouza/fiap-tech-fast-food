package repository_test

import (
	"testing"

	"github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	t.Run("Should find a customer by CPF", func(t *testing.T) {
		customerRepository := repository.NewCustomerRespository(nil)
		cpf := CPF("12345678900")

		customer, err := customerRepository.Find(cpf)

		assert.Nil(t, err)
		assert.NotNil(t, customer)
	})
	t.Run("Should save a customer", func(t *testing.T) {
		customerRepository := repository.NewCustomerRespository(nil)
		customer := &domain.Customer{
			Name:  "John Doe",
			CPF:   CPF("12345678900"),
			Email: "john.doe@email.com",
		}

		err := customerRepository.Save(customer)

		assert.Nil(t, err)
	})
}
