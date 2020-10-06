package bindings

import (
	"context"
	"github.com/golobby/container"
	"github.com/webdelo/tradebot/database"
	interfaces "github.com/webdelo/tradebot/interfaces/repositories/binance"
	repositories "github.com/webdelo/tradebot/repositories/markets/binance"
)

func BindRepositories(ctx context.Context) {
	container.Transient(func() interfaces.BinanceKlineRepository {
		return repositories.NewGormBinanceKlineRepo(database.Sqlite)
	})
}
