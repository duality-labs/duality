package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dualityapp "github.com/duality-labs/duality/app"
	"github.com/duality-labs/duality/x/dex/keeper"
	. "github.com/duality-labs/duality/x/dex/keeper"
	. "github.com/duality-labs/duality/x/dex/keeper/internal/testutils"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// / Test suite
type MsgServerTestSuite struct {
	suite.Suite
	app         *dualityapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
	alice       sdk.AccAddress
	bob         sdk.AccAddress
	carol       sdk.AccAddress
	dan         sdk.AccAddress
	goCtx       context.Context
	feeTiers    []types.FeeTier
}

var defaultPairId *types.PairId = &types.PairId{Token0: "TokenA", Token1: "TokenB"}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
}

func (s *MsgServerTestSuite) SetupTest() {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.DexKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	accAlice := app.AccountKeeper.NewAccountWithAddress(ctx, s.alice)
	app.AccountKeeper.SetAccount(ctx, accAlice)
	accBob := app.AccountKeeper.NewAccountWithAddress(ctx, s.bob)
	app.AccountKeeper.SetAccount(ctx, accBob)
	accCarol := app.AccountKeeper.NewAccountWithAddress(ctx, s.carol)
	app.AccountKeeper.SetAccount(ctx, accCarol)
	accDan := app.AccountKeeper.NewAccountWithAddress(ctx, s.dan)
	app.AccountKeeper.SetAccount(ctx, accDan)

	// add the fee tiers of 1, 3, 5, 10 ticks
	feeTiers := []types.FeeTier{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[0])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[1])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[2])
	app.DexKeeper.AppendFeeTier(ctx, feeTiers[3])

	s.app = app
	s.msgServer = keeper.NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
	s.feeTiers = feeTiers
}

/// Fund accounts

func (s *MsgServerTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance int64, bBalance int64) {
	aBalanceInt := sdk.NewInt(aBalance)
	bBalanceInt := sdk.NewInt(bBalance)
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalances(account, aBalance, bBalance)
}

func (s *MsgServerTestSuite) fundAliceBalances(a int64, b int64) {
	s.fundAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundBobBalances(a int64, b int64) {
	s.fundAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalances(a int64, b int64) {
	s.fundAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundDanBalances(a int64, b int64) {
	s.fundAccountBalances(s.dan, a, b)
}

/// Assert balances

func (s *MsgServerTestSuite) assertAccountBalancesInt(
	account sdk.AccAddress,
	aBalance sdk.Int,
	bBalance sdk.Int,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount
	s.Assert().True(aBalance.Equal(aActual), "expected %s != actual %s", aBalance, aActual)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount
	s.Assert().True(bBalance.Equal(bActual), "expected %s != actual %s", bBalance, bActual)
}

func (s *MsgServerTestSuite) assertAccountBalances(
	account sdk.AccAddress,
	aBalance int64,
	bBalance int64,
) {
	s.assertAccountBalancesInt(account, sdk.NewInt(aBalance), sdk.NewInt(bBalance))
}

func (s *MsgServerTestSuite) assertAliceBalances(a int64, b int64) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a int64, b int64) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalances(a int64, b int64) {
	s.assertAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a int64, b int64) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a int64, b int64) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) traceBalances() {
	aliceA := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenA")
	aliceB := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenB")
	bobA := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenA")
	bobB := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenB")
	carolA := s.app.BankKeeper.GetBalance(s.ctx, s.carol, "TokenA")
	carolB := s.app.BankKeeper.GetBalance(s.ctx, s.carol, "TokenB")
	danA := s.app.BankKeeper.GetBalance(s.ctx, s.dan, "TokenA")
	danB := s.app.BankKeeper.GetBalance(s.ctx, s.dan, "TokenB")
	fmt.Printf(
		"Alice: %+v %+v\nBob: %+v %+v\nCarol: %+v %+v\nDan: %+v %+v",
		aliceA, aliceB,
		bobA, bobB,
		carolA, carolB,
		danA, danB,
	)
}

/// Place limit order

