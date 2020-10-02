package cmd

import (
	_ "github.com/adshao/go-binance"
	"github.com/webdelo/tradebot/server"
)

func Run() error {

	// start web-server
	if err := server.Run(); err != nil {
		return err
	}

	return nil
}

