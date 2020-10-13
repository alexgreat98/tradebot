package binancews

import (
	"github.com/adshao/go-binance"
	binanceLocal "github.com/webdelo/tradebot/pkg/binance"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/utils"
)

// WSKlineToDTO convert binance wsKline to standard market DTO object
func WSKlineToDTO(wsKline binance.WsKline) (*market.KlineDTO, error) {
	dto := new(market.KlineDTO)

	symbol := binanceLocal.Symbols[wsKline.Symbol]

	dto.StartTime = wsKline.StartTime
	dto.EndTime = wsKline.EndTime
	dto.Symbol = symbol
	dto.Interval = binanceLocal.Intervals[wsKline.Interval]

	amount, err := utils.ToCoinsFromString(wsKline.Open, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.Open = amount

	amount, err = utils.ToCoinsFromString(wsKline.Close, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.Close = amount

	amount, err = utils.ToCoinsFromString(wsKline.High, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.High = amount

	amount, err = utils.ToCoinsFromString(wsKline.Low, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.Low = amount

	amount, err = utils.ToCoinsFromString(wsKline.Low, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.Low = amount

	amount, err = utils.ToCoinsFromString(wsKline.Volume, symbol.QuoteAssetPrecision)
	if err != nil {
		return dto, err
	}
	dto.Volume = amount

	dto.TradeNum = wsKline.TradeNum
	dto.IsFinal = wsKline.IsFinal

	return dto, nil
}
