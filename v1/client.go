package v1

import (
	"net/http"

	i "github.com/trading-peter/go-woo/go-woo/interfaces"
)

const BaseURL = "https://api.woo.org/v1"

func GetPublicInfo(c i.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", BaseURL+"/public/info", nil)
	if err != nil {
		return nil, err
	}
	return c.SendRequestV1(req, false)
}

func GetOrders(c i.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", BaseURL+"/orders", nil)
	if err != nil {
		return nil, err
	}
	return c.SendRequestV1(req, true)
}
