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

func (s *MsgServerTestSuite) fundAccountBalancesDec(account sdk.AccAddress, aBalance sdk.Dec, bBalance sdk.Dec) {
	aBalanceInt := sdk.NewIntFromBigInt(aBalance.BigInt())
	bBalanceInt := sdk.NewIntFromBigInt(bBalance.BigInt())
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalancesDec(account, aBalance, bBalance)
}

func (s *MsgServerTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.fundAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerTestSuite) fundAliceBalances(a int, b int) {
	s.fundAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerTestSuite) fundBobBalances(a int, b int) {
	s.fundAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalances(a int, b int) {
	s.fundAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundCarolBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.carol, a, b)
}

func (s *MsgServerTestSuite) fundDanBalances(a int, b int) {
	s.fundAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) fundDanBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.assertAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerTestSuite) assertAccountBalancesDec(
	account sdk.AccAddress,
	aBalance sdk.Dec,
	bBalance sdk.Dec,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA")
	aDec := sdk.NewDecFromBigIntWithPrec(aActual.Amount.BigInt(), 18)
	s.Assert().Equal(aBalance, aDec)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB")
	bDec := sdk.NewDecFromBigIntWithPrec(bActual.Amount.BigInt(), 18)
	s.Assert().Equal(bBalance, bDec)
}

func (s *MsgServerTestSuite) assertAliceBalances(a int, b int) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a int, b int) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalances(a int, b int) {
	s.assertAccountBalances(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertCarolBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a int, b int) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a int, b int) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) alicePlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placesLimitOrder(s.alice, wantsToken, tick, amountIn)
}

func (s *MsgServerTestSuite) bobPlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placesLimitOrder(s.bob, wantsToken, tick, amountIn)
}

func (s *MsgServerTestSuite) carolPlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placesLimitOrder(s.carol, wantsToken, tick, amountIn)
}

func (s *MsgServerTestSuite) danPlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placesLimitOrder(s.dan, wantsToken, tick, amountIn)
}

