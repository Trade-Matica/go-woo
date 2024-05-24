package v1

import (
	"mm/package/woox/shared"
	t "mm/package/woox/v1/types"
)

func (c *Client) GetSystemMaintenanceStatus() (*t.SystemMaintenanceStatusResult, error) {
	raw, err := c.SendRequest("GET", "public/system_info", nil, false)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.SystemMaintenanceStatusResult](raw)
}

func (c *Client) AvailableSymbols() (*t.AvailableSymbolsResult, error) {
	raw, err := c.SendRequest("GET", "public/info", nil, false)
	if err != nil {
		return nil, err
	}

	return shared.UnmarshalTo[t.AvailableSymbolsResult](raw)
}
