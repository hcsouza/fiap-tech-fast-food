package httpClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/qrCodeResponse"
)

type Gateway struct {
	client http.Client
}

func NewGateway(client http.Client) *Gateway {
	return &Gateway{client: client}
}

func (g *Gateway) Post(req *http.Request) (*http.Response, error) {
	bodyBytes, _ := io.ReadAll(req.Body)

	decodedBase64 := base64.StdEncoding.EncodeToString(bodyBytes)

	qrCodeResponseBody := &QRCodeResponse{
		QRCode: decodedBase64,
	}

	responseBodyBytes, err := json.Marshal(qrCodeResponseBody)
	if err != nil {
		return nil, err
	}

	response := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(bytes.NewBuffer(responseBodyBytes)),
	}

	return response, nil
}
