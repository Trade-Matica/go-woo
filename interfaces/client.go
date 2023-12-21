package interfaces

import "net/http"

type Client interface {
	SendRequestV1(req *http.Request, requiresAuth bool) ([]byte, error)
	SendRequestV3(req *http.Request, body []byte, requiresAuth bool) ([]byte, error)
}
