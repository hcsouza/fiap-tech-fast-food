package paymentGateway

import (
	"bytes"
	"encoding/json"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/gateway"
	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/qrCodeResponse"
	"io"
	"net/http"
)

type paymentGateway struct {
	gateway gateway.IGateway
}

const paymentGatewayURL = "www.payment-gateway.url"

type OrderPaymentGetQRCodeRequest struct {
	OrderId string  `json:"orderId"`
	Value   float64 `json:"value"`
}

func NewPaymentGateway(gateway gateway.IGateway) *paymentGateway {
	return &paymentGateway{gateway: gateway}
}

func (p *paymentGateway) GetQRCodeFromOrder(orderId string, value float64) (QRCodeResponse, error) {
	body := &OrderPaymentGetQRCodeRequest{
		OrderId: orderId,
		Value:   value,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)

	req, err := http.NewRequest(http.MethodPost, paymentGatewayURL, payloadBuf)
	if err != nil {
		return QRCodeResponse{}, err
	}

	resp, err := p.gateway.Post(req)
	if err != nil {
		return QRCodeResponse{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return QRCodeResponse{}, err
	}

	var qrCodeResponse QRCodeResponse
	err = json.Unmarshal(responseBody, &qrCodeResponse)
	if err != nil {
		return QRCodeResponse{}, err
	}

	return qrCodeResponse, nil
}
