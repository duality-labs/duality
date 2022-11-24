package keeper_test

import (
	"context"
	"fmt"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper/internal/testutils"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

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
	feeTiers    []types.FeeList
}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
}

func (s *MsgServerTestSuite) SetupTest() {
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

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
	feeTiers := []types.FeeList{
		{Id: 0, Fee: 1},
		{Id: 1, Fee: 3},
		{Id: 2, Fee: 5},
		{Id: 3, Fee: 10},
	}

	// Set Fee List
	app.DexKeeper.AppendFeeList(ctx, feeTiers[0])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[1])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[2])
	app.DexKeeper.AppendFeeList(ctx, feeTiers[3])

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

func (s *MsgServerTestSuite) assertAccountBalancesDec(account sdk.AccAddress, aBalance sdk.Dec, bBalance sdk.Dec) {
	s.assertAccountBalances(account, aBalance.RoundInt64(), bBalance.RoundInt64())
}

func (s *MsgServerTestSuite) assertAccountBalances(
	account sdk.AccAddress,
	aBalance int64,
	bBalance int64,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount.Int64()

	s.Assert().Equal(aActual, aBalance, "expected %s != actual %s", aBalance, aBalance)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount.Int64()
	s.Assert().Equal(bActual, bBalance, "expected %s != actual %s", bBalance, bBalance)
}

func (s *MsgServerTestSuite) assertAliceBalances(a int64, b int64) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a int64, b int64) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalances(a int64, b int64) {
	s.assertAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a int64, b int64) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a int64, b int64) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) aliceLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.alice, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) bobLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.bob, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) carolLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.carol, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) danLimitSells(selling string, tick int, amountIn int) {
	s.limitSells(s.dan, selling, tick, amountIn)
}

func (s *MsgServerTestSuite) limitSells(account sdk.AccAddress, tokenIn string, tick int, amountIn int) {
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  amountInDec,
	})
	s.Assert().Nil(err)
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
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  amountInDec,
	})
	s.Assert().ErrorIs(expectedErr, err)
}

type Deposit struct {
	AmountA   sdk.Dec
	AmountB   sdk.Dec
	TickIndex int64
	FeeIndex  uint64
}

func NewDeposit(amountA int, amountB int, tickIndex int, feeIndex int) *Deposit {
	return &Deposit{
		AmountA:   sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountA))),
		AmountB:   sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountB))),
		TickIndex: int64(tickIndex),
		FeeIndex:  uint64(feeIndex),
	}
}

func (s *MsgServerTestSuite) aliceDeposits(deposits ...*Deposit) {
	s.deposits(s.alice, deposits...)
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
	amountsA := make([]sdk.Dec, len(deposits))
	amountsB := make([]sdk.Dec, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
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
	})
	s.Assert().Nil(err)
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
	amountsA := make([]sdk.Dec, len(deposits))
	amountsB := make([]sdk.Dec, len(deposits))
	tickIndexes := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndexes[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
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
	amountsA []sdk.Dec
	amountsB []sdk.Dec
}

func (s *MsgServerTestSuite) withdraws(account sdk.AccAddress, withdrawls ...*Withdrawl) error {
	tickIndexes := make([]int64, len(withdrawls))
	feeIndexes := make([]uint64, len(withdrawls))
	sharesToRemove := make([]sdk.Dec, len(withdrawls))
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

	return err
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawals ...*Withdrawl) error {
	return s.withdraws(s.alice, withdrawals...)
}

func (s *MsgServerTestSuite) bobWithdraws(withdrawals ...*Withdrawl) error {
	return s.withdraws(s.bob, withdrawals...)
}

func (s *MsgServerTestSuite) carolWithdraws(withdrawals ...*Withdrawl) error {
	return s.withdraws(s.carol, withdrawals...)
}

func (s *MsgServerTestSuite) danWithdraws(withdrawals ...*Withdrawl) error {
	return s.withdraws(s.dan, withdrawals...)
}

func (s *MsgServerTestSuite) aliceCancelsLimitSell(keyToken string, tick int, key int, amountOut int) {
	s.cancelsLimitSell(s.alice, keyToken, tick, key, amountOut)
}

func (s *MsgServerTestSuite) bobCancelsLimitSell(keyToken string, tick int, key int, amountOut int) {
	s.cancelsLimitSell(s.bob, keyToken, tick, key, amountOut)
}

func (s *MsgServerTestSuite) carolCancelsLimitSell(keyToken string, tick int, key int, amountOut int) {
	s.cancelsLimitSell(s.carol, keyToken, tick, key, amountOut)
}

