package bindings

import (
	"context"
	"github.com/golobby/container"
	interfaces "github.com/webdelo/tradebot/interfaces/repositories/binance"
	repositories "github.com/webdelo/tradebot/repositories/markets/binance"
	"gorm.io/gorm"
)

func BindRepositories(ctx *context.Context, db *gorm.DB) {
	container.Transient(func() interfaces.BinanceKlineRepository {
		return repositories.NewGormBinanceKlineRepo(ctx, db)
	})
}
