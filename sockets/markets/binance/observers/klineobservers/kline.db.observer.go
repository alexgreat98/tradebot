package klineobservers

import (
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/golobby/container"
	"github.com/jinzhu/copier"
	interfaces "github.com/webdelo/tradebot/interfaces/repositories/binance"
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	web "github.com/webdelo/tradebot/sockets/web/binance"
	"strconv"
)

type klineToDB struct {
	key string
}

var currentKlineToDBKey int = 1

func NewKlineToDB() *klineToDB {
	observer := new(klineToDB)
	// set default key
	observer.SetKey("kline.to.db.observer." + strconv.Itoa(currentKlineToDBKey))
	currentKlineToDBKey++

	return observer
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
	var klineModel *binanceModels.Kline
	klineModel = new(binanceModels.Kline)
	err := copier.Copy(&klineModel, &kline)
	if err != nil {
		return err
	}

	var BinanceKlineRepo interfaces.BinanceKlineRepository
	container.Make(&BinanceKlineRepo)

	BinanceKlineRepo.Create(klineModel)
	go func() { web.Messages <- klineModel.Volume }()
	fmt.Println("Stored successfully! ", klineModel)
	return nil
}

func (obj *klineToDB) ObserverKey() string {
	return obj.key
}

// SetKey set needed observer key
func (obj *klineToDB) SetKey(key string) *klineToDB {
	obj.key = key
	return obj
}
