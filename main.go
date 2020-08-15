package main

import "fmt"

var isNextOperationBuy = true
var lastOpPrice = 100.00

// BUY Tresholds - Bot State is SELL
const BULLISH_TRESHOLD = 1.3
const BEARISH_TRESHOLD = -2.0

// SELL Tresholds - Bot State is BUY
const PROFIT_TRESHOLD = 0.7
const STOP_LOSS_TRESHOLD = -0.5

func main() {
	fmt.Println("Starting the bot-service...")

	fmt.Println("Exiting the bot-service...")
}
