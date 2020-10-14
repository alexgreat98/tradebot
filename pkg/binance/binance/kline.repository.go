package binance

import (
	"context"
	"gorm.io/gorm"
)

func NewGormBinanceKlineRepo(ctx *context.Context, dbDriver *gorm.DB) *GormBinanceKlineRepo {
	return &GormBinanceKlineRepo{
		db: dbDriver,
	}
}

type GormBinanceKlineRepo struct {
	db *gorm.DB
}

func (repo GormBinanceKlineRepo) Find() []Kline {
	var kline []Kline
	repo.db.Find(&kline)
	return kline
}

// Create store in DB the Kline model
func (repo GormBinanceKlineRepo) Create(kline *Kline) *Kline {
	repo.db.Create(kline)
	return kline
}

// Update store in DB the Kline model
func (repo GormBinanceKlineRepo) Update(kline *Kline) *Kline {
	repo.db.Save(kline)
	return kline
}

// Delete mark model in DB like deleted for the Kline model
func (repo GormBinanceKlineRepo) Delete(kline *Kline) *Kline {
	repo.db.Delete(kline)
	return kline
}

// First retrieve first model from table (ID order by Asc)
func (repo GormBinanceKlineRepo) First() *Kline {
	var kline *Kline
	repo.db.First(kline)
	return kline
}

// Last retrieve first model from table (ID order by DESC)
func (repo GormBinanceKlineRepo) Last() *Kline {
	var kline *Kline
	repo.db.Last(kline)
	return kline
}
