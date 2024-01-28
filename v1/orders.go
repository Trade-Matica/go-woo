package v1

import (
	"net/http"

	"github.com/trading-peter/go-woo/shared"
	t "github.com/trading-peter/go-woo/v1/types"
)

func (c *Client) SendOrder(o t.Order) (*t.SendOrderResult, error) {
	req, err := http.NewRequest("POST", BaseURL+"/order", nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.SendRequest(req, o, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.SendOrderResult](raw)
}

func (c *Client) GetOrders() (*t.GetOrdersResult, error) {
	req, err := http.NewRequest("GET", BaseURL+"/orders", nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.SendRequest(req, nil, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.GetOrdersResult](raw)
}
