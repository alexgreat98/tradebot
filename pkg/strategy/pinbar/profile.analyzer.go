package pinbar

import "github.com/webdelo/tradebot/pkg/market"

type VolumeProfileAnalyzer struct {
	next   Analyzer
	config Config
}

func (fa *VolumeProfileAnalyzer) Check(klineStorage *market.KlineStorage) bool {
	if fa.isVolumeProfileValid(klineStorage) {
		if fa.next != nil {
			return fa.next.Check(klineStorage)
		}
		return true
	}
	return false
}

func (fa *VolumeProfileAnalyzer) isVolumeProfileValid(klineStorage *market.KlineStorage) bool {
	// TODO: реализовать логику анализа распределения объёма
	return true
}

func (fa *VolumeProfileAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
