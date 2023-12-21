package main

import (
	"fmt"

	gowoo "github.com/trading-peter/go-woo/go-woo"
)

func main() {
	client := gowoo.NewClient(gowoo.WithAuth("xxx", "xxx"))

	// Get client info
	clientInfo, err := client.GetAccountInfoV3()
	if err != nil {
		fmt.Println("Error getting client info:", err)
	} else {
		fmt.Println("Client Info:", string(clientInfo))
	}
}
