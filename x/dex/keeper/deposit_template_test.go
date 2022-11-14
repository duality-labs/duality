package keeper_test

import (
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Template function for testing the execution of deposit.
func DepositTemplate(s *MsgServerTestSuite, denomA string, denomB string, amountsA []sdk.Dec, amountsB []sdk.Dec, acc sdk.AccAddress, tickIndexes []int64, feeTiers []uint64, expectedTxErr error) {
	app, ctx := s.app, s.ctx
	goCtx := sdk.WrapSDKContext(ctx)

	// validate inputs
	s.Assert().True(len(amountsA) == len(amountsB))
	denom0, denom1, amounts0, amounts1 := sortCoins(s, denomA, denomB, amountsA, amountsB)
	// TODO: create trueAmounts (the amounts passed aren't necessarily the amounts deposited)
	pairId := app.DexKeeper.CreatePairId(denom0, denom1)

	accBalance0Initial, accBalance1Initial := getBalance(s, acc, denom0), getBalance(s, acc, denom1)
	totalAmount0, totalAmount1 := getTotalAmount(amounts0), getTotalAmount(amounts1)

	// verify acc has sufficient balances in env
	s.Assert().True(accBalance0Initial.GTE(totalAmount0))
	s.Assert().True(accBalance1Initial.GTE(totalAmount1))

	// get Dex initial balance
	dexBalance0Initial, dexBalance1Initial := getDexBalance(s, denom0), getDexBalance(s, denom1)
	// get amount of shares before depositing
	initialShares, initialSharesFound := makeShares(s, acc, pairId, tickIndexes, feeTiers)
	// get initial pair info for 0to1 and 1to0 ticks
	pairInitial := makePair(s, pairId, tickIndexes[0], feeTiers[0], expectedTxErr)

	// deposit the specified amounts0, amounts1 and given fee tiers
	_, err := s.msgServer.Deposit(goCtx, &types.MsgDeposit{ // (discard message response because we don't need it)
		Creator:     acc.String(),
		TokenA:      denomA,
		TokenB:      denomB,
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes, // single deposit
		FeeIndexes:  feeTiers,
		Receiver:    acc.String(),
	})

	// verify no error
	if expectedTxErr != nil {
		s.Require().True(strings.Contains(err.Error(), expectedTxErr.Error()))
		return
	} else {
		s.Require().Nil(err)
	}

	// verify alice's resulting balances is aliceBalanceInitial - depositCoin
	accBalance0Final, accBalance1Final := accBalance0Initial.Sub(totalAmount0), accBalance1Initial.Sub(totalAmount1)
	s.Assert().True(getBalance(s, acc, denom0).Equal(accBalance0Final))
	s.Assert().True(getBalance(s, acc, denom1).Equal(accBalance1Final))

	// verify dex's resulting balances is dexBalanceInitial + depositCoin
	dexBalance0Final, dexBalance1Final := dexBalance0Initial.Add(totalAmount0), dexBalance1Initial.Add(totalAmount1)
	s.Assert().True(getDexBalance(s, denom0).Equal(dexBalance0Final))
	s.Assert().True(getDexBalance(s, denom1).Equal(dexBalance1Final))

	// accumulate shares and calculate final CurrentTicks
	expectedTick0to1, expectedTick1to0 := calculateFinalTicks(s, pairInitial, amounts0, amounts1, tickIndexes, feeTiers)
	sharesCalcAccum := calculateFinalShares(s, pairId, amounts0, amounts1, tickIndexes, feeTiers)
	// verify correct amount of shares in every tick
	for i := range amounts0 {
		expectedShares := sharesCalcAccum[tickIndexes[i]][feeTiers[i]]
		// verify amount of shares minted for acc
		finalShares, finalSharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[i], feeTiers[i])
		s.Assert().True(finalSharesFound)
		if !initialSharesFound[i] {
			// Handle the case when no shares held by account initially but mintedShares != accSharesCalc
			s.Assert().True(finalShares.SharesOwned.Equal(expectedShares))
		} else if initialSharesFound[i] {
			// Handle the case when account had an initial balance of shares but finalShares != initalShares + accSharesCalc
			s.Assert().True(finalShares.SharesOwned.Equal(initialShares[i].SharesOwned.Add(expectedShares)))
		}

		// verify fee tier of minted shares
		s.Assert().True(finalShares.FeeIndex == feeTiers[i])
	}

	// verify current ticks set properly
	pairFinal, pairFinalFound := app.DexKeeper.GetPairMap(ctx, pairId)
	s.Assert().True(pairFinalFound)
	s.Assert().True(pairFinal.TokenPair.CurrentTick0To1 == expectedTick0to1)
	s.Assert().True(pairFinal.TokenPair.CurrentTick1To0 == expectedTick1to0)
}
