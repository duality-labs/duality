package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func newCCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenC", amt)
}

func newDCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenD", amt)
}

// Handle checking for intentional failure of test
// func (s *MsgServerTestSuite) handleIntentionalFail(t *testing.T, format string, args ...interface{}) {
// 	if !env.intentionalFail {
// 		t.Fatalf(format, args...)
// 	} else {
// 		t.Skipf("Test intentionally failed, skipping rest of execution. Error: "+format, args)
// 	}
// }

func getBalance(s *MsgServerTestSuite, acc sdk.AccAddress, denom string) sdk.Dec {
	app, ctx := s.app, s.ctx
	return sdk.NewDecFromIntWithPrec(app.BankKeeper.GetBalance(ctx, acc, denom).Amount, 18)
}

func getDexBalance(s *MsgServerTestSuite, denom string) sdk.Dec {
	app, ctx := s.app, s.ctx
	return sdk.NewDecFromIntWithPrec(app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex")).AmountOf(denom), 18)
}

// Helper to convert coins into sorted amount0, amount1
func sortCoins(s *MsgServerTestSuite, denomA string, denomB string, amountsA []sdk.Dec, amountsB []sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec) {
	app, ctx := s.app, s.ctx
	denom0, denom1, err := app.DexKeeper.SortTokens(ctx, denomA, denomB)
	s.Require().Nil(err)
	// this corresponds to lines 45-54 of verification.go
	amounts0, amounts1 := amountsA, amountsB
	// flip amounts if denoms were flipped
	if denom0 != denomA {
		amounts0, amounts1 = amountsB, amountsA
	}
	return denom0, denom1, amounts0, amounts1
}

// TODO: this was taken from core.go, lines 287-294. should be moved to utils somewhere
func min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}