func (s *MsgServerTestSuite) aliceLimitSells(selling string, tick int, amountIn int) string {
	return s.limitSells(s.alice, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) bobLimitSells(selling string, tick int, amountIn int) string {
	return s.limitSells(s.bob, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) carolLimitSells(selling string, tick int, amountIn int) string {
	return s.limitSells(s.carol, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) danLimitSells(selling string, tick int, amountIn int) string {
	return s.limitSells(s.dan, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) limitSells(account sdk.AccAddress, tokenIn string, tick int, amountIn int) string {
	msg, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  sdk.NewInt(int64(amountIn)),
		OrderType: types.LimitOrderType_GOOD_TIL_CANCELLED,
	})
	s.Assert().Nil(err)
	return msg.TrancheKey
}

func (s *MsgServerTestSuite) assertAliceLimitSellFails(err error, selling string, tick int, amountIn int) {
	s.assertLimitSellFails(s.alice, err, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) assertBobLimitSellFails(err error, selling string, tick int, amountIn int) {
	s.assertLimitSellFails(s.bob, err, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) assertCarolLimitSellFails(err error, selling string, tick int, amountIn int) {
	s.assertLimitSellFails(s.carol, err, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) assertDanLimitSellFails(err error, selling string, tick int, amountIn int) {
	s.assertLimitSellFails(s.dan, err, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) assertLimitSellFails(account sdk.AccAddress, expectedErr error, tokenIn string, tick int, amountIn int) {
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  sdk.NewInt(int64(amountIn)),
	})
	s.Assert().ErrorIs(err, expectedErr)
}

// / Deposit
type Deposit struct {
	AmountA   sdk.Int
	AmountB   sdk.Int
	TickIndex int64
	FeeIndex  uint64
}

type DepositOptions struct {
	Autoswap bool
}

type DepositWithOptions struct {
	AmountA   sdk.Int
	AmountB   sdk.Int
	TickIndex int64
	FeeIndex  uint64
	Options   DepositOptions
}

func NewDeposit(amountA int, amountB int, tickIndex int, feeIndex int) *Deposit {
	return &Deposit{
		AmountA:   sdk.NewInt(int64(amountA)),
		AmountB:   sdk.NewInt(int64(amountB)),
		TickIndex: int64(tickIndex),
		FeeIndex:  uint64(feeIndex),
	}
}

func NewDepositWithOptions(amountA int, amountB int, tickIndex int, feeIndex int, options DepositOptions) *DepositWithOptions {
	return &DepositWithOptions{
		AmountA:   sdk.NewInt(int64(amountA)),
		AmountB:   sdk.NewInt(int64(amountB)),
		TickIndex: int64(tickIndex),
		FeeIndex:  uint64(feeIndex),
		Options:   options,
	}
}

func (s *MsgServerTestSuite) aliceDeposits(deposits ...*Deposit) {
	s.deposits(s.alice, deposits...)
}

func (s *MsgServerTestSuite) aliceDepositsWithOptions(deposits ...*DepositWithOptions) {
	s.depositsWithOptions(s.alice, deposits...)
}

func (s *MsgServerTestSuite) bobDeposits(deposits ...*Deposit) {
	s.deposits(s.bob, deposits...)
}

func (s *MsgServerTestSuite) carolDeposits(deposits ...*Deposit) {
	s.deposits(s.carol, deposits...)
}

func (s *MsgServerTestSuite) danDeposits(deposits ...*Deposit) {
	s.deposits(s.dan, deposits...)
}

func (s *MsgServerTestSuite) deposits(account sdk.AccAddress, deposits ...*Deposit) {
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		options[i] = &types.DepositOptions{false}
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     account.String(),
		Receiver:    account.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes,
		FeeIndexes:  feeIndexes,
		Options:     options,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) depositsWithOptions(account sdk.AccAddress, deposits ...*DepositWithOptions) {
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		options[i] = &types.DepositOptions{
			Autoswap: e.Options.Autoswap,
		}
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     account.String(),
		Receiver:    account.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes,
		FeeIndexes:  feeIndexes,
		Options:     options,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) getLiquidityAtTick(tickIndex int64, feeIndex uint64) (sdk.Int, sdk.Int) {
	pairId := CreatePairId("TokenA", "TokenB")
	feeTier := s.feeTiers[feeIndex]
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, pairId, tickIndex, feeTier)
	s.Assert().NoError(err)

	liquidityA := pool.LowerTick0.Reserves
	liquidityB := pool.UpperTick1.Reserves

	if &liquidityA == nil {
		liquidityA = sdk.ZeroInt()
	}

	if &liquidityB == nil {
		liquidityB = sdk.ZeroInt()
	}

	return liquidityA, liquidityB
}

func (s *MsgServerTestSuite) assertAliceDepositFails(err error, deposits ...*Deposit) {
	s.assertDepositFails(s.alice, err, deposits...)
}

func (s *MsgServerTestSuite) assertBobDepositFails(err error, deposits ...*Deposit) {
	s.assertDepositFails(s.bob, err, deposits...)
}

func (s *MsgServerTestSuite) assertCarolDepositFails(err error, deposits ...*Deposit) {
	s.assertDepositFails(s.carol, err, deposits...)
}

func (s *MsgServerTestSuite) assertDanDepositFails(err error, deposits ...*Deposit) {
	s.assertDepositFails(s.dan, err, deposits...)
}
func (s *MsgServerTestSuite) assertDepositFails(account sdk.AccAddress, expectedErr error, deposits ...*Deposit) {
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		options[i] = &types.DepositOptions{false}
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     account.String(),
		Receiver:    account.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndexes,
		FeeIndexes:  feeIndexes,
		Options:     options,
	})
	s.Assert().NotNil(err)
	s.Assert().ErrorIs(err, expectedErr)
}

func (s *MsgServerTestSuite) assertDepositReponse(depositResponse DepositReponse, expectedDepositResponse DepositReponse) {

	for i, _ := range expectedDepositResponse.amountsA {
		s.Assert().Equal(
			depositResponse.amountsA[i],
			expectedDepositResponse.amountsA[i],
			"Assertion failed for response.amountsA[%d]", i,
		)
		s.Assert().Equal(
			depositResponse.amountsB[i],
			expectedDepositResponse.amountsB[i],
			"Assertion failed for response.amountsB[%d]", i,
		)
	}
}

type DepositReponse struct {
	amountsA []sdk.Int
	amountsB []sdk.Int
}

// Withdraw
type Withdrawl struct {
	TickIndex int64
	FeeIndex  uint64
	Shares    sdk.Int
}

func NewWithdrawlInt(shares sdk.Int, tick int64, feeIndex uint64) *Withdrawl {
	return &Withdrawl{
		Shares:    shares,
		FeeIndex:  feeIndex,
		TickIndex: tick,
	}
}

func NewWithdrawl(shares int64, tick int64, feeIndex uint64) *Withdrawl {
	return NewWithdrawlInt(sdk.NewInt(shares), tick, feeIndex)
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.alice, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.bob, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.carol, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdraws(withdrawals ...*Withdrawl) {
	s.withdraws(s.dan, withdrawals...)
}

func (s *MsgServerTestSuite) withdraws(account sdk.AccAddress, withdrawls ...*Withdrawl) {
	tickIndexes := make([]int64, len(withdrawls))
	feeIndexes := make([]uint64, len(withdrawls))
	sharesToRemove := make([]sdk.Int, len(withdrawls))
	for i, e := range withdrawls {
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		sharesToRemove[i] = e.Shares
	}

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        account.String(),
		Receiver:       account.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndexes,
		FeeIndexes:     feeIndexes,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawFails(expectedErr error, withdrawals ...*Withdrawl) {
	s.withdrawFails(s.alice, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdrawFails(expectedErr error, withdrawals ...*Withdrawl) {
	s.withdrawFails(s.bob, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdrawFails(expectedErr error, withdrawals ...*Withdrawl) {
	s.withdrawFails(s.carol, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdrawFails(expectedErr error, withdrawals ...*Withdrawl) {
	s.withdrawFails(s.dan, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) withdrawFails(account sdk.AccAddress, expectedErr error, withdrawals ...*Withdrawl) {
	tickIndexes := make([]int64, len(withdrawals))
	feeIndexes := make([]uint64, len(withdrawals))
	sharesToRemove := make([]sdk.Int, len(withdrawals))
	for i, e := range withdrawals {
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
		sharesToRemove[i] = e.Shares
	}

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        account.String(),
		Receiver:       account.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndexes,
		FeeIndexes:     feeIndexes,
	})
	s.Assert().NotNil(err)
	s.Assert().ErrorIs(err, expectedErr)
}

/// Cancel limit order

func (s *MsgServerTestSuite) aliceCancelsLimitSell(keyToken string, tick int, trancheKey string) {
	s.cancelsLimitSell(s.alice, keyToken, tick, trancheKey)
}

func (s *MsgServerTestSuite) bobCancelsLimitSell(keyToken string, tick int, trancheKey string) {
	s.cancelsLimitSell(s.bob, keyToken, tick, trancheKey)
}

func (s *MsgServerTestSuite) carolCancelsLimitSell(keyToken string, tick int, trancheKey string) {
	s.cancelsLimitSell(s.carol, keyToken, tick, trancheKey)
}

func (s *MsgServerTestSuite) danCancelsLimitSell(keyToken string, tick int, trancheKey string) {
	s.cancelsLimitSell(s.dan, keyToken, tick, trancheKey)
}

func (s *MsgServerTestSuite) cancelsLimitSell(account sdk.AccAddress, selling string, tick int, trancheKey string) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:    account.String(),
		TokenA:     "TokenA",
		TokenB:     "TokenB",
		TickIndex:  int64(tick),
		KeyToken:   selling,
		TrancheKey: trancheKey,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceCancelsLimitSellFails(keyToken string, tick int, trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.alice, keyToken, tick, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) bobCancelsLimitSellFails(keyToken string, tick int, trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.bob, keyToken, tick, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) carolCancelsLimitSellFails(keyToken string, tick int, trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.carol, keyToken, tick, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) danCancelsLimitSellFails(keyToken string, tick int, trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.dan, keyToken, tick, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) cancelsLimitSellFails(account sdk.AccAddress, selling string, tick int, trancheKey string, expectedErr error) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:    account.String(),
		TokenA:     "TokenA",
		TokenB:     "TokenB",
		TickIndex:  int64(tick),
		KeyToken:   selling,
		TrancheKey: trancheKey,
	})
	s.Assert().ErrorIs(err, expectedErr)
}

/// Swap

func (s *MsgServerTestSuite) aliceMarketSells(selling string, amountIn int, minOut int) {
	s.marketSells(s.alice, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) bobMarketSells(selling string, amountIn int, minOut int) {
	s.marketSells(s.bob, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) carolMarketSells(selling string, amountIn int, minOut int) {
	s.marketSells(s.carol, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) danMarketSells(selling string, amountIn int, minOut int) {
	s.marketSells(s.dan, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) marketSells(account sdk.AccAddress, selling string, amountIn int, minOut int) {
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  selling,
		AmountIn: sdk.NewInt(int64(amountIn)),
		MinOut:   sdk.NewInt(int64(minOut)),
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceMarketSellFails(err error, selling string, amountIn int, minOut int) {
	s.marketSellFails(s.alice, err, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) bobMarketSellFails(err error, selling string, amountIn int, minOut int) {
	s.marketSellFails(s.bob, err, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) carolMarketSellFails(err error, selling string, amountIn int, minOut int) {
	s.marketSellFails(s.bob, err, selling, amountIn, minOut)
}

func (s *MsgServerTestSuite) danMarketSellFails(err error, selling string, amountIn int, minOut int) {
	s.marketSellFails(s.bob, err, selling, amountIn, minOut)
}
func (s *MsgServerTestSuite) marketSellFails(account sdk.AccAddress, expectedErr error, selling string, amountIn int, minOut int) {
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  selling,
		AmountIn: sdk.NewInt(int64(amountIn)),
		MinOut:   sdk.NewInt(int64(minOut)),
	})
	s.Assert().ErrorIs(err, expectedErr)
}

/// Withdraw filled limit order

func (s *MsgServerTestSuite) aliceWithdrawsLimitSell(selling string, tick int, trancheKey string) {
	s.withdrawsLimitSell(s.alice, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) bobWithdrawsLimitSell(selling string, tick int, trancheKey string) {
	s.withdrawsLimitSell(s.bob, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) carolWithdrawsLimitSell(selling string, tick int, trancheKey string) {
	s.withdrawsLimitSell(s.carol, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) danWithdrawsLimitSell(selling string, tick int, trancheKey string) {
	s.withdrawsLimitSell(s.dan, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) withdrawsLimitSell(account sdk.AccAddress, selling string, tick int, trancheKey string) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:    account.String(),
		TokenA:     "TokenA",
		TokenB:     "TokenB",
		TickIndex:  int64(tick),
		KeyToken:   selling,
		TrancheKey: trancheKey,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawLimitSellFails(expectedErr error, selling string, tick int, trancheKey string) {
	s.withdrawLimitSellFails(s.alice, expectedErr, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) bobWithdrawLimitSellFails(expectedErr error, selling string, tick int, trancheKey string) {
	s.withdrawLimitSellFails(s.bob, expectedErr, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) carolWithdrawLimitSellFails(expectedErr error, selling string, tick int, trancheKey string) {
	s.withdrawLimitSellFails(s.carol, expectedErr, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) danWithdrawLimitSellFails(expectedErr error, selling string, tick int, trancheKey string) {
	s.withdrawLimitSellFails(s.dan, expectedErr, selling, tick, trancheKey)
}

func (s *MsgServerTestSuite) withdrawLimitSellFails(account sdk.AccAddress, expectedErr error, selling string, tick int, trancheKey string) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:    account.String(),
		TokenA:     "TokenA",
		TokenB:     "TokenB",
		TickIndex:  int64(tick),
		KeyToken:   selling,
		TrancheKey: trancheKey,
	})
	s.Assert().ErrorIs(err, expectedErr)
}

// Shares
func (s *MsgServerTestSuite) getPoolShares(
	token0 string,
	token1 string,
	tick int64,
	feeIndex uint64,
) (shares sdk.Int) {
	sharesId := CreateSharesId(token0, token1, tick, feeIndex)
	return s.app.BankKeeper.GetSupply(s.ctx, sharesId).Amount
}

func (s *MsgServerTestSuite) assertPoolShares(
	tick int64,
	feeIndex uint64,
	sharesExpected uint64,
) {
	_, found := s.app.DexKeeper.GetFeeTier(s.ctx, feeIndex)
	if !found {
		s.Require().Fail("Invalid fee index given")
	}
	sharesExpectedInt := sdk.NewIntFromUint64(sharesExpected)
	sharesOwned := s.getPoolShares("TokenA", "TokenB", tick, feeIndex)
	s.Assert().Equal(sharesExpectedInt, sharesOwned)
}

func (s *MsgServerTestSuite) getAccountShares(
	account sdk.AccAddress,
	token0 string,
	token1 string,
	tick int64,
	feeIndex uint64,
) (shares sdk.Int) {
	_, found := s.app.DexKeeper.GetFeeTier(s.ctx, feeIndex)
	if !found {
		s.Require().Fail("Invalid fee index given")
		return sdk.ZeroInt()
	}
	sharesId := CreateSharesId(token0, token1, tick, feeIndex)
	return s.app.BankKeeper.GetBalance(s.ctx, account, sharesId).Amount
}

func (s *MsgServerTestSuite) assertAccountShares(
	account sdk.AccAddress,
	tick int64,
	feeIndex uint64,
	sharesExpected uint64,
) {
	sharesExpectedInt := sdk.NewIntFromUint64(sharesExpected)
	sharesOwned := s.getAccountShares(account, "TokenA", "TokenB", tick, feeIndex)
	s.Assert().Equal(sharesExpectedInt, sharesOwned, "expected %s != actual %s", sharesExpected, sharesOwned)
}

func (s *MsgServerTestSuite) assertAliceShares(tick int64, feeIndex uint64, sharesExpected uint64) {
	s.assertAccountShares(s.alice, tick, feeIndex, sharesExpected)
}
func (s *MsgServerTestSuite) assertBobShares(tick int64, fee uint64, sharesExpected uint64) {
	s.assertAccountShares(s.bob, tick, fee, sharesExpected)
}
func (s *MsgServerTestSuite) assertCarolShares(tick int64, fee uint64, sharesExpected uint64) {
	s.assertAccountShares(s.carol, tick, fee, sharesExpected)
}
func (s *MsgServerTestSuite) assertDanShares(tick int64, fee uint64, sharesExpected uint64) {
	s.assertAccountShares(s.dan, tick, fee, sharesExpected)
}

// Ticks
func (s *MsgServerTestSuite) assertCurrentTicks(
	expected1To0 int64,
	expected0To1 int64,
) {
	s.assertCurr0To1(expected0To1)
	s.assertCurr1To0(expected1To0)
}

func (s *MsgServerTestSuite) assertCurr0To1(curr0To1Expected int64) {
	pairId := CreatePairId("TokenA", "TokenB")
	curr0To1Actual, _ := s.app.DexKeeper.GetCurrTick0To1(s.ctx, pairId)
	s.Assert().Equal(curr0To1Expected, curr0To1Actual)
}

func (s *MsgServerTestSuite) assertCurr1To0(curr1To0Expected int64) {
	pairId := CreatePairId("TokenA", "TokenB")

	curr1to0Actual, _ := s.app.DexKeeper.GetCurrTick1To0(s.ctx, pairId)
	s.Assert().Equal(curr1To0Expected, curr1to0Actual)
}

// Pool liquidity (i.e. deposited rather than LO)
func (s *MsgServerTestSuite) assertLiquidityAtTick(amountA sdk.Int, amountB sdk.Int, tickIndex int64, feeIndex uint64) {

	liquidityA, liquidityB := s.getLiquidityAtTick(tickIndex, feeIndex)
	s.Assert().True(amountA.Equal(liquidityA), "liquidity A: actual %s, expected %s", liquidityA, amountA)
	s.Assert().True(amountB.Equal(liquidityB), "liquidity B: actual %s, expected %s", liquidityB, amountB)
}

func (s *MsgServerTestSuite) assertPoolLiquidity(amountA int, amountB int, tickIndex int64, feeIndex uint64) {
	s.assertLiquidityAtTick(sdk.NewInt(int64(amountA)), sdk.NewInt(int64(amountB)), tickIndex, feeIndex)
}

func (s *MsgServerTestSuite) assertNoLiquidityAtTick(tickIndex int64, feeIndex uint64) {
	s.assertLiquidityAtTick(sdk.ZeroInt(), sdk.ZeroInt(), tickIndex, feeIndex)
}

// Filled limit liquidity
func (s *MsgServerTestSuite) assertAliceLimitFilledAtTickAtIndex(selling string, amount int, tickIndex int64, trancheKey string) {
	s.assertLimitFilledAtTickAtIndex(s.alice, selling, amount, tickIndex, trancheKey)
}

func (s *MsgServerTestSuite) assertBobLimitFilledAtTickAtIndex(selling string, amount int, tickIndex int64, trancheKey string) {
	s.assertLimitFilledAtTickAtIndex(s.bob, selling, amount, tickIndex, trancheKey)
}

func (s *MsgServerTestSuite) assertCarolLimitFilledAtTickAtIndex(selling string, amount int, tickIndex int64, trancheKey string) {
	s.assertLimitFilledAtTickAtIndex(s.carol, selling, amount, tickIndex, trancheKey)
}

func (s *MsgServerTestSuite) assertDanLimitFilledAtTickAtIndex(selling string, amount int, tickIndex int64, trancheKey string) {
	s.assertLimitFilledAtTickAtIndex(s.dan, selling, amount, tickIndex, trancheKey)
}

func (s *MsgServerTestSuite) assertLimitFilledAtTickAtIndex(account sdk.AccAddress, selling string, amount int, tickIndex int64, trancheKey string) {
	filled := s.getLimitFilledLiquidityAtTickAtIndex(selling, tickIndex, trancheKey)
	amt := sdk.NewInt(int64(amount))
	s.Assert().True(amt.Equal(filled))
}

// Limit liquidity
func (s *MsgServerTestSuite) assertAliceLimitLiquidityAtTick(selling string, amount int, tickIndex int64) {
	s.assertAccountLimitLiquidityAtTick(s.alice, selling, amount, tickIndex)
}

func (s *MsgServerTestSuite) assertBobLimitLiquidityAtTick(selling string, amount int, tickIndex int64) {
	s.assertAccountLimitLiquidityAtTick(s.bob, selling, amount, tickIndex)
}

func (s *MsgServerTestSuite) assertCarolLimitLiquidityAtTick(selling string, amount int, tickIndex int64) {
	s.assertAccountLimitLiquidityAtTick(s.carol, selling, amount, tickIndex)
}

func (s *MsgServerTestSuite) assertDanLimitLiquidityAtTick(selling string, amount int, tickIndex int64) {
	s.assertAccountLimitLiquidityAtTick(s.dan, selling, amount, tickIndex)
}

func (s *MsgServerTestSuite) assertAccountLimitLiquidityAtTick(account sdk.AccAddress, selling string, amount int, tickIndex int64) {

	userShares, totalShares := s.getLimitUserSharesAtTick(account, selling, tickIndex), s.getLimitTotalSharesAtTick(selling, tickIndex)
	userRatio := userShares.ToDec().QuoInt(totalShares)
	userLiquidity := userRatio.MulInt64(int64(amount)).TruncateInt()

	s.assertLimitLiquidityAtTick(selling, tickIndex, userLiquidity.Int64())
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTick(selling string, tickIndex int64, amount int64) {
	s.assertLimitLiquidityAtTickInt(selling, tickIndex, sdk.NewInt(amount))
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTickInt(selling string, tickIndex int64, amount sdk.Int) {

	pairId := CreatePairId("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairId, selling, tickIndex)
	liquidity := sdk.ZeroInt()
	for _, t := range tranches {
		liquidity = liquidity.Add(t.ReservesTokenIn)
	}

	s.Assert().True(amount.Equal(liquidity), "Incorrect liquidity: expected %s, have %s", amount.String(), liquidity.String())
}

func (s *MsgServerTestSuite) assertFillAndPlaceTrancheKeys(selling string, tickIndex int64, expectedFill string, expectedPlace string) {
	pairId := CreatePairId("TokenA", "TokenB")
	placeTranche, foundPlace := s.app.DexKeeper.GetPlaceTranche(s.ctx, pairId, selling, tickIndex)
	fillTranche, foundFill := s.app.DexKeeper.GetFillTranche(s.ctx, pairId, selling, tickIndex)
	placeKey, fillKey := "", ""
	if foundPlace {
		placeKey = placeTranche.TrancheKey
	}

	if foundFill {
		fillKey = fillTranche.TrancheKey
	}
	s.Assert().Equal(expectedFill, fillKey)
	s.Assert().Equal(expectedPlace, placeKey)
}

// Limit order map helpers
func (s *MsgServerTestSuite) getLimitUserSharesAtTick(account sdk.AccAddress, selling string, tickIndex int64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairId, selling, tickIndex)
	fillTranche := tranches[0]
	// get user shares and total shares
	userShares := s.getLimitUserSharesAtTickAtIndex(account, selling, tickIndex, fillTranche.TrancheKey)
	if len(tranches) >= 2 {
		userShares = userShares.Add(s.getLimitUserSharesAtTickAtIndex(account, selling, tickIndex, tranches[1].TrancheKey))
	}
	return userShares
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTickAtIndex(account sdk.AccAddress, selling string, tickIndex int64, trancheKey string) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	userShares, userSharesFound := s.app.DexKeeper.GetLimitOrderTrancheUser(s.ctx, pairId, tickIndex, selling, trancheKey, account.String())
	s.Assert().True(userSharesFound, "Failed to get limit order user shares for index %s", trancheKey)
	return userShares.SharesOwned
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTick(selling string, tickIndex int64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairId, selling, tickIndex)
	// get user shares and total shares
	totalShares := sdk.ZeroInt()
	for _, t := range tranches {
		totalShares = totalShares.Add(t.TotalTokenIn)
	}
	return totalShares
}

func (s *MsgServerTestSuite) getLimitFilledLiquidityAtTickAtIndex(selling string, tickIndex int64, trancheKey string) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, _, found := s.app.DexKeeper.FindLimitOrderTranche(s.ctx, pairId, tickIndex, selling, trancheKey)
	s.Assert().True(found, "Failed to get limit order filled reserves for index %s", trancheKey)
	return tranche.ReservesTokenOut
}

func (s *MsgServerTestSuite) getLimitReservesAtTickAtKey(selling string, tickIndex int64, trancheKey string) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, _, found := s.app.DexKeeper.FindLimitOrderTranche(s.ctx, pairId, tickIndex, selling, trancheKey)
	s.Assert().True(found, "Failed to get limit order reserves for index %s", trancheKey)
	return tranche.ReservesTokenIn
}

// SingleLimitOrderFill() simulates the fill of a single limit order and returns the amount
// swapped into it, filling some of it (amount_in) and the amount swapped out (amount_out). It
// takes as input the amount that was placed for the limit order (amount_placed), the price the
// trader pays when filling it (price_filled_at) and the amount that they are swapping (amount_to_swap).
// The format of the return statement is (amount_in, amount_out).
func SingleLimitOrderFill(amount_placed sdk.Int,
	price_filled_at sdk.Dec,
	amount_to_swap sdk.Int) (sdk.Dec, sdk.Dec) {
	amount_out, amount_in := sdk.ZeroDec(), sdk.ZeroDec()
	amountPlacedDec := amount_placed.ToDec()
	amountPlacedForPrice := amountPlacedDec.Quo(price_filled_at)
	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amount_to_swap.ToDec().GT(amountPlacedForPrice) {
		amount_out = amount_placed.ToDec()
		amount_in = amountPlacedForPrice
	} else {
		amount_in = amount_to_swap.ToDec()
		amount_out = amount_in.Mul(price_filled_at)
	}

	return amount_in, amount_out
}

// Calls SingleLimitOrderFill() and updates the filled and unfilled reserves.
// Returns the unfilled reserves (unfilled_reserves), filled reserves (filled_reserves) and the amount left to swap
// (amount_to_swap_remaining)
func SingleLimitOrderFillAndUpdate(amount_placed sdk.Int,
	price_filled_at sdk.Dec,
	amount_to_swap sdk.Int,
	unfilled_reserves sdk.Int) (sdk.Dec, sdk.Dec, sdk.Dec) {
	amount_in, amount_out := SingleLimitOrderFill(amount_placed, price_filled_at, amount_to_swap)
	unfilled_reservesDec := unfilled_reserves.ToDec().Sub(amount_out)
	filled_reserves := amount_placed.ToDec().Add(amount_in)
	amount_to_swap_remaining := amount_to_swap.ToDec().Sub(amount_in)
	return unfilled_reservesDec, filled_reserves, amount_to_swap_remaining
}

// MultipleLimitOrderFills() simulates the fill of multiple consecutive limit orders and returns the
// total amount filled. It takes as input the amounts that were placed for the limit
// order (amount_placed), the pricesthe trader pays when filling the orders (price_filled_at)
// and the amount that they are swapping (amount_to_swap).
func MultipleLimitOrderFills(amounts_placed []sdk.Int, prices []sdk.Dec, amount_to_swap sdk.Int) sdk.Dec {
	total_out, amount_remaining := sdk.ZeroDec(), amount_to_swap

	// Loops through all of the limit orders that need to be filled
	for i := 0; i < len(amounts_placed); i++ {
		_, amount_out := SingleLimitOrderFill(amounts_placed[i], prices[i], amount_remaining)

		// amount_remaining = amount_remaining.Sub(amount_in)
		total_out = total_out.Add(amount_out)
	}
	return total_out
}

// SinglePoolSwap() simulates swapping through a single liquidity pool and returns the amount
// swapped into it (amount_in) and the amount swapped out, received by the swapper (amount_out). It
// takes as input the amount of liquidity in the pool (amount_liquidity), the price the
// trader pays when swapping through it (price_swapped_at) and the amount that they are
// swapping (amount_to_swap). The format of the return statement is (amount_in, amount_out).
// Same thing as SingleLimitOrderFill() except in naming.
func SinglePoolSwap(amount_liquidity sdk.Int, price_swapped_at sdk.Dec, amount_to_swap sdk.Int) (sdk.Dec, sdk.Dec) {
	amount_out, amount_in := sdk.ZeroDec(), sdk.ZeroDec()
	liquidityAtPrice := amount_liquidity.ToDec().Quo(price_swapped_at)
	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amount_to_swap.ToDec().GT(liquidityAtPrice) {
		amount_out = amount_liquidity.ToDec()
		amount_in = liquidityAtPrice
	} else {
		amount_in = amount_to_swap.ToDec()
		amount_out = amount_in.Mul(price_swapped_at)
	}
	return amount_in, amount_out
}

// SinglePoolSwapAndUpdate() simulates swapping through a single liquidity pool and updates that pool's
// liquidity. Takes in all of the same inputs as SinglePoolSwap(): amount_liquidity, price_swapped_at,
// and amount_to_swap; but has additional inputs, reservesOfInToken, reservesOfOutToken. It returns the
// updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format of
// (resulting_reserves_in_token, resulting_reserves_out_token, amount_in, amount_out)
func SinglePoolSwapAndUpdate(amount_liquidity sdk.Int,
	price_swapped_at sdk.Dec,
	amount_to_swap sdk.Int,
	reservesOfInToken sdk.Int,
	reservesOfOutToken sdk.Int) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec) {
	amount_in, amount_out := SinglePoolSwap(amount_liquidity, price_swapped_at, amount_to_swap)
	resulting_reserves_in_token := reservesOfInToken.ToDec().Add(amount_in)
	resulting_reserves_out_token := reservesOfOutToken.ToDec().Add(amount_out)
	return resulting_reserves_in_token, resulting_reserves_out_token, amount_in, amount_out
}

// SinglePoolSwapAndUpdateDirection() simulates swapping through a single liquidity pool and updates that pool's
// liquidity and specifies whether the in and out tokens are 0 or 1. Takes in all of the same inputs as
// SinglePoolSwapAndUpdate(): amount_liquidity, price_swapped_at, amount_to_swap, reservesOfToken0 sdk.Int,
// reservesOfToken1 but has an additional input inToken which is a bool indicating whether 0 or 1 is swapped into
// the pool. It returns the updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format
// of (reservesOfInToken,reservesOfOutToken).
func SinglePoolSwapAndUpdateDirectional(amount_liquidity sdk.Int,
	price_swapped_at sdk.Dec,
	amount_to_swap sdk.Int,
	reservesOfToken0 sdk.Int,
	reservesOfToken1 sdk.Int,
	inToken bool) (sdk.Dec, sdk.Dec) {
	resultingReservesOfToken0, resultingReservesOfToken1 := sdk.ZeroDec(), sdk.ZeroDec()
	if inToken {
		resultingReservesOfToken1, resultingReservesOfToken0, _, _ = SinglePoolSwapAndUpdate(amount_liquidity,
			price_swapped_at,
			amount_to_swap,
			reservesOfToken1,
			reservesOfToken0)
	} else {
		resultingReservesOfToken0, resultingReservesOfToken1, _, _ = SinglePoolSwapAndUpdate(amount_liquidity,
			price_swapped_at,
			amount_to_swap,
			reservesOfToken0,
			reservesOfToken1)
	}
	return resultingReservesOfToken0, resultingReservesOfToken1
}

// MultiplePoolSwapAndUpdate() simulates swapping through multiple liquidity pools and updates that pool's
// liquidity. Takes in similar inputs to SinglePoolSwapAndUpdate(): amount_liquidity, price_swapped_at,
// and amount_to_swap, reservesOfInToken, reservesOfOutToken; But they are held in arrays the size of how many
// pools are being swapped through. It returns the updated amounts for the reservesOfInToken and the
// reservesOfOutToken, in the format of (reservesOfInToken,reservesOfOutToken)
func MultiplePoolSwapAndUpdate(amounts_liquidity []sdk.Int,
	prices_swapped_at []sdk.Dec,
	amount_to_swap sdk.Int,
	reserves_in_token_array []sdk.Int,
	reserves_out_token_array []sdk.Int) ([]sdk.Dec, []sdk.Dec, sdk.Dec, sdk.Dec) {
	num_pools := len(amounts_liquidity)
	amountRemainingDec := amount_to_swap.ToDec()
	amount_out_total, amount_out_temp, amount_in := sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
	resulting_reserves_in_token := make([]sdk.Dec, num_pools, num_pools)
	resulting_reserves_out_token := make([]sdk.Dec, num_pools, num_pools)
	for i := 0; i < num_pools; i++ {
		resulting_reserves_in_token[i], resulting_reserves_out_token[i], amount_in, amount_out_temp = SinglePoolSwapAndUpdate(amounts_liquidity[i],
			prices_swapped_at[i],
			amount_to_swap,
			reserves_in_token_array[i],
			reserves_out_token_array[i])
		amount_out_total = amount_out_total.Add(amount_out_temp)
		amountRemainingDec = amountRemainingDec.Sub(amount_in)
		i++
	}

	return resulting_reserves_in_token, resulting_reserves_out_token, amountRemainingDec, amount_out_total
}

func SharesOnDeposit(existing_shares sdk.Dec, existing_amount0 sdk.Int, existing_amount1 sdk.Int, new_amount0 sdk.Int, new_amount1 sdk.Int, tickIndex int64) (shares_minted sdk.Int) {
	price1To0 := types.MustNewPrice(-1 * tickIndex)
	newAmount0Dec := sdk.NewDecFromInt(new_amount0)
	new_value := newAmount0Dec.Add(price1To0.MulInt(new_amount1))

	if existing_amount0.Add(existing_amount1).GT(sdk.ZeroInt()) {
		existing_value := existing_amount0.ToDec().Add(price1To0.MulInt(existing_amount1))
		shares_minted = shares_minted.ToDec().Mul(new_value.Quo(existing_value)).TruncateInt()
	} else {
		shares_minted = new_value.TruncateInt()
	}

	return shares_minted
}

func (s *MsgServerTestSuite) calcAutoswapSharesMinted(centerTick int64, feeIndex uint64, _residual0 int64, _residual1 int64, _balanced0 int64, _balanced1 int64, _totalShares int64, _valuePool int64) sdk.Int {
	residual0, residual1, balanced0, balanced1, totalShares, valuePool := sdk.NewInt(_residual0), sdk.NewInt(_residual1), sdk.NewInt(_balanced0), sdk.NewInt(_balanced1), sdk.NewInt(_totalShares), sdk.NewInt(_valuePool)

	// residualValue = 1.0001^-f * residualAmount0 + 1.0001^{i-f} * residualAmount1
	// balancedValue = balancedAmount0 + 1.0001^{i} * balancedAmount1
	// value = residualValue + balancedValue
	// shares minted = value * totalShares / valuePool
	fee := s.feeTiers[feeIndex].Fee

	centerPrice := types.MustNewPrice(-1 * centerTick)
	leftPrice := types.MustNewPrice(-1 * (centerTick - int64(fee)))
	discountPrice := types.MustNewPrice(-1 * int64(fee))

	balancedValue := balanced0.ToDec().Add(centerPrice.MulInt(balanced1)).TruncateInt()
	residualValue := discountPrice.MulInt(residual0).Add(leftPrice.Mul(residual1.ToDec())).TruncateInt()
	valueMint := balancedValue.Add(residualValue)

	return valueMint.Mul(totalShares).Quo(valuePool)
}

func (s *MsgServerTestSuite) calcSharesMinted(centerTick int64, feeIndex uint64, _amount0 int64, _amount1 int64) sdk.Int {
	amount0, amount1 := sdk.NewInt(_amount0), sdk.NewInt(_amount1)
	centerPrice := types.MustNewPrice(-1 * centerTick)

	return amount0.ToDec().Add(centerPrice.Mul(amount1.ToDec())).TruncateInt()
}

func (s *MsgServerTestSuite) calcExpectedBalancesAfterWithdrawOnePool(sharesMinted sdk.Int, account sdk.AccAddress, tickIndex int64, feeIndex uint64) (sdk.Int, sdk.Int, sdk.Int, sdk.Int) {
	dexCurrentBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenA").Amount
	dexCurrentBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenB").Amount
	currentBalance0 := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount
	currentBalance1 := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount
	amountPool0, amountPool1 := s.getLiquidityAtTick(tickIndex, feeIndex)
	poolShares := s.getPoolShares("TokenA", "TokenB", tickIndex, feeIndex)

	amountOut0 := amountPool0.Mul(sharesMinted).Quo(poolShares)
	amountOut1 := amountPool1.Mul(sharesMinted).Quo(poolShares)

	expectedBalance0 := currentBalance0.Add(amountOut0)
	expectedBalance1 := currentBalance1.Add(amountOut1)
	dexExpectedBalance0 := dexCurrentBalance0.Sub(amountOut0)
	dexExpectedBalance1 := dexCurrentBalance1.Sub(amountOut1)

	return expectedBalance0, expectedBalance1, dexExpectedBalance0, dexExpectedBalance1
}

// Swap helpers (use for writing the tests, but replace with actual values before finishing!)
func (s *MsgServerTestSuite) calculateSingleSwapNoLOAToB(tick int64, tickLiqudity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOAToB(tick int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapAToB(tick int64, tickLiqudidty int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOBToA(tick int64, tickLiqudity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOBToA(tick int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapBToA(tick int64, tickLiqudidty int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwapNoLO(price *types.Price, tickLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, tickLiquidity, 0, amountIn)
}

func calculateSingleSwapOnlyLO(price *types.Price, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, 0, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwap(price *types.Price, tickLiquidity int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	// swap against CSMM liquidity
	amountLeft, amountOut := calculateSwap(price, tickLiquidity, amountIn)
	// fmt.Printf("left %s out %s\n", amountLeft, amountOut)

	// swap against limit orders
	if amountLeft.GT(sdk.ZeroInt()) {
		tmpAmountLeft, tmpAmountOut := calculateSwap(price, tickLimitOrderLiquidity, amountLeft.Int64())
		amountLeft = tmpAmountLeft
		amountOut = amountOut.Add(tmpAmountOut)
	}
	return amountLeft, amountOut
}

func calculateSwap(price *types.Price, liquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	amountInInt := sdk.NewInt(amountIn)
	liquidityInt := sdk.NewInt(liquidity)
	if tmpAmountOut := price.MulInt(amountInInt); tmpAmountOut.LT(liquidityInt.ToDec()) {
		// fmt.Printf("sufficient tmpOut %s\n", tmpAmountOut)
		// sufficient liquidity
		return sdk.ZeroInt(), tmpAmountOut.TruncateInt()
	} else {
		// only sufficient for part of amountIn
		tmpAmountIn := price.Inv().MulInt(liquidityInt).TruncateInt()
		// fmt.Printf("insufficient tmpIn %s\n", tmpAmountIn)
		return amountInInt.Sub(tmpAmountIn), liquidityInt
	}
}

func (s *MsgServerTestSuite) calculateMultipleSwapsAToB(
	tickIndexes []int64,
	tickLiquidities []int64,
	tickLimitOrderLiquidities []int64,
	amountIn int64,
) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i] = types.MustNewPrice(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOAToB(tickIndexes []int64, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(tickIndexes[i])
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOAToB(tickIndexes []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(tickIndexes[i])
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsBToA(tickIndexes []int64, tickLiquidities []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(-1 * tickIndexes[i])
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOBToA(tickIndexes []int64, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(-1 * tickIndexes[i])
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOBToA(tickIndexes []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(-1 * tickIndexes[i])
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLO(prices []*types.Price, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	// zero array for tickLimitOrders
	tickLimitOrderLiquidities := make([]int64, len(prices))
	for i := range tickLimitOrderLiquidities {
		tickLimitOrderLiquidities[i] = 0
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLO(prices []*types.Price, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	// zero array for tickLimitOrders
	tickLiquidities := make([]int64, len(prices))
	for i := range tickLiquidities {
		tickLiquidities[i] = 0
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwaps(prices []*types.Price, tickLiquidities []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	amountLeft, amountOut := sdk.NewInt(amountIn), sdk.ZeroInt()
	for i := 0; i < len(prices); i++ {
		tmpAmountLeft, tmpAmountOut := calculateSingleSwap(prices[i], tickLiquidities[i], tickLimitOrderLiquidities[i], amountLeft.Int64())
		amountLeft, amountOut = tmpAmountLeft, amountOut.Add(tmpAmountOut)
	}
	return amountLeft, amountOut
}
