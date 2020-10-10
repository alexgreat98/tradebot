package markets

import (
	"container/list"
	"fmt"
	baseinterfaces "github.com/webdelo/tradebot/interfaces"
	interfaces "github.com/webdelo/tradebot/interfaces/market"
)

type KlineStorage struct {
	interval     string
	size         int
	currentKline interfaces.Kline
	list         *list.List
	subscribers  []baseinterfaces.Observer
}

// NewKlineStorage instance new
func NewKlineStorage(interval string, size int) *KlineStorage {
	return &KlineStorage{
		interval: interval,
		size:     size,
		list:     list.New(),
	}
}

// SetKline put new kline to the list and remove last kline
func (s *KlineStorage) SetKline(kline interfaces.Kline) interfaces.KlineStorage {
	// Check if this first kline adding to storage
	if s.currentKline != nil {
		// If current kline is final move it to the last-list
		if s.currentKline.IsCompleted() {
			s.list.PushFront(s.currentKline)
			// check if the list is full truncate the last kline
			if s.list.Len() > s.size {
				s.list.Remove(
					s.list.Back(),
				)
			}
		}
	}
	s.currentKline = kline

	// Notify observers about new kline
	for _, subscriver := range s.subscribers {
		// TODO: use named type for event title
		subscriver.HandleEvent("KlineStorageUpdated", s)
	}
	return s
}

// GetInterval retrieve interval for stored kline-list
func (s *KlineStorage) GetInterval() string {
	return s.interval
}

// GetCurrent retrieve the current (last) market kline (could be NOT finished!)
func (s *KlineStorage) GetCurrent() interfaces.Kline {
	return s.currentKline
}

// GetCurrentVolume retrieve volume for the current (last) kline
func (s *KlineStorage) GetCurrentVolume() string {
	return s.currentKline.GetVolume()
}

// GetCurrentTradeNum retrieve trades numbers for the current (last) kline
func (s *KlineStorage) GetCurrentTradeNum() int64 {
	return s.currentKline.GetTradeNum()
}

// GetLastList retrieve last market kline list
func (s *KlineStorage) GetLastList() *list.List {
	return s.list
}

// GetLastListVolume retrieve volume summary for the last market kline-list
func (s *KlineStorage) GetLastListVolume() string {
	// TODO: convert string type to numeric
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	return ""
}

// GetLastListTradeNum retrieve trades numbers summary for the last market kline-list
func (s *KlineStorage) GetLastListTradeNum() int64 {
	var sum int64 = 0
	for e := s.list.Front(); e != nil; e = e.Next() {
		k := e.Value.(interfaces.Kline)
		sum += k.GetTradeNum()
	}
	return sum
}

// GetStorageSize retrieve how many kline could be logged in the list (current + last list)
func (s *KlineStorage) GetStorageSize() int {
	return s.size
}

// CountLastList retrieve how many kline was stored in the last-list
func (s *KlineStorage) CountLastList() int {
	return s.list.Len()
}

// IsStorageReady detect is last-list is full (if that is so important to work with full list for some analytics modules)
func (s *KlineStorage) IsStorageReady() bool {
	return s.GetStorageSize() == s.CountLastList()
}

// Observable interface implementation

//Subscribe new observer and will notify about new events
func (s *KlineStorage) Subscribe(subscriber baseinterfaces.Observer) {
	s.subscribers = append(s.subscribers, subscriber)
}

//Unsubscribe observer and will not send new events
func (s *KlineStorage) Unsubscribe(subscriber baseinterfaces.Observer) {
	s.subscribers = removeFromSlice(s.subscribers, subscriber)
}

func removeFromSlice(subscribers []baseinterfaces.Observer, subscriber baseinterfaces.Observer) []baseinterfaces.Observer {
	observerListLength := len(subscribers)
	for i, observer := range subscribers {
		if subscriber.ObserverKey() == observer.ObserverKey() {
			subscribers[observerListLength-1], subscribers[i] = subscribers[i], subscribers[observerListLength-1]
			return subscribers[:observerListLength-1]
		}
	}
	return subscribers
}
