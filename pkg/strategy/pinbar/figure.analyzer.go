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
	if fa.isFigurePinbar(klineStorage) {
		if fa.next != nil {
			return fa.next.Check(klineStorage)
		}
		return true
	}
	return false
}

func (fa *FigureAnalyzer) isFigurePinbar(klineStorage *market.KlineStorage) bool {
	current := klineStorage.GetCurrent()
	open := float64(current.Open())
	close := float64(current.Close())
	low := float64(current.Low())
	high := float64(current.High())
	//--- Define of it bullish or bearish
	var bull = open < close
	//--- Get the absolute value of the candlestick body size
	bodySize := math.Abs(open - close)

	//--- Get the size of shadows
	shadeLow := close - low
	shadeHigh := high - open
	if bull {
		shadeLow = open - low
		shadeHigh = high - close
	}
	//--- Calculate the average body size of previous candlesticks
	current.Volume()

	// is "Hammer"
	if shadeLow > bodySize*2 && shadeHigh < (bodySize+shadeLow)*0.1 {
		return true
	}
	// is "invert Hammer"
	if shadeLow < (bodySize+shadeHigh)*0.1 && shadeHigh > bodySize*2 {
		return true
	}

	// TODO: реализовать логику анализа формы фигуры свечи
	// klineStorage.GetCurrent() - анализируй форму текущей свечи из хранилища на соответствие пинбару
	// граничные условия бери из конфиг-структуры так, чтобы их можно было менять на стадии запуска анализатора
	// там пока только одно свойство, можешь добавить туда те, которые будут тут.
	return false
}

func (fa *FigureAnalyzer) isTrue(a int, b int) bool {
	return a == b
}

func (fa *FigureAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
