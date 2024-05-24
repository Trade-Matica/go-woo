package v1

import (
	"mm/package/woox/shared"
	t "mm/package/woox/v1/types"
	"time"
)

func (c *Client) SendOrder(o t.Order) (*t.SendOrderResult, error) {
	if _, exists := c.limiter.ordersPlacing[o.Symbol]; !exists {
		c.limiter.ordersPlacing[o.Symbol] = shared.NewRateLimiter(5, time.Second)
	}
	c.limiter.ordersPlacing[o.Symbol].Wait()

	raw, err := c.SendRequest("POST", "order", o, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.SendOrderResult](raw)
}

func (c *Client) GetOrders() (*t.GetOrdersResult, error) {
	c.limiter.getOrder.Wait()

	raw, err := c.SendRequest("GET", "orders", nil, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.GetOrdersResult](raw)
}

func (c *Client) CancelOrder(o t.CancelOrder) (*t.CancelOrderResult, error) {
	c.limiter.cancel.Wait()

	raw, err := c.SendRequest("DELETE", "order", o, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.CancelOrderResult](raw)
}
