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
	// TODO: create trueamounts
	pairId := app.DexKeeper.CreatePairId(denom0, denom1)

	// calculate total trade amounts
	totalAmount0, totalAmount1 := sdk.NewDecFromInt(convInt("0")), sdk.NewDecFromInt(convInt("0"))
	for i := range amounts0 {
		totalAmount0 = totalAmount0.Add(amounts0[i])
		totalAmount1 = totalAmount1.Add(amounts1[i])
	}

	// verify acc has sufficient balances in env
	accBalance0Initial, accBalance1Initial := sdk.NewDecFromIntWithPrec(app.BankKeeper.GetBalance(ctx, acc, denom0).Amount, 18), sdk.NewDecFromIntWithPrec(app.BankKeeper.GetBalance(ctx, acc, denom1).Amount, 18)
	if !(accBalance0Initial.GTE(totalAmount0)) {
		t.Errorf("%s has insufficient balance of %s. Has %s, expected at least %s.", acc, denomA, accBalance0Initial, totalAmount0)
	}
	if !(accBalance1Initial.GTE(totalAmount1)) {
		t.Errorf("%s has insufficient balance of %s. Has %s, expected at least %s.", acc, denomB, accBalance1Initial, totalAmount1)
	}

	// get Dex initial balance
	dexAllCoinsInitial := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalance0Initial, dexBalance1Initial := sdk.NewDecFromIntWithPrec(dexAllCoinsInitial.AmountOf(denom0), 18), sdk.NewDecFromIntWithPrec(dexAllCoinsInitial.AmountOf(denom1), 18)

	// get amount of shares before depositing
	initialShares, initialSharesFound := env.makeShares(acc, pairId, tickIndexes, feeTiers)

	// get initial pair info for 0to1 and 1to0 ticks
	pairInitial := env.makePair(t, pairId, tickIndexes[0], feeTiers[0])

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
	if balance := sdk.NewDecFromIntWithPrec(app.BankKeeper.GetBalance(ctx, acc, denom0).Amount, 18); !(balance.Equal(accBalance0Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom0, balance)
	}
	if balance := sdk.NewDecFromIntWithPrec(app.BankKeeper.GetBalance(ctx, acc, denom1).Amount, 18); !(balance.Equal(accBalance1Final)) {
		env.handleIntentionalFail(t, "%s's final balance of %s (%s) does not reflect deposit", acc, denom1, balance)
	}

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexAllCoinsFinal := app.BankKeeper.GetAllBalances(ctx, app.AccountKeeper.GetModuleAddress("dex"))
	dexBalance0Final, dexBalance1Final := dexBalance0Initial.Add(totalAmount0), dexBalance1Initial.Add(totalAmount1)
	if balance := sdk.NewDecFromIntWithPrec(dexAllCoinsFinal.AmountOf(denom0), 18); !(balance.Equal(dexBalance0Final)) {
		env.handleIntentionalFail(t, "Dex module's final balance of %s (%s) does not reflect deposit", denom0, balance)
	}
	if balance := sdk.NewDecFromIntWithPrec(dexAllCoinsFinal.AmountOf(denom1), 18); !(balance.Equal(dexBalance1Final)) {
		env.handleIntentionalFail(t, "Dex module's final balance of %s (%s) does not reflect deposit", denom1, balance)
	}

	// accumulate shares and calculate final CurrentTicks
	sharesCalcAccum := make(map[int64]map[uint64]sdk.Dec) // map from tickIndex->feeTier->sharesCalc
	expectedTick0to1, expectedTick1to0 := pairInitial.TokenPair.CurrentTick0To1, pairInitial.TokenPair.CurrentTick1To0
	for i := range amounts0 {
		// get expected amount of minted shares and increase accum on both sides of spread
		accSharesCalc := env.calculateShares(t, amounts0[i], amounts1[i], pairId, tickIndexes[i], feeTiers[i])
		// accumulate minted shares
		if shares, ok := sharesCalcAccum[tickIndexes[i]][feeTiers[i]]; ok {
			// if already exists, add to previous value
			sharesCalcAccum[tickIndexes[i]][feeTiers[i]] = shares.Add(accSharesCalc)
		} else {
			// if inner map hasn't been initialized, init
			if _, ok := sharesCalcAccum[tickIndexes[i]]; !ok {
				sharesCalcAccum[tickIndexes[i]] = make(map[uint64]sdk.Dec)
			}
			// else add value to map
			sharesCalcAccum[tickIndexes[i]][feeTiers[i]] = accSharesCalc
		}

		// move expected current ticks
		tick0to1Calc, tick1to0Calc := env.calculateNewCurrentTicks(amounts0[i], amounts1[i], tickIndexes[i], feeTiers[i], pairInitial)
		if tick0to1Calc < expectedTick0to1 {
			expectedTick0to1 = tick0to1Calc
		}
		if tick1to0Calc > expectedTick1to0 {
			expectedTick1to0 = tick1to0Calc
		}

	}

	for i := range amounts0 {
		expectedShares := sharesCalcAccum[tickIndexes[i]][feeTiers[i]]
		// verify amount of shares minted for acc
		finalShares, finalSharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[i], feeTiers[i])
		if !finalSharesFound {
			env.handleIntentionalFail(t, "Shares resulting from deposit by %s have not been minted (not found by getter).", acc)
		} else if !initialSharesFound[i] && !(finalShares.SharesOwned.Equal(expectedShares)) {
			// Handle the case when no shares held by account initially but mintedShares != accSharesCalc
			env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit (no current shares) by %s of %s (%v), %s (%v). Needed %s, final %s.", acc, denom0, amounts0[i], denom1, amounts1[i], expectedShares, finalShares.SharesOwned)
		} else if initialSharesFound[i] && !finalShares.SharesOwned.Equal(initialShares[i].SharesOwned.Add(expectedShares)) {
			// Handle the case when account had an initial balance of shares but finalShares != initalShares + accSharesCalc
			env.handleIntentionalFail(t, "Incorrect amount of shares minted after deposit (current shares exist) by %s of %s (%v), %s (%v). Needed %s, final %s.", acc, denom0, amounts0[i], denom1, amounts1[i], expectedShares, finalShares.SharesOwned)
		}

		// verify fee tier of minted shares
		if finalShares.FeeIndex != feeTiers[i] {
			env.handleIntentionalFail(t, "Shares minted in the wrong fee tier. Needed %d, final %d", feeTiers[i], finalShares.FeeIndex)
		}
	}

	// verify current ticks set properly
	pairFinal, pairFinalFound := app.DexKeeper.GetPairMap(ctx, pairId)
	if !pairFinalFound {
		env.handleIntentionalFail(t, "TODO: handle pairFinal not found")
	}
	if pairFinal.TokenPair.CurrentTick0To1 != expectedTick0to1 {
		env.handleIntentionalFail(t, "Invalid CurrentTick0To1 resulted from deposit. Needed %d, final %d", expectedTick0to1, pairFinal.TokenPair.CurrentTick0To1)
	}
	if pairFinal.TokenPair.CurrentTick1To0 != expectedTick1to0 {
		env.handleIntentionalFail(t, "Invalid CurrentTick1To0 resulted from deposit. Needed %d, final %d", expectedTick1to0, pairFinal.TokenPair.CurrentTick1To0)
	}
}
