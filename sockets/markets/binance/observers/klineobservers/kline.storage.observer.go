package klineobservers

import (
	"github.com/adshao/go-binance"
	"github.com/jinzhu/copier"
	"github.com/webdelo/tradebot/interfaces/market"
	binanceModels "github.com/webdelo/tradebot/models/markets/binance"
	"strconv"
)

type klineToStorage struct {
	key     string
	storage interfaces.KlineStorage
}

var currentKlineToStorageKey int = 1

func NewKlineToStorage(storage interfaces.KlineStorage) *klineToStorage {
	observer := new(klineToStorage)
	observer.storage = storage

	// set default key
	observer.SetKey("kline.to.storage.observer." + strconv.Itoa(currentKlineToStorageKey))
	currentKlineToStorageKey++

	return observer
}

func (o *klineToStorage) HandleEvent(event string, data interface{}) error {
	// TODO: use named type for event detection
	if event == "KlineIssued" {
		binanceKlineEvent, ok := data.(*binance.WsKlineEvent)
		if ok {
			err := o.ProcessKline(binanceKlineEvent.Kline)
			if err != nil {
				return err
			}
		} else {
			//TODO: log alert!
		}
	}
	return nil
}

// ProcessKline method store in DB new Kline
func (o *klineToStorage) ProcessKline(kline binance.WsKline) error {
	var klineModel *binanceModels.Kline
	klineModel = new(binanceModels.Kline)
	err := copier.Copy(&klineModel, &kline)
	if err != nil {
		return err
	}

	o.storage.SetKline(klineModel)

	return nil
}

func (o *klineToStorage) ObserverKey() string {
	return o.key
}

// SetKey set needed observer key
func (o *klineToStorage) SetKey(key string) *klineToStorage {
	o.key = key
	return o
}
