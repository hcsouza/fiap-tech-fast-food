package order

import (
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
)

type orderUseCase struct {
	repository      repository.OrderRepository
	productUseCase  product.IProductUseCase
	customerUseCase customer.ICustomerUseCase
}

func NewOrderUseCase(repo repository.OrderRepository, productUseCase product.IProductUseCase, customerUseCase customer.ICustomerUseCase) IOrderUseCase {
	return &orderUseCase{
		repository:      repo,
		productUseCase:  productUseCase,
		customerUseCase: customerUseCase,
	}
}

func (o *orderUseCase) FindById(id string) (*domain.Order, error) {
	order, err := o.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderUseCase) GetAllByStatus(status OrderStatus) ([]domain.Order, error) {
	orders, err := o.repository.FindAllByStatus(status)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderUseCase) CreateOrder(order OrderCreateDTO) (string, error) {
	var amount float64
	var orderItems []domain.OrderItem
	var err error
	orderItemsDto := order.OrderItemsDTO
	cpf := order.Cpf

	amount, orderItems, err = processProductsAndAmountFromOrderItemDTO(orderItemsDto, o, amount, orderItems)

	if err != nil {
		return "", err
	}

	customer := findCustomerByCpf(order, cpf, o)

	orderToCreate := domain.Order{
		OrderStatus: ORDER_STARTED,
		OrderItems:  orderItems,
		Value:       amount,
		Customer:    customer,
	}

	orderId, err := o.repository.Save(&orderToCreate)

	if err != nil {
		return "", err
	}

	return orderId, nil
}

func findCustomerByCpf(order OrderCreateDTO, cpf string, o *orderUseCase) domain.Customer {
	var customer domain.Customer
	if len(cpf) > 0 {
		cpfMap := map[string]string{
			"cpf": order.Cpf,
		}
		foundCustomer, _ := o.customerUseCase.GetCustomer(nil, cpfMap)

		if foundCustomer != nil {
			customer = domain.Customer{
				CPF:   foundCustomer.CPF,
				Email: foundCustomer.Email,
				Name:  foundCustomer.Name,
			}
		}
	}
	return customer
}

func (o *orderUseCase) UpdateOrder(orderId string, order OrderUpdateDTO) error {
	var amount float64
	var orderItems []domain.OrderItem
	var err error
	orderItemsDto := order.OrderItemsDTO

	amount, orderItems, err = processProductsAndAmountFromOrderItemDTO(orderItemsDto, o, amount, orderItems)
	if err != nil {
		return err
	}

	orderToUpdate := domain.Order{
		ID:          orderId,
		OrderStatus: order.OrderStatus,
		OrderItems:  orderItems,
		Value:       amount,
	}

	err = o.repository.Update(&orderToUpdate)

	if err != nil {
		return err
	}

	return nil
}

func processProductsAndAmountFromOrderItemDTO(orderItemsDto []OrderItemDTO, o *orderUseCase, amount float64, orderItems []domain.OrderItem) (float64, []domain.OrderItem, error) {
	for _, item := range orderItemsDto {
		prod, err := o.productUseCase.FindById(item.ProductId)

		if err != nil {
			return 0, nil, err
		}

		amount += prod.Price * float64(item.Amount)

		itemInOrder := domain.OrderItem{
			Product: *prod,
			Amount:  item.Amount,
		}

		orderItems = append(orderItems, itemInOrder)
	}
	return amount, orderItems, nil
}