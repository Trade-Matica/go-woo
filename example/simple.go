package main

import (
	"fmt"

	v1 "github.com/trading-peter/go-woo/v1"
	v1Types "github.com/trading-peter/go-woo/v1/types"
)

func main() {
	client := v1.NewClient(v1.WithAuth("your api key", "your api secret"))

	// Get client info
	symbols, _ := client.AvailableSymbols()

	for _, asr := range symbols.Rows {
		fmt.Printf("%+v\n", asr.Symbol)
	}

	orderResult, err := client.SendOrder(v1Types.Order{
		Symbol:        "PERP_SOL_USDT",
		OrderType:     v1Types.LIMIT,
		OrderPrice:    90.5,
		OrderQuantity: 1,
		Side:          v1Types.BUY,
	})
	fmt.Printf("%+v\n", orderResult)
	fmt.Printf("%+v\n", err)
}
