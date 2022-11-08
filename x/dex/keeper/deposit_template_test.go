package keeper_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Template function for testing the execution of deposit.
func (env *TestEnv) TestDeposit(t *testing.T, denomA string, denomB string, amountsA []sdk.Dec, amountsB []sdk.Dec, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64) {
	app, ctx := env.cosmos.app, env.cosmos.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// GIVEN inital balances
	if len(amountsA) != len(amountsB) {
		t.Errorf("Need equal lengths of coinsA and coinsB for multi-deposit")
	}
	denom0, denom1, amounts0, amounts1 := env.sortCoins(t, denomA, denomB, amountsA, amountsB)
	pairId := app.DexKeeper.CreatePairId(denom0, denom1)

	// calculate total trade amounts
	accBalance0Initial, accBalance1Initial := app.BankKeeper.GetBalance(ctx, acc, denom0).Amount.ToDec(), app.BankKeeper.GetBalance(ctx, acc, denom1).Amount.ToDec()
	totalAmount0, totalAmount1 := sdk.NewDecFromInt(convInt("0")), sdk.NewDecFromInt(convInt("0"))
	for i := range amounts0 {
		totalAmount0 = totalAmount0.Add(amounts0[i])
		totalAmount1 = totalAmount1.Add(amounts1[i])
	}

	// verify acc has sufficient balances in env
	if !(accBalance0Initial.GTE(totalAmount0)) {
		t.Errorf("%s has insufficient balance of %s. Has %s, expected at least %s.", acc, denomA, accBalance0Initial, totalAmount0)
	}
	if !(accBalance1Initial.GTE(totalAmount1)) {
		t.Errorf("%s has insufficient balance of %s. Has %s, expected at least %s.", acc, denomB, accBalance1Initial, totalAmount1)
	}

	// get Dex initial balance
	dexAllCoinsInitial := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalanceAInitial, dexBalanceBInitial := dexAllCoinsInitial.AmountOf(denomA).ToDec(), dexAllCoinsInitial.AmountOf(denomB).ToDec()

	// get amount of shares before depositing
	initialShares, initialSharesFound := env.getShares(acc, pairId, tickIndexes, feeTiers)

	// TODO: this should be array
	// get initial pair info for 0to1 and 1to0 ticks
	pairsInitial := env.makePairs(t, pairId, tickIndexes, feeTiers)

	// WHEN depositing the specified amounts0, amounts1 and given fee tiers
	_, err := env.cosmos.msgServer.Deposit(goCtx, &types.MsgDeposit{ // (discard message response because we don't need it)
		Creator:     acc.String(),
		TokenA:      denomA,
		TokenB:      denomB,
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes, // single deposit
		FeeIndexes:  feeTiers,
		Receiver:    acc.String(),
	})

	// THEN no error, alice's balances changed only by the amount depoisited, funds transfered to dex module, and position minted with appropriate fee tier
	// verify no error
	if err != nil {
		env.handleIntentionalFail(t, "Deposit of %s %v, %s %v by %s failed:\n\t%s", denom0, amounts0, denom1, amounts1, acc, err)
	}

	// verify alice's resulting balances is aliceBalanceInitial - depositCoin
	accBalance0Final, accBalance1Final := accBalance0Initial.Sub(totalAmount0), accBalance1Initial.Sub(totalAmount1)
	t.Log(accBalance0Final, accBalance1Final)
	if balance := app.BankKeeper.GetBalance(ctx, acc, denom0).Amount.ToDec(); !(balance.Equal(accBalance0Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom0, balance)
	}
	if balance := app.BankKeeper.GetBalance(ctx, acc, denom1).Amount.ToDec(); !(balance.Equal(accBalance1Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom1, balance)
	}

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexAllCoinsFinal := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalance0Final, dexBalance1Final := dexBalanceAInitial.Add(totalAmount0), dexBalanceBInitial.Add(totalAmount1)
	if balance := newACoin(dexAllCoinsFinal.AmountOf(denom0)).Amount.ToDec(); !(balance.Equal(dexBalance0Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom0, balance)
	}
	if balance := newBCoin(dexAllCoinsFinal.AmountOf(denom1)).Amount.ToDec(); !(balance.Equal(dexBalance1Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom1, balance)
	}

	for i := 0; i < len(amounts0); i++ {
		// verify amount of shares minted for acc
		accSharesCalc := env.calculateShares(t, amounts0[i], amounts1[i], pairId, tickIndexes[i], feeTiers[i])
		finalShares, finalSharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[i], feeTiers[i])
		if !finalSharesFound {
			env.handleIntentionalFail(t, "Shares resulting from deposit by %s have not been minted (not found by getter).", acc)
		} else if !initialSharesFound[i] && !(finalShares.SharesOwned.Equal(accSharesCalc)) {
			// Handle the case when no shares held by account initially but mintedShares != accSharesCalc
			env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit by %s of %s (%v), %s (%v). Needed %s, final %s.", acc, denom0, amounts0[i], denom1, amounts1[i], accSharesCalc, finalShares.SharesOwned)
		} else if initialSharesFound[i] && !finalShares.SharesOwned.Equal(initialShares[i].SharesOwned.Add(accSharesCalc)) {
			// Handle the case when account had an initial balance of shares but finalShares != initalShares + accSharesCalc
			env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit by %s of %s (%v), %s (%v). Needed %s, final %s.", acc, denom0, amounts0[i], denom1, amounts1[i], accSharesCalc, finalShares.SharesOwned)
		}

		// verify fee tier of minted shares
		if finalShares.FeeIndex != feeTiers[i] {
			env.handleIntentionalFail(t, "Shares minted in the wrong fee tier. Needed %d, final %d", feeTiers[i], finalShares.FeeIndex)
		}

		// verify current ticks set properly
		tick0to1Calc, tick1to0Calc := env.calculateNewCurrentTicks(amounts0[i], amounts1[i], tickIndexes[i], feeTiers[i], pairsInitial[i])
		pairFinal, pairFinalFound := app.DexKeeper.GetPairMap(ctx, pairId)
		if !pairFinalFound {
			env.handleIntentionalFail(t, "TODO: handle pairFinal not found")
		}
		if pairFinal.TokenPair.CurrentTick0To1 != tick0to1Calc {
			env.handleIntentionalFail(t, "Invalid CurrentTick0To1 resulted from deposit. Needed %d, final %d", tick0to1Calc, pairFinal.TokenPair.CurrentTick0To1)
		}
		if pairFinal.TokenPair.CurrentTick1To0 != tick1to0Calc {
			env.handleIntentionalFail(t, "Invalid CurrentTick1To0 resulted from deposit. Needed %d, final %d", tick1to0Calc, pairFinal.TokenPair.CurrentTick1To0)
		}
	}
}
