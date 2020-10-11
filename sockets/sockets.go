package sockets

import (
	"context"
	binanceSockets "github.com/webdelo/tradebot/sockets/markets/binance"
	web "github.com/webdelo/tradebot/sockets/web/binance"
)

func Run(ctx *context.Context) error {
	// Run Binance sockets listening
	err := binanceSockets.Run(ctx)
	web.Run()
	return err
}
