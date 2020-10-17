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

	// is "Hammer"
	if shadeLow > bodySize*2 && shadeHigh < (bodySize+shadeLow)*0.1 {
		return true, false
	}
	// is "invert Hammer"
	if shadeLow < (bodySize+shadeHigh)*0.1 && shadeHigh > bodySize*2 {
		return true, true
	}

	// TODO: реализовать логику анализа формы фигуры свечи
	// klineStorage.GetCurrent() - анализируй форму текущей свечи из хранилища на соответствие пинбару
	// граничные условия бери из конфиг-структуры так, чтобы их можно было менять на стадии запуска анализатора
	// там пока только одно свойство, можешь добавить туда те, которые будут тут.
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
