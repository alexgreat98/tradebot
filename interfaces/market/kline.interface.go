package interfaces

// Kline interface is the data contract for all kline entities
type Kline interface {
	GetStartTime() int64
	GetEndTime() int64

	GetSymbol() string
	GetInterval() string

	GetOpen() string
	GetClose() string
	GetHigh() string
	GetLow() string

	GetVolume() string
	GetTradeNum() int64

	IsCompleted() bool
	InProgress() bool
}