func (s *MsgServerTestSuite) danCancelsLimitSell(keyToken string, tick int, key int, amountOut int) {
	s.cancelsLimitSell(s.dan, keyToken, tick, key, amountOut)
}

func (s *MsgServerTestSuite) cancelsLimitSell(account sdk.AccAddress, selling string, tick int, key int, sharesOut int) {
	sharesOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(sharesOut)))
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(key),
		SharesOut: sharesOutDec,
	})
	s.Assert().Nil(err)
}

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
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	minOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(minOut)))
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  selling,
		AmountIn: amountInDec,
		MinOut:   minOutDec,
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
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	minOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(minOut)))
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  selling,
		AmountIn: amountInDec,
		MinOut:   minOutDec,
	})
	s.Assert().ErrorIs(expectedErr, err)
}

func (s *MsgServerTestSuite) withdrawsLimitSell(account sdk.AccAddress, selling string, tick int, tranche int) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(tranche),
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawsLimitSell(selling string, tick int, tranche int) {
	s.withdrawsLimitSell(s.alice, selling, tick, tranche)
}

func (s *MsgServerTestSuite) bobWithdrawsLimitSell(selling string, tick int, tranche int) {
	s.withdrawsLimitSell(s.bob, selling, tick, tranche)
}

func (s *MsgServerTestSuite) carolWithdrawsLimitSell(selling string, tick int, tranche int) {
	s.withdrawsLimitSell(s.carol, selling, tick, tranche)
}

