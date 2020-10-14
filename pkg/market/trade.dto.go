package market

func NewTradeDto(
	time int64,
	tradeTime int64,
	symbol Symbol,
	tradeID int64,
	price int64,
	quantity int64,
	buyerOrderID int64,
	sellerOrderID int64,
	isBuyerMaker bool,
) *TradeDto {
	return &TradeDto{
		time:          time,
		tradeTime:     tradeTime,
		symbol:        symbol,
		tradeID:       tradeID,
		price:         price,
		quantity:      quantity,
		buyerOrderID:  buyerOrderID,
		sellerOrderID: sellerOrderID,
		isBuyerMaker:  isBuyerMaker,
	}
}

type TradeDto struct {
	time          int64
	tradeTime     int64
	symbol        Symbol
	tradeID       int64
	price         int64
	quantity      int64
	buyerOrderID  int64
	sellerOrderID int64
	isBuyerMaker  bool
}

// Time retrieve time when the trade was received from the market ws
func (t TradeDto) Time() int64 {
	return t.time
}

// TradeTime retrieve time when the trade was completed in the market
func (t TradeDto) TradeTime() int64 {
	return t.tradeTime
}

// Symbol retrieve symbol for the trade
func (t TradeDto) Symbol() Symbol {
	return t.symbol
}

// Price retrieve trade's price
func (t TradeDto) Price() int64 {
	return t.price
}

// Quantity retrieve trade's quantity
func (t TradeDto) Quantity() int64 {
	return t.quantity
}

// BuyerOrderID retrieve buyers order ID in the market
func (t TradeDto) BuyerOrderID() int64 {
	return t.buyerOrderID
}

// SellerOrderID retrieve seller order ID in the market
func (t TradeDto) SellerOrderID() int64 {
	return t.sellerOrderID
}

// IsBuyerMaker retrieve seller order ID in the market
func (t TradeDto) IsBuyerMaker() bool {
	return t.isBuyerMaker
}
