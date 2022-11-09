package keeper_test

import (
	// stdlib

	"testing"

	// cosmos SDK

	sdk "github.com/cosmos/cosmos-sdk/types"
	// duality
	// "github.com/NicholasDotSol/duality/x/dex/types"
)

func TestMultiMinFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/MinFee")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceB, BA = 1/3 * balanceA, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit with min fee tier:
	tickIndexes := []int64{0, 0}
	minFeeTiers := []uint64{0, 0}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, minFeeTiers)

	// THEN the transaction should execute successfully
	// validity assertions are done inside env.TestDeposit
}

func TestMultiMaxFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/MaxFee")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceA, BA = 1/3 * balanceB, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit with max fee tier
	tickIndexes := []int64{0, 0}
	maxFeeTiers := []uint64{uint64(len(env.feeTiers) - 1), uint64(len(env.feeTiers) - 1)}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, maxFeeTiers)

	// THEN the transaction should execute successfully
	// validity assertions are done inside env.TestDeposit
}

func TestMultiInvalidFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/Invalid Fee Tier")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, true)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceA, BA = 1/3 * balanceB, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndexes := []int64{0, 0}
	invalidFeeTiers := []uint64{uint64(len(env.feeTiers)), uint64(len(env.feeTiers))}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, invalidFeeTiers)

	// THEN the transaction should fail midway (using SkipNow)
}

func TestMultiInitPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/InitPair")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceB, BA = 1/3 * balanceA, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit at tick 0, feeIndex 0 and tick 0 feeIndex 2
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 2}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers)

	// THEN the transaction should execute successfully
	// validity assertions are done inside env.TestDeposit
}

func TestMultiInitTick(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/InitTick")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceB, BA = 1/3 * balanceA, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit at tick 0 feeIndex 0, tick 3 feeIndex 0
	tickIndexes := []int64{0, 3}
	feeTiers := []uint64{0, 0}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers)

	// THEN the transaction should execute successfully
}

func TestMultiInitFeeTier(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/InitFeeTier")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceA, BA = 1/3 * balanceB, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit at tick 0 fee, tick 0 feeIndex 2
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 2}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers)

	// THEN the transaction should execute successfully
}

func TestMultiExistingPair(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/ExistingPair")

	// GIVEN initial balances and fee tiers from the setup
	env := EnvSetup(t, false)

	// WHEN alice deposits her setup balance of tokenA and tokenB into the minimal fee tier
	// prep deposit args
	acc := env.addrs[0]
	// AA = 1/5 * balanceA, AB = 1/5 * balanceA, BA = 1/3 * balanceB, BB = 1/3 * balanceB
	coinA, coinB := env.balances[acc.String()][0], env.balances[acc.String()][1]
	amountAA, amountAB, amountBA, amountBB := coinA.Amount.Quo(convInt("5")), coinA.Amount.Quo(convInt("5")), coinB.Amount.Quo(convInt("3")), coinB.Amount.Quo(convInt("3"))
	denomA, denomB := coinA.Denom, coinB.Denom
	amountsA, amountsB := []sdk.Dec{sdk.NewDecFromIntWithPrec(amountAA, 18), sdk.NewDecFromIntWithPrec(amountAB, 18)}, []sdk.Dec{sdk.NewDecFromIntWithPrec(amountBA, 18), sdk.NewDecFromIntWithPrec(amountBB, 18)}

	// deposit with invalid fee tier: maxFeeTier + 1 > maxFeeTier, i.e. invalid
	tickIndexes := []int64{0, 0}
	feeTiers := []uint64{0, 2}

	env.TestDeposit(t, denomA, denomB, amountsA, amountsB, acc, tickIndexes, feeTiers)

	// THEN the transaction should execute successfully
}

func TestMultiBehindEnemyLines(t *testing.T) {
	t.Logf("[ UnitTests|Keeper ] Starting test: MultiPool/BehindEnemyLines")

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
