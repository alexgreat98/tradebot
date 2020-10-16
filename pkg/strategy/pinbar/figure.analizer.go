package pinbar

import "github.com/webdelo/tradebot/pkg/market"

type FigureAnalyzer struct {
	next Analyzer
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
	// TODO: реализовать логику анализа формы фигуры свечи
	return true
}

func (fa *FigureAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
