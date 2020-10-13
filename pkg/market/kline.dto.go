package market

type KlineDTO struct {
	StartTime int64
	EndTime   int64
	Symbol    Symbol
	Interval  Interval
	Open      int64
	Close     int64
	High      int64
	Low       int64
	Volume    int64
	TradeNum  int64
	IsFinal   bool
}

func (k *KlineDTO) GetStartTime() int64 {
	return k.StartTime
}

func (k *KlineDTO) GetEndTime() int64 {
	return k.EndTime
}

func (k *KlineDTO) GetSymbol() Symbol {
	return k.Symbol
}

func (k *KlineDTO) GetInterval() Interval {
	return k.Interval
}

func (k *KlineDTO) GetOpen() int64 {
	return k.Open
}

func (k *KlineDTO) GetClose() int64 {
	return k.Close
}

func (k *KlineDTO) GetHigh() int64 {
	return k.High
}

func (k *KlineDTO) GetLow() int64 {
	return k.Low
}

func (k *KlineDTO) GetVolume() int64 {
	return k.Volume
}

func (k *KlineDTO) GetTradeNum() int64 {
	return k.TradeNum
}

func (k *KlineDTO) IsCompleted() bool {
	return k.IsFinal
}

func (k *KlineDTO) InProgress() bool {
	return !k.IsFinal
}
