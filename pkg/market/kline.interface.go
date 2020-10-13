package market

// Kline interface is the data contract for all kline entities
type Kline interface {
	GetStartTime() int64
	GetEndTime() int64

	GetSymbol() Symbol
	GetInterval() Interval

	GetOpen() int64
	GetClose() int64
	GetHigh() int64
	GetLow() int64

	GetVolume() int64
	GetTradeNum() int64

	IsCompleted() bool
	InProgress() bool
}
