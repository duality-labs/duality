package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func GetBlockGasUsed(ctx sdk.Context) (gasUsed uint64, err error) {
	switch {
	case ctx.BlockGasMeter() != nil:
		return ctx.BlockGasMeter().GasConsumed(), nil
	case ctx.IsCheckTx():
		return 0, nil
	default:
		return 0, sdkerrors.Wrap(sdkerrors.ErrAppConfig, "Block Gas Meter is not initialized")
	}
}

func MustGetBlockGasUsed(ctx sdk.Context) uint64 {
	gasUsed, err := GetBlockGasUsed(ctx)
	if err != nil {
		panic(err)
	}

	return gasUsed
}
