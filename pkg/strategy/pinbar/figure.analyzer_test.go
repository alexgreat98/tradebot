package pinbar

import (
	binance2 "github.com/webdelo/tradebot/pkg/binance/binance"
	"github.com/webdelo/tradebot/pkg/market"
	"testing"
	"time"
)

func TestIsFigurePinbar(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage market.KlineStorage = market.KlineStorage{}
	var klineDto = market.NewKlineDto(
		time.Now().Unix(),
		time.Now().Add(1).Unix(),
		binance2.Symbols["BTCUSDT"],
		binance2.Intervals["1m"],
		15,
		25,
		70,
		10,
		4,
		10203,
		true,
	)
	klineStorage.SetKline(klineDto)
	ans := analyzer.isFigurePinbar(&klineStorage)

	if !ans {
		t.Error("Candle is Hammer want true")
	}
}
