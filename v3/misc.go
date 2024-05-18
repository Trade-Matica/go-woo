package v3

import (
	"net/http"

	"github.com/trading-peter/go-woo/shared"
	t "github.com/trading-peter/go-woo/v3/types"
)

func (c *Client) GetAccountInfo() (*t.GetAccountInfoResult, error) {
	req, err := http.NewRequest("GET", "/accountinfo", nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.SendRequest(req, nil, true)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.GetAccountInfoResult](raw)
}
