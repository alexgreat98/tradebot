package binancerepository

import (
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	"gorm.io/gorm"
)

type KlineRepo struct {
	db gorm.DB
}

// Create store in DB the Kline model
func (repo KlineRepo) Create(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Create(kline)
	return kline
}

// Update store in DB the Kline model
func (repo KlineRepo) Update(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Save(kline)
	return kline
}

// Delete mark model in DB like deleted for the Kline model
func (repo KlineRepo) Delete(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Delete(kline)
	return kline
}
