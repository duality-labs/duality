package keeper_test

import (
	"testing"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Helpers for calculating the amount of shares that should be minted
// These are here for the sake of more concise unit tests and the corresponding code in core.go
// should eventually be refactored so that core.go is modularized for easier unit testing and readability

// Calculation of shares when depositing the initial amount (no reserves)
func calculateSharesEmpty(amount0 sdk.Dec, amount1 sdk.Dec, price sdk.Dec) sdk.Dec {
	return amount0.Add(amount1.Mul(price))
}

// Calculation of shares when there are pre-existing reserves
func calculateSharesNonEmpty(amount sdk.Dec, reserve sdk.Dec, totalShares sdk.Dec) sdk.Dec {
	return amount.Quo(reserve).Mul(totalShares)
}

// Pure func that takes all the parameters requires to compute the amount of minted shares and handles the different cases accordingly.
// This is probably excessive as keeping only the calculation pure is reasonable enough, but it's here for posterity.
func calculateSharesPure(
	amount0 sdk.Dec,
	trueAmount0 sdk.Dec,
	amount1 sdk.Dec,
	trueAmount1 sdk.Dec,
	price sdk.Dec,
	feeIndex uint64,
	lowerTickFound bool,
	lowerReserve1 sdk.Dec,
	upperTickFound bool,
	upperReserve0 sdk.Dec,
	upperTotalShares sdk.Dec,
) sdk.Dec {
	// calculating shares minted in DepositHelper
	if !lowerTickFound || !upperTickFound || upperTotalShares.Equal(sdk.ZeroDec()) {
		// this case corresponds to lines 129-132 in function DepositHelper of core.go
		return calculateSharesEmpty(amount0, amount1, price)
	} else {
		// these cases correspond to lines 228-234 in function DepositHelper of core.go
		if trueAmount0.GT(sdk.ZeroDec()) {
			return calculateSharesNonEmpty(trueAmount0, upperReserve0, upperTotalShares)
		} else {
			return calculateSharesNonEmpty(trueAmount1, lowerReserve1, upperTotalShares)
		}
	}
}

// Pure helper function to balance amounts to pool ratio
func trueAmounts(amount0 sdk.Dec, amount1 sdk.Dec, lowerReserve1 sdk.Dec, upperReserve0 sdk.Dec) (sdk.Dec, sdk.Dec) {
	trueAmount0, trueAmount1 := amount0, amount1
	if upperReserve0.GT(sdk.ZeroDec()) {
		// this corresponds to lines 217-221 in function DepositHelper of core.go
		trueAmount1 = min(amount1, lowerReserve1.Mul(amount0).Quo(upperReserve0))
	}
	if lowerReserve1.GT(sdk.ZeroDec()) {
		// this corresponds to lines 223-226 in function DepositHelper of core.go
		trueAmount0 = min(amount0, upperReserve0.Mul(amount1).Quo(lowerReserve1))
	}
	return trueAmount0, trueAmount1
}

// Impure function that pulls all the state variables required for calculating the amount of shares to mint.
func (env *TestEnv) calculateShares(t *testing.T, amount0 sdk.Dec, amount1 sdk.Dec, pairId string, tickIndex int64, feeIndex uint64) sdk.Dec {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx

	price, err := k.Calc_price(tickIndex, false)
	if err != nil {
		env.handleIntentionalFail(t, "TODO: calc price error format")
	}

	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee

	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-fee)
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+fee)
	lowerReserve1 := lowerTick.TickData.Reserve1[feeIndex]
	upperReserve0, upperTotalShares := upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares

	trueAmount0, trueAmount1 := trueAmounts(amount0, amount1, lowerReserve1, upperReserve0)

	return calculateSharesPure(
		amount0,
		trueAmount0,
		amount1,
		trueAmount1,
		price,
		feeIndex,
		lowerTickFound,
		lowerReserve1,
		upperTickFound,
		upperReserve0,
		upperTotalShares,
	)
}

// Helper function to calculate if current ticks change
func calculateNewCurrentTicksPure(amount0 sdk.Dec, amount1 sdk.Dec, tickIndex int64, fee int64, curr0to1 int64, curr1to0 int64) (int64, int64) {
	// this corresponds to lines 245-253 in function DepositHelper of core.go
	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	new0to1, new1to0 := curr0to1, curr1to0
	if amount0.GT(sdk.ZeroDec()) && ((tickIndex+fee > curr0to1) && (tickIndex+fee < curr1to0)) {
		new1to0 = tickIndex + fee
	}
	if amount1.GT(sdk.ZeroDec()) && ((tickIndex-fee > curr0to1) && (tickIndex-fee < curr1to0)) {
		new0to1 = tickIndex - fee
	}
	return new0to1, new1to0
}

func (env *TestEnv) calculateNewCurrentTicks(amount0 sdk.Dec, amount1 sdk.Dec, tickIndex int64, feeIndex uint64, pair types.PairMap) (new0to1 int64, new1to0 int64) {
	k, ctx := env.cosmos.app.DexKeeper, env.cosmos.ctx
	feelist := k.GetAllFeeList(ctx)
	fee := feelist[feeIndex].Fee
	return calculateNewCurrentTicksPure(amount0, amount1, tickIndex, fee, pair.TokenPair.CurrentTick0To1, pair.TokenPair.CurrentTick1To0)
}

// Helper for getting a pair id
func (env *TestEnv) makePairs(t *testing.T, pairId string, tickIndexes []int64, feeTiers []uint64) []types.PairMap {
	// TODO: this really should be cleaned up
	app, ctx, k := env.cosmos.app, env.cosmos.ctx, env.cosmos.app.DexKeeper

	// this corresponds to line 16 in function DepositVerification of verification.go
	feeList := k.GetAllFeeList(ctx)

	var pairs []types.PairMap

	for i, _ := range feeTiers {
		// handle invalid fee index
		if feeTiers[i] >= uint64(len(feeList)) {
			env.handleIntentionalFail(t, "Fee index (%d) > fee tier count (%d)", feeTiers[i], len(feeList))
		}

		pair, pairFound := app.DexKeeper.GetPairMap(ctx, pairId)
		if !pairFound {
			pair = types.PairMap{
				PairId: pairId,
				TokenPair: &types.TokenPairType{
					CurrentTick0To1: tickIndexes[i] - feeList[feeTiers[i]].Fee,
					CurrentTick1To0: tickIndexes[i] + feeList[feeTiers[i]].Fee,
				},
				TotalTickCount: 0,
			}
		}
		pairs = append(pairs, pair)
	}

	return pairs
}

// Helper for getting shares. If no shares object exists, marks it as not found and fills in with an empty object. Must handle using the "found" bools.
func (env *TestEnv) getShares(acc sdk.AccAddress, pairId string, tickIndexes []int64, feeTiers []uint64) ([]types.Shares, []bool) {
	app, ctx := env.cosmos.app, env.cosmos.ctx
	var sharesSlice []types.Shares
	var sharesFoundSlice []bool
	for i, _ := range tickIndexes {
		shares, sharesFound := app.DexKeeper.GetShares(ctx, acc.String(), pairId, tickIndexes[i], feeTiers[i])
		if !sharesFound {
			// if shares not found verification will handle, so add empty object that will be ignored (to maintain length)
			shares = types.Shares{}
		}
		sharesSlice = append(sharesSlice, shares)
		sharesFoundSlice = append(sharesFoundSlice, sharesFound)
	}

	return sharesSlice, sharesFoundSlice
}
