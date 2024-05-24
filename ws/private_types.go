package ws

type ExecutionType int

var (
	ExecutionNew             ExecutionType = 0
	ExecutionReject          ExecutionType = 1
	ExecutionCancelReject    ExecutionType = 2
	ExecutionCancelAllReject ExecutionType = 3
)

type PrivateRequest struct {
	ID     string        `json:"id"`
	Event  string        `json:"event"`
	Topic  string        `json:"topic"`
	Params RequestParams `json:"params"`
}

type RequestParams struct {
	Apikey    string `json:"apikey"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}

type BalanceEvent struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Balances map[string]struct {
			Holding          float64 `json:"holding"`
			Frozen           float64 `json:"frozen"`
			Interest         float64 `json:"interest"`
			PendingShortQty  float64 `json:"pendingShortQty"`
			PendingLongQty   float64 `json:"pendingLongQty"`
			Version          float64 `json:"version"`
			Staked           float64 `json:"staked"`
			Unbonding        float64 `json:"unbonding"`
			Vault            float64 `json:"vault"`
			LaunchpadVault   float64 `json:"launchpadVault"`
			Earn             float64 `json:"earn"`
			AverageOpenPrice float64 `json:"averageOpenPrice"`
			Pnl24H           float64 `json:"pnl24H"`
			Fee24H           float64 `json:"fee24H"`
			MarkPrice        float64 `json:"markPrice"`
		} `json:"balances"`
	} `json:"data"`
}

func (e BalanceEvent) GetTopic() string {
	return e.Topic
}

type ExecutionReportEvent struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		MsgType               ExecutionType `json:"msgType"`
		Symbol                string        `json:"symbol"`
		ClientOrderID         int           `json:"clientOrderId"`
		OrderID               int           `json:"orderId"`
		Type                  string        `json:"type"`
		Side                  string        `json:"side"`
		Quantity              float64       `json:"quantity"`
		Price                 float64       `json:"price"`
		TradeID               int           `json:"tradeId"`
		ExecutedPrice         float64       `json:"executedPrice"`
		ExecutedQuantity      float64       `json:"executedQuantity"`
		Fee                   float64       `json:"fee"`
		FeeAsset              string        `json:"feeAsset"`
		TotalExecutedQuantity float64       `json:"totalExecutedQuantity"`
		AvgPrice              float64       `json:"avgPrice"`
		Status                string        `json:"status"`
		Reason                string        `json:"reason"`
		OrderTag              string        `json:"orderTag"`
		TotalFee              float64       `json:"totalFee"`
		Visible               float64           `json:"visible"`
		Timestamp             int64         `json:"timestamp"`
		ReduceOnly            bool          `json:"reduceOnly"`
		Maker                 bool          `json:"maker"`
	} `json:"data"`
}

func (e ExecutionReportEvent) GetTopic() string {
	return e.Topic
}

type PositionPushEvent struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Positions map[string]struct {
			Holding          int     `json:"holding"`
			PendingLongQty   float64 `json:"pendingLongQty"`
			PendingShortQty  float64 `json:"pendingShortQty"`
			AverageOpenPrice float64 `json:"averageOpenPrice"`
			Pnl24H           float64 `json:"pnl24H"`
			Fee24H           float64 `json:"fee24H"`
			SettlePrice      float64 `json:"settlePrice"`
			MarkPrice        float64 `json:"markPrice"`
			Version          int     `json:"version"`
			OpeningTime      int     `json:"openingTime"`
			Pnl24HPercentage float64 `json:"pnl24HPercentage"`
			AdlQuantile      float64 `json:"adlQuantile"`
		} `json:"positions"`
	} `json:"data"`
}

func (e PositionPushEvent) GetTopic() string {
	return e.Topic
}