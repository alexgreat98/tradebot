package bindings

import (
	"context"
	"github.com/golobby/container"
	"github.com/webdelo/tradebot/pkg/binance"
	"gorm.io/gorm"
)

func BindRepositories(ctx *context.Context, db *gorm.DB) {
	container.Transient(func() binance.BinanceKlineRepository {
		return binance.NewGormBinanceKlineRepo(ctx, db)
	})
}
