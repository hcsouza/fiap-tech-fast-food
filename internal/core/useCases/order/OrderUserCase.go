package order

import (
	"sort"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	paymentGateway "github.com/hcsouza/fiap-tech-fast-food/internal/core/gateway"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/repository"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	qrCodeResp "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/qrCodeResponse"
)

type orderUseCase struct {
	repository      repository.OrderRepository
	productUseCase  product.IProductUseCase
	customerUseCase customer.ICustomerUseCase
	paymentGateway  paymentGateway.PaymentGateway
}

func NewOrderUseCase(repo repository.OrderRepository, productUseCase product.IProductUseCase, customerUseCase customer.ICustomerUseCase, paymentGateway paymentGateway.PaymentGateway) IOrderUseCase {
	return &orderUseCase{
		repository:      repo,
		productUseCase:  productUseCase,
		customerUseCase: customerUseCase,
		paymentGateway:  paymentGateway,
	}
}

func (o *orderUseCase) FindAll() ([]domain.Order, error) {
	orders, err := o.repository.FindAll()

	if err != nil {
		return nil, err
	}

	sort.Slice(orders, func(secondIndex, firstIndex int) bool {
		return sortByCreatedAt(orders[firstIndex], orders[secondIndex])
	})

	sort.Slice(orders, func(secondIndex, firstIndex int) bool {
		return sortByStatus(orders[firstIndex], orders[secondIndex])
	})

	var filtredOrders []domain.Order

	for _, order := range orders {
		if order.OrderStatus != ORDER_COMPLETED {
			filtredOrders = append(filtredOrders, order)
		}
	}

	return filtredOrders, nil
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

		amount += prod.Price * float64(item.Quantity)

		itemInOrder := domain.OrderItem{
			Product:  *prod,
			Quantity: item.Quantity,
		}

		orderItems = append(orderItems, itemInOrder)
	}
	return amount, orderItems, nil
}

func (o *orderUseCase) Checkout(orderId string) (qrCodeResp.QRCodeResponse, error) {
	order, err := o.FindById(orderId)
	if err != nil {
		return qrCodeResp.QRCodeResponse{}, err
	}

	qrcode, err := o.paymentGateway.GetQRCodeFromOrder(orderId, order.Value)
	if err != nil {
		return qrCodeResp.QRCodeResponse{}, err
	}

	err = o.updateOrderStatus(*order, ORDER_WAITING_PAYMENT)
	if err != nil {
		return qrCodeResp.QRCodeResponse{}, err
	}

	return qrcode, nil
}

func (o *orderUseCase) ConfirmPayment(orderId string) error {
	return o.UpdateOrderStatus(orderId, ORDER_PAYMENT_RECEIVED)
}

func (o *orderUseCase) UpdateOrderStatus(orderId string, status OrderStatus) error {
	order, err := o.FindById(orderId)
	if err != nil {
		return err
	}

	err = o.updateOrderStatus(*order, status)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderUseCase) updateOrderStatus(order domain.Order, newStatus OrderStatus) error {
	order.OrderStatus = newStatus
	return o.repository.Update(&order)
}

func sortByStatus(firstOrder domain.Order, secondOrder domain.Order) bool {
	return (secondOrder.OrderStatus == ORDER_READY ||
		(secondOrder.OrderStatus == ORDER_BEING_PREPARED && firstOrder.OrderStatus != ORDER_READY)) &&
		secondOrder.OrderStatus != firstOrder.OrderStatus
}

func sortByCreatedAt(firstOrder domain.Order, secondOrder domain.Order) bool {
	return !secondOrder.CreatedAt.Equal(firstOrder.CreatedAt.Time) &&
		secondOrder.CreatedAt.Before(firstOrder.CreatedAt.Time)
}
