package types

type SystemMaintenanceStatusResult struct {
	Success   bool                        `json:"success"`
	Timestamp int64                       `json:"timestamp"`
	Data      SystemMaintenanceStatusData `json:"data"`
}

type SystemMaintenanceStatusData struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type AvailableSymbolsResult struct {
	Success bool                   `json:"success"`
	Rows    []AvailableSymbolsRows `json:"rows"`
}

type AvailableSymbolsRows struct {
	CreatedTime string  `json:"created_time"`
	UpdatedTime string  `json:"updated_time"`
	Symbol      string  `json:"symbol"`
	QuoteMin    float64 `json:"quote_min"`
	QuoteMax    float64 `json:"quote_max"`
	QuoteTick   float64 `json:"quote_tick"`
	BaseMin     float64 `json:"base_min"`
	BaseMax     float64 `json:"base_max"`
	BaseTick    float64 `json:"base_tick"`
	MinNotional float64 `json:"min_notional"`
	PriceRange  float64 `json:"price_range"`
	PriceScope  float64 `json:"price_scope"`
	Precisions  []int   `json:"precisions"`
}
