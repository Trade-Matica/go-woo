package ws

type WSEventI interface {
	GetTopic() string
}

type WSRequest struct {
	ID    string `json:"id"`
	Event string `json:"event"`
	Topic string `json:"topic"`
}

type BestBookOfferEvent struct {
	Topic string `json:"topic"`
	Ts    int64  `json:"ts"`
	Data  struct {
		Symbol  string  `json:"symbol"`
		Ask     float64 `json:"ask"`
		AskSize float64 `json:"askSize"`
		Bid     float64 `json:"bid"`
		BidSize float64 `json:"bidSize"`
	} `json:"data"`
}

func (e BestBookOfferEvent) GetTopic() string {
	return e.Topic
}
