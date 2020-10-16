package pinbar

import "github.com/webdelo/tradebot/pkg/market"

type Analyzer interface {
	// Check that is pattern condition valid
	Check(klineStorage *market.KlineStorage) bool

	// SetNext analyzer will be execute in chain
	SetNext(analyzer Analyzer) Analyzer
}
