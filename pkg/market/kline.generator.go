package market

import (
	"errors"
	"github.com/webdelo/tradebot/pkg/observer"
	"github.com/webdelo/tradebot/pkg/utils"
	"sync"
	"time"
)

const KlineTimeoutSeconds = 2 // kline will be automatically closed after that period + kline.endTime

// NewKlineGenerator constructor for kline generation
func NewKlineGenerator(symbol Symbol, interval Interval) *KlineGenerator {
	kg := KlineGenerator{
		symbol:   symbol,
		interval: interval,

		firstKlineFlag: true,
	}
	kg.setZeroKline(time.Now()).startTimeoutHandler()
	return &kg
}

type KlineGenerator struct {
	observer.ObservableImpl

	symbol   Symbol
	interval Interval

	trades    []Trade
	lastTrade Trade
	kline     *KlineDto

	firstKlineFlag bool

	mutex sync.Mutex
}

// startTimeoutHandler start goroutine for every second checking current kline timeout
func (kg *KlineGenerator) startTimeoutHandler() *KlineGenerator {
	checkFunc := func(timer time.Time) {
		if kg.firstKlineFlag == false {
			if kg.kline.endTime >= (time.Now().Unix() + KlineTimeoutSeconds) {
				kg.closeKline().notifySubscribers()
			}
		}
	}
	go utils.DoEvery(time.Second, checkFunc)
	return kg
}

// Symbol retrieve symbol for that generator
func (kg *KlineGenerator) Symbol() Symbol {
	return kg.symbol
}

// Interval retrieve interval for that generator
func (kg *KlineGenerator) Interval() Interval {
	return kg.interval
}

// Kline retrieve current built kline
func (kg *KlineGenerator) Kline() Kline {
	return kg.kline
}

// SetTrade insert new trade to generator and rebuild current kline
func (kg *KlineGenerator) SetTrade(trade Trade) error {
	kg.mutex.Lock()
	// check is trade from next kline
	if trade.TradeTime() > kg.kline.endTime {
		kg.closeKline().notifySubscribers()
	}

	kg.trades = append(kg.trades, trade)
	// check is trade valid
	if trade.Symbol().Code != kg.symbol.Code {
		return errors.New("Try to set Trade with symbol" + trade.Symbol().Code + " to " + kg.symbol.Code + "-KlineGenerator!")
	}
	kg.rebuildKline()

	// notify about new kline generation
	kg.notifySubscribers()
	kg.mutex.Unlock()
	return nil
}

// closeKline complete current kline
func (kg *KlineGenerator) closeKline() *KlineGenerator {
	kg.kline.isFinal = true
	return kg
}

// notifySubscribers send kline to subscribers
func (kg *KlineGenerator) notifySubscribers() *KlineGenerator {
	kg.NotifySubscribers(KlineIssuedEvent, kg.kline)
	return kg
}

// rebuildKline generate new kline DTO according trades list
func (kg *KlineGenerator) rebuildKline() {

	if kg.kline.isFinal {
		kg.setZeroKline(time.Now())
	}

	kg.kline.tradeNum = kg.kline.tradeNum + 1
	kg.kline.volume = kg.lastTrade.Quantity() + kg.kline.Volume()

	if kg.kline.open == 0 {
		kg.kline.open = kg.lastTrade.Price()
		kg.kline.close = kg.lastTrade.Price()
		kg.kline.high = kg.lastTrade.Price()
		kg.kline.low = kg.lastTrade.Price()
	}

	if kg.lastTrade.Price() > kg.kline.high {
		kg.kline.high = kg.lastTrade.Price()
	}

	if kg.lastTrade.Price() < kg.kline.low || kg.kline.low == 0 {
		kg.kline.low = kg.lastTrade.Price()
	}

	if kg.lastTrade.TradeTime() >= kg.kline.EndTime() {
		kg.kline.low = kg.lastTrade.Price()
	}
}

// setZeroKline generate zero-kline and set it
func (kg *KlineGenerator) setZeroKline(time time.Time) *KlineGenerator {
	kg.kline = NewKlineDto(
		kg.interval.Start(time).Unix(),
		kg.interval.Start(time).Unix(),
		kg.symbol,
		kg.interval,
		0,
		0,
		0,
		0,
		0,
		0,
		false,
	)
	return kg
}
