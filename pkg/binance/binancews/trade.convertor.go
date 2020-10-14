package binancews

import (
	"github.com/adshao/go-binance"
	binanceLocal "github.com/webdelo/tradebot/pkg/binance/binance"
	"github.com/webdelo/tradebot/pkg/market"
	"github.com/webdelo/tradebot/pkg/utils"
)

// ConvertWsTradeToDto convert binance wsTrade to standard market DTO object
func ConvertWsTradeToDto(wsTrade *binance.WsTradeEvent) (*market.TradeDto, error) {
	var tradeDto *market.TradeDto

	symbol := binanceLocal.Symbols[wsTrade.Symbol]

	price, err := utils.ToCoinsFromString(wsTrade.Price, symbol.QuoteAssetPrecision)
	if err != nil {
		return tradeDto, err
	}

	quantity, err := utils.ToCoinsFromString(wsTrade.Quantity, symbol.QuoteAssetPrecision)
	if err != nil {
		return tradeDto, err
	}

	tradeDto = market.NewTradeDto(
		wsTrade.Time,
		wsTrade.TradeTime,
		symbol,
		wsTrade.TradeID,
		price,
		quantity,
		wsTrade.BuyerOrderID,
		wsTrade.SellerOrderID,
		wsTrade.IsBuyerMaker,
	)

	return tradeDto, nil
}
