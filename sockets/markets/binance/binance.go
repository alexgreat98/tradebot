package binance

import (
	"context"
	"github.com/webdelo/tradebot/sockets/markets/binance/channels"
)

func Run(ctx *context.Context) error {

	channels.ListenTrades(ctx)

	channels.ListenKlines(ctx)

	channels.ListenDepth(ctx)

	return nil
}
