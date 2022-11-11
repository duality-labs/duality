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
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func (s *MsgServerLimitTestSuite) TestSingle() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)
}

func (s *MsgServerLimitTestSuite) TestMultiple() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 50)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 400)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 100)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(100),
	})
	s.Require().Nil(err)

	s.assertAliceBalances(100, 400)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 200)
}

func (s *MsgServerLimitTestSuite) TestDifferentReceiverAndCreator() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(100),
	})
	s.Require().Nil(err)

	s.assertAliceBalances(100, 500)
	s.assertBobBalances(100, 100)
	s.assertDexBalances(0, 100)
}

func (s *MsgServerLimitTestSuite) TestFailUnrecognizedToken() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenC",
		AmountIn:  NewDec(100),
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestFailInsufficientBalance() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.PlaceLimitOrder(s.goCtx, &types.MsgPlaceLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.bob.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		TokenIn:   "TokenB",
		AmountIn:  NewDec(1000),
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestMultiTickLimitOrder1to0WithWithdraw() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenA", 0, 25)
	s.alicePlacesLimitOrder("TokenA", -1, 25)
	s.bobPlacesSwapOrder("TokenB", 40, 30)

	s.assertAliceBalances(100, 450)
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))

	s.aliceWithdrawsFilledLimitOrder("TokenB", 0)

	s.assertAliceBalances(125, 450)
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))

	s.aliceWithdrawsFilledLimitOrder("TokenB", -1)

	s.assertAliceBalancesDec(sdk.MustNewDecFromStr("133.999460032398056116"), NewDec(450))
	s.assertBobBalancesDec(NewDec(60), sdk.MustNewDecFromStr("239.9985"))
}

func (s *MsgServerLimitTestSuite) TestMultiTickLimitOrder0to1WithWithdraw() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)
	s.alicePlacesLimitOrder("TokenB", 1, 25)
	s.bobPlacesSwapOrder("TokenA", 40, 30)

	s.assertBobBalancesDec(sdk.MustNewDecFromStr("140.001500000000000000"), NewDec(160))

	s.aliceWithdrawsFilledLimitOrder("TokenA", 0)

	s.assertAliceBalances(99950, 525)

	s.aliceWithdrawsFilledLimitOrder("TokenA", 1)

	// TODO: Figure out if this is correct... maybe fees are involved?
	// One would expect the output to be 539.99850015
	// 525 + (15 / 1.0001) = 539.99850015
	// Which gives an effective price of 1.66656666667
	// 15 / (534.000540032401944116 - 525) = 1.66656666667
	// not an integer tick!
	// log(1.66656666667) / log(1.0001) = 5107.91159823
	s.assertAliceBalancesDec(NewDec(99950), sdk.MustNewDecFromStr("534.000540032401944116"))
}

func (s *MsgServerLimitTestSuite) TestWithdrawFailsWhenNothingToWithdraw() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestFailsWhenWithdrawNotCalledByOwner() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)

	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.bob.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       0,
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestFailsWhenWrongKeyToken() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)

	// Errors because of wrong KeyToken
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenA",
		Key:       0,
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestFailsWhenWrongKey() {
	s.fundAliceBalances(100000, 500)
	s.fundBobBalances(100, 200)

	s.alicePlacesLimitOrder("TokenB", 0, 25)

	// errors because of wrong key
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   s.alice.String(),
		Receiver:  s.alice.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: 0,
		KeyToken:  "TokenB",
		Key:       1,
	})
	s.Require().Error(err)
}

func (s *MsgServerLimitTestSuite) TestCancelSingle() {
	s.fundAliceBalances(100, 500)

	s.assertDexBalances(0, 0)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertDexBalances(0, 50)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 50)

	s.assertAliceBalances(100, 500)
	s.assertDexBalances(0, 0)
}

func (s *MsgServerLimitTestSuite) TestCancelPartial() {
	s.fundAliceBalances(100, 500)

	s.assertDexBalances(0, 0)

	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 450)
	s.assertDexBalances(0, 50)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 25)

	s.assertAliceBalances(100, 475)
	s.assertDexBalances(0, 25)

	s.aliceCancelsLimitOrder("TokenB", 0, 0, 25)

	s.assertAliceBalances(100, 500)
	s.assertDexBalances(0, 0)
}

func (s *MsgServerLimitTestSuite) TestProgressiveLimitOrderFill() {
	s.fundAliceBalances(100, 500)
	s.fundBobBalances(100, 200)

	s.aliceDeposits(NewDeposit(0, 10, 0, 0))
	s.alicePlacesLimitOrder("TokenA", 0, 50)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(100, 200)
	s.assertDexBalances(0, 60)

	s.bobPlacesSwapOrder("TokenB", 10, 10)

	s.assertAliceBalances(100, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(10, 50)

	s.aliceWithdrawsFilledLimitOrder("TokenB", 0)

	// Limit order is filled progressively
	s.assertAliceBalances(102, 440)
	s.assertBobBalances(90, 210)
	s.assertDexBalances(8, 50)

	// TODO: How to verify current tick?
}

func TestMsgServerLimitTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerLimitTestSuite))
}

type MsgServerLimitTestSuite struct {
	suite.Suite

	app         *dualityapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
	alice       sdk.AccAddress
	bob         sdk.AccAddress
	goCtx       context.Context
}

