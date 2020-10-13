package cmd

import (
	"context"
	"github.com/webdelo/tradebot/providers"
	"github.com/webdelo/tradebot/server"
	"github.com/webdelo/tradebot/sockets"
)

func Run() error {
	ctx := context.Background()

	// Boot all dependencies
	if err := providers.Boot(&ctx); err != nil {
		return err
	}

	// start web-sockets listening
	if err := sockets.Run(&ctx); err != nil {
		return err
	}

	// start web-server
	if err := server.Run(&ctx); err != nil {
		return err
	}

	return nil
}
