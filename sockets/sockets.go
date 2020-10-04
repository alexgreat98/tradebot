package sockets

import (
	"context"
	binanceSocket "github.com/webdelo/tradebot/sockets/markets/binance"
)

func Run(ctx *context.Context) error {
	// Run Binance sockets listening
	err := binanceSocket.Run(ctx)
	return err
}
