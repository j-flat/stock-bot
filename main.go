package main

import (
	"belfort-bot/api"
	"fmt"
)

// BUY Tresholds - Bot State is SELL
const BULLISH_TRESHOLD = 1.3
const BEARISH_TRESHOLD = -2.0

// SELL Tresholds - Bot State is BUY
const PROFIT_TRESHOLD = 0.7
const STOP_LOSS_TRESHOLD = -0.5

var lastOpPrice *int
var isNextOperationBuy *bool

func evaluateTresholds(changePercent, lowTreshold, highTreshold float32) bool {
	if changePercent >= highTreshold || changePercent <= lowTreshold {
		return true
	}

	return false
}

func getPercentageDiff(currentPrice, lastOpPrice int16) float32 {
	return float32((currentPrice - lastOpPrice) / lastOpPrice * 100)
}

func main() {
	fmt.Println("Starting the bot-service...")

	startingPrice := 10000
	startWithBuy := true

	// Assign pointers the initial values
	lastOpPrice = &startingPrice
	isNextOperationBuy = &startWithBuy

	currentBalance := api.GetBalance()

	fmt.Sprintln("Making trade with balance of:%d", currentBalance)

	fmt.Println("Exiting the bot-service...")

}
