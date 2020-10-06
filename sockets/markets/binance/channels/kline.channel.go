package channels

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/interfaces"
)

// NewKlineChannel is constructor instances new channel
func NewKlineChannel() *KlineChannel {
	return new(KlineChannel)
}

// KlineChannel get info from sockets and notify observers
type KlineChannel struct {
	subscribers []interfaces.Observer
}

// Listen method starts the socket listening
func (kl *KlineChannel) Listen(ctx *context.Context, symbol string, interval string) (doneC, stopC chan struct{}, err error) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		for _, observer := range kl.subscribers {
			go observer.HandleEvent("KlineIssued", event)
		}
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

func (kl *KlineChannel) Subscribe(subscriber interfaces.Observer) {
	kl.subscribers = append(kl.subscribers, subscriber)
}

func (kl *KlineChannel) Unsubscribe(subscriber interfaces.Observer) {
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
