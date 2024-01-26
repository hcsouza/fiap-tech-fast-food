package gateway

import (
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/qrCodeResponse"
)

type PaymentGateway interface {
	GetQRCodeFromOrder(orderId string, value float64) (QRCodeResponse, error)
}
