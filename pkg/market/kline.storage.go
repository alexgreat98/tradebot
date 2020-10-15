package market

import (
	"container/list"
	"github.com/webdelo/tradebot/pkg/observer"
)

type KlineStorage struct {
	observer.ObservableImpl

	interval     Interval
	size         int
	currentKline Kline
	list         *list.List
	subscribers  []observer.Observer
}

// NewKlineStorage instance new
func NewKlineStorage(interval Interval, size int) *KlineStorage {
	return &KlineStorage{
		interval: interval,
		size:     size,
		list:     list.New(),
	}
}

// SetKline put new kline to the list and remove last kline
func (s *KlineStorage) SetKline(kline Kline) *KlineStorage {
	// Check if this first kline adding to storage
	if s.currentKline != nil {
		// If current kline is final move it to the last-list
		if s.currentKline.IsFinal() {
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

	// Notify subscribers about new kline
	s.NotifySubscribers("KlineStorageUpdated", s)

	return s
}

// GetInterval retrieve interval for stored kline-list
func (s *KlineStorage) GetInterval() Interval {
	return s.interval
}

// GetCurrent retrieve the current (last) market kline (could be NOT finished!)
func (s *KlineStorage) GetCurrent() Kline {
	return s.currentKline
}

// GetCurrentVolume retrieve volume for the current (last) kline
func (s *KlineStorage) GetCurrentVolume() int64 {
	return s.currentKline.Volume()
}

// GetCurrentTradeNum retrieve trades numbers for the current (last) kline
func (s *KlineStorage) GetCurrentTradeNum() int64 {
	return s.currentKline.TradeNum()
}

// GetLastList retrieve last market kline list
func (s *KlineStorage) GetLastList() *list.List {
	return s.list
}

// GetLastListVolume retrieve volume summary for the last market kline-list
func (s *KlineStorage) GetLastListVolume() int64 {
	var sum int64 = 0
	for e := s.list.Front(); e != nil; e = e.Next() {
		k := e.Value.(Kline)
		sum += k.Volume()
	}
	return sum
}

// GetLastListTradeNum retrieve trades numbers summary for the last market kline-list
func (s *KlineStorage) GetLastListTradeNum() int64 {
	var sum int64 = 0
	for e := s.list.Front(); e != nil; e = e.Next() {
		k := e.Value.(Kline)
		sum += k.TradeNum()
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
