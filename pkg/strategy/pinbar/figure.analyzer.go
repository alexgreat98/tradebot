package pinbar

import (
	"github.com/webdelo/tradebot/pkg/market"
	"math"
)

type FigureAnalyzer struct {
	next   Analyzer
	config Config
}

func (fa *FigureAnalyzer) Check(klineStorage *market.KlineStorage) bool {
	var isHammerFigure, _ = fa.isHammerFigure(klineStorage.GetCurrent())
	if isHammerFigure {
		if fa.next != nil {
			return fa.next.Check(klineStorage)
		}
		return true
	}
	return false
}

func (fa *FigureAnalyzer) isHammerFigure(kline market.Kline) (bool, bool) {
	openValue := float64(kline.Open())
	closeValue := float64(kline.Close())
	lowValue := float64(kline.Low())
	highValue := float64(kline.High())
	//--- Define of it bullish or bearish
	var bull = fa.isBullishPattern(kline)
	//--- Get the absolute value of the candlestick body size
	bodySize := math.Abs(openValue - closeValue)

	//--- Get the size of shadows
	shadeLow := closeValue - lowValue
	shadeHigh := highValue - openValue
	if bull {
		shadeLow = openValue - lowValue
		shadeHigh = highValue - closeValue
	}

	fullSize := bodySize + shadeLow + shadeHigh

	// is "Hammer"
	if shadeLow >= bodySize*2 && shadeHigh <= fullSize*0.1 && shadeHigh <= fullSize*0.3 {
		return true, false
	}
	// is "invert Hammer"
	if shadeLow <= fullSize*0.1 && shadeHigh >= bodySize*2 && shadeLow <= fullSize*0.3 {
		return true, true
	}

	return false, false
}

func (fa *FigureAnalyzer) isHangingManPattern(kline market.Kline) bool {
	isHammer, isInvert := fa.isHammerFigure(kline)

	return (isHammer && !isInvert) && !fa.isBullishPattern(kline)
}

func (fa *FigureAnalyzer) isShootingStarPattern(kline market.Kline) bool {
	isHammer, isInvert := fa.isHammerFigure(kline)

	return (isHammer && isInvert) && !fa.isBullishPattern(kline)
}

func (fa *FigureAnalyzer) isBullishPattern(kline market.Kline) bool {
	return kline.Open() < kline.Close()
}

func (fa *FigureAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
