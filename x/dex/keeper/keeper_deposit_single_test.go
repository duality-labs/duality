package keeper_test

import (
	// stdlib

	"testing"

	// cosmos SDK

	sdk "github.com/cosmos/cosmos-sdk/types"
	// duality
	// "github.com/NicholasDotSol/duality/x/dex/types"
)

func TestSingleMinFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/MinFee")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	minFeeTier := []uint64{0}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, minFeeTier)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestSingleMaxFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/MaxFee")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	maxFeeTier := []uint64{uint64(len(env.feeTiers) - 1)}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, maxFeeTier)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestSingleInvalidFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/Invalid Fee Tier")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, true)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndex := []int64{0}
	invalidFeeTier := []uint64{uint64(len(env.feeTiers))}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, invalidFeeTier)

	// THEN the transaction should fail midway (SkipNow)
}

func TestSingleInitPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitPair")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// THEN the transaction should execute successfully
	// validity assertions are done inside testSingleDeposit
}

func TestSingleInitTick(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitTick")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := EnvSetup(t, false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// WHEN alice deposits at tick 1 in fee tier 0
	newTickIndex := []int64{1}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, newTickIndex, feeTier)

	// THEN the transaction should execute successfully
}

func TestSingleInitFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/InitFeeTier")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := EnvSetup(t, false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// WHEN alice deposits at tick 0 in fee tier 1
	newFeeTier := []uint64{1}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, newFeeTier)

	// THEN the transaction should execute successfully
}

func TestSingleExistingPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/ExistingPair")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := EnvSetup(t, false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// WHEN deposit in the same pair, tick and fee tier again
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// THEN the transaction should execute successfully
}

func TestSingleBehindEnemyLines(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: SinglePool/BehindEnemyLines")

	// GIVEN initial balances and fee tiers from the setup, and pair already has liquidity at tick 0 fee tier 0
	env := EnvSetup(t, false)
	acc := env.addrs[0]
	// fifth of acc's balance of each coin
	coinA, coinB := newACoin(env.balances[acc.String()][0].Amount.Quo(convInt("5"))), newBCoin(env.balances[acc.String()][1].Amount.Quo(convInt("5")))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(coinA.Amount, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(coinB.Amount, 18)}

	// deposit at tick 0 in fee tier 0
	tickIndex := []int64{0}
	feeTier := []uint64{0}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndex, feeTier)

	// WHEN alice deposits at tick 0 in fee tier 1
	newTickIndex := []int64{-3}
	newFeeTier := []uint64{1}
	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, newTickIndex, newFeeTier)

	// THEN the transaction should execute successfully
}
