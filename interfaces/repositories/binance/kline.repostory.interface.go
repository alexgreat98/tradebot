package interfaces

import binanceModels "github.com/webdelo/tradebot/models/markets/binance"

type BinanceKlineRepository interface {
	// Create store in DB the Kline model
	Create(kline *binanceModels.Kline) *binanceModels.Kline

	// Update store in DB the Kline model
	Update(kline *binanceModels.Kline) *binanceModels.Kline

	// Delete mark model in DB like deleted for the Kline model
	Delete(kline *binanceModels.Kline) *binanceModels.Kline

	// First retrieve first model from table (ID order by Asc)
	First() *binanceModels.Kline

	// Last retrieve first model from table (ID order by DESC)
	Last() *binanceModels.Kline
}
