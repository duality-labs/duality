package keeper_test

import (
	"context"
	"testing"

	"github.com/duality-labs/duality/x/mev/keeper"
	"github.com/duality-labs/duality/x/mev/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dualityapp "github.com/duality-labs/duality/app"
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
	types.RegisterQueryServer(queryHelper, app.MevKeeper)
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
	s.msgServer = keeper.NewMsgServerImpl(app.MevKeeper)
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.queryClient = queryClient
	s.alice = sdk.AccAddress([]byte("alice"))
	s.bob = sdk.AccAddress([]byte("bob"))
	s.carol = sdk.AccAddress([]byte("carol"))
	s.dan = sdk.AccAddress([]byte("dan"))
}

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, types.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amounts)
}

func (s *MsgServerTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance, bBalance int64) {
	aBalanceInt := sdk.NewInt(aBalance)
	bBalanceInt := sdk.NewInt(bBalance)
	balances := sdk.NewCoins(NewACoin(aBalanceInt), NewBCoin(bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalances(account, aBalance, bBalance)
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

func (s *MsgServerTestSuite) assertMEVDummyBalancesEpsilon(
	aBalance sdk.Int,
	bBalance sdk.Int,
) {
	// Checks that Dex account balances are within 1 of arithmetically calculated amount
	// and are strictly greater that expected amount
	allowableError := sdk.NewInt(2)
	aActual := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("mev"), "TokenA").Amount

	aBalanceDelta := aBalance.Sub(aActual)

	s.Assert().True(aBalanceDelta.Abs().LTE(allowableError), "expected %s != actual %s", aBalance, aActual)
	s.Assert().True(aActual.GTE(aBalance), "Actual balance A (%s), is greater than expected balance (%s)", aActual, aBalance)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, s.app.AccountKeeper.GetModuleAddress("mev"), "TokenB").Amount
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

func (s *MsgServerTestSuite) assertAliceBalances(a, b int64) {
	s.assertAccountBalances(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertAliceBalancesEpsilon(a, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.alice, a, b)
}

func (s *MsgServerTestSuite) assertBobBalances(a, b int64) {
	s.assertAccountBalances(s.bob, a, b)
}

func (s *MsgServerTestSuite) assertBobBalancesEpsilon(a, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.bob, a, b)
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

func (s *MsgServerTestSuite) assertCarolBalancesEpsilon(a, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.carol, a, b)
}

func (s *MsgServerTestSuite) assertDanBalances(a, b int64) {
	s.assertAccountBalances(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertDanBalancesEpsilon(a, b sdk.Int) {
	s.assertAccountBalancesEpsilon(s.dan, a, b)
}

func (s *MsgServerTestSuite) assertMEVDummyBalances(a, b int64) {
	s.assertAccountBalances(s.app.AccountKeeper.GetModuleAddress("mev"), a, b)
}

func (s *MsgServerTestSuite) assertMEVDummyBalancesInt(a, b sdk.Int) {
	s.assertAccountBalancesInt(s.app.AccountKeeper.GetModuleAddress("mev"), a, b)
}

func (s *MsgServerTestSuite) msgSend(account sdk.AccAddress, amountIn int, tokenIn string) {
	amountInInt := sdk.NewInt(int64(amountIn))
	_, err := s.msgServer.Send(s.goCtx, &types.MsgSend{
		Creator:  account.String(),
		AmountIn: amountInInt,
		TokenIn:  tokenIn,
	})
	s.Assert().Nil(err)
}

// UTILS //
func NewACoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenA", amt)
}

func NewBCoin(amt sdk.Int) sdk.Coin {
	return sdk.NewCoin("TokenB", amt)
}

// TESTS //

func (s *MsgServerTestSuite) TestValidTransaction() {
	s.fundAliceBalances(10, 10)
	s.assertAliceBalances(10, 10)
	s.msgSend(s.alice, 5, "TokenA")
	s.assertAliceBalances(5, 10)
	s.assertMEVDummyBalances(5, 0)

	//amt := sdk.Coins{
	//sdk.Coin{
	//"TokenA",
	//sdk.NewInt(5),
	//},
	//}

	// err := s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, s.alice, amt)
}
