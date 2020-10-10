package pinbar

import (
	"fmt"
	"github.com/webdelo/tradebot/repositories/markets"
)

func NewPinbarObserver() *PinbarObserver {
	return new(PinbarObserver)
}

type PinbarObserver struct {
}

//HandleEvent method for processing events from observable objects
func (k PinbarObserver) HandleEvent(event string, data interface{}) error {
	fmt.Println("--------- PINBAR --------")
	fmt.Println(event)
	var storage *markets.KlineStorage
	storage = data.(*markets.KlineStorage)
	fmt.Println(storage.GetCurrent())
	fmt.Println("-------------------------")
	return nil
}

//ObserverKey retrieve unique key (using for unsubscribe method)
func (k PinbarObserver) ObserverKey() string {
	return "pinbar.observer"
}
