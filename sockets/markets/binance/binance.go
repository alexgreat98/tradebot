package binance

import (
	"context"
	"github.com/webdelo/tradebot/sockets/markets/binance/channels"
)

func Run(ctx *context.Context) error {

	channels.ListenTrade(ctx)

	channels.ListenKline(ctx)

	channels.ListenDepth(ctx)

	return nil
}
