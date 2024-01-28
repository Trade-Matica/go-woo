package v1

import (
	"net/http"

	"github.com/trading-peter/go-woo/shared"
	t "github.com/trading-peter/go-woo/v1/types"
)

func (c *Client) GetSystemMaintenanceStatus() (*t.SystemMaintenanceStatusResult, error) {
	req, err := http.NewRequest("GET", BaseURL+"/public/system_info", nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.SendRequest(req, nil, false)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.SystemMaintenanceStatusResult](raw)
}

func (c *Client) AvailableSymbols() (*t.AvailableSymbolsResult, error) {
	req, err := http.NewRequest("GET", BaseURL+"/public/info", nil)
	if err != nil {
		return nil, err
	}

	raw, err := c.SendRequest(req, nil, false)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.AvailableSymbolsResult](raw)
}
