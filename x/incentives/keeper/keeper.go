package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/duality-labs/duality/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper provides a way to manage incentives module storage.
type Keeper struct {
	storeKey    sdk.StoreKey
	paramSpace  paramtypes.Subspace
	hooks       types.IncentiveHooks
	ak          types.AccountKeeper
	bk          types.BankKeeper
	ek          types.EpochKeeper
	dk          types.DexKeeper
	distributor Distributor
}

// NewKeeper returns a new instance of the incentive module keeper struct.
func NewKeeper(
	storeKey sdk.StoreKey,
	paramSpace paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ek types.EpochKeeper,
	dk types.DexKeeper,
) *Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	keeper := &Keeper{
		storeKey:   storeKey,
		paramSpace: paramSpace,
		ak:         ak,
		bk:         bk,
		ek:         ek,
		dk:         dk,
	}
	keeper.distributor = NewDistributor(keeper)
	return keeper
}

// SetHooks sets the incentives hooks.
func (k *Keeper) SetHooks(ih types.IncentiveHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set incentive hooks twice")
	}

	k.hooks = ih

	return k
}

// Logger returns a logger instance for the incentives module.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleBalance returns full balance of the module.
func (k Keeper) GetModuleBalance(ctx sdk.Context) sdk.Coins {
	acc := k.ak.GetModuleAccount(ctx, types.ModuleName)
	return k.bk.GetAllBalances(ctx, acc.GetAddress())
}

// GetModuleLockedCoins Returns locked balance of the module.
func (k Keeper) GetModuleLockedCoins(ctx sdk.Context) sdk.Coins {
	// all not unlocking + not finished unlocking
	notUnlockingLocksCoins := k.getFullLocks(ctx).GetCoins()
	unlockingLocksNotMaturedCoins := k.getUnlockingLocksNotMatured(ctx).GetCoins()
	return notUnlockingLocksCoins.Add(unlockingLocksNotMaturedCoins...)
}
