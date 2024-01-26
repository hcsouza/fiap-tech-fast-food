package orderStatus

import (
	"fmt"
	"strings"
)

type OrderStatus string

const (
	ORDER_STARTED          OrderStatus = "STARTED"
	ORDER_WAITING_PAYMENT  OrderStatus = "WAITING_PAYMENT"
	ORDER_PAYMENT_RECEIVED OrderStatus = "PAYMENT_RECEIVED"
	ORDER_RECEIVED         OrderStatus = "RECEIVED"
	ORDER_BEING_PREPARED   OrderStatus = "PREPARING"
	ORDER_READY            OrderStatus = "READY"
	ORDER_COMPLETED        OrderStatus = "COMPLETED"
)

func (o OrderStatus) String() string {
	return string(o)
}

func ParseOrderStatus(s string) (o OrderStatus, err error) {
	statuses := map[OrderStatus]struct{}{
		ORDER_STARTED:          {},
		ORDER_WAITING_PAYMENT:  {},
		ORDER_PAYMENT_RECEIVED: {},
		ORDER_RECEIVED:         {},
		ORDER_BEING_PREPARED:   {},
		ORDER_READY:            {},
		ORDER_COMPLETED:        {},
	}

	orderStatus := OrderStatus(strings.ToUpper(s))
	_, ok := statuses[orderStatus]

	if !ok {
		return o, fmt.Errorf(`cannot parse:[%s] as order status`, s)
	}
	return orderStatus, nil
}