func (s *MsgServerTestSuite) danWithdrawsLimitSell(selling string, tick int, tranche int) {
	s.withdrawsLimitSell(s.dan, selling, tick, tranche)
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

type Withdrawl struct {
	TickIndex int64
	FeeIndex  uint64
	Shares    sdk.Dec
}

func NewWithdrawl(shares int64, tick int64, feeIndex uint64) *Withdrawl {
	return NewWithdrawlDec(sdk.NewDec(shares), tick, feeIndex)
}

func NewWithdrawlDec(shares sdk.Dec, tick int64, feeIndex uint64) *Withdrawl {
	return &Withdrawl{
		Shares:    shares,
		FeeIndex:  feeIndex,
		TickIndex: tick,
	}
}

func (s *MsgServerTestSuite) getShares(
	account sdk.AccAddress,
	pairId string,
	tick int64,
	fee uint64,
) (shares sdk.Dec) {

	sharesData, sharesFound := s.app.DexKeeper.GetShares(s.ctx, account.String(), pairId, tick, fee)
	s.Assert().True(sharesFound)
	return sharesData.SharesOwned
}

func (s *MsgServerTestSuite) assertAccountShares(
	account sdk.AccAddress,
	pairId string,
	tick int64,
	fee uint64,
	sharesExpected sdk.Dec,
) {
	sharesOwned := s.getShares(account, pairId, tick, fee)
	s.Assert().Equal(sharesExpected, sharesOwned)
}

func (s *MsgServerTestSuite) assertAliceShares(
	tick int64,
	fee uint64,
	sharesExpected sdk.Dec,
) {
	s.assertAccountShares(s.alice, "TokenA<>TokenB", tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertCurrentTicks(
	expected1To0 int64,
	expected0To1 int64,
) {
	tickMap, found := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA<>TokenB")
	s.Assert().NotNil(found)
	s.Assert().Equal(expected1To0, tickMap.TokenPair.CurrentTick1To0)
	s.Assert().Equal(expected0To1, tickMap.TokenPair.CurrentTick0To1)
}

func (s *MsgServerTestSuite) assertCurr0To1(curr0To1Expected int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetPairMap(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	curr0To1Actual := pair.TokenPair.CurrentTick0To1
	s.Assert().Equal(curr0To1Expected, curr0To1Actual)
}

func (s *MsgServerTestSuite) assertCurr1To0(curr1To0Expected int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetPairMap(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	curr1to0Actual := pair.TokenPair.CurrentTick1To0
	s.Assert().Equal(curr1To0Expected, curr1to0Actual)
}

func (s *MsgServerTestSuite) assertMinTick(minTickExpected int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetPairMap(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	minTickActual := pair.MinTick
	s.Assert().Equal(minTickExpected, minTickActual)
}

func (s *MsgServerTestSuite) assertMaxTick(maxTickExpected int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetPairMap(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	maxTickActual := pair.MaxTick
	s.Assert().Equal(maxTickExpected, maxTickActual)
}

func (s *MsgServerTestSuite) printTicks() {
	tickMap, _ := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA<>TokenB")
	fmt.Printf("\nTick0To1: %v, Tick1To0: %v", tickMap.TokenPair.CurrentTick0To1, tickMap.TokenPair.CurrentTick1To0)
}

func (s *MsgServerTestSuite) assertLiquidityAtTick(amountA int64, amountB int64, tickIndex int64, feeIndex uint64) {
	amtA, amtB := sdk.NewDec(amountA), sdk.NewDec(amountB)
	s.assertLiquidityAtTickDec(amtA, amtB, tickIndex, feeIndex)
}
func (s *MsgServerTestSuite) assertLiquidityAtTickDec(amountA sdk.Dec, amountB sdk.Dec, tickIndex int64, feeIndex uint64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fee := s.feeTiers[feeIndex].Fee
	lowerTick, lowerTickFound := s.app.DexKeeper.GetTickMap(s.ctx, pairId, tickIndex-fee)
	if !lowerTickFound {
		s.Require().Fail("Invalid tick %d and fee %d", tickIndex, fee)
	}
	upperTick, upperTickFound := s.app.DexKeeper.GetTickMap(s.ctx, pairId, tickIndex+fee)
	if !upperTickFound {
		s.Require().Fail("Invalid tick %d and fee %d", tickIndex, fee)
	}

	liquidityA := lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
	liquidityB := upperTick.TickData.Reserve1[feeIndex]
	s.Assert().Equal(amountA, liquidityA)
	s.Assert().Equal(amountB, liquidityB)
}

func (s *MsgServerTestSuite) assertNoLiquidityAtTick(tickIndex int64, feeIndex uint64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fee := s.feeTiers[feeIndex].Fee

	lowerTick, lowerTickFound := s.app.DexKeeper.GetTickMap(s.ctx, pairId, tickIndex-fee)
	if !lowerTickFound {
		s.Assert().True(!lowerTickFound)
		return
	}
	// in case tick was initialized, assert no liquidity in it
	amtA := NewDec(0)
	liquidityA := lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
	s.Assert().Equal(amtA, liquidityA)

	upperTick, upperTickFound := s.app.DexKeeper.GetTickMap(s.ctx, pairId, tickIndex+fee)
	if !upperTickFound {
		s.Assert().True(!upperTickFound)
		return
	}
	// in case tick was initialized, assert no liquidity in it
	amtB := NewDec(0)
	liquidityB := upperTick.TickData.Reserve1[feeIndex]
	s.Assert().Equal(amtB, liquidityB)
}

func (s *MsgServerTestSuite) assertAliceLimitFilledAtTickAtKey(selling string, amount int, tickIndex int64, key uint64) {
	s.assertLimitFilledAtTickAtKey(s.alice, selling, amount, tickIndex, key)
}

func (s *MsgServerTestSuite) assertBobLimitFilledAtTickAtKey(selling string, amount int, tickIndex int64, key uint64) {
	s.assertLimitFilledAtTickAtKey(s.bob, selling, amount, tickIndex, key)
}

func (s *MsgServerTestSuite) assertCarolLimitFilledAtTickAtKey(selling string, amount int, tickIndex int64, key uint64) {
	s.assertLimitFilledAtTickAtKey(s.carol, selling, amount, tickIndex, key)
}

func (s *MsgServerTestSuite) assertDanLimitFilledAtTickAtKey(selling string, amount int, tickIndex int64, key uint64) {
	s.assertLimitFilledAtTickAtKey(s.dan, selling, amount, tickIndex, key)
}

func (s *MsgServerTestSuite) assertLimitFilledAtTickAtKey(account sdk.AccAddress, selling string, amount int, tickIndex int64, key uint64) {
	filled := s.getLimitFilledLiquidityAtTickAtKey(selling, tickIndex, key)
	amt := NewDec(amount)
	s.Assert().True(amt.Equal(filled))
}

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
	amt := NewDec(amount)
	s.assertAccountLimitLiquidityAtTickDec(account, selling, amt, tickIndex)
}

func (s *MsgServerTestSuite) assertAccountLimitLiquidityAtTickDec(account sdk.AccAddress, selling string, amount sdk.Dec, tickIndex int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")

	// get tick liquidity
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get liquidity from fill
	liquidity := s.getLimitReservesAtTickAtKey(selling, tickIndex, fillTranche)
	// if fill == place - 1, get liquidity from place
	if fillTranche == placeTranche-1 {
		liquidity = liquidity.Add(s.getLimitReservesAtTickAtKey(selling, tickIndex, placeTranche))
	}
	// get user liquidity
	userShares, totalShares := s.getLimitUserSharesAtTick(account, selling, tickIndex), s.getLimitTotalSharesAtTick(selling, tickIndex)
	userRatio := userShares.Quo(totalShares)
	// assert enough liq
	userLiquidity := amount.Mul(userRatio)
	s.assertLimitLiquidityAtTickDec(selling, userLiquidity, tickIndex)
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTick(selling string, tickIndex int64, amount int) {
	amt := NewDec(amount)
	s.assertLimitLiquidityAtTickDec(selling, amt, tickIndex)
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTickDec(selling string, amount sdk.Dec, tickIndex int64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get liquidity from fill
	liquidity := s.getLimitReservesAtTickAtKey(selling, tickIndex, fillTranche)
	// if fill == place - 1, get liquidity from place
	if fillTranche == placeTranche-1 {
		liquidity = liquidity.Add(s.getLimitReservesAtTickAtKey(selling, tickIndex, placeTranche))
	}

	s.Assert().True(amount.Equal(liquidity), "Incorrect liquidity: expected %s, have %s", amount.String(), liquidity.String())
}

func (s *MsgServerTestSuite) assertFillAndPlaceTrancheKeys(selling string, tickIndex int64, expectedFill uint64, expectedPlace uint64) {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fill, place := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	s.Assert().Equal(expectedFill, fill)
	s.Assert().Equal(expectedPlace, place)
}

func (s *MsgServerTestSuite) getFillAndPlaceTrancheKeys(selling string, pairId string, tickIndex int64) (uint64, uint64) {
	// grab current fill and place tranches
	tick, tickFound := s.app.DexKeeper.GetTickMap(s.ctx, pairId, tickIndex)
	s.Assert().True(tickFound, "Invalid tickIndex for pair %s", pairId)

	// handle correct limit order pool
	if selling == "TokenA" {
		return tick.LimitOrderTranche0To1.FillTrancheIndex, tick.LimitOrderTranche0To1.PlaceTrancheIndex
	} else {
		return tick.LimitOrderTranche1To0.FillTrancheIndex, tick.LimitOrderTranche1To0.PlaceTrancheIndex
	}
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTick(account sdk.AccAddress, selling string, tickIndex int64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get user shares and total shares
	userShares := s.getLimitUserSharesAtTickAtKey(account, selling, tickIndex, fillTranche)
	if fillTranche == placeTranche-1 {
		userShares = userShares.Add(s.getLimitUserSharesAtTickAtKey(account, selling, tickIndex, placeTranche))
	}
	return userShares
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTickAtKey(account sdk.AccAddress, selling string, tickIndex int64, key uint64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	userShares, userSharesFound := s.app.DexKeeper.GetLimitOrderTrancheUser(s.ctx, pairId, tickIndex, selling, key, account.String())
	s.Assert().True(userSharesFound, "Failed to get limit order user shares for key %s", key)
	return userShares.SharesOwned
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTick(selling string, tickIndex int64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get user shares and total shares
	totalShares := s.getLimitTotalSharesAtTickAtKey(selling, tickIndex, fillTranche)
	if fillTranche == placeTranche-1 {
		totalShares = totalShares.Add(s.getLimitTotalSharesAtTickAtKey(selling, tickIndex, placeTranche))
	}
	return totalShares
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order total shares for key %s", key)
	return tranche.TotalTokenIn
}

func (s *MsgServerTestSuite) getLimitFilledLiquidityAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order filled reserves for key %s", key)
	return tranche.ReservesTokenOut
}

func (s *MsgServerTestSuite) getLimitReservesAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Dec {
	pairId := s.app.DexKeeper.CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order reserves for key %s", key)
	return tranche.ReservesTokenIn
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOAToB(tick int64, tickLiqudity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice0To1(tick)

	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOAToB(tick int64, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice0To1(tick)

	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapAToB(tick int64, tickLiqudidty sdk.Dec, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice0To1(tick)

	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOBToA(tick int64, tickLiqudity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice1To0(tick)

	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOBToA(tick int64, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice1To0(tick)

	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapBToA(tick int64, tickLiqudidty sdk.Dec, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	price := keeper.CalcPrice1To0(tick)

	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwapNoLO(price sdk.Dec, tickLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	return calculateSingleSwap(price, tickLiquidity, sdk.ZeroDec(), amountIn)
}

func calculateSingleSwapOnlyLO(price sdk.Dec, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	return calculateSingleSwap(price, sdk.ZeroDec(), tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwap(price sdk.Dec, tickLiquidity sdk.Dec, tickLimitOrderLiquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	// swap against CSMM liquidity
	amountLeft, amountOut := calculateSwap(price, tickLiquidity, amountIn)
	// fmt.Printf("left %s out %s\n", amountLeft, amountOut)

	// swap against limit orders
	if amountLeft.GT(sdk.ZeroDec()) {
		tmpAmountLeft, tmpAmountOut := calculateSwap(price, tickLimitOrderLiquidity, amountLeft)
		amountLeft = tmpAmountLeft
		amountOut = amountOut.Add(tmpAmountOut)
	}
	return amountLeft, amountOut
}

func calculateSwap(price sdk.Dec, liquidity sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	if tmpAmountOut := price.Mul(amountIn); tmpAmountOut.LT(liquidity) {
		// fmt.Printf("sufficient tmpOut %s\n", tmpAmountOut)
		// sufficient liquidity
		return sdk.ZeroDec(), tmpAmountOut
	} else {
		// only sufficient for part of amountIn
		tmpAmountIn := liquidity.Quo(price)
		// fmt.Printf("insufficient tmpIn %s\n", tmpAmountIn)
		return amountIn.Sub(tmpAmountIn), liquidity
	}
}

func (s *MsgServerTestSuite) calculateMultipleSwapsAToB(tickIndexes []int64, tickLiquidities []sdk.Dec, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice0To1(tickIndexes[i])
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOAToB(tickIndexes []int64, tickLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice0To1(tickIndexes[i])
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOAToB(tickIndexes []int64, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice0To1(tickIndexes[i])
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsBToA(tickIndexes []int64, tickLiquidities []sdk.Dec, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice1To0(tickIndexes[i])
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOBToA(tickIndexes []int64, tickLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice1To0(tickIndexes[i])
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOBToA(tickIndexes []int64, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	prices := make([]sdk.Dec, len(tickIndexes))
	for i := range prices {
		prices[i] = keeper.CalcPrice1To0(tickIndexes[i])
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLO(prices []sdk.Dec, tickLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	// zero array for tickLimitOrders
	tickLimitOrderLiquidities := make([]sdk.Dec, len(prices))
	for i := range tickLimitOrderLiquidities {
		tickLimitOrderLiquidities[i] = sdk.ZeroDec()
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLO(prices []sdk.Dec, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	// zero array for tickLimitOrders
	tickLiquidities := make([]sdk.Dec, len(prices))
	for i := range tickLiquidities {
		tickLiquidities[i] = sdk.ZeroDec()
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwaps(prices []sdk.Dec, tickLiquidities []sdk.Dec, tickLimitOrderLiquidities []sdk.Dec, amountIn sdk.Dec) (sdk.Dec, sdk.Dec) {
	amountLeft, amountOut := amountIn, sdk.ZeroDec()
	for i := 0; i < len(prices); i++ {
		tmpAmountLeft, tmpAmountOut := calculateSingleSwap(prices[i], tickLiquidities[i], tickLimitOrderLiquidities[i], amountLeft)
		amountLeft, amountOut = tmpAmountLeft, amountOut.Add(tmpAmountOut)
	}
	return amountLeft, amountOut
}

func (s *MsgServerTestSuite) addTickWithFee0Tokens(tickIndex int64, amountA int, amountB int) types.TickMap {

	tick := types.TickMap{
		PairId:    "TokenA/TokenB",
		TickIndex: tickIndex,
		TickData: &types.TickDataType{
			Reserve0AndShares: make([]*types.Reserve0AndSharesType, 1),
			Reserve1:          make([]sdk.Dec, 1),
		},
		LimitOrderTranche0To1: &types.LimitOrderTrancheTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitOrderTrancheTrancheIndexes{0, 0},
	}

	tick.TickData.Reserve0AndShares[0] = &types.Reserve0AndSharesType{NewDec(amountA), NewDec(amountA)}
	tick.TickData.Reserve1[0] = NewDec(amountB)

	s.app.DexKeeper.SetTickMap(s.ctx, "TokenA/TokenB", tick)
	return tick
}

func (s *MsgServerTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) (lowerTick types.TickMap, upperTick types.TickMap) {
	pairId := "TokenA<>TokenB"
	lowerTick = s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex-1)
	upperTick = s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex+1)
	priceCenter1To0 := keeper.CalcPrice0To1(tickIndex)
	amountADec := NewDec(amountA)
	amountBDec := NewDec(amountB)
	lowerTick.TickData.Reserve0AndShares[0].Reserve0 = amountADec
	lowerTick.TickData.Reserve0AndShares[0].TotalShares = keeper.CalcShares(amountADec, amountBDec, priceCenter1To0)
	upperTick.TickData.Reserve1[0] = amountBDec
	s.app.DexKeeper.SetTickMap(s.ctx, pairId, lowerTick)
	s.app.DexKeeper.SetTickMap(s.ctx, pairId, upperTick)
	return lowerTick, upperTick
}
