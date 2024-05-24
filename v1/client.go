package v1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"mm/package/woox/shared"

	"github.com/google/go-querystring/query"
)

const BaseURL = "https://api.woo.org/v1"

var httpClient = &http.Client{}

type limiter struct {
	cancel        *shared.RateLimiter
	getOrder      *shared.RateLimiter
	ordersPlacing map[string]*shared.RateLimiter
}

type Client struct {
	apiKey     string
	apiSecret  string
	httpClient *http.Client
	limiter    limiter
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
		limiter: limiter{
			cancel:        shared.NewRateLimiter(10, time.Second),
			getOrder:      shared.NewRateLimiter(10, time.Second),
			ordersPlacing: make(map[string]*shared.RateLimiter),
		},
	}

	for _, optionsFn := range options {
		optionsFn(c)
	}

	return c
}

func (c *Client) SendRequest(method, endpoint string, data any, requiresAuth bool) ([]byte, error) {
	var payloadData io.Reader
	var encodedParams string

	if data != nil {
		params, err := query.Values(data)
		if err != nil {
			return nil, err
		}

		encodedParams = params.Encode()
		payloadData = strings.NewReader(encodedParams)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", BaseURL, endpoint), payloadData)
	if err != nil {
		return nil, err
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if requiresAuth {
		signature, timestamp := c.generateSignature(encodedParams)

		req.Header.Set("x-api-key", c.apiKey)
		req.Header.Set("x-api-timestamp", timestamp)
		req.Header.Set("x-api-signature", signature)
	}

	resp, err := httpClient.Do(req)
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

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(fmt.Sprintf("%s|%s", paramStr, timestamp)))

	return hex.EncodeToString(h.Sum(nil)), timestamp
}
