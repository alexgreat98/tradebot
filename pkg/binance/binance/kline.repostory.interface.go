package binance

type BinanceKlineRepository interface {

	//Find retrieving all objects
	Find() []Kline

	// Create store in DB the Kline model
	Create(kline *Kline) *Kline

	// Update store in DB the Kline model
	Update(kline *Kline) *Kline

	// Delete mark model in DB like deleted for the Kline model
	Delete(kline *Kline) *Kline

	// First retrieve first model from table (ID order by Asc)
	First() *Kline

	// Last retrieve first model from table (ID order by DESC)
	Last() *Kline
}
