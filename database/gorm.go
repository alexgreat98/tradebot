package database

import (
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	sqliteGorm "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Sqlite *gorm.DB = new(gorm.DB)
)

func init() {
	Sqlite, err := gorm.Open(sqliteGorm.Open("bot.sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("Db opening failed! " + err.Error())
	}

	err = Sqlite.AutoMigrate(&binanceModels.Kline{})
	if err != nil {
		panic("Db migrating failed! " + err.Error())
	}
}
