# go-woo

API client library for the crypto exchange x.woo.org

Very much wip. I wouldn't really use it just yet.

## Sending orders

```go
import (
	v1 "github.com/trading-peter/go-woo/v1"
	v1Types "github.com/trading-peter/go-woo/v1/types"
)

client := v1.NewClient(v1.WithAuth("your api key", "your api secret"))

result, err := client.SendOrder(v1Types.Order{
  Symbol:        "PERP_SOL_USDT",
  OrderType:     v1Types.LIMIT,
  OrderPrice:    90.5,
  OrderQuantity: 1,
  Side:          v1Types.BUY,
})
```
