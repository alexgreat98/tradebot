package utils

import (
	"fmt"
	"math"
	"strconv"
)

// Precision lists for currency types
var Precision = map[string]int{
	"btc":  8,
	"eth":  8,
	"usd":  2,
	"usdt": 2,
	"rub":  2,
	"eur":  2,
}

// ToCoins convert float amount to coins according precision for the currency (ex. precision: USD = 2, BTC = 8, ETH = 16)
func ToCoins(amount float64, precision int) int64 {
	return int64(math.Round(amount * math.Pow10(precision)))
}

// ToCoinsFromString convert string-float (ex. "14.54") amount to coins according precision for the currency (ex. precision: USD = 2, BTC = 8, ETH = 16)
func ToCoinsFromString(amount string, precision int) (int64, error) {
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}
	return ToCoins(f, precision), nil
}

// FromCoins convert int amount in coins to float according precision for the currency (ex. precision: USD = 2, BTC = 8, ETH = 16)
func FromCoins(amount int64, precision int) float64 {
	return float64(amount) / math.Pow10(precision)
}

// FromCoinsToString convert int amount in coins to string according precision for the currency (ex. precision: USD = 2, BTC = 8, ETH = 16)
func FromCoinsToString(amount int64, precision int) string {
	return fmt.Sprintf("%f", FromCoins(amount, precision))
}
