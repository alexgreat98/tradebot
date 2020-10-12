package klineobservers

import (
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/golobby/container"
	"github.com/jinzhu/copier"
	binance2 "github.com/webdelo/tradebot/pkg/binance"
	"strconv"
)

var currentKlineToDBKey int = 1

func NewKlineToDB() *klineToDB {
	observer := new(klineToDB)
	// set default key
	observer.key = "kline.to.db.observer." + strconv.Itoa(currentKlineToDBKey)
	currentKlineToDBKey++

	return observer
}

type klineToDB struct {
	key string
}

func (obj *klineToDB) ObserverKey() string {
	return obj.key
}

func (obj *klineToDB) HandleEvent(event string, data interface{}) error {
	// TODO: use named type for event detection
	if event == "KlineIssued" {
		binanceKlineEvent, ok := data.(*binance.WsKlineEvent)
		if ok {
			err := obj.StoreKline(binanceKlineEvent.Kline)
			if err != nil {
				return err
			}
		} else {
			//TODO: log alert!
		}
	}
	return nil
}

// StoreKline method store in DB new Kline
func (obj *klineToDB) StoreKline(kline binance.WsKline) error {
	var klineModel *binance2.Kline
	klineModel = new(binance2.Kline)
	err := copier.Copy(&klineModel, &kline)
	if err != nil {
		return err
	}

	var BinanceKlineRepo binance2.BinanceKlineRepository
	container.Make(&BinanceKlineRepo)

	BinanceKlineRepo.Create(klineModel)

	fmt.Println("Stored successfully! ", klineModel)
	return nil
}
