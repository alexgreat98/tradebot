package binance

import (
	"github.com/webdelo/tradebot/pkg/market"
	"gorm.io/gorm"
)

type Kline struct {
	gorm.Model

	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Symbol       string `json:"symbol"`
	Interval     string `json:"interval"`
	FirstTradeID int64  `json:"first_trade_id"`
	LastTradeID  int64  `json:"last_trade_id"`
	Open         int64  `json:"open"`
	Close        int64  `json:"close"`
	High         int64  `json:"high"`
	Low          int64  `json:"low"`
	Volume       int64  `json:"volume"`
	TradeNum     int64  `json:"trade_numbers"`
	IsFinal      bool   `json:"is_final"`
	//QuoteVolume          int64  `json:"quote_volume"`
	//ActiveBuyVolume      int64  `json:"active_buy_volume"`
	//ActiveBuyQuoteVolume int64  `json:"active_buy_quote_volume"`
}

func (k *Kline) GetStartTime() int64 {
	return k.StartTime
}

func (k *Kline) GetEndTime() int64 {
	return k.EndTime
}

func (k *Kline) GetSymbol() market.Symbol {
	return Symbols[k.Symbol]
}

func (k *Kline) GetInterval() market.Interval {
	return Intervals[k.Interval]
}

func (k *Kline) GetOpen() int64 {
	return k.Open
}

func (k *Kline) GetClose() int64 {
	return k.Close
}

func (k *Kline) GetHigh() int64 {
	return k.High
}

func (k *Kline) GetLow() int64 {
	return k.Low
}

func (k *Kline) GetVolume() int64 {
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
