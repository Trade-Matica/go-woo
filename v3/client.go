package v3

import (
	"net/http"

	i "github.com/trading-peter/go-woo/go-woo/interfaces"
)

const BaseURL = "https://api.woo.org/v3"

func GetAccountInfo(c i.Client) ([]byte, error) {
	req, err := http.NewRequest("GET", BaseURL+"/accountinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.SendRequestV3(req, nil, true)
}
