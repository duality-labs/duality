package cli_test

import (
	"context"
	"fmt"
	"testing"

	dualityapp "github.com/NicholasDotSol/duality/app"
	"github.com/NicholasDotSol/duality/x/dex/client/cli"
	"github.com/NicholasDotSol/duality/x/dex/types"
	dexmoduletypes "github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type ClientTestSuite struct {
	suite.Suite
	app      *dualityapp.App
	ctx      sdk.Context
	goCtx    context.Context
	cfg      network.Config
	network  *network.Network
	feeTiers []types.FeeList
}

func (s *ClientTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	s.cfg = cfg
	s.network = network.New(s.T(), s.cfg)
	app := dualityapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
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
	s.ctx = ctx
	s.goCtx = sdk.WrapSDKContext(ctx)
	s.feeTiers = feeTiers

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

}
func (s *ClientTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	// This is important and must be called to ensure other tests can create
	// a network!
	s.network.Cleanup()
}
func (s *ClientTestSuite) fundAccountBalances(account sdk.AccAddress, aBalance int64, bBalance int64) {
	aBalanceInt := sdk.NewInt(aBalance)
	bBalanceInt := sdk.NewInt(bBalance)
	balances := sdk.NewCoins(sdk.NewCoin("TokenA", aBalanceInt), sdk.NewCoin("TokenB", bBalanceInt))
	err := FundAccount(s.app.BankKeeper, s.ctx, account, balances)
	s.Assert().NoError(err)
	s.assertAccountBalances(account, aBalance, bBalance)
}

func (s *ClientTestSuite) assertAccountBalances(
	account sdk.AccAddress,
	aBalance int64,
	bBalance int64,
) {
	aActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenA").Amount.Int64()

	s.Assert().Equal(aActual, aBalance, "expected %s != actual %s", aBalance, aBalance)

	bActual := s.app.BankKeeper.GetBalance(s.ctx, account, "TokenB").Amount.Int64()
	s.Assert().Equal(bActual, bBalance, "expected %s != actual %s", bBalance, bBalance)
}

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, dexmoduletypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, dexmoduletypes.ModuleName, addr, amounts)
}

func (s ClientTestSuite) TestDepositCmd() {
	val := s.network.Validators[0]
	s.fundAccountBalances(val.Address, 10, 10)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			//[receiver] [token-a] [token-b] [list of amount-0] [list of amount-1] [list of tick-index] [list of fee]
			"Deposit with invalid balane",
			[]string{
				fmt.Sprintf("%s", val.Address.String()),
				fmt.Sprintf("%s", "TokenA"),
				fmt.Sprintf("%s", "TokenB"),
				fmt.Sprintf("%s", "15a"),
				fmt.Sprintf("%s", "15"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=test", flags.FlagKeyringBackend),
				//fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				//fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			true, &sdk.TxResponse{}, 0,
		},
		{
			//[receiver] [token-a] [token-b] [list of amount-0] [list of amount-1] [list of tick-index] [list of fee]
			"Deposit with sufficient balane",
			[]string{
				fmt.Sprintf("%s", val.Address.String()),
				fmt.Sprintf("%s", "TokenA"),
				fmt.Sprintf("%s", "TokenB"),
				fmt.Sprintf("%s", "5"),
				fmt.Sprintf("%s", "5"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=test", flags.FlagKeyringBackend),
				//fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				//fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.CmdDeposit()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.T().Log(out)
			fmt.Println(tc)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}
	s.T().Log(s.app.DexKeeper.GetAllTickMap(s.ctx))
}

func (s ClientTestSuite) TestPlaceLimitOrderCmd() {
	val := s.network.Validators[0]
	s.fundAccountBalances(val.Address, 10, 10)

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			//[place-limit-order [receiver] [token-a] [token-b] [tick-index] [token-in] [amount-in]"
			"Place Limit Order with sufficient balane",
			[]string{
				fmt.Sprintf("%s", val.Address.String()),
				fmt.Sprintf("%s", "TokenA"),
				fmt.Sprintf("%s", "TokenB"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("%s", "0"),
				fmt.Sprintf("%s", "5"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				// common args
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=test", flags.FlagKeyringBackend),
				//fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				//fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.CmdPlaceLimitOrder()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.T().Log(out)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err, out.String())
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code, out.String())
			}
		})
	}

}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
