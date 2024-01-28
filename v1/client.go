package v1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/trading-peter/go-woo/shared"
)

const BaseURL = "https://api.woo.org/v1"

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

func (c *Client) SendRequest(req *http.Request, data any, requiresAuth bool) ([]byte, error) {
	var params url.Values
	var err error

	if data != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		params, err = query.Values(data)
		if err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(strings.NewReader(params.Encode()))
	}

	if requiresAuth {
		signature, timestamp := c.generateSignature(params.Encode())
		req.Header.Set("x-api-key", c.apiKey)
		req.Header.Set("x-api-signature", signature)
		req.Header.Set("x-api-timestamp", timestamp)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, shared.NewWooError(resp.StatusCode, body)
	}

	return body, nil
}

func (c *Client) generateSignature(paramStr string) (signature string, timestamp string) {
	timestamp = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	payload := paramStr + "|" + timestamp

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(payload))
	signature = hex.EncodeToString(h.Sum(nil))
	return
}
