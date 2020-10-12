package binance

import "github.com/webdelo/tradebot/pkg/market"

var (
	Intervals = map[string]market.Interval{
		"1m":  market.Interval{Code: "1m", Seconds: 60},
		"3m":  market.Interval{Code: "3m", Seconds: 180},
		"5m":  market.Interval{Code: "3m", Seconds: 300},
		"15m": market.Interval{Code: "15m", Seconds: 900},
		"30m": market.Interval{Code: "30m", Seconds: 1800},
		"1h":  market.Interval{Code: "1h", Seconds: 3600},
		"2h":  market.Interval{Code: "2h", Seconds: 7200},
		"4h":  market.Interval{Code: "4h", Seconds: 14400},
		"6h":  market.Interval{Code: "6h", Seconds: 21600},
		"8h":  market.Interval{Code: "8h", Seconds: 28800},
		"12h": market.Interval{Code: "12h", Seconds: 43200},
		"1d":  market.Interval{Code: "1d", Seconds: 86400},
		"3d":  market.Interval{Code: "3d", Seconds: 259200},
		"1w":  market.Interval{Code: "1w", Seconds: 604800},
		"1M":  market.Interval{Code: "1M", Seconds: 2592000},
	}

	Symbols = map[string]market.Symbol{
		"btcusdt": market.Symbol{Code: "BTCUSDT"},
		"ethusdt": market.Symbol{Code: "ETHUSDT"},
	}
)
