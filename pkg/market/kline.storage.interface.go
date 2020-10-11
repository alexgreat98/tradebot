package market

//
//import (
//	"container/list"
//)
//
//type KlineStorage interface {
//	// SetKline append new kline to the list
//	SetKline(kline Kline) KlineStorage
//
//	// GetInterval retrieve interval for stored kline-list
//	GetInterval() string
//
//	// GetCurrent retrieve the current (last) market kline (could be NOT finished!)
//	GetCurrent() Kline
//
//	// GetCurrentVolume retrieve volume for the current (last) kline
//	GetCurrentVolume() string
//
//	// GetCurrentTradeNum retrieve trades numbers for the current (last) kline
//	GetCurrentTradeNum() int64
//
//	// GetLastList retrieve last market kline list
//	GetLastList() *list.List
//
//	// GetLastListVolume retrieve volume summary for the last market kline-list
//	GetLastListVolume() string
//
//	// GetLastListTradeNum retrieve trades numbers summary for the last market kline-list
//	GetLastListTradeNum() int64
//
//	// GetStorageSize retrieve how many kline could be logged in the list (current + last list)
//	GetStorageSize() int
//
//	// CountLastList retrieve how many kline was stored in the last-list
//	CountLastList() int
//
//	// IsStorageReady detect is last-list is full (if that is so important to work with full list for some analytics modules)
//	IsStorageReady() bool
//}
