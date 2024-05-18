package types

type OrderType string

const (
	LIMIT     = OrderType("LIMIT")
	MARKET    = OrderType("MARKET")
	IOC       = OrderType("IOC")
	FOK       = OrderType("FOK")
	POST_ONLY = OrderType("POST_ONLY")
	ASK       = OrderType("ASK")
	BID       = OrderType("BID")
)

type OrderSide string

const (
	BUY  = OrderSide("BUY")
	SELL = OrderSide("SELL")
)

type PositionSide string

const (
	LONG  = PositionSide("LONG")
	SHORT = PositionSide("SHORT")
)

type Order struct {
	Symbol          string       `json:"symbol,omitempty" url:"symbol,omitempty"`
	ClientOrderId   int64        `json:"client_order_id,omitempty" url:"client_order_id,omitempty"`
	OrderTag        string       `json:"order_tag,omitempty" url:"order_tag,omitempty"`
	OrderType       OrderType    `json:"order_type,omitempty" url:"order_type,omitempty"`
	OrderPrice      float64      `json:"order_price,omitempty" url:"order_price,omitempty"`
	OrderQuantity   float64      `json:"order_quantity,omitempty" url:"order_quantity,omitempty"`
	OrderAmount     float64      `json:"order_amount,omitempty" url:"order_amount,omitempty"`
	ReduceOnly      bool         `json:"reduce_only,omitempty" url:"reduce_only,omitempty"`
	VisibleQuantity float64      `json:"visible_quantity,omitempty" url:"visible_quantity,omitempty"`
	Side            OrderSide    `json:"side,omitempty" url:"side,omitempty"`
	PositionSide    PositionSide `json:"position_side,omitempty" url:"position_side,omitempty"`
}

type SendOrderResult struct {
	Success       bool      `json:"success"`
	OrderId       int64     `json:"order_id"`
	ClientOrderId int64     `json:"client_order_id"`
	OrderType     OrderType `json:"order_type"`
	OrderPrice    float64   `json:"order_price"`
	OrderQuantity float64   `json:"order_quantity"`
	OrderAmount   float64   `json:"order_amount"`
	ReduceOnly    bool      `json:"reduce_only"`
	Timestamp     string    `json:"timestamp"`
}

type OrderStatusResult struct {
	Success              bool    `json:"success"`
	CreatedTime          string  `json:"created_time"`
	Side                 string  `json:"side"`
	Status               string  `json:"status"`
	Symbol               string  `json:"symbol"`
	ClientOrderID        int     `json:"client_order_id"`
	ReduceOnly           bool    `json:"reduce_only"`
	OrderID              int     `json:"order_id"`
	OrderTag             string  `json:"order_tag"`
	Type                 string  `json:"type"`
	Price                int     `json:"price"`
	Quantity             float64 `json:"quantity"`
	Amount               any     `json:"amount"`
	Visible              float64 `json:"visible"`
	Executed             float64 `json:"executed"`
	TotalFee             float64 `json:"total_fee"`
	FeeAsset             string  `json:"fee_asset"`
	AverageExecutedPrice int     `json:"average_executed_price"`
	RealizedPnl          any     `json:"realized_pnl"`
	PositionSide         string  `json:"position_side"`
	Transactions         []struct {
		ID                int     `json:"id"`
		Symbol            string  `json:"symbol"`
		Fee               float64 `json:"fee"`
		FeeAsset          string  `json:"fee_asset"`
		Side              string  `json:"side"`
		OrderID           int     `json:"order_id"`
		ExecutedPrice     int     `json:"executed_price"`
		ExecutedQuantity  float64 `json:"executed_quantity"`
		ExecutedTimestamp string  `json:"executed_timestamp"`
		IsMaker           int     `json:"is_maker"`
	} `json:"Transactions"`
}

type GetOrdersResult struct {
	Success bool            `json:"success"`
	Meta    GetOrdersMeta   `json:"meta"`
	Rows    []GetOrdersRows `json:"rows"`
}

type GetOrdersMeta struct {
	Total          int `json:"total"`
	RecordsPerPage int `json:"records_per_page"`
	CurrentPage    int `json:"current_page"`
}

type GetOrdersRows struct {
	Side                 OrderSide    `json:"side"`
	Status               string       `json:"status"`
	Symbol               string       `json:"symbol"`
	ClientOrderID        int64        `json:"client_order_id"`
	ReduceOnly           bool         `json:"reduce_only"`
	OrderID              int64        `json:"order_id"`
	OrderTag             string       `json:"order_tag"`
	Type                 OrderType    `json:"type"`
	Price                float64      `json:"price"`
	Quantity             float64      `json:"quantity"`
	Amount               float64      `json:"amount"`
	Visible              float64      `json:"visible"`
	Executed             float64      `json:"executed"`
	TotalFee             float64      `json:"total_fee"`
	FeeAsset             string       `json:"fee_asset"`
	CreatedTime          string       `json:"created_time"`
	UpdatedTime          string       `json:"updated_time"`
	AverageExecutedPrice float64      `json:"average_executed_price"`
	PositionSide         PositionSide `json:"position_side"`
	RealizedPnl          float64      `json:"realized_pnl"`
}

type CancelOrder struct {
	Symbol  string `json:"symbol,omitempty" url:"symbol,omitempty"`
	OrderId int64  `json:"order_id,omitempty" url:"order_id,omitempty"`
}

type CancelOrderResult struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
}
