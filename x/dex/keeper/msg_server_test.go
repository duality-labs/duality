package keeper_test

import (
	"context"
	"fmt"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/keeper"
	. "github.com/NicholasDotSol/duality/x/dex/keeper"
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
	feeTiers    []types.FeeTier
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

func (s *MsgServerTestSuite) assertAccountBalancesEpsilon(
	account sdk.AccAddress,
	aBalance sdk.Int,
	bBalance sdk.Int,
) {
	// Checks that user account balances are within 1 of arithmetically calculated amount
	// and are strictly less that expected amount
	allowableError := sdk.NewInt(2)
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount

	aBalanceDelta := aBalance.Sub(aActual)

	s.Assert().True(aBalanceDelta.Abs().LTE(allowableError), "expected %s != actual %s", aBalance, aActual)
	s.Assert().True(aActual.LTE(aBalance), "Actual balance A (%s), is greater than expected balance (%s)", aActual, aBalance)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount
	bBalanceDelta := bBalance.Sub(bActual)

	s.Assert().True(bBalanceDelta.Abs().LTE(allowableError), "expected %s != actual %s", bBalance, bActual)
	s.Assert().True(bActual.LTE(bBalance), "Actual balance A (%s), is greater than expected balance (%s)", bActual, bBalance)
}

func (s *MsgServerTestSuite) assertDexBalancesEpsilon(
	aBalance sdk.Int,
	bBalance sdk.Int,
) {
	// Checks that Dex account balances are within 1 of arithmetically calculated amount
	// and are strictly greater that expected amount
	allowableError := sdk.NewInt(2)
	aActual := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenA").Amount

	aBalanceDelta := aBalance.Sub(aActual)

	s.Assert().True(aBalanceDelta.Abs().LTE(allowableError), "expected %s != actual %s", aBalance, aActual)
	s.Assert().True(aActual.GTE(aBalance), "Actual balance A (%s), is greater than expected balance (%s)", aActual, aBalance)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("dex"), "TokenB").Amount
	bBalanceDelta := bBalance.Sub(bActual)

	s.Assert().True(bBalanceDelta.Abs().LTE(allowableError), "expected %s != actual %s", bBalance, bActual)
	s.Assert().True(bActual.GTE(bBalance), "Actual balance A (%s), is less than expected balance (%s)", bActual, bBalance)
}

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

func (s *MsgServerTestSuite) assertAliceBalancesEpsilon(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a int64, b int64) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesEpsilon(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.bob, a, b)
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

func (s *MsgServerTestSuite) assertCarolBalancesEpsilon(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a int64, b int64) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesEpsilon(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDexBalances(a int64, b int64) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerTestSuite) assertDexBalancesInt(a sdk.Int, b sdk.Int) {
	s.assertAccountBalancesInt(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
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
	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		TokenIn:   tokenIn,
		AmountIn:  sdk.NewInt(int64(amountIn)),
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

type Deposit struct {
	AmountA   sdk.Int
	AmountB   sdk.Int
	TickIndex int64
	FeeIndex  uint64
}

func NewDeposit(amountA int, amountB int, tickIndex int, feeIndex int) *Deposit {
	return &Deposit{
		AmountA:   sdk.NewInt(int64(amountA)),
		AmountB:   sdk.NewInt(int64(amountB)),
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
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
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
	amountsA := make([]sdk.Int, len(deposits))
	amountsB := make([]sdk.Int, len(deposits))
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
	amountsA []sdk.Int
	amountsB []sdk.Int
}

func (s *MsgServerTestSuite) withdraws(account sdk.AccAddress, withdrawls ...*Withdrawl) error {
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

func (s *MsgServerTestSuite) aliceCancelsLimitSell(keyToken string, tick int, key int) {
	s.cancelsLimitSell(s.alice, keyToken, tick, key)
}

func (s *MsgServerTestSuite) bobCancelsLimitSell(keyToken string, tick int, key int) {
	s.cancelsLimitSell(s.bob, keyToken, tick, key)
}

func (s *MsgServerTestSuite) carolCancelsLimitSell(keyToken string, tick int, key int) {
	s.cancelsLimitSell(s.carol, keyToken, tick, key)
}

func (s *MsgServerTestSuite) danCancelsLimitSell(keyToken string, tick int, key int) {
	s.cancelsLimitSell(s.dan, keyToken, tick, key)
}

func (s *MsgServerTestSuite) aliceCancelsLimitSellFails(keyToken string, tick int, key int, expectedErr error) {
	s.cancelsLimitSellFails(s.alice, keyToken, tick, key, expectedErr)
}

func (s *MsgServerTestSuite) bobCancelsLimitSellFails(keyToken string, tick int, key int, expectedErr error) {
	s.cancelsLimitSellFails(s.bob, keyToken, tick, key, expectedErr)
}

func (s *MsgServerTestSuite) carolCancelsLimitSellFails(keyToken string, tick int, key int, expectedErr error) {
	s.cancelsLimitSellFails(s.carol, keyToken, tick, key, expectedErr)
}

func (s *MsgServerTestSuite) danCancelsLimitSellFails(keyToken string, tick int, key int, expectedErr error) {
	s.cancelsLimitSellFails(s.dan, keyToken, tick, key, expectedErr)
}

func (s *MsgServerTestSuite) cancelsLimitSell(account sdk.AccAddress, selling string, tick int, key int) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(key),
	})
	s.Assert().Nil(err)
}

func (s *MsgServerTestSuite) cancelsLimitSellFails(account sdk.AccAddress, selling string, tick int, key int, expectedErr error) {
	_, err := s.msgServer.CancelLimitOrder(s.goCtx, &types.MsgCancelLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(key),
	})
	s.Assert().ErrorIs(err, expectedErr)
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

func (s *MsgServerTestSuite) withdrawLimitSellFails(account sdk.AccAddress, expectedErr error, selling string, tick int, tranche int) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  selling,
		Key:       uint64(tranche),
	})
	s.Assert().ErrorIs(err, expectedErr)
}

