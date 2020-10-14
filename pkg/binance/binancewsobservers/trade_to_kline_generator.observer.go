package binancewsobservers

import (
	"fmt"
	"github.com/webdelo/tradebot/pkg/binance/binancews"
	"github.com/webdelo/tradebot/pkg/market"
	"strconv"
)

type TradeToKlineGenerator struct {
	key string
}

var currentTradeToKlineGeneratorKey int = 1

func NewTradeToKlineGenerator() *TradeToKlineGenerator {
	observer := new(TradeToKlineGenerator)

	// set default key
	observer.SetKey("trade.to.kline.generator.observer." + strconv.Itoa(currentTradeToKlineGeneratorKey))
	currentTradeToKlineGeneratorKey++

	return observer
}

func (o *TradeToKlineGenerator) HandleEvent(event string, data interface{}) error {
	if event == binancews.TradeReceivedEvent {
		trade, ok := data.(*market.TradeDto)
		if ok {
			err := o.ProcessKline(trade)
			if err != nil {
				return err
			}
		} else {
			//TODO: log alert!
		}
	}
	return nil
}

// ProcessKline method store in DB new Kline
func (o *TradeToKlineGenerator) ProcessKline(trade *market.TradeDto) error {
	fmt.Println("--------- Trade -----------")
	fmt.Println(trade)

	return nil
}

func (o *TradeToKlineGenerator) ObserverKey() string {
	return o.key
}

// SetKey set needed observer key
func (o *TradeToKlineGenerator) SetKey(key string) *TradeToKlineGenerator {
	o.key = key
	return o
}
