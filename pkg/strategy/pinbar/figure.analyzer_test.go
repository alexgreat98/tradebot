package pinbar

import (
	binance2 "github.com/webdelo/tradebot/pkg/binance/binance"
	"github.com/webdelo/tradebot/pkg/market"
	"testing"
	"time"
)

func TestIsHammerFigure(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage = market.KlineStorage{}
	klineStorage.SetKline(getBullishHammerKline())
	ans, _ := analyzer.isHammerFigure(klineStorage.GetCurrent())

	if !ans {
		t.Error("Candle has Hammer figure")
	}
}

func TestIsBullishPattern(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage = market.KlineStorage{}
	klineStorage.SetKline(getBullishHammerKline())
	ans := analyzer.isBullishPattern(klineStorage.GetCurrent())

	if !ans {
		t.Error("Candle has Bullish pattern")
	}
}

func TestIsInvertHammerFigure(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage = market.KlineStorage{}
	klineStorage.SetKline(getBullishInvertHammerKline())
	ans, isInvert := analyzer.isHammerFigure(klineStorage.GetCurrent())

	if !ans {
		t.Error("Candle has Hammer figure")
	}

	if !isInvert {
		t.Error("Candle has Invert Hammer figure")
	}
}

func TestIsHangingManFigure(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage = market.KlineStorage{}
	klineStorage.SetKline(getHangingManKline())
	ans := analyzer.isHangingManPattern(klineStorage.GetCurrent())

	if !ans {
		t.Error("Candle has Hanging Man figure")
	}
}

func TestIsShootingStarFigure(t *testing.T) {
	var analyzer = FigureAnalyzer{}
	var klineStorage = market.KlineStorage{}
	klineStorage.SetKline(getShootingStarKline())
	ans := analyzer.isShootingStarPattern(klineStorage.GetCurrent())

	if !ans {
		t.Error("Candle has Shooting Star figure")
	}
}

func getBullishHammerKline() *market.KlineDto {
	return market.NewKlineDto(
		time.Now().Unix(),
		time.Now().Add(1).Unix(),
		binance2.Symbols["BTCUSDT"],
		binance2.Intervals["1m"],
		50,
		65,
		70,
		10,
		4,
		10203,
		true,
	)
}

func getBullishInvertHammerKline() *market.KlineDto {
	return market.NewKlineDto(
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
}

func getHangingManKline() *market.KlineDto {
	return market.NewKlineDto(
		time.Now().Unix(),
		time.Now().Add(1).Unix(),
		binance2.Symbols["BTCUSDT"],
		binance2.Intervals["1m"],
		65,
		50,
		70,
		10,
		4,
		10203,
		true,
	)
}

func getShootingStarKline() *market.KlineDto {
	return market.NewKlineDto(
		time.Now().Unix(),
		time.Now().Add(1).Unix(),
		binance2.Symbols["BTCUSDT"],
		binance2.Intervals["1m"],
		25,
		10,
		70,
		10,
		4,
		10203,
		true,
	)
}
