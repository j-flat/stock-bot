package trading

import (
	"belfort-bot/trading/api"
	"fmt"
)

var isNextOperationBuy = true
var lastOpPrice *float32
var defaultPrice = 100.00
lastOpPrice = &defaultPrice

// BUY Tresholds - Bot State is SELL
const BULLISH_TRESHOLD = 1.3
const BEARISH_TRESHOLD = -2.0

// SELL Tresholds - Bot State is BUY
const PROFIT_TRESHOLD = 0.7
const STOP_LOSS_TRESHOLD = -0.5

func sell(pctDiff float32) {
	if pctDiff >= PROFIT_TRESHOLD || pctDiff <= STOP_LOSS_TRESHOLD {
		if api.PlaceSellOrder() != nil {
			fmt.Println("")
		}
	}
}
