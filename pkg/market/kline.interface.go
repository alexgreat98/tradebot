package market

// Kline interface is the data contract for all kline entities
type Kline interface {
	GetStartTime() int64
	GetEndTime() int64

	GetSymbol() Symbol
	GetInterval() Interval

	GetOpen() string
	GetClose() string
	GetHigh() string
	GetLow() string

	GetVolume() string
	GetTradeNum() int

	IsCompleted() bool
	InProgress() bool
}
