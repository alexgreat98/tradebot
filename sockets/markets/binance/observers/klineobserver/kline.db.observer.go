package klineobserver

import (
	"fmt"
	"strconv"
)

type klineToDB struct {
	key string
}

var currentKey int = 1

func NewKlineToDB() *klineToDB {
	observer := new(klineToDB)
	// set default key
	observer.SetKey("kline.to.db.observer." + strconv.Itoa(currentKey))
	currentKey++

	return observer
}

func (obj *klineToDB) HandleEvent(event string, data interface{}) error {
	if event == "KlineIssued" {
		err := obj.StoreKline(data)
		if err != nil {
			return err
		}
	}
	return nil
}

// StoreKline method store in DB new Kline
func (obj *klineToDB) StoreKline(kline interface{}) error {
	fmt.Println("Tring to store Kline!", kline)
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
