package market

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

func (k *KlineDto) IsCompleted() bool {
	return k.isFinal
}

func (k *KlineDto) InProgress() bool {
	return !k.isFinal
}
