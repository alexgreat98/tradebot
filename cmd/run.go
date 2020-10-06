package cmd

import (
	"context"
	_ "github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/providers"
	"github.com/webdelo/tradebot/server"
	"github.com/webdelo/tradebot/sockets"
)

func Run() error {
	ctx := context.Background()

	// Bind all dependencies
	providers.Bind(ctx)

	// start sockets listening
	if err := sockets.Run(&ctx); err != nil {
		return err
	}

	// start web-server
	if err := server.Run(&ctx); err != nil {
		return err
	}

	return nil
}
