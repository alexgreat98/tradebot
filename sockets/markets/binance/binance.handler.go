package binance

import (
	"context"
	"github.com/webdelo/tradebot/sockets/markets/binance/channels"
	"github.com/webdelo/tradebot/sockets/markets/binance/observers/klineobserver"
)

func Run(ctx *context.Context) error {

	listenKlineChannel(ctx)

	//TODO: move listening starting to handler
	//channels.ListenKline(ctx)
	//
	//channels.ListenDepth(ctx)

	return nil
}

// listenKlineChannel starts listening for all needed Kline channels
func listenKlineChannel(ctx *context.Context) error {
	klineChannel := channels.NewKlineChannel()

	// Attach observers for that channel
	observer := klineobserver.NewKlineToDB()
	klineChannel.Subscribe(observer)

	// Start channel listening
	_, _, err := klineChannel.Listen(
		ctx,
		Symbols()["BTCUSDT"],
		KlineIntervals()["1m"],
	)

	if err != nil {
		return err
	}

	// TODO: save done&stop channels

	return nil
}
