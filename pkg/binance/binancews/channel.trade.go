package binancews

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/observer"
)

// NewTradeChannel is constructor instances new channel
func NewTradeChannel() *TradeChannel {
	return new(TradeChannel)
}

// TradeChannel get info from sockets and notify subscribers
type TradeChannel struct {
	observer.ObservableImpl
}

// Listen method starts the socket listening
func (kl *TradeChannel) Listen(
	ctx *context.Context,
	symbol market.Symbol,
) (doneC, stopC chan struct{}, err error) {
	wsTradeHandler := func(event *binance.WsTradeEvent) {
		tradeDto, err := ConvertWsTradeToDto(event)
		if err != nil {
			// TODO: log error
			fmt.Println(err)
		}
		kl.NotifySubscribers(TradeReceivedEvent, tradeDto)
	}

	errHandler := func(err error) {
		// TODO: log error
		fmt.Println(err)
	}

	doneC, stopC, err = binance.WsTradeServe(symbol.Code, wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return doneC, stopC, nil

}
