package pinbar

import "github.com/webdelo/tradebot/pkg/market"

type VolumeAnalyzer struct {
	next Analyzer
}

func (fa *VolumeAnalyzer) Check(klineStorage *market.KlineStorage) bool {
	if fa.isVolumeValid(klineStorage) {
		if fa.next != nil {
			return fa.next.Check(klineStorage)
		}
		return true
	}
	return false
}

func (fa *VolumeAnalyzer) isVolumeValid(klineStorage *market.KlineStorage) bool {
	// TODO: реализовать логику анализа общего объёма
	return true
}

func (fa *VolumeAnalyzer) SetNext(analyzer Analyzer) Analyzer {
	fa.next = analyzer
	return fa
}
