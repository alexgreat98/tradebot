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

	IsFinal() bool
	InProgress() bool

	IsRed() bool
	IsGreen() bool

	VolumeProfile() VolumeProfile
}

// VolumeProfile interface retrieve methods for analyzing volume range
type VolumeProfile interface {
	// VolumeSum retrieve summary of volume between prices
	VolumeSum(startPrice int64, endPrice int64) int64

	// VolumePercent retrieve percentage of volume between prices in comparison with total kline's volume
	VolumePercent(startPrice int64, endPrice int64) int
}
