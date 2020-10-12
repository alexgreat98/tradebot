package ws

import (
	"context"
	web "github.com/webdelo/tradebot/sockets/web/binance"
)

func Run(ctx *context.Context) error {
	// Run Binance ws listening
	err := BinanceRun(ctx)
	web.Run()
	return err
}
