package repositories

import (
	"context"
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	"gorm.io/gorm"
)

func NewGormBinanceKlineRepo(ctx *context.Context, dbDriver *gorm.DB) *GormBinanceKlineRepo {
	repo := &GormBinanceKlineRepo{
		db: dbDriver,
	}
	return repo
}

type GormBinanceKlineRepo struct {
	db *gorm.DB
}

// Create store in DB the Kline model
func (repo GormBinanceKlineRepo) Create(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Create(kline)
	return kline
}

// Update store in DB the Kline model
func (repo GormBinanceKlineRepo) Update(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Save(kline)
	return kline
}

// Delete mark model in DB like deleted for the Kline model
func (repo GormBinanceKlineRepo) Delete(kline *binanceModels.Kline) *binanceModels.Kline {
	repo.db.Delete(kline)
	return kline
}

// First retrieve first model from table (ID order by Asc)
func (repo GormBinanceKlineRepo) First() *binanceModels.Kline {
	var kline *binanceModels.Kline
	repo.db.First(kline)
	return kline
}

// Last retrieve first model from table (ID order by DESC)
func (repo GormBinanceKlineRepo) Last() *binanceModels.Kline {
	var kline *binanceModels.Kline
	repo.db.Last(kline)
	return kline
}
