package binancewsobservers

import (
	"fmt"
	"github.com/golobby/container"
	"github.com/jinzhu/copier"
	"github.com/webdelo/tradebot/pkg/binance"
	"github.com/webdelo/tradebot/pkg/market"
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
		kline, ok := data.(*market.KlineDTO)
		if ok {
			err := obj.StoreKline(kline)
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
func (obj *klineToDB) StoreKline(kline *market.KlineDTO) error {
	var klineModel *binance.Kline
	klineModel = new(binance.Kline)
	err := copier.Copy(&klineModel, &kline)
	if err != nil {
		return err
	}

	klineModel.Symbol = kline.Symbol.Code
	klineModel.Interval = kline.Interval.Code

	var BinanceKlineRepo binance.BinanceKlineRepository
	container.Make(&BinanceKlineRepo)

	BinanceKlineRepo.Create(klineModel)

	fmt.Println("Stored successfully! ", klineModel)
	return nil
}
