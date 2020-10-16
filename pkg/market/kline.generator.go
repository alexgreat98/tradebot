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
	}
	kg.setZeroKline(time.Now()).startTimeoutHandler()
	return &kg
}

type KlineGenerator struct {
	observer.ObservableImpl

	symbol   Symbol
	interval Interval

	kline *KlineDto

	mutex sync.Mutex
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

	// check is trade valid
	if trade.Symbol().Code != kg.symbol.Code {
		return errors.New("Try to set Trade with symbol" + trade.Symbol().Code + " to " + kg.symbol.Code + "-KlineGenerator!")
	}

	if kg.kline.isFinal {
		kg.setZeroKline(time.Now())
	}

	kg.kline.tradeNum += 1
	kg.kline.volume += trade.Quantity()

	// if it is a first trade in the kline
	if kg.kline.open == 0 {
		kg.kline.open = trade.Price()
		kg.kline.high = trade.Price()
		kg.kline.low = trade.Price()
	} else {
		if trade.Price() > kg.kline.high {
			kg.kline.high = trade.Price()
		}

		if trade.Price() < kg.kline.low {
			kg.kline.low = trade.Price()
		}
	}

	// last trade always will be close-price
	kg.kline.close = trade.Price()

	kg.kline.volumeProfile.AddTrade(trade)

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

// startTimeoutHandler start goroutine for every second checking current kline timeout
func (kg *KlineGenerator) startTimeoutHandler() *KlineGenerator {
	checkFunc := func(timer time.Time) {
		if kg.kline.endTime >= (time.Now().Unix() + KlineTimeoutSeconds) {
			kg.closeKline().notifySubscribers()
		}
	}
	go utils.DoEvery(time.Second, checkFunc)
	return kg
}
