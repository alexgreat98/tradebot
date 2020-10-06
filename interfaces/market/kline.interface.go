package market

// Kline interface is the data contract for all kline entities
type Kline interface {
	getStartTime() int64
	getEndTime() int64
	getSymbol() string
	getInterval() string
	getOpen() string
	getClose() string
	getHigh() string
	getLow() string
	getVolume() string
	getTradeNum() int64
	IsCompleted() bool
	InProgress() bool
}
