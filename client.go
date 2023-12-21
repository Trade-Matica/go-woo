package gowoo

import (
	"net/http"

	v1 "github.com/trading-peter/go-woo/go-woo/v1"
	v3 "github.com/trading-peter/go-woo/go-woo/v3"
)

type Client struct {
	apiKey     string
	apiSecret  string
	httpClient *http.Client
}

type ClientOption func(c *Client)

func WithAuth(apiKey string, apiSecret string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
		c.apiSecret = apiSecret
	}
}

func WithCustomHttpClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

func NewClient(options ...ClientOption) *Client {
	c := &Client{
		httpClient: &http.Client{},
	}

	for _, optionsFn := range options {
		optionsFn(c)
	}

	return c
}

func (c *Client) GetPublicInfoV1() ([]byte, error) {
	return v1.GetPublicInfo(c)
}

func (c *Client) GetOrdersV1() ([]byte, error) {
	return v1.GetOrders(c)
}

func (c *Client) GetAccountInfoV3() ([]byte, error) {
	return v3.GetAccountInfo(c)
}
