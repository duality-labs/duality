package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type LiquidityType int

const (
	LiquidityType_PoolReserves LiquidityType = iota
	LiquidityType_LimitOrder
)

type TickLiquidityKey interface {
	KeyMarshal() []byte
	KeyUnmarshal([]byte) error
	PriceTakerToMaker() (priceTakerToMaker sdk.Dec, err error)
}
