package types

type TickLiquidityKey interface {
	KeyMarshal() []byte
	PriceTakerToMaker() (priceTakerToMaker sdk.Dec, err error)
}
