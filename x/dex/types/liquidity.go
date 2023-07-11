package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Liquidity interface {
	Swap(maxAmountTakerIn sdk.Int, maxAmountMakerOut *sdk.Int) (inAmount, outAmount sdk.Int)
	Price() sdk.Dec
}
