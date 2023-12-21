package gowoo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

func (c *Client) SendRequestV3(req *http.Request, body []byte, requiresAuth bool) ([]byte, error) {
	if requiresAuth {
		timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
		signature := c.generateSignatureV3(req, timestamp, body)
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

func (c *Client) SendRequestV1(req *http.Request, requiresAuth bool) ([]byte, error) {
	if requiresAuth {
		timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
		signature := c.generateSignatureV1(req, timestamp)
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

	return body, nil
}

func (c *Client) generateSignatureV1(req *http.Request, timestamp string) string {
	var params []string
	query := req.URL.Query()
	for k, v := range query {
		params = append(params, fmt.Sprintf("%s=%s", k, v[0]))
	}
	sort.Strings(params)
	payload := strings.Join(params, "&") + "|" + timestamp

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
