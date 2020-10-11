package channels

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/pkg/observer"
)

// NewKlineChannel is constructor instances new channel
func NewKlineChannel() *KlineChannel {
	return new(KlineChannel)
}

// KlineChannel get info from ws and notify observers
type KlineChannel struct {
	observer.ObservableImpl
}

// Listen method starts the socket listening
func (kl *KlineChannel) Listen(ctx *context.Context, symbol string, interval string) (doneC, stopC chan struct{}, err error) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		// TODO: use named type for event detection
		kl.NotifySubscribers("KlineIssued", event)
	}

	errHandler := func(err error) {
		// TODO: log error
		fmt.Println(err)
	}

	doneC, stopC, err = binance.WsKlineServe(symbol, interval, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return doneC, stopC, nil

}
