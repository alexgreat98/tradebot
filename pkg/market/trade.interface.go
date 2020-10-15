package market

type Trade interface {
	// Time retrieve time when the trade was received from the market ws
	Time() int64

	// TradeTime retrieve time when the trade was completed in the market
	TradeTime() int64

	// Symbol retrieve symbol for the trade
	Symbol() Symbol

	// Price retrieve trade's price
	Price() int64

	// Quantity retrieve trade's quantity
	Quantity() int64

	// BuyerOrderID retrieve buyers order ID in the market
	BuyerOrderID() int64

	// SellerOrderID retrieve seller order ID in the market
	SellerOrderID() int64

	// IsBuyerMaker retrieve seller order ID in the market
	IsBuyerMaker() bool
}
