package binance

import "github.com/webdelo/tradebot/pkg/market"

var (
	Symbols = map[string]market.Symbol{
		"btcusdt": Symbol{code: "BTCUSDT"},
		"ethusdt": Symbol{code: "ETHUSDT"},
	}
)

type Symbol struct {
	code string
}

func (s Symbol) Code() string {
	return s.code
}
