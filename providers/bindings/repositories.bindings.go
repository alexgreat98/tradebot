package bindings

import (
	"context"
	"github.com/golobby/container"
	binance2 "github.com/webdelo/tradebot/pkg/binance/binance"
	"gorm.io/gorm"
)

func BindRepositories(ctx *context.Context, db *gorm.DB) {
	container.Transient(func() binance2.BinanceKlineRepository {
		return binance2.NewGormBinanceKlineRepo(ctx, db)
	})
}
