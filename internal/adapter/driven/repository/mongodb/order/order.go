package repository

import (
	"fmt"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type orderRepository struct {
	databaseAdapter repository.IDatabaseAdapter
}

func NewOrderRepository(databaseAdapter repository.IDatabaseAdapter) *orderRepository {
	return &orderRepository{databaseAdapter: databaseAdapter}
}

func (or orderRepository) FindById(id string) (*domain.Order, error) {
	order, err := or.databaseAdapter.FindOne("_id", id)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	foundOrder := order.(*domain.Order)

	return foundOrder, nil
}

func (or orderRepository) FindAllByStatus(status OrderStatus) ([]domain.Order, error) {
	orders, err := or.databaseAdapter.FindAll("orderStatus", string(status))

	if err != nil {
		return nil, err
	}

	foundOrders := []domain.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(domain.Order))
	}

	return foundOrders, nil
}

func (or orderRepository) Save(order *domain.Order) (string, error) {
	insertResult, err := or.databaseAdapter.Save(
		order.ToSaveMongo(),
	)

	if err != nil {
		return "", err
	}

	fmt.Println(insertResult)

	orderInserted := insertResult.(string)
	return orderInserted, nil
}

func (or orderRepository) Update(order *domain.Order) error {
	_, err := or.databaseAdapter.Update(
		order.ID,
		order.ToUpdateMongo())

	if err != nil {
		return err
	}

	return nil
}
