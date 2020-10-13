package market

type Symbol struct {
	Code string

	BaseAsset          string
	BaseAssetPrecision int

	QuoteAsset          string
	QuoteAssetPrecision int
}

type Interval struct {
	Code    string
	Seconds int
}
