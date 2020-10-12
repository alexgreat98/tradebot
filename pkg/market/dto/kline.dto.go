package dto

type Kline struct {
	StartTime uint64 `json:"start_time"`
	EndTime   uint64 `json:"end_time"`
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Open      uint64 `json:"open"`
	Close     uint64 `json:"close"`
	High      uint64 `json:"high"`
	Low       uint64 `json:"low"`
	Volume    uint64 `json:"volume"`
	TradeNum  uint   `json:"trade_numbers"`
	IsFinal   bool   `json:"is_final"`
}
