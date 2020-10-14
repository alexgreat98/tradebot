package sockets

import (
	"context"
	"github.com/webdelo/tradebot/pkg/binance/binance"
	"github.com/webdelo/tradebot/pkg/binance/binancews"
	"github.com/webdelo/tradebot/pkg/binance/binancewsobservers"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/strategy/pinbar"
)

// BinanceRun binancews sockets listeners
func BinanceRun(ctx *context.Context) error {

	if err := listenKlineChannel(ctx); err != nil {
		return err
	}

	if err := listenTradeChannel(ctx); err != nil {
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
	klineChannel := binancews.NewKlineChannel()

	// Attach subscribers that store kline to DB
	klineChannel.Subscribe(
		binancewsobservers.NewKlineToDB(),
	)

	klineStorage := market.NewKlineStorage(binance.Intervals["1m"], 5)
	// Attache Pinbar strategy for storage subscribers
	klineStorage.Subscribe(
		pinbar.NewPinbarObserver(),
	)

	// Attach observer that put kline to fast storage
	klineChannel.Subscribe(
		binancewsobservers.NewKlineToStorage(klineStorage),
	)

	// Start channel listening
	_, _, err := klineChannel.Listen(
		ctx,
		binance.Symbols["BTCUSDT"],
		binance.Intervals["1m"],
	)

	if err != nil {
		return err
	}

	// TODO: save done&stop channels

	return nil
}

func listenTradeChannel(ctx *context.Context) error {
	channel := binancews.NewTradeChannel()

	tradeToKline := binancewsobservers.NewTradeToKlineGenerator()

	channel.Subscribe(tradeToKline)

	_, _, err := channel.Listen(ctx, binance.Symbols["BTCUSDT"])

	if err != nil {
		return err
	}

	// TODO: save done&stop channels

	return nil
}
