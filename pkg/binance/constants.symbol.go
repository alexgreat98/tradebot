package binance

import "github.com/webdelo/tradebot/pkg/market"

var (
	Symbols = map[string]market.Symbol{
		"btcusdt": market.Symbol{Code: "BTCUSDT"},
		"ethusdt": market.Symbol{Code: "ETHUSDT"},
	}
)