func (s *MsgServerLimitTestSuite) SetupTest() {
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

	// Set Fee List
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{0, 1})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{1, 2})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{2, 3})
	app.DexKeeper.AppendFeeList(ctx, types.FeeList{3, 4})

	s.app = app
	s.msgServer = keeper.NewMsgServerImpl(app.DexKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
}

func (s *MsgServerLimitTestSuite) fundAccountBalancesDec(account sdk.AccAddress, aBalance sdk.Dec, bBalance sdk.Dec) {
	aBalanceInt := sdk.NewIntFromBigInt(aBalance.BigInt())
	bBalanceInt := sdk.NewIntFromBigInt(bBalance.BigInt())
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := simapp.FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Require().NoError(err)
	s.assertAccountBalancesDec(account, aBalance, bBalance)
}

func (s *MsgServerLimitTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.fundAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerLimitTestSuite) fundAliceBalances(a int, b int) {
	s.fundAccountBalances(s.alice, a, b)
}

func (s *MsgServerLimitTestSuite) fundAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerLimitTestSuite) fundBobBalances(a int, b int) {
	s.fundAccountBalances(s.bob, a, b)
}

func (s *MsgServerLimitTestSuite) fundBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.fundAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerLimitTestSuite) assertAccountBalances(account sdk.AccAddress, aBalance int, bBalance int) {
	s.assertAccountBalancesDec(account, NewDec(aBalance), NewDec(bBalance))
}

func (s *MsgServerLimitTestSuite) assertAccountBalancesDec(
	account sdk.AccAddress,
	aBalance sdk.Dec,
	bBalance sdk.Dec,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA")
	aDec := sdk.NewDecFromBigIntWithPrec(aActual.Amount.BigInt(), 18)
	s.Require().Equal(aBalance, aDec)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB")
	bDec := sdk.NewDecFromBigIntWithPrec(bActual.Amount.BigInt(), 18)
	s.Require().Equal(bBalance, bDec)
}

func (s *MsgServerLimitTestSuite) assertAliceBalances(a int, b int) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerLimitTestSuite) assertBobBalances(a int, b int) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerLimitTestSuite) assertAliceBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.alice, a, b)
}

func (s *MsgServerLimitTestSuite) assertBobBalancesDec(a sdk.Dec, b sdk.Dec) {
	s.assertAccountBalancesDec(s.bob, a, b)
}

func (s *MsgServerLimitTestSuite) assertDexBalances(a int, b int) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("dex"), a, b)
}

func (s *MsgServerLimitTestSuite) alicePlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placeLimitOrder(s.alice, wantsToken, tick, amountIn)
}

func (s *MsgServerLimitTestSuite) bobPlacesLimitOrder(wantsToken string, tick int, amountIn int) {
	s.placeLimitOrder(s.bob, wantsToken, tick, amountIn)
}

func (s *MsgServerLimitTestSuite) placeLimitOrder(account sdk.AccAddress, wantsToken string, tick int, amountIn int) {
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
	s.Require().Nil(err)
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

func (s *MsgServerLimitTestSuite) aliceDeposits(deposits ...*Deposit) {
	s.deposits(s.alice, deposits...)
}

func (s *MsgServerLimitTestSuite) bobDeposits(deposits ...*Deposit) {
	s.deposits(s.bob, deposits...)
}

func (s *MsgServerLimitTestSuite) deposits(account sdk.AccAddress, deposits ...*Deposit) {
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
	s.Require().Nil(err)
}

func (s *MsgServerLimitTestSuite) aliceCancelsLimitOrder(keyToken string, tick int, key int, sharesOut int) {
	s.cancelsLimitOrder(s.alice, keyToken, tick, key, sharesOut)
}

func (s *MsgServerLimitTestSuite) cancelsLimitOrder(account sdk.AccAddress, keyToken string, tick int, key int, sharesOut int) {
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
	s.Require().Nil(err)
}

func (s *MsgServerLimitTestSuite) bobPlacesSwapOrder(wantsToken string, amountIn int, minOut int) {
	s.placesSwapOrder(s.bob, wantsToken, amountIn, minOut)
}

func (s *MsgServerLimitTestSuite) placesSwapOrder(account sdk.AccAddress, wantsToken string, amountIn int, minOut int) {
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
	s.Require().Nil(err)
}

func (s *MsgServerLimitTestSuite) aliceWithdrawsFilledLimitOrder(withdrawToken string, tick int) {
	s.withdrawsFilledLimitOrder(s.alice, withdrawToken, tick)
}

func (s *MsgServerLimitTestSuite) withdrawsFilledLimitOrder(account sdk.AccAddress, withdrawToken string, tick int) {
	_, err := s.msgServer.WithdrawFilledLimitOrder(s.goCtx, &types.MsgWithdrawFilledLimitOrder{
		Creator:   account.String(),
		Receiver:  account.String(),
		TokenA:    "TokenA",
		TokenB:    "TokenB",
		TickIndex: int64(tick),
		KeyToken:  withdrawToken,
		Key:       0,
	})
	s.Require().Nil(err)
}

func (s *MsgServerLimitTestSuite) traceBalances() {
	aliceA := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenA")
	aliceB := s.app.BankKeeper.GetBalance(s.ctx, s.alice, "TokenB")
	bobA := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenA")
	bobB := s.app.BankKeeper.GetBalance(s.ctx, s.bob, "TokenB")
	fmt.Printf("Alice: %+v %+v, Bob: %+v %+v", aliceA, aliceB, bobA, bobB)
}
