package binancews

import (
	"github.com/adshao/go-binance"
	binance2 "github.com/webdelo/tradebot/pkg/binance/binance"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/utils"
)

// ConvertWSKlineToDto convert binance wsKline to standard market DTO object
func ConvertWSKlineToDto(wsKline binance.WsKline) (*market.KlineDto, error) {
	var dto *market.KlineDto

	symbol := binance2.Symbols[wsKline.Symbol]
	interval := binance2.Intervals[wsKline.Interval]

	openPrice, err := utils.ToCoinsFromString(wsKline.Open, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}

	closePrice, err := utils.ToCoinsFromString(wsKline.Close, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}

	highPrice, err := utils.ToCoinsFromString(wsKline.High, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}

	lowPrice, err := utils.ToCoinsFromString(wsKline.Low, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}

	volume, err := utils.ToCoinsFromString(wsKline.Volume, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}

	dto = market.NewKlineDto(
		wsKline.StartTime,
		wsKline.EndTime,
		symbol,
		interval,
		openPrice,
		closePrice,
		highPrice,
		lowPrice,
		volume,
		wsKline.TradeNum,
		wsKline.IsFinal,
	)

	return dto, nil
}
