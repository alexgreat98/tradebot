package ws

import (
	"context"
)

func Run(ctx *context.Context) error {
	// Run Binance ws listening
	err := BinanceRun(ctx)
	return err
}
