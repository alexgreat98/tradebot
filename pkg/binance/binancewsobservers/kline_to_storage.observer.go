package binancewsobservers

import (
	"github.com/webdelo/tradebot/pkg/binance/binancews"
	"github.com/webdelo/tradebot/pkg/market"
	"strconv"
)

type klineToStorage struct {
	key     string
	storage *market.KlineStorage
}

var currentKlineToStorageKey int = 1

func NewKlineToStorage(storage *market.KlineStorage) *klineToStorage {
	observer := new(klineToStorage)
	observer.storage = storage

	// set default key
	observer.SetKey("kline.to.storage.observer." + strconv.Itoa(currentKlineToStorageKey))
	currentKlineToStorageKey++

	return observer
}

func (o *klineToStorage) HandleEvent(event string, data interface{}) error {
	if event == binancews.KlineIssuedEvent {
		kline, ok := data.(*market.KlineDto)
		if ok {
			err := o.ProcessKline(kline)
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
func (o *klineToStorage) ProcessKline(kline *market.KlineDto) error {
	o.storage.SetKline(kline)

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
