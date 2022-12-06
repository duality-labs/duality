package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getTotalAmount(amounts []sdk.Int) sdk.Int {
	// calculate total trade amounts
	totalAmount := sdk.NewInt(0)
	for i := range amounts {
		totalAmount = totalAmount.Add(amounts[i])
	}
	return totalAmount
}

// Pure helper function to balance amounts to pool ratio
func trueAmounts(amount0 sdk.Int, amount1 sdk.Int, lowerReserve1 sdk.Int, upperReserve0 sdk.Int) (sdk.Int, sdk.Int) {
	trueAmount0, trueAmount1 := amount0, amount1
	if upperReserve0.GT(sdk.ZeroInt()) {
		// this corresponds to lines 217-221 in function DepositHelper of core.go
		trueAmount1 = min(amount1, lowerReserve1.Mul(amount0).ToDec().Quo(upperReserve0.ToDec()).TruncateInt())
	}
	if lowerReserve1.GT(sdk.ZeroInt()) {
		// this corresponds to lines 223-226 in function DepositHelper of core.go
		trueAmount0 = min(amount0, upperReserve0.Mul(amount1).Quo(lowerReserve1))
	}
	return trueAmount0, trueAmount1
}

// Pure helper function for calculateNewCurrentTicks
func calculateNewCurrentTicksPure(amount0 sdk.Int, amount1 sdk.Int, tickIndex int64, fee int64, curr0to1 int64, curr1to0 int64) (int64, int64) {
	// this corresponds to lines 245-253 in function DepositHelper of core.go
	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	new0to1, new1to0 := curr0to1, curr1to0
	if amount0.GT(sdk.ZeroInt()) && ((tickIndex+fee > curr0to1) && (tickIndex+fee < curr1to0)) {
		new1to0 = tickIndex + fee
	}
	if amount1.GT(sdk.ZeroInt()) && ((tickIndex-fee > curr0to1) && (tickIndex-fee < curr1to0)) {
		new0to1 = tickIndex - fee
	}
	return new0to1, new1to0
}

// Helper function to calculate if current ticks change
func calculateNewCurrentTicks(s *MsgServerTestSuite, amount0 sdk.Int, amount1 sdk.Int, tickIndex int64, feeIndex uint64, pair types.TradingPair) (new0to1 int64, new1to0 int64) {
	k, ctx := s.app.DexKeeper, s.ctx
	FeeTier := k.GetAllFeeTier(ctx)
	fee := FeeTier[feeIndex].Fee
	return calculateNewCurrentTicksPure(amount0, amount1, tickIndex, fee, pair.CurrentTick0To1, pair.CurrentTick1To0)
}

// Helper for getting a pair id. If pair hasn't been initialized, defaults to pair with tickIndex and feeTier for CurrentTick
func makePair(s *MsgServerTestSuite, pairId string, tickIndex int64, feeTier uint64, expectedTxErr error) types.TradingPair {
	// TODO: this really should be cleaned up
	app, ctx, k := s.app, s.ctx, s.app.DexKeeper

	// this corresponds to line 16 in function DepositVerification of verification.go
	FeeTier := k.GetAllFeeTier(ctx)

	var fee int64
	// handle invalid fee index
	if feeTier >= uint64(len(FeeTier)) {
		s.Assert().True(expectedTxErr == types.ErrValidFeeIndexNotFound)
		fee = 0
	} else {
		fee = FeeTier[feeTier].Fee
	}

	pair, pairFound := app.DexKeeper.GetTradingPair(ctx, pairId)
	if !pairFound {
		pair = types.TradingPair{
			PairId:          pairId,
			MinTick:         tickIndex - fee,
			MaxTick:         tickIndex + fee,
			CurrentTick0To1: tickIndex - fee,
			CurrentTick1To0: tickIndex + fee,
		}
	}

	return pair
}

// Helpers for calculating the amount of shares that should be minted
// These are here for the sake of more concise unit tests and the corresponding code in core.go
// should eventually be refactored so that core.go is modularized for easier unit testing and readability

// Calculation of shares when depositing the initial amount (no reserves)
func calculateSharesEmpty(amount0 sdk.Int, amount1 sdk.Int, price sdk.Dec) sdk.Dec {
	return price.MulInt(amount1).Add(amount0.ToDec())
}

// Calculation of shares when there are pre-existing reserves
func calculateSharesNonEmpty(amount sdk.Int, reserve sdk.Int, totalShares sdk.Dec) sdk.Dec {
	return amount.ToDec().QuoInt(reserve).Mul(totalShares)
}

