package binance

var klineIntervals = map[string]string{
	"1m":  "1m",
	"3m":  "3m",
	"5m":  "5m",
	"15m": "15m",
	"30m": "30m",
	"1h":  "1h",
	"2h":  "2h",
	"4h":  "4h",
	"6h":  "6h",
	"8h":  "8h",
	"12h": "12h",
	"1d":  "1d",
	"3d":  "3d",
	"1w":  "1w",
	"1M":  "1M",
}

// KlineIntervals retrieve constant-map with Kline's klineIntervals values
func KlineIntervals() map[string]string {
	return klineIntervals
}

var symbols = map[string]string{
	"BTCUSDT": "BTCUSDT",
	"ETHUSDT": "ETHUSDT",
}

// Symbols retrieve constant-map with symbols values
func Symbols() map[string]string {
	return symbols
}
