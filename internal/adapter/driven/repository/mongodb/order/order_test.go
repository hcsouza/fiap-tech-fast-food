package repository_test

import (
	"errors"
	"testing"

	repository "github.com/hcsouza/fiap-tech-fast-food/internal/adapter/driven/repository/mongodb/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	"github.com/hcsouza/fiap-tech-fast-food/test/mocks"
	"github.com/stretchr/testify/assert"
)

var databaseAdapter *mocks.MockIDatabaseAdapter

func TestMain(m *testing.M) {
	databaseAdapter = &mocks.MockIDatabaseAdapter{}
}

func TestOrderRepository_FindAll(t *testing.T) {
	t.Run("Orders Found", func(t *testing.T) {
		expectedOrders := []interface{}{&domain.Order{ID: "1"}, &domain.Order{ID: "2"}}
		databaseAdapter.On("FindAll", "", "").Return(expectedOrders, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAll()

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("No Orders Found", func(t *testing.T) {
		databaseAdapter.On("FindAll", "", "").Return([]interface{}{}, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAll()

		assert.NoError(t, err)
		assert.Empty(t, result)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Error in DatabaseAdapter", func(t *testing.T) {
		databaseAdapter.On("FindAll", "", "").Return(nil, errors.New("database error"))

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAll()

		assert.Error(t, err)
		assert.Nil(t, result)
		databaseAdapter.AssertExpectations(t)
	})
}

func TestOrderRepository_FindById(t *testing.T) {
	t.Run("Existing Order", func(t *testing.T) {
		expectedOrder := &domain.Order{ID: "123"}
		databaseAdapter.On("FindOne", "_id", "123").Return(expectedOrder, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindById("123")

		assert.NoError(t, err)
		assert.Equal(t, expectedOrder, result)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Order Not Found", func(t *testing.T) {
		databaseAdapter.On("FindOne", "_id", "456").Return(nil, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindById("456")

		assert.NoError(t, err)
		assert.Nil(t, result)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Error in DatabaseAdapter", func(t *testing.T) {
		databaseAdapter.On("FindOne", "_id", "789").Return(nil, errors.New("database error"))

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindById("789")

		assert.Error(t, err)
		assert.Nil(t, result)
		databaseAdapter.AssertExpectations(t)
	})
}

func TestOrderRepository_FindAllByStatus(t *testing.T) {
	t.Run("Orders Found", func(t *testing.T) {
		expectedOrders := []interface{}{&domain.Order{ID: "1"}, &domain.Order{ID: "2"}}
		databaseAdapter.On("FindAll", "orderStatus", "STARTED").Return(expectedOrders, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAllByStatus(ORDER_STARTED)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("No Orders Found", func(t *testing.T) {
		databaseAdapter.On("FindAll", "orderStatus", "COMPLETED").Return([]interface{}{}, nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAllByStatus(ORDER_COMPLETED)

		assert.NoError(t, err)
		assert.Empty(t, result)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Error in DatabaseAdapter", func(t *testing.T) {
		databaseAdapter.On("FindAll", "orderStatus", "PREPARING").Return(nil, errors.New("database error"))

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.FindAllByStatus(ORDER_BEING_PREPARED)

		assert.Error(t, err)
		assert.Nil(t, result)
		databaseAdapter.AssertExpectations(t)
	})
}

func TestOrderRepository_Save(t *testing.T) {
	t.Run("Save Order Successfully", func(t *testing.T) {
		orderToSave := &domain.Order{ID: "123"}
		databaseAdapter.On("Save", orderToSave.ToSaveMongo()).Return("123", nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.Save(orderToSave)

		assert.NoError(t, err)
		assert.Equal(t, "123", result)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Error in DatabaseAdapter", func(t *testing.T) {
		orderToSave := &domain.Order{ID: "123"}
		databaseAdapter.On("Save", orderToSave.ToSaveMongo()).Return(nil, errors.New("database error"))

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		result, err := orderRepository.Save(orderToSave)

		assert.Error(t, err)
		assert.Empty(t, result)
		databaseAdapter.AssertExpectations(t)
	})
}

func TestOrderRepository_Update(t *testing.T) {

	t.Run("Update Order Successfully", func(t *testing.T) {
		orderToUpdate := &domain.Order{ID: "123"}
		databaseAdapter.On("Update", "123", orderToUpdate.ToUpdateMongo()).Return(nil)

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		err := orderRepository.Update(orderToUpdate)

		assert.NoError(t, err)
		databaseAdapter.AssertExpectations(t)
	})

	t.Run("Error in DatabaseAdapter", func(t *testing.T) {
		orderToUpdate := &domain.Order{ID: "456"}
		databaseAdapter.On("Update", "456", orderToUpdate.ToUpdateMongo()).Return(errors.New("database error"))

		orderRepository := repository.NewOrderRepository(databaseAdapter)
		err := orderRepository.Update(orderToUpdate)

		assert.Error(t, err)
		databaseAdapter.AssertExpectations(t)
	})
}
