package binance

import (
	"github.com/webdelo/tradebot/pkg/market"
	"gorm.io/gorm"
)

type Kline struct {
	gorm.Model

	startTime    int64  `json:"start_time"`
	endTime      int64  `json:"end_time"`
	symbol       string `json:"symbol"`
	interval     string `json:"interval"`
	firstTradeID int64  `json:"first_trade_id"`
	lastTradeID  int64  `json:"last_trade_id"`
	open         int64  `json:"open"`
	close        int64  `json:"close"`
	high         int64  `json:"high"`
	low          int64  `json:"low"`
	volume       int64  `json:"volume"`
	tradeNum     int64  `json:"trade_numbers"`
	isFinal      bool   `json:"is_final"`
	//QuoteVolume          int64  `json:"quote_volume"`
	//ActiveBuyVolume      int64  `json:"active_buy_volume"`
	//ActiveBuyQuoteVolume int64  `json:"active_buy_quote_volume"`
}

func (k *Kline) StartTime() int64 {
	return k.startTime
}

func (k *Kline) EndTime() int64 {
	return k.endTime
}

func (k *Kline) Symbol() market.Symbol {
	return Symbols[k.symbol]
}

func (k *Kline) Interval() market.Interval {
	return Intervals[k.interval]
}

func (k *Kline) Open() int64 {
	return k.open
}

func (k *Kline) Close() int64 {
	return k.close
}

func (k *Kline) High() int64 {
	return k.high
}

func (k *Kline) Low() int64 {
	return k.low
}

func (k *Kline) Volume() int64 {
	return k.volume
}

func (k *Kline) TradeNum() int64 {
	return k.tradeNum
}

func (k *Kline) IsCompleted() bool {
	return k.isFinal
}

func (k *Kline) InProgress() bool {
	return !k.isFinal
}
