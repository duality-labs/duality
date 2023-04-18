package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dualityapp "github.com/duality-labs/duality/app"
	. "github.com/duality-labs/duality/x/dex/keeper"
	. "github.com/duality-labs/duality/x/dex/keeper/internal/testutils"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
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
}

var defaultPairID *types.PairID = &types.PairID{Token0: "TokenA", Token1: "TokenB"}

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

	s.app = app
	s.msgServer = NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
}

/// Fund accounts

func (s *MsgServerTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance, bBalance int64) {
	aBalanceInt := sdk.NewInt(aBalance)
	bBalanceInt := sdk.NewInt(bBalance)
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalances(account, aBalance, bBalance)
}

func (s *MsgServerTestSuite) fundAccountBalancesWithDenom(addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, amounts); err != nil {
		return err
	}

	return s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, amounts)
}

func (s *MsgServerTestSuite) fundAliceBalances(a, b int64) {
	s.fundAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundBobBalances(a, b int64) {
	s.fundAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalances(a, b int64) {
	s.fundAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundDanBalances(a, b int64) {
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

func (s *MsgServerTestSuite) assertAccountBalanceWithDenom(account sdk.AccAddress, denom string, expBalance int64) {
	actualBalance := s.app.BankKeeper.GetBalance(s.ctx, account, denom).Amount
	expBalanceInt := sdk.NewInt(expBalance)
	s.Assert().True(expBalanceInt.Equal(actualBalance), "expected %s != actual %s", expBalance, actualBalance)
}

func (s *MsgServerTestSuite) assertAliceBalances(a, b int64) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a, b int64) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalances(a, b int64) {
	s.assertAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a, b int64) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a, b int64) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalanceWithDenom(denom string, expectedAmount int64) {
	s.assertAccountBalanceWithDenom(s.app.AccountKeeper.GetModuleAddress("dex"), denom, expectedAmount)
}

func (s *MsgServerTestSuite) assertDexBalancesInt(a, b sdk.Int) {
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
	s.T().Logf(
		"Alice: %+v %+v\nBob: %+v %+v\nCarol: %+v %+v\nDan: %+v %+v",
		aliceA, aliceB,
		bobA, bobB,
		carolA, carolB,
		danA, danB,
	)
}

/// Place limit order

func (s *MsgServerTestSuite) aliceLimitSells(selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) string {
	return s.limitSellsSuccess(s.alice, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) bobLimitSells(selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) string {
	return s.limitSellsSuccess(s.bob, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) carolLimitSells(selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) string {
	return s.limitSellsSuccess(s.carol, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) danLimitSells(selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) string {
	return s.limitSellsSuccess(s.dan, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) limitSellsSuccess(account sdk.AccAddress, tokenIn string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) string {
	trancheKey, err := s.limitSells(account, tokenIn, tick, amountIn, orderTypeOpt...)
	s.Assert().Nil(err)
	return trancheKey
}

func (s *MsgServerTestSuite) aliceLimitSellsGoodTil(selling string, tick, amountIn int, goodTil time.Time) string {
	return s.limitSellsGoodTil(s.alice, selling, tick, amountIn, goodTil)
}

func (s *MsgServerTestSuite) bobLimitSellsGoodTil(selling string, tick, amountIn int, goodTil time.Time) string {
	return s.limitSellsGoodTil(s.bob, selling, tick, amountIn, goodTil)
}

func (s *MsgServerTestSuite) carolLimitSellsGoodTil(selling string, tick, amountIn int, goodTil time.Time) string {
	return s.limitSellsGoodTil(s.carol, selling, tick, amountIn, goodTil)
}

func (s *MsgServerTestSuite) danLimitSellsGoodTil(selling string, tick, amountIn int, goodTil time.Time) string {
	return s.limitSellsGoodTil(s.dan, selling, tick, amountIn, goodTil)
}

func (s *MsgServerTestSuite) assertAliceLimitSellFails(err error, selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) {
	s.assertLimitSellFails(s.alice, err, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) assertBobLimitSellFails(err error, selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) {
	s.assertLimitSellFails(s.bob, err, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) assertCarolLimitSellFails(err error, selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) {
	s.assertLimitSellFails(s.carol, err, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) assertDanLimitSellFails(err error, selling string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) {
	s.assertLimitSellFails(s.dan, err, selling, tick, amountIn, orderTypeOpt...)
}

func (s *MsgServerTestSuite) assertLimitSellFails(account sdk.AccAddress, expectedErr error, tokenIn string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) {
	_, err := s.limitSells(account, tokenIn, tick, amountIn, orderTypeOpt...)
	s.Assert().ErrorIs(err, expectedErr)
}

func (s *MsgServerTestSuite) limitSells(account sdk.AccAddress, tokenIn string, tick, amountIn int, orderTypeOpt ...types.LimitOrderType) (string, error) {
	var orderType types.LimitOrderType
	if len(orderTypeOpt) == 0 {
		orderType = types.LimitOrderType_GOOD_TIL_CANCELLED
	} else {
		orderType = orderTypeOpt[0]
	}
	tokenIn, tokenOut := GetInOutTokens(tokenIn, "TokenA", "TokenB")

	msg, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		TickIndex: int64(tick),
		AmountIn:  sdk.NewInt(int64(amountIn)),
		OrderType: orderType,
	})

	return msg.TrancheKey, err
}

func (s *MsgServerTestSuite) limitSellsGoodTil(account sdk.AccAddress, tokenIn string, tick, amountIn int, goodTil time.Time) string {
	tokenIn, tokenOut := GetInOutTokens(tokenIn, "TokenA", "TokenB")

	msg, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:        account.String(),
		Receiver:       account.String(),
		TokenIn:        tokenIn,
		TokenOut:       tokenOut,
		TickIndex:      int64(tick),
		AmountIn:       sdk.NewInt(int64(amountIn)),
		OrderType:      types.LimitOrderType_GOOD_TIL_TIME,
		ExpirationTime: &goodTil,
	})

	s.Assert().NoError(err)

	return msg.TrancheKey
}

// / Deposit
type Deposit struct {
	AmountA   sdk.Int
	AmountB   sdk.Int
	TickIndex int64
	Fee       uint64
}

type DepositOptions struct {
	Autoswap bool
}

type DepositWithOptions struct {
	AmountA   sdk.Int
	AmountB   sdk.Int
	TickIndex int64
	Fee       uint64
	Options   DepositOptions
}

func NewDeposit(amountA, amountB, tickIndex, fee int) *Deposit {
	return &Deposit{
		AmountA:   sdk.NewInt(int64(amountA)),
		AmountB:   sdk.NewInt(int64(amountB)),
		TickIndex: int64(tickIndex),
		Fee:       uint64(fee),
	}
}

func NewDepositWithOptions(amountA, amountB, tickIndex, fee int, options DepositOptions) *DepositWithOptions {
	return &DepositWithOptions{
		AmountA:   sdk.NewInt(int64(amountA)),
		AmountB:   sdk.NewInt(int64(amountB)),
		TickIndex: int64(tickIndex),
		Fee:       uint64(fee),
		Options:   options,
	}
}

func (s *MsgServerTestSuite) aliceDeposits(deposits ...*Deposit) {
	s.deposits(s.alice, deposits)
}

func (s *MsgServerTestSuite) aliceDepositsWithOptions(deposits ...*DepositWithOptions) {
	s.depositsWithOptions(s.alice, deposits...)
}

func (s *MsgServerTestSuite) bobDeposits(deposits ...*Deposit) {
	s.deposits(s.bob, deposits)
}

func (s *MsgServerTestSuite) carolDeposits(deposits ...*Deposit) {
	s.deposits(s.carol, deposits)
}

func (s *MsgServerTestSuite) danDeposits(deposits ...*Deposit) {
	s.deposits(s.dan, deposits)
}

func (s *MsgServerTestSuite) deposits(account sdk.AccAddress, deposits []*Deposit, pairID ...types.PairID) {
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	fees := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		fees[i] = e.Fee
		options[i] = &types.DepositOptions{Autoswap: false}
	}

	var tokenA, tokenB string
	switch {
	case len(pairID) == 0:
		tokenA = "TokenA"
		tokenB = "TokenB"
	case len(pairID) == 1:
		tokenA = pairID[0].Token0
		tokenB = pairID[0].Token1
	case len(pairID) > 1:
		s.Assert().Fail("Only 1 pairID can be provided")
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:         account.String(),
		Receiver:        account.String(),
		TokenA:          tokenA,
		TokenB:          tokenB,
		AmountsA:        amountsA,
		AmountsB:        amountsB,
		TickIndexesAToB: tickIndexes,
		Fees:            fees,
		Options:         options,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) depositsWithOptions(account sdk.AccAddress, deposits ...*DepositWithOptions) {
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	fees := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		fees[i] = e.Fee
		options[i] = &types.DepositOptions{
			Autoswap: e.Options.Autoswap,
		}
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:         account.String(),
		Receiver:        account.String(),
		TokenA:          "TokenA",
		TokenB:          "TokenB",
		AmountsA:        amountsA,
		AmountsB:        amountsB,
		TickIndexesAToB: tickIndexes,
		Fees:            fees,
		Options:         options,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) getLiquidityAtTick(tickIndex int64, fee uint64) (sdk.Int, sdk.Int) {
	pairID := CreatePairID("TokenA", "TokenB")
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, pairID, tickIndex, fee)
	s.Assert().NoError(err)

	liquidityA := pool.LowerTick0.Reserves
	liquidityB := pool.UpperTick1.Reserves

	return liquidityA, liquidityB
}

func (s *MsgServerTestSuite) getLiquidityAtTickWithDenom(pairID *types.PairID, tickIndex int64, fee uint64) (sdk.Int, sdk.Int) {
	pool, err := s.app.DexKeeper.GetOrInitPool(s.ctx, pairID, tickIndex, fee)
	s.Assert().NoError(err)

	liquidityA := pool.LowerTick0.Reserves
	liquidityB := pool.UpperTick1.Reserves

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
	fees := make([]uint64, len(deposits))
	options := make([]*types.DepositOptions, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		fees[i] = e.Fee
		options[i] = &types.DepositOptions{Autoswap: false}
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:         account.String(),
		Receiver:        account.String(),
		TokenA:          "TokenA",
		TokenB:          "TokenB",
		AmountsA:        amountsA,
		AmountsB:        amountsB,
		TickIndexesAToB: tickIndexes,
		Fees:            fees,
		Options:         options,
	})
	s.Assert().NotNil(err)
	s.Assert().ErrorIs(err, expectedErr)
}

func (s *MsgServerTestSuite) assertDepositReponse(depositResponse, expectedDepositResponse DepositReponse) {
	for i := range expectedDepositResponse.amountsA {
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
type Withdrawal struct {
	TickIndex int64
	Fee       uint64
	Shares    sdk.Int
}

func NewWithdrawalInt(shares sdk.Int, tick int64, fee uint64) *Withdrawal {
	return &Withdrawal{
		Shares:    shares,
		Fee:       fee,
		TickIndex: tick,
	}
}

func NewWithdrawal(shares, tick int64, fee uint64) *Withdrawal {
	return NewWithdrawalInt(sdk.NewInt(shares), tick, fee)
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawals ...*Withdrawal) {
	s.withdraws(s.alice, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdraws(withdrawals ...*Withdrawal) {
	s.withdraws(s.bob, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdraws(withdrawals ...*Withdrawal) {
	s.withdraws(s.carol, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdraws(withdrawals ...*Withdrawal) {
	s.withdraws(s.dan, withdrawals...)
}

func (s *MsgServerTestSuite) withdraws(account sdk.AccAddress, withdrawals ...*Withdrawal) {
	tickIndexes := make([]int64, len(withdrawals))
	fee := make([]uint64, len(withdrawals))
	sharesToRemove := make([]sdk.Int, len(withdrawals))
	for i, e := range withdrawals {
		tickIndexes[i] = e.TickIndex
		fee[i] = e.Fee
		sharesToRemove[i] = e.Shares
	}

	_, err := s.msgServer.Withdrawal(s.goCtx, &types.MsgWithdrawal{
		Creator:         account.String(),
		Receiver:        account.String(),
		TokenA:          "TokenA",
		TokenB:          "TokenB",
		SharesToRemove:  sharesToRemove,
		TickIndexesAToB: tickIndexes,
		Fees:            fee,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawFails(expectedErr error, withdrawals ...*Withdrawal) {
	s.withdrawFails(s.alice, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdrawFails(expectedErr error, withdrawals ...*Withdrawal) {
	s.withdrawFails(s.bob, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdrawFails(expectedErr error, withdrawals ...*Withdrawal) {
	s.withdrawFails(s.carol, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdrawFails(expectedErr error, withdrawals ...*Withdrawal) {
	s.withdrawFails(s.dan, expectedErr, withdrawals...)
}

func (s *MsgServerTestSuite) withdrawFails(account sdk.AccAddress, expectedErr error, withdrawals ...*Withdrawal) {
	tickIndexes := make([]int64, len(withdrawals))
	fee := make([]uint64, len(withdrawals))
	sharesToRemove := make([]sdk.Int, len(withdrawals))
	for i, e := range withdrawals {
		tickIndexes[i] = e.TickIndex
		fee[i] = e.Fee
		sharesToRemove[i] = e.Shares
	}

	_, err := s.msgServer.Withdrawal(s.goCtx, &types.MsgWithdrawal{
		Creator:         account.String(),
		Receiver:        account.String(),
		TokenA:          "TokenA",
		TokenB:          "TokenB",
		SharesToRemove:  sharesToRemove,
		TickIndexesAToB: tickIndexes,
		Fees:            fee,
	})
	s.Assert().NotNil(err)
	s.Assert().ErrorIs(err, expectedErr)
}

/// Cancel limit order

func (s *MsgServerTestSuite) aliceCancelsLimitSell(trancheKey string) {
	s.cancelsLimitSell(s.alice, trancheKey)
}

func (s *MsgServerTestSuite) bobCancelsLimitSell(trancheKey string) {
	s.cancelsLimitSell(s.bob, trancheKey)
}

func (s *MsgServerTestSuite) carolCancelsLimitSell(trancheKey string) {
	s.cancelsLimitSell(s.carol, trancheKey)
}

func (s *MsgServerTestSuite) danCancelsLimitSell(trancheKey string) {
	s.cancelsLimitSell(s.dan, trancheKey)
}

func (s *MsgServerTestSuite) cancelsLimitSell(account sdk.AccAddress, trancheKey string) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:    account.String(),
		TrancheKey: trancheKey,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceCancelsLimitSellFails(trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.alice, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) bobCancelsLimitSellFails(trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.bob, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) carolCancelsLimitSellFails(trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.carol, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) danCancelsLimitSellFails(trancheKey string, expectedErr error) {
	s.cancelsLimitSellFails(s.dan, trancheKey, expectedErr)
}

func (s *MsgServerTestSuite) cancelsLimitSellFails(account sdk.AccAddress, trancheKey string, expectedErr error) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:    account.String(),
		TrancheKey: trancheKey,
	})
	s.Assert().ErrorIs(err, expectedErr)
}

/// Swap

func (s *MsgServerTestSuite) aliceMarketSells(selling string, amountIn int) {
	s.marketSells(s.alice, selling, amountIn)
}

func (s *MsgServerTestSuite) bobMarketSells(selling string, amountIn int) {
	s.marketSells(s.bob, selling, amountIn)
}

func (s *MsgServerTestSuite) carolMarketSells(selling string, amountIn int) {
	s.marketSells(s.carol, selling, amountIn)
}

func (s *MsgServerTestSuite) danMarketSells(selling string, amountIn int) {
	s.marketSells(s.dan, selling, amountIn)
}

func (s *MsgServerTestSuite) marketSells(account sdk.AccAddress, selling string, amountIn int) {
	tokenIn, tokenOut := GetInOutTokens(selling, "TokenA", "TokenB")
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		AmountIn: sdk.NewInt(int64(amountIn)),
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceMarketSellsWithMaxOut(selling string, amountIn, maxAmountOut int) {
	s.marketSellsWithMaxOut(s.alice, selling, amountIn, maxAmountOut)
}

func (s *MsgServerTestSuite) bobMarketSellsWithMaxOut(selling string, amountIn, maxAmountOut int) {
	s.marketSellsWithMaxOut(s.bob, selling, amountIn, maxAmountOut)
}

func (s *MsgServerTestSuite) carolMarketSellsWithMaxOut(selling string, amountIn, maxAmountOut int) {
	s.marketSellsWithMaxOut(s.carol, selling, amountIn, maxAmountOut)
}

func (s *MsgServerTestSuite) danMarketSellsWithMaxOut(selling string, amountIn, maxAmountOut int) {
	s.marketSellsWithMaxOut(s.dan, selling, amountIn, maxAmountOut)
}

func (s *MsgServerTestSuite) marketSellsWithMaxOut(account sdk.AccAddress, selling string, amountIn int, maxAmountOut int) {
	tokenIn, tokenOut := GetInOutTokens(selling, "TokenA", "TokenB")
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:      account.String(),
		Receiver:     account.String(),
		TokenIn:      tokenIn,
		TokenOut:     tokenOut,
		AmountIn:     sdk.NewInt(int64(amountIn)),
		MaxAmountOut: sdk.NewInt(int64(maxAmountOut)),
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceMarketSellFails(err error, selling string, amountIn int) {
	s.marketSellFails(s.alice, err, selling, amountIn)
}

func (s *MsgServerTestSuite) bobMarketSellFails(err error, selling string, amountIn int) {
	s.marketSellFails(s.bob, err, selling, amountIn)
}

func (s *MsgServerTestSuite) carolMarketSellFails(err error, selling string, amountIn int) {
	s.marketSellFails(s.bob, err, selling, amountIn)
}

func (s *MsgServerTestSuite) danMarketSellFails(err error, selling string, amountIn int) {
	s.marketSellFails(s.bob, err, selling, amountIn)
}

func (s *MsgServerTestSuite) marketSellFails(account sdk.AccAddress, expectedErr error, selling string, amountIn int) {
	tokenIn, tokenOut := GetInOutTokens(selling, "TokenA", "TokenB")
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		AmountIn: sdk.NewInt(int64(amountIn)),
	})
	s.Assert().ErrorIs(err, expectedErr)
}

/// MultiHopSwap

func (s *MsgServerTestSuite) aliceMultiHopSwaps(routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwaps(s.alice, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) bobMultiHopSwaps(routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwaps(s.bob, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) carolMultiHopSwaps(routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwaps(s.carol, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) danMultiHopSwaps(routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwaps(s.dan, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) multiHopSwaps(account sdk.AccAddress, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	msg := types.NewMsgMultiHopSwap(
		account.String(),
		account.String(),
		routes,
		sdk.NewInt(int64(amountIn)),
		exitLimitPrice,
		pickBest,
	)
	_, err := s.msgServer.MultiHopSwap(s.goCtx, msg)
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceMultiHopSwapFails(err error, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwapFails(s.alice, err, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) bobMultiHopSwapFails(err error, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwapFails(s.bob, err, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) carolMultiHopSwapFails(err error, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwapFails(s.carol, err, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) danMultiHopSwapFails(err error, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	s.multiHopSwapFails(s.dan, err, routes, amountIn, exitLimitPrice, pickBest)
}

func (s *MsgServerTestSuite) multiHopSwapFails(account sdk.AccAddress, expectedErr error, routes [][]string, amountIn int, exitLimitPrice sdk.Dec, pickBest bool) {
	msg := types.NewMsgMultiHopSwap(
		account.String(),
		account.String(),
		routes,
		sdk.NewInt(int64(amountIn)),
		exitLimitPrice,
		pickBest,
	)
	_, err := s.msgServer.MultiHopSwap(s.goCtx, msg)
	s.Assert().ErrorIs(err, expectedErr)
}

/// Withdraw filled limit order

func (s *MsgServerTestSuite) aliceWithdrawsLimitSell(trancheKey string) {
	s.withdrawsLimitSell(s.alice, trancheKey)
}

func (s *MsgServerTestSuite) bobWithdrawsLimitSell(trancheKey string) {
	s.withdrawsLimitSell(s.bob, trancheKey)
}

func (s *MsgServerTestSuite) carolWithdrawsLimitSell(trancheKey string) {
	s.withdrawsLimitSell(s.carol, trancheKey)
}

func (s *MsgServerTestSuite) danWithdrawsLimitSell(trancheKey string) {
	s.withdrawsLimitSell(s.dan, trancheKey)
}

func (s *MsgServerTestSuite) withdrawsLimitSell(account sdk.AccAddress, trancheKey string) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:    account.String(),
		TrancheKey: trancheKey,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawLimitSellFails(expectedErr error, trancheKey string) {
	s.withdrawLimitSellFails(s.alice, expectedErr, trancheKey)
}

func (s *MsgServerTestSuite) bobWithdrawLimitSellFails(expectedErr error, trancheKey string) {
	s.withdrawLimitSellFails(s.bob, expectedErr, trancheKey)
}

func (s *MsgServerTestSuite) carolWithdrawLimitSellFails(expectedErr error, trancheKey string) {
	s.withdrawLimitSellFails(s.carol, expectedErr, trancheKey)
}

func (s *MsgServerTestSuite) danWithdrawLimitSellFails(expectedErr error, trancheKey string) {
	s.withdrawLimitSellFails(s.dan, expectedErr, trancheKey)
}

func (s *MsgServerTestSuite) withdrawLimitSellFails(account sdk.AccAddress, expectedErr error, trancheKey string) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:    account.String(),
		TrancheKey: trancheKey,
	})
	s.Assert().ErrorIs(err, expectedErr)
}

// Shares
func (s *MsgServerTestSuite) getPoolShares(
	token0 string,
	token1 string,
	tick int64,
	fee uint64,
) (shares sdk.Int) {
	sharesID := types.NewDepositDenom(&types.PairID{Token0: token0, Token1: token1}, tick, fee).String()
	return s.app.BankKeeper.GetSupply(s.ctx, sharesID).Amount
}

func (s *MsgServerTestSuite) assertPoolShares(
	tick int64,
	fee uint64,
	sharesExpected uint64,
) {
	sharesExpectedInt := sdk.NewIntFromUint64(sharesExpected)
	sharesOwned := s.getPoolShares("TokenA", "TokenB", tick, fee)
	s.Assert().Equal(sharesExpectedInt, sharesOwned)
}

func (s *MsgServerTestSuite) getAccountShares(
	account sdk.AccAddress,
	token0 string,
	token1 string,
	tick int64,
	fee uint64,
) (shares sdk.Int) {
	sharesID := types.NewDepositDenom(&types.PairID{Token0: token0, Token1: token1}, tick, fee).String()
	return s.app.BankKeeper.GetBalance(s.ctx, account, sharesID).Amount
}

func (s *MsgServerTestSuite) assertAccountShares(
	account sdk.AccAddress,
	tick int64,
	fee uint64,
	sharesExpected uint64,
) {
	sharesExpectedInt := sdk.NewIntFromUint64(sharesExpected)
	sharesOwned := s.getAccountShares(account, "TokenA", "TokenB", tick, fee)
	s.Assert().Equal(sharesExpectedInt, sharesOwned, "expected %s != actual %s", sharesExpected, sharesOwned)
}

func (s *MsgServerTestSuite) assertAliceShares(tick int64, fee, sharesExpected uint64) {
	s.assertAccountShares(s.alice, tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertBobShares(tick int64, fee, sharesExpected uint64) {
	s.assertAccountShares(s.bob, tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertCarolShares(tick int64, fee, sharesExpected uint64) {
	s.assertAccountShares(s.carol, tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertDanShares(tick int64, fee, sharesExpected uint64) {
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
	pairID := CreatePairID("TokenA", "TokenB")
	curr0To1Actual, _ := s.app.DexKeeper.GetCurrTick0To1(s.ctx, pairID)
	s.Assert().Equal(curr0To1Expected, curr0To1Actual)
}

func (s *MsgServerTestSuite) assertCurr1To0(curr1To0Expected int64) {
	pairID := CreatePairID("TokenA", "TokenB")

	curr1to0Actual, _ := s.app.DexKeeper.GetCurrTick1To0(s.ctx, pairID)
	s.Assert().Equal(curr1To0Expected, curr1to0Actual)
}

// Pool liquidity (i.e. deposited rather than LO)
func (s *MsgServerTestSuite) assertLiquidityAtTick(amountA, amountB sdk.Int, tickIndex int64, fee uint64) {
	liquidityA, liquidityB := s.getLiquidityAtTick(tickIndex, fee)
	s.Assert().True(amountA.Equal(liquidityA), "liquidity A: actual %s, expected %s", liquidityA, amountA)
	s.Assert().True(amountB.Equal(liquidityB), "liquidity B: actual %s, expected %s", liquidityB, amountB)
}

func (s *MsgServerTestSuite) assertLiquidityAtTickWithDenom(pairID *types.PairID, expected0, expected1 sdk.Int, tickIndex int64, fee uint64) {
	liquidity0, liquidity1 := s.getLiquidityAtTickWithDenom(pairID, tickIndex, fee)
	s.Assert().True(expected0.Equal(liquidity0), "liquidity 0: actual %s, expected %s", liquidity0, expected0)
	s.Assert().True(expected1.Equal(liquidity1), "liquidity 1: actual %s, expected %s", liquidity1, expected1)
}

func (s *MsgServerTestSuite) assertPoolLiquidity(amountA, amountB int, tickIndex int64, fee uint64) {
	s.assertLiquidityAtTick(sdk.NewInt(int64(amountA)), sdk.NewInt(int64(amountB)), tickIndex, fee)
}

func (s *MsgServerTestSuite) assertNoLiquidityAtTick(tickIndex int64, fee uint64) {
	s.assertLiquidityAtTick(sdk.ZeroInt(), sdk.ZeroInt(), tickIndex, fee)
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
	userShares, totalShares := s.getLimitUserSharesAtTick(account, selling, tickIndex), s.getLimitTotalSharesAtTick(selling, tickIndex)
	userRatio := userShares.ToDec().QuoInt(totalShares)
	filled := s.getLimitFilledLiquidityAtTickAtIndex(selling, tickIndex, trancheKey)
	amt := sdk.NewInt(int64(amount))
	userFilled := userRatio.MulInt(filled).RoundInt()
	s.Assert().True(amt.Equal(userFilled))
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

func (s *MsgServerTestSuite) assertLimitLiquidityAtTick(selling string, tickIndex, amount int64) {
	s.assertLimitLiquidityAtTickInt(selling, tickIndex, sdk.NewInt(amount))
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTickInt(selling string, tickIndex int64, amount sdk.Int) {
	pairID := CreatePairID("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairID, selling, tickIndex)
	liquidity := sdk.ZeroInt()
	for _, t := range tranches {
		if !t.IsExpired(s.ctx) {
			liquidity = liquidity.Add(t.ReservesTokenIn)
		}
	}

	s.Assert().True(amount.Equal(liquidity), "Incorrect liquidity: expected %s, have %s", amount.String(), liquidity.String())
}

func (s *MsgServerTestSuite) assertFillAndPlaceTrancheKeys(selling string, tickIndex int64, expectedFill, expectedPlace string) {
	pairID := CreatePairID("TokenA", "TokenB")
	placeTranche, foundPlace := s.app.DexKeeper.GetPlaceTranche(s.ctx, pairID, selling, tickIndex)
	fillTranche, foundFill := s.app.DexKeeper.GetFillTranche(s.ctx, pairID, selling, tickIndex)
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
	pairID := CreatePairID("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairID, selling, tickIndex)
	fillTranche := tranches[0]
	// get user shares and total shares
	userShares := s.getLimitUserSharesAtTickAtIndex(account, fillTranche.TrancheKey)
	if len(tranches) >= 2 {
		userShares = userShares.Add(s.getLimitUserSharesAtTickAtIndex(account, tranches[1].TrancheKey))
	}

	return userShares
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTickAtIndex(account sdk.AccAddress, trancheKey string) sdk.Int {
	userShares, userSharesFound := s.app.DexKeeper.GetLimitOrderTrancheUser(s.ctx, account.String(), trancheKey)
	s.Assert().True(userSharesFound, "Failed to get limit order user shares for index %s", trancheKey)
	return userShares.SharesOwned
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTick(selling string, tickIndex int64) sdk.Int {
	pairID := CreatePairID("TokenA", "TokenB")
	tranches := s.app.DexKeeper.GetAllLimitOrderTrancheAtIndex(s.ctx, pairID, selling, tickIndex)
	// get user shares and total shares
	totalShares := sdk.ZeroInt()
	for _, t := range tranches {
		totalShares = totalShares.Add(t.TotalTokenIn)
	}

	return totalShares
}

func (s *MsgServerTestSuite) getLimitFilledLiquidityAtTickAtIndex(selling string, tickIndex int64, trancheKey string) sdk.Int {
	pairID := CreatePairID("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, _, found := s.app.DexKeeper.FindLimitOrderTranche(s.ctx, pairID, tickIndex, selling, trancheKey)
	s.Assert().True(found, "Failed to get limit order filled reserves for index %s", trancheKey)

	return tranche.ReservesTokenOut
}

func (s *MsgServerTestSuite) getLimitReservesAtTickAtKey(selling string, tickIndex int64, trancheKey string) sdk.Int {
	pairID := CreatePairID("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, _, found := s.app.DexKeeper.FindLimitOrderTranche(s.ctx, pairID, tickIndex, selling, trancheKey)
	s.Assert().True(found, "Failed to get limit order reserves for index %s", trancheKey)

	return tranche.ReservesTokenIn
}

func (s *MsgServerTestSuite) assertNLimitOrderExpiration(expected int) {
	exps := s.app.DexKeeper.GetAllLimitOrderExpiration(s.ctx)
	s.Assert().Equal(expected, len(exps))
}

// SingleLimitOrderFill() simulates the fill of a single limit order and returns the amount
// swapped into it, filling some of it (amount_in) and the amount swapped out (amount_out). It
// takes as input the amount that was placed for the limit order (amount_placed), the price the
// trader pays when filling it (price_filled_at) and the amount that they are swapping (amount_to_swap).
// The format of the return statement is (amount_in, amount_out).
func SingleLimitOrderFill(amountPlaced sdk.Int,
	priceFilledAt sdk.Dec,
	amountToSwap sdk.Int,
) (amountIn sdk.Dec, amountOut sdk.Dec) {
	amountPlacedDec := amountPlaced.ToDec()
	amountPlacedForPrice := amountPlacedDec.Quo(priceFilledAt)
	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amountToSwap.ToDec().GT(amountPlacedForPrice) {
		amountOut = amountPlaced.ToDec()
		amountIn = amountPlacedForPrice
	} else {
		amountIn = amountToSwap.ToDec()
		amountOut = amountIn.Mul(priceFilledAt)
	}

	return amountIn, amountOut
}

// Calls SingleLimitOrderFill() and updates the filled and unfilled reserves.
// Returns the unfilled reserves (unfilled_reserves), filled reserves (filled_reserves) and the amount left to swap
// (amount_to_swap_remaining)
func SingleLimitOrderFillAndUpdate(amountPlaced sdk.Int,
	priceFilledAt sdk.Dec,
	amountToSwap sdk.Int,
	unfilledReserves sdk.Int,
) (sdk.Dec, sdk.Dec, sdk.Dec) {
	amountIn, amountOut := SingleLimitOrderFill(amountPlaced, priceFilledAt, amountToSwap)
	unfilledReservesDec := unfilledReserves.ToDec().Sub(amountOut)
	filledReserves := amountPlaced.ToDec().Add(amountIn)
	amountToSwapRemaining := amountToSwap.ToDec().Sub(amountIn)

	return unfilledReservesDec, filledReserves, amountToSwapRemaining
}

// MultipleLimitOrderFills() simulates the fill of multiple consecutive limit orders and returns the
// total amount filled. It takes as input the amounts that were placed for the limit
// order (amount_placed), the pricesthe trader pays when filling the orders (price_filled_at)
// and the amount that they are swapping (amount_to_swap).
func MultipleLimitOrderFills(amountsPlaced []sdk.Int, prices []sdk.Dec, amountToSwap sdk.Int) sdk.Dec {
	totalOut, amountRemaining := sdk.ZeroDec(), amountToSwap

	// Loops through all of the limit orders that need to be filled
	for i := 0; i < len(amountsPlaced); i++ {
		_, amountOut := SingleLimitOrderFill(amountsPlaced[i], prices[i], amountRemaining)

		// amount_remaining = amount_remaining.Sub(amount_in)
		totalOut = totalOut.Add(amountOut)
	}

	return totalOut
}

// SinglePoolSwap() simulates swapping through a single liquidity pool and returns the amount
// swapped into it (amount_in) and the amount swapped out, received by the swapper (amount_out). It
// takes as input the amount of liquidity in the pool (amount_liquidity), the price the
// trader pays when swapping through it (price_swapped_at) and the amount that they are
// swapping (amount_to_swap). The format of the return statement is (amount_in, amount_out).
// Same thing as SingleLimitOrderFill() except in naming.
func SinglePoolSwap(amountLiquidity sdk.Int, priceSwappedAt sdk.Dec, amountToSwap sdk.Int) (sdk.Dec, sdk.Dec) {
	var amountOut, amountIn sdk.Dec
	liquidityAtPrice := amountLiquidity.ToDec().Quo(priceSwappedAt)
	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amountToSwap.ToDec().GT(liquidityAtPrice) {
		amountOut = amountLiquidity.ToDec()
		amountIn = liquidityAtPrice
	} else {
		amountIn = amountToSwap.ToDec()
		amountOut = amountIn.Mul(priceSwappedAt)
	}

	return amountIn, amountOut
}

// SinglePoolSwapAndUpdate() simulates swapping through a single liquidity pool and updates that pool's
// liquidity. Takes in all of the same inputs as SinglePoolSwap(): amount_liquidity, price_swapped_at,
// and amount_to_swap; but has additional inputs, reservesOfInToken, reservesOfOutToken. It returns the
// updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format of
// (resulting_reserves_in_token, resulting_reserves_out_token, amount_in, amount_out)
func SinglePoolSwapAndUpdate(amountLiquidity sdk.Int,
	priceSwappedAt sdk.Dec,
	amountToSwap sdk.Int,
	reservesOfInToken sdk.Int,
	reservesOfOutToken sdk.Int,
) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec) {
	amountIn, amountOut := SinglePoolSwap(amountLiquidity, priceSwappedAt, amountToSwap)
	resultingReservesInToken := reservesOfInToken.ToDec().Add(amountIn)
	resultingReservesOutToken := reservesOfOutToken.ToDec().Add(amountOut)

	return resultingReservesInToken, resultingReservesOutToken, amountIn, amountOut
}

// SinglePoolSwapAndUpdateDirection() simulates swapping through a single liquidity pool and updates that pool's
// liquidity and specifies whether the in and out tokens are 0 or 1. Takes in all of the same inputs as
// SinglePoolSwapAndUpdate(): amount_liquidity, price_swapped_at, amount_to_swap, reservesOfToken0 sdk.Int,
// reservesOfToken1 but has an additional input inToken which is a bool indicating whether 0 or 1 is swapped into
// the pool. It returns the updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format
// of (reservesOfInToken,reservesOfOutToken).
func SinglePoolSwapAndUpdateDirectional(amountLiquidity sdk.Int,
	priceSwappedAt sdk.Dec,
	amountToSwap sdk.Int,
	reservesOfToken0 sdk.Int,
	reservesOfToken1 sdk.Int,
	inToken bool,
) (resultingReservesOfToken0 sdk.Dec, resultingReservesOfToken1 sdk.Dec) {
	if inToken {
		resultingReservesOfToken1, resultingReservesOfToken0, _, _ = SinglePoolSwapAndUpdate(amountLiquidity,
			priceSwappedAt,
			amountToSwap,
			reservesOfToken1,
			reservesOfToken0)
	} else {
		resultingReservesOfToken0, resultingReservesOfToken1, _, _ = SinglePoolSwapAndUpdate(amountLiquidity,
			priceSwappedAt,
			amountToSwap,
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
func MultiplePoolSwapAndUpdate(amountsLiquidity []sdk.Int,
	pricesSwappedAt []sdk.Dec,
	amountToSwap sdk.Int,
	reservesInTokenArray []sdk.Int,
	reservesOutTokenArray []sdk.Int,
) ([]sdk.Dec, []sdk.Dec, sdk.Dec, sdk.Dec) {
	numPools := len(amountsLiquidity)
	amountRemainingDec := amountToSwap.ToDec()
	var amountOutTotal, amountIn, amountOutTemp sdk.Dec
	resultingReservesInToken := make([]sdk.Dec, numPools)
	resultingReservesOutToken := make([]sdk.Dec, numPools)
	for i := 0; i < numPools; i++ {
		resultingReservesInToken[i], resultingReservesOutToken[i], amountIn, amountOutTemp = SinglePoolSwapAndUpdate(amountsLiquidity[i],
			pricesSwappedAt[i],
			amountToSwap,
			reservesInTokenArray[i],
			reservesOutTokenArray[i])
		amountOutTotal = amountOutTotal.Add(amountOutTemp)
		amountRemainingDec = amountRemainingDec.Sub(amountIn)
		i++
	}

	return resultingReservesInToken, resultingReservesOutToken, amountRemainingDec, amountOutTotal
}

func SharesOnDeposit(existingAmount0, existingAmount1, newAmount0, newAmount1 sdk.Int, tickIndex int64) (sharesMinted sdk.Int) {
	price1To0 := types.MustNewPrice(-1 * tickIndex)
	newAmount0Dec := sdk.NewDecFromInt(newAmount0)
	newValue := newAmount0Dec.Add(price1To0.MulInt(newAmount1))

	if existingAmount0.Add(existingAmount1).GT(sdk.ZeroInt()) {
		existingValue := existingAmount0.ToDec().Add(price1To0.MulInt(existingAmount1))
		sharesMinted = sharesMinted.ToDec().Mul(newValue.Quo(existingValue)).TruncateInt()
	} else {
		sharesMinted = newValue.TruncateInt()
	}

	return sharesMinted
}

func (s *MsgServerTestSuite) calcAutoswapSharesMinted(centerTick int64, fee uint64, residual0, residual1, balanced0, balanced1, totalShares, valuePool int64) sdk.Int {
	residual0Int, residual1Int, balanced0Int, balanced1Int, totalSharesInt, valuePoolInt := sdk.NewInt(residual0), sdk.NewInt(residual1), sdk.NewInt(balanced0), sdk.NewInt(balanced1), sdk.NewInt(totalShares), sdk.NewInt(valuePool)

	// residualValue = 1.0001^-f * residualAmount0 + 1.0001^{i-f} * residualAmount1
	// balancedValue = balancedAmount0 + 1.0001^{i} * balancedAmount1
	// value = residualValue + balancedValue
	// shares minted = value * totalShares / valuePool

	centerPrice := types.MustNewPrice(-1 * centerTick)
	leftPrice := types.MustNewPrice(-1 * (centerTick - int64(fee)))
	discountPrice := types.MustNewPrice(-1 * int64(fee))

	balancedValue := balanced0Int.ToDec().Add(centerPrice.MulInt(balanced1Int)).TruncateInt()
	residualValue := discountPrice.MulInt(residual0Int).Add(leftPrice.Mul(residual1Int.ToDec())).TruncateInt()
	valueMint := balancedValue.Add(residualValue)

	return valueMint.Mul(totalSharesInt).Quo(valuePoolInt)
}

func (s *MsgServerTestSuite) calcSharesMinted(centerTick, amount0Int, amount1Int int64) sdk.Int {
	amount0, amount1 := sdk.NewInt(amount0Int), sdk.NewInt(amount1Int)
	centerPrice := types.MustNewPrice(-1 * centerTick)

	return amount0.ToDec().Add(centerPrice.Mul(amount1.ToDec())).TruncateInt()
}

func (s *MsgServerTestSuite) calcExpectedBalancesAfterWithdrawOnePool(sharesMinted sdk.Int, account sdk.AccAddress, tickIndex int64, fee uint64) (sdk.Int, sdk.Int, sdk.Int, sdk.Int) {
	dexCurrentBalance0 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenA").Amount
	dexCurrentBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenB").Amount
	currentBalance0 := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount
	currentBalance1 := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount
	amountPool0, amountPool1 := s.getLiquidityAtTick(tickIndex, fee)
	poolShares := s.getPoolShares("TokenA", "TokenB", tickIndex, fee)

	amountOut0 := amountPool0.Mul(sharesMinted).Quo(poolShares)
	amountOut1 := amountPool1.Mul(sharesMinted).Quo(poolShares)

	expectedBalance0 := currentBalance0.Add(amountOut0)
	expectedBalance1 := currentBalance1.Add(amountOut1)
	dexExpectedBalance0 := dexCurrentBalance0.Sub(amountOut0)
	dexExpectedBalance1 := dexCurrentBalance1.Sub(amountOut1)

	return expectedBalance0, expectedBalance1, dexExpectedBalance0, dexExpectedBalance1
}

// Swap helpers (use for writing the tests, but replace with actual values before finishing!)
func (s *MsgServerTestSuite) calculateSingleSwapNoLOAToB(tick, tickLiqudity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOAToB(tick, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapAToB(tick, tickLiqudidty, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(tick)
	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOBToA(tick, tickLiqudity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOBToA(tick, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapBToA(tick, tickLiqudidty, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	price := types.MustNewPrice(-1 * tick)
	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwapNoLO(price *types.Price, tickLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, tickLiquidity, 0, amountIn)
}

func calculateSingleSwapOnlyLO(price *types.Price, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, 0, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwap(price *types.Price, tickLiquidity, tickLimitOrderLiquidity, amountIn int64) (sdk.Int, sdk.Int) {
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

func calculateSwap(price *types.Price, liquidity, amountIn int64) (sdk.Int, sdk.Int) {
	amountInInt := sdk.NewInt(amountIn)
	liquidityInt := sdk.NewInt(liquidity)
	if tmpAmountOut := price.MulInt(amountInInt); tmpAmountOut.LT(liquidityInt.ToDec()) {
		// fmt.Printf("sufficient tmpOut %s\n", tmpAmountOut)
		// sufficient liquidity
		return sdk.ZeroInt(), tmpAmountOut.TruncateInt()
	}
	// only sufficient for part of amountIn
	tmpAmountIn := price.Inv().MulInt(liquidityInt).TruncateInt()

	return amountInInt.Sub(tmpAmountIn), liquidityInt
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

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOAToB(tickIndexes, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(tickIndexes[i])
	}

	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOAToB(tickIndexes, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(tickIndexes[i])
	}

	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsBToA(tickIndexes, tickLiquidities, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(-1 * tickIndexes[i])
	}

	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOBToA(tickIndexes, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]*types.Price, len(tickIndexes))
	for i := range prices {
		prices[i] = types.MustNewPrice(-1 * tickIndexes[i])
	}

	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOBToA(tickIndexes, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
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

func (s *MsgServerTestSuite) calculateMultipleSwaps(prices []*types.Price, tickLiquidities, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	amountLeft, amountOut := sdk.NewInt(amountIn), sdk.ZeroInt()
	for i := 0; i < len(prices); i++ {
		tmpAmountLeft, tmpAmountOut := calculateSingleSwap(prices[i], tickLiquidities[i], tickLimitOrderLiquidities[i], amountLeft.Int64())
		amountLeft, amountOut = tmpAmountLeft, amountOut.Add(tmpAmountOut)
	}

	return amountLeft, amountOut
}

func (s *MsgServerTestSuite) nextBlockWithTime(blockTime time.Time) {
	newCtx := s.ctx.WithBlockTime(blockTime)
	s.ctx = newCtx
	s.goCtx = sdk.WrapSDKContext(newCtx)
	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash,
		Time: blockTime,
	}})
}
