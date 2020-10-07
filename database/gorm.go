package database

import (
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	sqliteGorm "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Sqlite instance new DB object with sqlite driver
func Sqlite() (*gorm.DB, error) {
	sqlite, err := gorm.Open(sqliteGorm.Open("bot.sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// start sockets listening
	if err := autoMigrate(sqlite); err != nil {
		return nil, err
	}

	return sqlite, nil
}

// autoMigrate create tables for needed models
func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&binanceModels.Kline{})
	if err != nil {
		return err
	}

	return nil
}