func (s *MsgServerTestSuite) placesLimitOrder(account sdk.AccAddress, wantsToken string, tick int, amountIn int) {
	var tokenIn string
	if wantsToken == "TokenA" {
		tokenIn = "TokenB"
	} else {
		tokenIn = "TokenA"
	}
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
	tickIndicies := make([]int64, len(deposits))
	feeIndexes := make([]uint64, len(deposits))
	for i, e := range deposits {
		amountsA[i] = e.AmountA
		amountsB[i] = e.AmountB
		tickIndicies[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
	}

	_, err := s.msgServer.Deposit(s.goCtx, &types.MsgDeposit{
		Creator:     account.String(),
		Receiver:    account.String(),
		TokenA:      "TokenA",
		TokenB:      "TokenB",
		AmountsA:    amountsA,
		AmountsB:    amountsB,
		TickIndexes: tickIndicies,
		FeeIndexes:  feeIndexes,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceCancelsLimitOrder(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitOrder(s.alice, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) bobCancelsLimitOrder(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitOrder(s.bob, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) carolCancelsLimitOrder(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitOrder(s.carol, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) danCancelsLimitOrder(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitOrder(s.dan, keyToken, tick, key, sharesOut)
}

func (s *MsgServerTestSuite) cancelsLimitOrder(account sdk.AccAddress, keyToken string, tick int, key int, sharesOut int) {
	sharesOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(sharesOut)))
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  keyToken,
		Key:       uint64(key),
		SharesOut: sharesOutDec,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) alicePlacesSwapOrder(wantsToken string, amountIn int, minOut int) {
	s.placesSwapOrder(s.alice, wantsToken, amountIn, minOut)
}

func (s *MsgServerTestSuite) bobPlacesSwapOrder(wantsToken string, amountIn int, minOut int) {
	s.placesSwapOrder(s.bob, wantsToken, amountIn, minOut)
}

func (s *MsgServerTestSuite) carolPlacesSwapOrder(wantsToken string, amountIn int, minOut int) {
	s.placesSwapOrder(s.bob, wantsToken, amountIn, minOut)
}

func (s *MsgServerTestSuite) danPlacesSwapOrder(wantsToken string, amountIn int, minOut int) {
	s.placesSwapOrder(s.bob, wantsToken, amountIn, minOut)
}

func (s *MsgServerTestSuite) placesSwapOrder(account sdk.AccAddress, wantsToken string, amountIn int, minOut int) {
	var tokenIn string
	if wantsToken == "TokenA" {
		tokenIn = "TokenB"
	} else {
		tokenIn = "TokenA"
	}
	amountInDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(amountIn)))
	minOutDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(uint64(minOut)))
	_, err := s.msgServer.Swap(s.goCtx, &types.MsgSwap{
		Creator:  account.String(),
		Receiver: account.String(),
		TokenA:   "TokenA",
		TokenB:   "TokenB",
		TokenIn:  tokenIn,
		AmountIn: amountInDec,
		MinOut:   minOutDec,
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) aliceWithdrawsFilledLimitOrder(withdrawToken string, tick int) {
	s.withdrawsFilledLimitOrder(s.alice, withdrawToken, tick)
}

func (s *MsgServerTestSuite) bobWithdrawsFilledLimitOrder(withdrawToken string, tick int) {
	s.withdrawsFilledLimitOrder(s.bob, withdrawToken, tick)
}

func (s *MsgServerTestSuite) carolWithdrawsFilledLimitOrder(withdrawToken string, tick int) {
	s.withdrawsFilledLimitOrder(s.carol, withdrawToken, tick)
}

func (s *MsgServerTestSuite) danWithdrawsFilledLimitOrder(withdrawToken string, tick int) {
	s.withdrawsFilledLimitOrder(s.dan, withdrawToken, tick)
}

func (s *MsgServerTestSuite) withdrawsFilledLimitOrder(account sdk.AccAddress, withdrawToken string, tick int) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  withdrawToken,
		Key:       0,
	})
	s.Assert().Nil(err)
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

func NewWithdrawl(shares int64, tick int64, fee uint64) *Withdrawl {
	return &Withdrawl{
		Shares:    sdk.NewDec(shares),
		FeeIndex:  fee,
		TickIndex: tick,
	}
}

func (s *MsgServerTestSuite) aliceWithdraws(withdrawls ...*Withdrawl) error {
	sharesToRemove := make([]sdk.Dec, len(withdrawls))
	tickIndicies := make([]int64, len(withdrawls))
	feeIndexes := make([]uint64, len(withdrawls))
	for i, e := range withdrawls {
		sharesToRemove[i] = e.Shares
		tickIndicies[i] = e.TickIndex
		feeIndexes[i] = e.FeeIndex
	}

	_, err := s.msgServer.Withdrawl(s.goCtx, &types.MsgWithdrawl{
		Creator:        s.alice.String(),
		Receiver:       s.alice.String(),
		TokenA:         "TokenA",
		TokenB:         "TokenB",
		SharesToRemove: sharesToRemove,
		TickIndexes:    tickIndicies,
		FeeIndexes:     feeIndexes,
	})

	return err
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
	sharesExpected sdk.Dec) {

	s.assertAccountShares(s.alice, "TokenA/TokenB", tick, fee, sharesExpected)
}

func (s *MsgServerTestSuite) assertCurrentTicks(
	expected0To1 int64,
	expected1To0 int64,
) {
	tickMap, found := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA/TokenB")
	s.Assert().NotNil(found)
	s.Assert().Equal(expected0To1, tickMap.TokenPair.CurrentTick0To1)
	s.Assert().Equal(expected1To0, tickMap.TokenPair.CurrentTick1To0)
}

func (s *MsgServerTestSuite) assertTickCount(tickCount int64) {
	tickMap, found := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA/TokenB")
	s.Assert().NotNil(found)
	s.Assert().Equal(tickCount, tickMap.TotalTickCount)
}

func (s *MsgServerTestSuite) printTicks() {
	tickMap, _ := s.app.DexKeeper.GetPairMap(s.ctx, "TokenA/TokenB")
	fmt.Printf("\nTick0To1: %v, Tick1To0: %v", tickMap.TokenPair.CurrentTick0To1, tickMap.TokenPair.CurrentTick1To0)
}
