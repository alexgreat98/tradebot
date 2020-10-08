package sockets

import (
	"context"
	binanceSockets "github.com/webdelo/tradebot/sockets/markets/binance"
)

func Run(ctx *context.Context) error {
	// Run Binance sockets listening
	err := binanceSockets.Run(ctx)
	return err
}
