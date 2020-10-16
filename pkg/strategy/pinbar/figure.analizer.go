package pinbar

import "github.com/webdelo/tradebot/pkg/market"

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
	// TODO: реализовать логику анализа формы фигуры свечи
	// klineStorage.GetCurrent() - анализируй форму текущей свечи из хранилища на соответствие пинбару
	// граничные условия бери из конфиг-структуры так, чтобы их можно было менять на стадии запуска анализатора
	// там пока только одно свойство, можешь добавить туда те, которые будут тут.
	return true
}

func (fa *FigureAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