// Pure func that takes all the parameters requires to compute the amount of minted shares and handles the different cases accordingly.
// This is probably excessive as keeping only the calculation pure is reasonable enough, but it's here for posterity.
func calculateSharesPure(
	amount0 sdk.Int,
	trueAmount0 sdk.Int,
	amount1 sdk.Int,
	trueAmount1 sdk.Int,
	price sdk.Dec,
	feeIndex uint64,
	lowerTickFound bool,
	lowerReserve1 sdk.Int,
	upperTickFound bool,
	upperReserve0 sdk.Int,
	upperTotalShares sdk.Dec,
) sdk.Dec {
	// calculating shares minted in DepositHelper
	if !lowerTickFound || !upperTickFound || upperTotalShares.Equal(sdk.ZeroDec()) {
		// this case corresponds to lines 129-132 in function DepositHelper of core.go
		return calculateSharesEmpty(amount0, amount1, price)
	} else {
		// these cases correspond to lines 228-234 in function DepositHelper of core.go
		if trueAmount0.GT(sdk.ZeroInt()) {
			return calculateSharesNonEmpty(trueAmount0, upperReserve0, upperTotalShares)
		} else {
			return calculateSharesNonEmpty(trueAmount1, lowerReserve1, upperTotalShares)
		}
	}
}

// Impure function that pulls all the state variables required for calculating the amount of shares to mint.
func calculateShares(s *MsgServerTestSuite, amount0 sdk.Int, amount1 sdk.Int, pairId string, tickIndex int64, feeIndex uint64) sdk.Dec {
	k, ctx := s.app.DexKeeper, s.ctx

	price1To0 := keeper.CalcPrice1To0(tickIndex)

	FeeTier := k.GetAllFeeTier(ctx)
	fee := FeeTier[feeIndex].Fee

	lowerTick, lowerTickFound := k.GetTick(ctx, pairId, tickIndex-fee)
	upperTick, upperTickFound := k.GetTick(ctx, pairId, tickIndex+fee)
	lowerReserve1 := lowerTick.TickData.Reserve1[feeIndex]
	upperReserve0, upperTotalShares := upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares

	trueAmount0, trueAmount1 := trueAmounts(amount0, amount1, lowerReserve1, upperReserve0)

	return calculateSharesPure(
		amount0,
		trueAmount0,
		amount1,
		trueAmount1,
		price1To0,

		feeIndex,
		lowerTickFound,
		lowerReserve1,
		upperTickFound,
		upperReserve0,
		upperTotalShares,
	)
}

// Helper for getting shares. If no shares object exists, marks it as not found and fills in with an empty object. Must handle using the "found" bools.
func makeShares(s *MsgServerTestSuite, acc sdk.AccAddress, pairId string, tickIndexes []int64, feeTiers []uint64) ([]types.Shares, []bool) {
	app, ctx := s.app, s.ctx
	var sharesSlice []types.Shares
	var sharesFoundSlice []bool
	for i := range tickIndexes {
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

type SharesMap = map[int64]map[uint64]sdk.Dec

func calculateFinalShares(s *MsgServerTestSuite, pairId string, amounts0 []sdk.Int, amounts1 []sdk.Int, tickIndexes []int64, feeTiers []uint64) SharesMap {
	accum := make(map[int64]map[uint64]sdk.Dec) // map from tickIndex->feeTier->sharesCalc
	for i := range amounts0 {
		// get expected amount of minted shares and increase accum on both sides of spread
		accSharesCalc := calculateShares(s, amounts0[i], amounts1[i], pairId, tickIndexes[i], feeTiers[i])
		// accumulate minted shares
		if shares, ok := accum[tickIndexes[i]][feeTiers[i]]; ok {
			// if already exists, add to previous value
			accum[tickIndexes[i]][feeTiers[i]] = shares.Add(accSharesCalc)
		} else {
			// if inner map hasn't been initialized, init
			if _, ok := accum[tickIndexes[i]]; !ok {
				accum[tickIndexes[i]] = make(map[uint64]sdk.Dec)
			}
			// else add value to map
			accum[tickIndexes[i]][feeTiers[i]] = accSharesCalc
		}
	}
	return accum
}

func calculateFinalTicks(s *MsgServerTestSuite, pair types.TradingPair, amounts0 []sdk.Int, amounts1 []sdk.Int, tickIndexes []int64, feeTiers []uint64) (int64, int64) {
	expectedTick0to1, expectedTick1to0 := pair.CurrentTick0To1, pair.CurrentTick1To0
	for i := range amounts0 {
		// move expected current ticks
		tick0to1Calc, tick1to0Calc := calculateNewCurrentTicks(s, amounts0[i], amounts1[i], tickIndexes[i], feeTiers[i], pair)
		if tick0to1Calc > expectedTick0to1 {
			expectedTick0to1 = tick0to1Calc
		}
		if tick1to0Calc < expectedTick1to0 {
			expectedTick1to0 = tick1to0Calc
		}
	}
	return expectedTick0to1, expectedTick1to0
}
