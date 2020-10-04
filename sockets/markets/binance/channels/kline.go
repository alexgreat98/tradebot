package channels

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/interfaces"
)

func ListenKline(ctx *context.Context) error {
	fmt.Println("Listen Klines channel")

	var klineListener KlineListener

	//TODO: subscribe observers

	//TODO: create binance constants for symbols and intervals
	klineListener.Listen(ctx, "LTCBTC", "1m")

	return nil

}

// KlineListener get info from sockets and notify observers
type KlineListener struct {
	subscribers []interfaces.Observer
}

func (kl *KlineListener) Listen(ctx *context.Context, symbol string, interval string) (doneC, stopC chan struct{}, err error) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		//TODO: notify observers
		fmt.Println(event)
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

func (kl *KlineListener) Subscribe(subscriber interfaces.Observer) {
	kl.subscribers = append(kl.subscribers, subscriber)
}

func (kl *KlineListener) Unsubscribe(subscriber interfaces.Observer) {
	kl.subscribers = removeFromSlice(kl.subscribers, subscriber)
}

func removeFromSlice(subscribers []interfaces.Observer, subscriber interfaces.Observer) []interfaces.Observer {
	observerListLength := len(subscribers)
	for i, observer := range subscribers {
		if subscriber.ObserverKey() == observer.ObserverKey() {
			subscribers[observerListLength-1], subscribers[i] = subscribers[i], subscribers[observerListLength-1]
			return subscribers[:observerListLength-1]
		}
	}
	return subscribers
}
