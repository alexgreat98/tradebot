package binance

import "github.com/webdelo/tradebot/pkg/market"

var (
	Intervals = map[string]market.Interval{
		"1m":  Interval{code: "1m", seconds: 60},
		"3m":  Interval{code: "3m", seconds: 180},
		"5m":  Interval{code: "3m", seconds: 300},
		"15m": Interval{code: "15m", seconds: 900},
		"30m": Interval{code: "30m", seconds: 1800},
		"1h":  Interval{code: "1h", seconds: 3600},
		"2h":  Interval{code: "2h", seconds: 7200},
		"4h":  Interval{code: "4h", seconds: 14400},
		"6h":  Interval{code: "6h", seconds: 21600},
		"8h":  Interval{code: "8h", seconds: 28800},
		"12h": Interval{code: "12h", seconds: 43200},
		"1d":  Interval{code: "1d", seconds: 86400},
		"3d":  Interval{code: "3d", seconds: 259200},
		"1w":  Interval{code: "1w", seconds: 604800},
		"1M":  Interval{code: "1M", seconds: 2592000},
	}
)

type Interval struct {
	code    string
	seconds int
}

func (i Interval) String() string {
	return i.code
}

func (i Interval) Seconds() int {
	return i.seconds
}
