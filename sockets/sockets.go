package sockets

import (
	"context"
	web "github.com/webdelo/tradebot/sockets/web/binance"
)

func Run(ctx *context.Context) error {
	// Run Binance sockets listening
	err := BinanceRun(ctx)
	go web.Run()
	return err
}
