package binancewsobservers

import (
	"github.com/webdelo/tradebot/pkg/binance/binancews"
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
	if event == binancews.KlineIssuedEvent {
		kline, ok := data.(*market.KlineDto)
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
func (obj *klineToDB) StoreKline(kline *market.KlineDto) error {
	//var klineModel *binance2.Kline
	//klineModel = new(binance2.Kline)
	// TODO: use convertor! Copier does not works now!
	//err := copier.Copy(&klineModel, &kline)
	//if err != nil {
	//	return err
	//}
	//
	//klineModel.symbol = kline.Symbol().Code
	//klineModel.Interval = kline.Interval().Code
	//
	//var BinanceKlineRepo binance2.BinanceKlineRepository
	//container.Make(&BinanceKlineRepo)
	//
	//BinanceKlineRepo.Create(klineModel)

	//fmt.Println("Stored successfully! ", klineModel)
	return nil
}
