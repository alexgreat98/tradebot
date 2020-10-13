package binancews

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/observer"
)

// NewKlineChannel is constructor instances new channel
func NewKlineChannel() *KlineChannel {
	return new(KlineChannel)
}

// KlineChannel get info from sockets and notify subscribers
type KlineChannel struct {
	observer.ObservableImpl
}

// Listen method starts the socket listening
func (kl *KlineChannel) Listen(
	ctx *context.Context,
	symbol market.Symbol,
	interval market.Interval,
) (doneC, stopC chan struct{}, err error) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		// TODO: use named type for event detection
		klineDTO, err := WSKlineToDTO(event.Kline)
		if err != nil {
			// TODO: log error
			fmt.Println(err)
		}
		kl.NotifySubscribers("KlineIssued", klineDTO)
	}

	errHandler := func(err error) {
		// TODO: log error
		fmt.Println(err)
	}

	doneC, stopC, err = binance.WsKlineServe(symbol.Code, interval.Code, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return doneC, stopC, nil

}
