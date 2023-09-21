package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type TickLiquidityKey interface {
	KeyMarshal() []byte
	PriceTakerToMaker() (priceTakerToMaker sdk.Dec, err error)
}
