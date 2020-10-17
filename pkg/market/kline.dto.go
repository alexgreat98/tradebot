package market

import "math"

// NewKlineDto constructor for new dto creation
func NewKlineDto(
	startTime int64,
	endTime int64,
	symbol Symbol,
	interval Interval,
	open int64,
	close int64,
	high int64,
	low int64,
	volume int64,
	tradeNum int64,
	isFinal bool,

) *KlineDto {
	return &KlineDto{
		startTime: startTime,
		endTime:   endTime,
		symbol:    symbol,
		interval:  interval,
		open:      open,
		close:     close,
		high:      high,
		low:       low,
		volume:    volume,
		tradeNum:  tradeNum,
		isFinal:   isFinal,

		volumeProfile: NewMapVolumeProfile(),
	}
}

type KlineDto struct {
	startTime int64
	endTime   int64
	symbol    Symbol
	interval  Interval
	open      int64
	close     int64
	high      int64
	low       int64
	volume    int64
	tradeNum  int64
	isFinal   bool

	volumeProfile *MapVolumeProfile
}

func (k *KlineDto) StartTime() int64 {
	return k.startTime
}

func (k *KlineDto) EndTime() int64 {
	return k.endTime
}

func (k *KlineDto) Symbol() Symbol {
	return k.symbol
}

func (k *KlineDto) Interval() Interval {
	return k.interval
}

func (k *KlineDto) Open() int64 {
	return k.open
}

func (k *KlineDto) Close() int64 {
	return k.close
}

func (k *KlineDto) High() int64 {
	return k.high
}

func (k *KlineDto) Low() int64 {
	return k.low
}

func (k *KlineDto) Volume() int64 {
	return k.volume
}

func (k *KlineDto) TradeNum() int64 {
	return k.tradeNum
}

func (k *KlineDto) IsFinal() bool {
	return k.isFinal
}

func (k *KlineDto) InProgress() bool {
	return !k.isFinal
}

func (k *KlineDto) IsRed() bool {
	return k.Open() > k.Close()
}

func (k *KlineDto) IsGreen() bool {
	return !k.IsRed()
}

func (k *KlineDto) VolumeProfile() VolumeProfile {
	return k.volumeProfile
}

func NewMapVolumeProfile() *MapVolumeProfile {
	return &MapVolumeProfile{
		volume: make(map[int64]int64),
		total:  0,
	}
}

// MapVolumeProfile implement VolumeProfile using map storage
type MapVolumeProfile struct {
	volume map[int64]int64
	total  int64
}

// AddTrade attach volume to the map
func (mvp *MapVolumeProfile) AddTrade(trade Trade) *MapVolumeProfile {
	amount, _ := mvp.volume[trade.Price()]
	amount = amount + trade.Quantity()
	mvp.volume[trade.Price()] = amount
	mvp.total += trade.Quantity()
	return mvp
}

// VolumeSum retrieve summary of volume between prices
func (mvp MapVolumeProfile) VolumeSum(startPrice int64, endPrice int64) int64 {
	var sum int64
	for price, volume := range mvp.volume {
		if price >= startPrice && price <= endPrice {
			sum += volume
		}
	}
	return sum
}

// VolumePercent retrieve percentage of volume between prices in comparison with total kline's volume
func (mvp MapVolumeProfile) VolumePercent(startPrice int64, endPrice int64) int {
	sum := mvp.VolumeSum(startPrice, endPrice)
	return int(math.RoundToEven(float64(sum) / float64(mvp.total) * 100))
}
