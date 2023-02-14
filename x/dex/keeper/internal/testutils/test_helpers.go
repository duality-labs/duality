package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	dexmoduletypes "github.com/duality-labs/duality/x/dex/types"
)

func NewACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func NewBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

func ConvInt(amt string) sdk.Int {
	IntAmt, err := sdk.NewIntFromString(amt)

	_ = err
	return IntAmt
}

func NewDec(a int) sdk.Dec {
	return sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(a)))
}

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, dexmoduletypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, dexmoduletypes.ModuleName, addr, amounts)
}
