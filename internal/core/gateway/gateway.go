package gateway

import "net/http"

type IGateway interface {
	Post(req *http.Request) (*http.Response, error)
}
