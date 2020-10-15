package market

import "time"

type Symbol struct {
	Code string

	BaseAsset          string
	BaseAssetPrecision int

	QuoteAsset          string
	QuoteAssetPrecision int
}

type Interval struct {
	Code     string
	Seconds  int
	Location time.Location
}

// Start retrieve time for the start of the current period
func (i Interval) Start(t time.Time) time.Time {
	// TODO: реализовать метод для других интервалов
	return t.Add(time.Duration(-1*t.Second()) * time.Second)
}

// Start retrieve time for the ent of the current period
func (i Interval) End(t time.Time) time.Time {
	// TODO: рассчитать время для интервалов
	return t.Add(time.Duration(60-t.Second()) * time.Second)
}
