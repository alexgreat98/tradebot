package binance

import "gorm.io/gorm"

type Kline struct {
	gorm.Model

	StartTime            int64  `json:"start_time"`
	EndTime              int64  `json:"end_time"`
	Symbol               string `json:"symbol"`
	Interval             string `json:"interval"`
	FirstTradeID         int64  `json:"first_trade_id"`
	LastTradeID          int64  `json:"last_trade_id"`
	Open                 string `json:"open"`
	Close                string `json:"close"`
	High                 string `json:"high"`
	Low                  string `json:"low"`
	Volume               string `json:"volume"`
	TradeNum             int64  `json:"trade_numbers"`
	IsFinal              bool   `json:"is_final"`
	QuoteVolume          string `json:"quote_volume"`
	ActiveBuyVolume      string `json:"active_buy_volume"`
	ActiveBuyQuoteVolume string `json:"active_buy_quote_volume"`
}

func (k *Kline) GetStartTime() int64 {
	return k.StartTime
}

func (k *Kline) GetEndTime() int64 {
	return k.EndTime
}

func (k *Kline) GetSymbol() string {
	return k.Symbol
}

func (k *Kline) GetInterval() string {
	return k.Interval
}

func (k *Kline) GetOpen() string {
	return k.Open
}

func (k *Kline) GetClose() string {
	return k.Close
}

func (k *Kline) GetHigh() string {
	return k.High
}

func (k *Kline) GetLow() string {
	return k.Low
}

func (k *Kline) GetVolume() string {
	return k.Volume
}

func (k *Kline) GetTradeNum() int64 {
	return k.TradeNum
}

func (k *Kline) IsCompleted() bool {
	return k.IsFinal
}

func (k *Kline) InProgress() bool {
	return !k.IsFinal
}
