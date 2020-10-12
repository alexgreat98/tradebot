package sockets

import (
	"context"
)

func Run(ctx *context.Context) error {
	// Run Binance sockets listening
	err := BinanceRun(ctx)
	return err
}
