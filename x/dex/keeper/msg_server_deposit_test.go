package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Returns -1 if error
func newDec(amount string) sdk.Dec {
	amt, err := sdk.NewDecFromStr(amount)

	if err != nil {
		return sdk.NewDec(-1)
	}

	return amt

}

func newACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func newBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

func convInt(amt string) sdk.Int {
	IntAmt, err := sdk.NewIntFromString(amt)

	_ = err
	return IntAmt
}

func (suite *IntegrationTestSuite) TestHasBalance() {
	app, ctx := suite.app, suite.ctx
	addr := sdk.AccAddress([]byte("addr1_______________"))

	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	app.AccountKeeper.SetAccount(ctx, acc)

	balances := sdk.NewCoins(newACoin(sdk.NewInt(100)))

	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(99))))

	suite.Require().NoError(simapp.FundAccount(app.BankKeeper, ctx, addr, balances))
	suite.Require().False(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(101))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(100))))
	suite.Require().True(app.BankKeeper.HasBalance(ctx, addr, newACoin(sdk.NewInt(1))))
}
