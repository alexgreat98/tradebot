package pinbar

import (
	"fmt"
	"github.com/webdelo/tradebot/pkg/market"
)

func NewPinbarObserver(config Config) *Observer {
	return &Observer{
		config: config,
	}
}

type Observer struct {
	config Config
}

//HandleEvent method for processing events from observable objects
func (k Observer) HandleEvent(event string, data interface{}) error {
	fmt.Println("--------- PINBAR --------")
	fmt.Println(event)
	var storage *market.KlineStorage
	storage = data.(*market.KlineStorage)
	fmt.Println(storage.GetCurrent())
	fmt.Println("-------------------------")
	return nil
}

//ObserverKey retrieve unique key (using for unsubscribe method)
func (k Observer) ObserverKey() string {
	return "pinbar.observer"
}