func (s *MsgServerTestSuite) aliceWithdrawLimitSellFails(expectedErr error, selling string, tick int, tranche int) {
	s.withdrawLimitSellFails(s.alice, expectedErr, selling, tick, tranche)
}

func (s *MsgServerTestSuite) bobWithdrawLimitSellFails(expectedErr error, selling string, tick int, tranche int) {
	s.withdrawLimitSellFails(s.bob, expectedErr, selling, tick, tranche)
}

func (s *MsgServerTestSuite) carolWithdrawLimitSellFails(expectedErr error, selling string, tick int, tranche int) {
	s.withdrawLimitSellFails(s.carol, expectedErr, selling, tick, tranche)
}

func (s *MsgServerTestSuite) danWithdrawLimitSellFails(expectedErr error, selling string, tick int, tranche int) {
	s.withdrawLimitSellFails(s.dan, expectedErr, selling, tick, tranche)
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

func (s *MsgServerTestSuite) assertCurrentTicks(
	expected1To0 int64,
	expected0To1 int64,
) {
	tick, found := s.app.DexKeeper.GetTradingPair(s.ctx, "TokenA<>TokenB")
	s.Assert().NotNil(found)
	s.Assert().Equal(expected1To0, tick.CurrentTick1To0)
	s.Assert().Equal(expected0To1, tick.CurrentTick0To1)
}

func (s *MsgServerTestSuite) assertCurr0To1(curr0To1Expected int64) {
	pairId := CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetTradingPair(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	curr0To1Actual := pair.CurrentTick0To1
	s.Assert().Equal(curr0To1Expected, curr0To1Actual)
}

func (s *MsgServerTestSuite) assertCurr1To0(curr1To0Expected int64) {
	pairId := CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetTradingPair(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	curr1to0Actual := pair.CurrentTick1To0
	s.Assert().Equal(curr1To0Expected, curr1to0Actual)
}

func (s *MsgServerTestSuite) assertMinTick(minTickExpected int64) {
	pairId := CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetTradingPair(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	minTickActual := pair.MinTick
	s.Assert().Equal(minTickExpected, minTickActual)
}

func (s *MsgServerTestSuite) assertMaxTick(maxTickExpected int64) {
	pairId := CreatePairId("TokenA", "TokenB")
	pair, pairFound := s.app.DexKeeper.GetTradingPair(s.ctx, pairId)
	if !pairFound {
		s.Require().Fail("Invalid GetPair in assertCurr0to1")
	}

	maxTickActual := pair.MaxTick
	s.Assert().Equal(maxTickExpected, maxTickActual)
}

func (s *MsgServerTestSuite) printTicks() {
	tickMap, _ := s.app.DexKeeper.GetTradingPair(s.ctx, "TokenA<>TokenB")
	fmt.Printf("\nTick0To1: %v, Tick1To0: %v", tickMap.CurrentTick0To1, tickMap.CurrentTick1To0)
}

func (s *MsgServerTestSuite) assertLiquidityAtTickInt(amountA sdk.Int, amountB sdk.Int, tickIndex int64, feeIndex uint64) {
	pairId := CreatePairId("TokenA", "TokenB")
	fee := s.feeTiers[feeIndex].Fee
	lowerTick, lowerTickFound := s.app.DexKeeper.GetTick(s.ctx, pairId, tickIndex-fee)
	liquidityA, liquidityB := sdk.ZeroInt(), sdk.ZeroInt()
	if lowerTickFound {
		liquidityA = lowerTick.TickData.Reserve0[feeIndex]
	} else {
		// noop, since liquidity was set to 0 already
		// s.Require().Fail("Invalid tick %d and fee %d", tickIndex, fee)
	}
	upperTick, upperTickFound := s.app.DexKeeper.GetTick(s.ctx, pairId, tickIndex+fee)
	if upperTickFound {
		liquidityB = upperTick.TickData.Reserve1[feeIndex]
	} else {
		// noop, since liquidity was set to 0 already
		// s.Require().Fail("Invalid tick %d and fee %d", tickIndex, fee)
	}

	s.Assert().Equal(amountA, liquidityA)
	s.Assert().Equal(amountB, liquidityB)
}

func (s *MsgServerTestSuite) assertPoolLiquidity(amountA int, amountB int, tickIndex int64, feeIndex uint64) {
	s.assertLiquidityAtTickInt(sdk.NewInt(int64(amountA)), sdk.NewInt(int64(amountB)), tickIndex, feeIndex)
}

func (s *MsgServerTestSuite) assertNoLiquidityAtTick(tickIndex int64, feeIndex uint64) {
	pairId := CreatePairId("TokenA", "TokenB")
	fee := s.feeTiers[feeIndex].Fee

	lowerTick, lowerTickFound := s.app.DexKeeper.GetTick(s.ctx, pairId, tickIndex-fee)
	if !lowerTickFound {
		s.Assert().True(!lowerTickFound)
		return
	}
	// in case tick was initialized, assert no liquidity in it
	amtA := sdk.NewInt(0)
	liquidityA := lowerTick.TickData.Reserve0[feeIndex]
	s.Assert().Equal(amtA, liquidityA)

	upperTick, upperTickFound := s.app.DexKeeper.GetTick(s.ctx, pairId, tickIndex+fee)
	if !upperTickFound {
		s.Assert().True(!upperTickFound)
		return
	}
	// in case tick was initialized, assert no liquidity in it
	amtB := sdk.NewInt(0)
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
	amt := sdk.NewInt(int64(amount))
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
	pairId := CreatePairId("TokenA", "TokenB")

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
	userRatio := userShares.ToDec().QuoInt(totalShares)
	// assert enough liq
	userLiquidity := userRatio.MulInt64(int64(amount)).TruncateInt()
	s.assertLimitLiquidityAtTick(selling, tickIndex, userLiquidity.Int64())
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTickInt(selling string, tickIndex int64, amount sdk.Int) {
	pairId := CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get liquidity from fill
	liquidity := s.getLimitReservesAtTickAtKey(selling, tickIndex, fillTranche)
	// if fill == place - 1, get liquidity from place
	if fillTranche == placeTranche-1 {
		liquidity = liquidity.Add(s.getLimitReservesAtTickAtKey(selling, tickIndex, placeTranche))
	}

	s.Assert().True(amount.Equal(liquidity), "Incorrect liquidity: expected %s, have %s", amount.String(), liquidity.String())
}

func (s *MsgServerTestSuite) assertLimitLiquidityAtTick(selling string, tickIndex int64, amount int64) {
	s.assertLimitLiquidityAtTickInt(selling, tickIndex, sdk.NewInt(amount))
}

func (s *MsgServerTestSuite) assertFillAndPlaceTrancheKeys(selling string, tickIndex int64, expectedFill uint64, expectedPlace uint64) {
	pairId := CreatePairId("TokenA", "TokenB")
	fill, place := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	s.Assert().Equal(expectedFill, fill)
	s.Assert().Equal(expectedPlace, place)
}

func (s *MsgServerTestSuite) getFillAndPlaceTrancheKeys(selling string, pairId string, tickIndex int64) (uint64, uint64) {
	// grab current fill and place tranches
	tick, tickFound := s.app.DexKeeper.GetTick(s.ctx, pairId, tickIndex)
	s.Assert().True(tickFound, "Invalid tickIndex for pair %s", pairId)

	// handle correct limit order pool
	if selling == "TokenA" {
		return tick.LimitOrderTranche0To1.FillTrancheIndex, tick.LimitOrderTranche0To1.PlaceTrancheIndex
	} else {
		return tick.LimitOrderTranche1To0.FillTrancheIndex, tick.LimitOrderTranche1To0.PlaceTrancheIndex
	}
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTick(account sdk.AccAddress, selling string, tickIndex int64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get user shares and total shares
	userShares := s.getLimitUserSharesAtTickAtKey(account, selling, tickIndex, fillTranche)
	if fillTranche == placeTranche-1 {
		userShares = userShares.Add(s.getLimitUserSharesAtTickAtKey(account, selling, tickIndex, placeTranche))
	}
	return userShares
}

func (s *MsgServerTestSuite) getLimitUserSharesAtTickAtKey(account sdk.AccAddress, selling string, tickIndex int64, key uint64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	userShares, userSharesFound := s.app.DexKeeper.GetLimitOrderTrancheUser(s.ctx, pairId, tickIndex, selling, key, account.String())
	s.Assert().True(userSharesFound, "Failed to get limit order user shares for key %s", key)
	return userShares.SharesOwned
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTick(selling string, tickIndex int64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	fillTranche, placeTranche := s.getFillAndPlaceTrancheKeys(selling, pairId, tickIndex)
	// get user shares and total shares
	totalShares := s.getLimitTotalSharesAtTickAtKey(selling, tickIndex, fillTranche)
	if fillTranche == placeTranche-1 {
		totalShares = totalShares.Add(s.getLimitTotalSharesAtTickAtKey(selling, tickIndex, placeTranche))
	}
	return totalShares
}

func (s *MsgServerTestSuite) getLimitTotalSharesAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order total shares for key %s", key)
	return tranche.TotalTokenIn
}

func (s *MsgServerTestSuite) getLimitFilledLiquidityAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order filled reserves for key %s", key)
	return tranche.ReservesTokenOut
}

func (s *MsgServerTestSuite) getLimitReservesAtTickAtKey(selling string, tickIndex int64, key uint64) sdk.Int {
	pairId := CreatePairId("TokenA", "TokenB")
	// grab fill tranche reserves and shares
	tranche, found := s.app.DexKeeper.GetLimitOrderTranche(s.ctx, pairId, tickIndex, selling, key)
	s.Assert().True(found, "Failed to get limit order reserves for key %s", key)
	return tranche.ReservesTokenIn
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOAToB(tick int64, tickLiqudity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice0To1(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOAToB(tick int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice0To1(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapAToB(tick int64, tickLiqudidty int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice0To1(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapNoLOBToA(tick int64, tickLiqudity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice1To0(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwapNoLO(price, tickLiqudity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapOnlyLOBToA(tick int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice1To0(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwapOnlyLO(price, tickLimitOrderLiquidity, amountIn)
}

func (s *MsgServerTestSuite) calculateSingleSwapBToA(tick int64, tickLiqudidty int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	price, err := keeper.CalcPrice1To0(tick)
	if err != nil {
		panic(err)
	}

	return calculateSingleSwap(price, tickLiqudidty, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwapNoLO(price sdk.Dec, tickLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, tickLiquidity, 0, amountIn)
}

func calculateSingleSwapOnlyLO(price sdk.Dec, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	return calculateSingleSwap(price, 0, tickLimitOrderLiquidity, amountIn)
}

func calculateSingleSwap(price sdk.Dec, tickLiquidity int64, tickLimitOrderLiquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
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

func calculateSwap(price sdk.Dec, liquidity int64, amountIn int64) (sdk.Int, sdk.Int) {
	amountInInt := sdk.NewInt(amountIn)
	liquidityInt := sdk.NewInt(liquidity)
	if tmpAmountOut := price.MulInt(amountInInt); tmpAmountOut.LT(liquidityInt.ToDec()) {
		// fmt.Printf("sufficient tmpOut %s\n", tmpAmountOut)
		// sufficient liquidity
		return sdk.ZeroInt(), tmpAmountOut.TruncateInt()
	} else {
		// only sufficient for part of amountIn
		tmpAmountIn := liquidityInt.ToDec().Quo(price).TruncateInt()
		// fmt.Printf("insufficient tmpIn %s\n", tmpAmountIn)
		return amountInInt.Sub(tmpAmountIn), liquidityInt
	}
}

func (s *MsgServerTestSuite) calculateMultipleSwapsAToB(tickIndexes []int64, tickLiquidities []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice0To1(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOAToB(tickIndexes []int64, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice0To1(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOAToB(tickIndexes []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice0To1(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsBToA(tickIndexes []int64, tickLiquidities []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice1To0(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLOBToA(tickIndexes []int64, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice1To0(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwapsNoLO(prices, tickLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLOBToA(tickIndexes []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	prices := make([]sdk.Dec, len(tickIndexes))
	var err error
	for i := range prices {
		prices[i], err = keeper.CalcPrice1To0(tickIndexes[i])
		if err != nil {
			panic(err)
		}
	}
	return s.calculateMultipleSwapsOnlyLO(prices, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsNoLO(prices []sdk.Dec, tickLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	// zero array for tickLimitOrders
	tickLimitOrderLiquidities := make([]int64, len(prices))
	for i := range tickLimitOrderLiquidities {
		tickLimitOrderLiquidities[i] = 0
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwapsOnlyLO(prices []sdk.Dec, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	// zero array for tickLimitOrders
	tickLiquidities := make([]int64, len(prices))
	for i := range tickLiquidities {
		tickLiquidities[i] = 0
	}
	return s.calculateMultipleSwaps(prices, tickLiquidities, tickLimitOrderLiquidities, amountIn)
}

func (s *MsgServerTestSuite) calculateMultipleSwaps(prices []sdk.Dec, tickLiquidities []int64, tickLimitOrderLiquidities []int64, amountIn int64) (sdk.Int, sdk.Int) {
	amountLeft, amountOut := sdk.NewInt(amountIn), sdk.ZeroInt()
	for i := 0; i < len(prices); i++ {
		tmpAmountLeft, tmpAmountOut := calculateSingleSwap(prices[i], tickLiquidities[i], tickLimitOrderLiquidities[i], amountLeft.Int64())
		amountLeft, amountOut = tmpAmountLeft, amountOut.Add(tmpAmountOut)
	}
	return amountLeft, amountOut
}

func (s *MsgServerTestSuite) addTickWithFee0Tokens(tickIndex int64, amountA int, amountB int) types.Tick {

	tick := types.Tick{
		PairId:    "TokenA/TokenB",
		TickIndex: tickIndex,
		TickData: &types.TickDataType{
			Reserve0: make([]sdk.Int, 1),
			Reserve1: make([]sdk.Int, 1),
		},
		LimitOrderTranche0To1: &types.LimitTrancheIndexes{0, 0},
		LimitOrderTranche1To0: &types.LimitTrancheIndexes{0, 0},
	}

	tick.TickData.Reserve0[0] = sdk.NewInt(int64(amountA))
	tick.TickData.Reserve1[0] = sdk.NewInt(int64(amountB))

	s.app.DexKeeper.SetTick(s.ctx, "TokenA/TokenB", tick)
	return tick
}

func (s *MsgServerTestSuite) setLPAtFee0Pool(tickIndex int64, amountA int, amountB int) (lowerTick types.Tick, upperTick types.Tick) {
	pairId := "TokenA<>TokenB"
	// sharesId := fmt.Sprintf("%s%st%df%d", "TokenA", "TokenB", tickIndex, 1)
	sharesId := CreateSharesId("TokenA", "TokenB", tickIndex, 0)
	lowerTick, err := s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex-1)
	s.Assert().NoError(err)
	upperTick, err = s.app.DexKeeper.GetOrInitTick(s.goCtx, pairId, tickIndex+1)
	s.Assert().NoError(err)
	priceCenter1To0, err := keeper.CalcPrice0To1(tickIndex)
	if err != nil {
		panic(err)
	}

	amountAInt := sdk.NewInt(int64(amountA))
	amountBInt := sdk.NewInt(int64(amountB))
	lowerTick.TickData.Reserve0[0] = amountAInt
	totalShares := keeper.CalcShares(amountAInt, amountBInt, priceCenter1To0).TruncateInt()
	s.app.DexKeeper.MintShares(s.ctx, s.alice, totalShares, sharesId)
	upperTick.TickData.Reserve1[0] = amountBInt
	s.app.DexKeeper.SetTick(s.ctx, pairId, lowerTick)
	s.app.DexKeeper.SetTick(s.ctx, pairId, upperTick)
	return lowerTick, upperTick
}
