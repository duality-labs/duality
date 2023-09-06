package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/duality-labs/duality/x/cwhooks/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		tKey       storetypes.StoreKey
		paramstore paramtypes.Subspace
		wasmKeeper types.WasmKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	tKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	wasmKeeper types.WasmKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		tKey:       tKey,
		paramstore: ps,
		wasmKeeper: wasmKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
