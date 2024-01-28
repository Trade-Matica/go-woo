package v3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/trading-peter/go-woo/shared"
)

const BaseURL = "https://api.woo.org/v3"

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

func (c *Client) SendRequest(req *http.Request, body []byte, requiresAuth bool) ([]byte, error) {
	if requiresAuth {
		timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
		signature := c.generateSignatureV3(req, timestamp, body)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", c.apiKey)
		req.Header.Set("x-api-signature", signature)
		req.Header.Set("x-api-timestamp", timestamp)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, shared.NewWooError(resp.StatusCode, respBody)
	}

	return respBody, nil
}

func (c *Client) generateSignatureV3(req *http.Request, timestamp string, body []byte) string {
	signString := timestamp + req.Method + req.URL.Path
	if len(body) > 0 {
		signString += string(body)
	}

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(signString))
	return hex.EncodeToString(h.Sum(nil))
}
