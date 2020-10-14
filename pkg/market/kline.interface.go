package market

// Kline interface is the data contract for all kline entities
type Kline interface {
	StartTime() int64
	EndTime() int64

	Symbol() Symbol
	Interval() Interval

	Open() int64
	Close() int64
	High() int64
	Low() int64

	Volume() int64
	TradeNum() int64

	IsCompleted() bool
	InProgress() bool
}
