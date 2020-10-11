package binance

import (
	"context"
	"github.com/webdelo/tradebot/repositories/markets"
	"github.com/webdelo/tradebot/sockets/markets/binance/channels"
	"github.com/webdelo/tradebot/sockets/markets/binance/observers/klineobservers"
	"github.com/webdelo/tradebot/strategy/pinbar"
)

// Run binance sockets listeners
func Run(ctx *context.Context) error {

	if err := listenKlineChannel(ctx); err != nil {
		return err
	}

	//TODO: move listening starting to handler
	//channels.ListenKline(ctx)
	//
	//channels.ListenDepth(ctx)

	return nil
}

// listenKlineChannel starts listening for all needed Kline channels
func listenKlineChannel(ctx *context.Context) error {
	klineChannel := channels.NewKlineChannel()

	// Attach observers that store kline to DB
	klineChannel.Subscribe(
		klineobservers.NewKlineToDB(),
	)

	klineStorage := markets.NewKlineStorage("1m", 5)
	// Attache Pinbar strategy for storage subscribers
	klineStorage.Subscribe(
		pinbar.NewPinbarObserver(),
	)

	// Attach observer that put kline to fast storage
	klineChannel.Subscribe(
		klineobservers.NewKlineToStorage(klineStorage),
	)

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
