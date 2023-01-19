package cli_test

import (
	"fmt"
	"testing"

	network "github.com/NicholasDotSol/duality/testutil/network"
	dexClient "github.com/NicholasDotSol/duality/x/dex/client/cli"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	addr1 sdk.AccAddress
	addr2 sdk.AccAddress
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	nw := network.NewCLITest(s.T())
	s.network = nw

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)

	info1, _, err := s.network.Validators[0].ClientCtx.Keyring.NewMnemonic("acc1", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	info2, _, err := s.network.Validators[0].ClientCtx.Keyring.NewMnemonic("acc2", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	pk := info1.GetPubKey()
	s.addr1 = sdk.AccAddress(pk.Address())
	pk = info2.GetPubKey()
	s.addr2 = sdk.AccAddress(pk.Address())

	// var commonFlags = []string{
	// 	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	// 	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	// 	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
	// 	fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
	// 	fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	// }

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	// args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0", "1"}, commonFlags...)
	// cmd := dexClient.CmdDeposit()
	// _, err = cli.ExecTestCLICmd(clientCtx, cmd, args)
	// require.NoError(s.T(), err)

	// args = append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "20", "TokenB", "10"}, commonFlags...)
	// cmd = dexClient.CmdPlaceLimitOrder()
	// _, err = cli.ExecTestCLICmd(clientCtx, cmd, args)
	// require.NoError(s.T(), err)

	s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr1, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))

	s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr2, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))

	// _, err = s.network.WaitForHeight(1)
	// s.Require().NoError(err)

	// cmd = bankClient.GetBalancesCmd()
	// out, err := cli.ExecTestCLICmd(clientCtx, cmd, []string{s.network.Validators[0].Address.String()})
	// fmt.Printf("out: %v\n", out)
}

func (s *IntegrationTestSuite) fundAccount(clientCtx client.Context, from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) {
	require := s.Require()

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
	}

	out, err := banktestutil.MsgSendExec(
		clientCtx,
		from,
		to,
		coins,
		commonFlags...,
	)
	require.NoError(err)

	var res sdk.TxResponse
	require.NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
	require.Zero(res.Code, res.RawLog)
}

func (s *IntegrationTestSuite) TestQueryCmdShowFeeTier() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	//clientCtx.OutputFormat = outputFormat
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.FeeTier
	}{
		{
			name:      "missing args",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name:      "too many args",
			args:      []string{"foo", "bar"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 2",
		},
		{
			name: "valid",
			args: []string{"0"},
			expOutput: types.FeeTier{
				Id:  0,
				Fee: 1,
			},
		},
		{
			name: "valid 2",
			args: []string{"1"},
			expOutput: types.FeeTier{
				Id:  1,
				Fee: 3,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowFeeTier()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetFeeTierResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res.FeeTier)
				require.Equal(s.T(), tc.expOutput, res.FeeTier)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListFeeTier() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.FeeTier
	}{
		{
			name: "valid",
			args: []string{},
			expOutput: []types.FeeTier{
				types.FeeTier{
					Id:  0,
					Fee: 1,
				},
				types.FeeTier{
					Id:  1,
					Fee: 3,
				},
				types.FeeTier{
					Id:  2,
					Fee: 5,
				},
				types.FeeTier{
					Id:  3,
					Fee: 10,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListFeeTier()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllFeeTierResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.FeeTier)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowTick() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	//clientCtx.OutputFormat = outputFormat

	Index3Price0to1 := sdk.MustNewDecFromStr("0.999700059990001500")
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.Tick
	}{
		{
			name: "valid",
			args: []string{`3`, "TokenA<>TokenB"},
			expOutput: types.Tick{
				PairId: &types.PairId{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TickIndex: 3,
				TickData: &types.TickDataType{
					Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
					Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(10), sdk.ZeroInt(), sdk.ZeroInt()},
				},
				LimitOrderTranche0To1: &types.LimitTrancheIndexes{
					PlaceTrancheIndex: 0,
					FillTrancheIndex:  0,
				},
				LimitOrderTranche1To0: &types.LimitTrancheIndexes{
					PlaceTrancheIndex: 0,
					FillTrancheIndex:  0,
				},
				Price0To1: &Index3Price0to1,
			},
		},
		{
			name:      "uninitialized tick",
			args:      []string{"1", "TokenA<>TokenB"},
			expErr:    true,
			expErrMsg: "key not found",
		},

		{
			name:      "pair not specified",
			args:      []string{"1"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 1",
		},

		{
			name:      "tick not specified",
			args:      []string{"TokenA<>TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 1",
		},

		{
			name:      "too many ticks",
			args:      []string{"1", "1"},
			expErr:    true,
			expErrMsg: "PairId does not conform to pattern",
		},

		{
			name:      "multiple pairIds",
			args:      []string{"TokenA<>TokenB", "TokenA<>stake"},
			expErr:    true,
			expErrMsg: "key not found",
		},

		{
			name:      "too many arguments",
			args:      []string{"1", " '-1' ", "TokenA<>TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 3",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowTick()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetTickResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res.Tick)
				require.Equal(s.T(), tc.expOutput, res.Tick)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListTick() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	Index3Price0to1 := sdk.MustNewDecFromStr("0.999700059990001500")
	IndexNeg3Price0to1 := sdk.MustNewDecFromStr("1.000300030001000000")
	Index20Price0to1 := sdk.MustNewDecFromStr("0.998002098460885074")
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.Tick
	}{
		{
			name: "valid",
			args: []string{},
			expOutput: []types.Tick{
				types.Tick{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TickIndex: -3,
					TickData: &types.TickDataType{
						Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(10), sdk.ZeroInt(), sdk.ZeroInt()},
						Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
					},
					LimitOrderTranche0To1: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					LimitOrderTranche1To0: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					Price0To1: &IndexNeg3Price0to1,
				},
				types.Tick{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TickIndex: 3,
					TickData: &types.TickDataType{
						Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
						Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.NewInt(10), sdk.ZeroInt(), sdk.ZeroInt()},
					},
					LimitOrderTranche0To1: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					LimitOrderTranche1To0: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					Price0To1: &Index3Price0to1,
				},
				types.Tick{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TickIndex: 20,
					TickData: &types.TickDataType{
						Reserve0: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
						Reserve1: []sdk.Int{sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()},
					},
					LimitOrderTranche0To1: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					LimitOrderTranche1To0: &types.LimitTrancheIndexes{
						PlaceTrancheIndex: 0,
						FillTrancheIndex:  0,
					},
					Price0To1: &Index20Price0to1,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListTick()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllTickResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.Tick)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowTradingPair() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.TradingPair
	}{
		{
			name: "valid",
			args: []string{"TokenA<>TokenB"},
			expOutput: types.TradingPair{
				PairId: &types.PairId{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				CurrentTick0To1: 3,
				CurrentTick1To0: -3,
				MaxTick:         20,
				MinTick:         -3,
			},
		},
		{
			name:      "invalid pair",
			args:      []string{"3"},
			expErr:    true,
			expErrMsg: "PairId does not conform to pattern",
		},
		{
			name:      "too many pairs",
			args:      []string{"TokenA<>TokenB", "TokenA<>stake"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received",
		},
		{
			name:      "no pair",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowTradingPair()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetTradingPairResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.TradingPair)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListTradingPair() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.TradingPair
	}{
		{
			name: "valid",
			args: []string{},
			expOutput: []types.TradingPair{
				types.TradingPair{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					CurrentTick0To1: 3,
					CurrentTick1To0: -3,
					MaxTick:         20,
					MinTick:         -3,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListTradingPair()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllTradingPairResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.TradingPair)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowUserPosition() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.UserPositions
	}{
		// FIX ME
		{
			name: "valid",
			args: []string{s.network.Validators[0].Address.String()},
			expOutput: types.UserPositions{
				PoolDeposits: []types.DepositRecord{
					types.DepositRecord{
						PairId: &types.PairId{
							Token0: "TokenA",
							Token1: "TokenB",
						},
						SharesOwned:     sdk.NewInt(20),
						CenterTickIndex: 0,
						LowerTickIndex:  -3,
						UpperTickIndex:  3,
						FeeIndex:        1,
					},
				},
				LimitOrders: []types.LimitOrderTrancheUser{types.LimitOrderTrancheUser{}},
			},
		},
		{
			name:      "invalid address",
			args:      []string{"0x0"},
			expErr:    true,
			expErrMsg: "invalid bech32 string length",
		},
		{
			name:      "too many addresses",
			args:      []string{s.network.Validators[0].Address.String(), s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received",
		},
		{
			name:      "no address",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
		{
			name: "empty valid",
			args: []string{s.addr1.String()},
			expOutput: types.UserPositions{
				PoolDeposits: []types.DepositRecord{},
				LimitOrders:  []types.LimitOrderTrancheUser{},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowUserPositions()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			fmt.Printf("out: %v\n", out)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetUserPositionsResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.UserPositions)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowTokens() {

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.Tokens
	}{
		{
			name:      "valid",
			args:      []string{"0"},
			expOutput: types.Tokens{Id: 0, Address: "TokenA"},
		},
		{
			name:      "invalid token",
			args:      []string{"3"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many tokens",
			args:      []string{"0", "1"},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received",
		},
		{
			name:      "no token",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 1 arg(s), received 0",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowTokens()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetTokensResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.Tokens)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListTokens() {
	val := s.network.Validators[0]

	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.Tokens
	}{
		{
			name: "valid",
			args: []string{},
			expOutput: []types.Tokens{
				types.Tokens{
					Id:      0,
					Address: "TokenA",
				},
				types.Tokens{
					Id:      1,
					Address: "TokenB",
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListTokens()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllTokensResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.Tokens)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.LimitOrderTranche
	}{
		{
			name: "valid",
			args: []string{"TokenA<>TokenB", "20", "TokenB", "0"},
			expOutput: types.LimitOrderTranche{
				PairId: &types.PairId{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:          "TokenB",
				TickIndex:        20,
				TrancheIndex:     0,
				ReservesTokenIn:  sdk.NewInt(10),
				ReservesTokenOut: sdk.NewInt(0),
				TotalTokenIn:     sdk.NewInt(10),
				TotalTokenOut:    sdk.NewInt(0),
			},
		},
		{
			name:      "invalid pair",
			args:      []string{"TokenC<>TokenB", "20", "TokenB", "1"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many parameters",
			args:      []string{"TokenA<>B", "20", "TokenB", "1", "10"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 5",
		},
		{
			name:      "no parameters",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 0",
		},
		{
			name:      "too few parameters",
			args:      []string{"TokenA<>B", "20", "TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 3",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowLimitOrderTranche()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetLimitOrderTrancheResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.LimitOrderTranche)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.LimitOrderTranche
	}{
		// show-limit-order-pool-total-shares-map [pairId] [tickIndex] [tokenIn] [TrancheIndex]
		{
			name: "valid",
			args: []string{},
			expOutput: []types.LimitOrderTranche{
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenA",
					TickIndex:        -3,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenB",
					TickIndex:        -3,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenA",
					TickIndex:        20,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenB",
					TickIndex:        20,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(10),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(10),
					TotalTokenOut:    sdk.NewInt(0),
				},
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenA",
					TickIndex:        3,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
				types.LimitOrderTranche{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:          "TokenB",
					TickIndex:        3,
					TrancheIndex:     0,
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListLimitOrderTranche()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllLimitOrderTrancheResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				fmt.Printf("res.LimitOrderTranche: %v\n", res.LimitOrderTranche)
				require.Equal(s.T(), tc.expOutput, res.LimitOrderTranche)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdShowLimitOrderTrancheUser() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.LimitOrderTrancheUser
	}{
		// "show-limit-order-pool-user-share-map [pairId] [tickIndex] [tokenIn] [trancheIndex] [address]"
		{
			name: "valid",
			args: []string{"TokenA<>TokenB", "20", "TokenB", "0", s.network.Validators[0].Address.String()},
			expOutput: types.LimitOrderTrancheUser{
				PairId: &types.PairId{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				Token:           "TokenB",
				TickIndex:       20,
				Count:           0,
				Address:         s.network.Validators[0].Address.String(),
				SharesOwned:     sdk.NewInt(10),
				SharesWithdrawn: sdk.NewInt(0),
				SharesCancelled: sdk.NewInt(0),
			},
		},
		{
			name:      "invalid pair",
			args:      []string{"TokenB<>TokenC", "20", "TokenB", "0", s.network.Validators[0].Address.String()},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many parameters",
			args:      []string{"TokenA<>TokenB", "20", "TokenB", "0", "1", s.network.Validators[0].Address.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 6",
		},
		{
			name:      "no parameters",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 0",
		},
		{
			name:      "too few parameters",
			args:      []string{"TokenA<>TokenB", "20", "TokenB", "0"},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 4",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowLimitOrderTrancheUser()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetLimitOrderTrancheUserResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.LimitOrderTrancheUser)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListLimitOrderTrancheUser() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.LimitOrderTrancheUser
	}{
		// show-limit-order-pool-total-shares-map [pairId] [tickIndex] [tokenIn] [TrancheIndex]
		{
			name: "valid",
			args: []string{},
			expOutput: []types.LimitOrderTrancheUser{
				types.LimitOrderTrancheUser{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					Token:           "TokenB",
					TickIndex:       20,
					Count:           0,
					Address:         s.network.Validators[0].Address.String(),
					SharesOwned:     sdk.NewInt(10),
					SharesWithdrawn: sdk.NewInt(0),
					SharesCancelled: sdk.NewInt(0),
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListLimitOrderTrancheUser()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllLimitOrderTrancheUserResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.LimitOrderTrancheUser)

			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryCmdListUserLimitOrders() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.LimitOrderTrancheUser
	}{
		{
			name: "valid",
			args: []string{s.network.Validators[0].Address.String()},
			expOutput: []types.LimitOrderTrancheUser{
				types.LimitOrderTrancheUser{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					Token:           "TokenB",
					TickIndex:       20,
					Count:           0,
					Address:         s.network.Validators[0].Address.String(),
					SharesOwned:     sdk.NewInt(10),
					SharesWithdrawn: sdk.NewInt(0),
					SharesCancelled: sdk.NewInt(0),
				},
			},
		},
	}

	fmt.Printf("testCases[0].args: %v\n", testCases[0].args)
	fmt.Printf("testCases[0].expOutput: %v\n", testCases[0].expOutput)
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListUserLimitOrders()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllUserLimitOrdersResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.LimitOrders)

			}
		})
	}
}
func (s *IntegrationTestSuite) TestQueryCmdListUserDeposits() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.DepositRecord
	}{
		{
			name: "valid",
			args: []string{s.network.Validators[0].Address.String()},
			expOutput: []types.DepositRecord{
				types.DepositRecord{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					SharesOwned:     sdk.NewInt(20),
					CenterTickIndex: 0,
					LowerTickIndex:  -3,
					UpperTickIndex:  3,
					FeeIndex:        1,
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListUserDeposits()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllUserDepositsResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				fmt.Printf("res.Deposits: %v\n", res.Deposits)
				require.Equal(s.T(), tc.expOutput, res.Deposits)

			}
		})
	}
}

// func (s *IntegrationTestSuite) TestQuery11CmdShowTokenMap() {
// 	val := s.network.Validators[0]
// 	clientCtx := val.ClientCtx
// 	testCases := []struct {
// 		name      string
// 		args      []string
// 		expErr    bool
// 		expErrMsg string
// 		expOutput types.TokenMap
// 	}{
// 		{
// 			name:      "valid",
// 			args:      []string{"TokenA"},
// 			expOutput: types.TokenMap{Address: "TokenA", Index: 0},
// 		},
// 		{
// 			name:      "invalid address",
// 			args:      []string{"TokenC"},
// 			expErr:    true,
// 			expErrMsg: "key not found",
// 		},
// 		{
// 			name:      "too many addresses",
// 			args:      []string{"TokenA", "TokenB"},
// 			expErr:    true,
// 			expErrMsg: "Error: accepts 1 arg(s), received 2",
// 		},
// 		{
// 			name:      "no token address",
// 			args:      []string{},
// 			expErr:    true,
// 			expErrMsg: "Error: accepts 1 arg(s), received 0",
// 		},
// 	}
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			cmd := dexClient.CmdShowTokenMap()
// 			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
// 			if tc.expErr {
// 				require.Error(s.T(), err)
// 				require.Contains(s.T(), out.String(), tc.expErrMsg)
// 			} else {
// 				require.NoError(s.T(), err)
// 				var res types.QueryGetTokenMapResponse
// 				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
// 				require.NotEmpty(s.T(), res)
// 				require.Equal(s.T(), tc.expOutput, res.TokenMap)
// 			}
// 		})
// 	}
// }

// func (s *IntegrationTestSuite) TestQuery12CmdListTokenMap() {
// 	val := s.network.Validators[0]
// 	clientCtx := val.ClientCtx
// 	testCases := []struct {
// 		name      string
// 		args      []string
// 		expErr    bool
// 		expErrMsg string
// 		expOutput []types.TokenMap
// 	}{
// 		{
// 			name: "valid",
// 			args: []string{},
// 			expOutput: []types.TokenMap{
// 				types.TokenMap{
// 					Address: "TokenA",
// 					Index:   0,
// 				},
// 				types.TokenMap{
// 					Address: "TokenB",
// 					Index:   1,
// 				},
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			cmd := dexClient.CmdListTokenMap()
// 			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
// 			if tc.expErr {
// 				require.Error(s.T(), err)
// 				require.Contains(s.T(), out.String(), tc.expErrMsg)
// 			} else {
// 				require.NoError(s.T(), err)
// 				var res types.QueryAllTokenMapResponse
// 				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
// 				require.NotEmpty(s.T(), res)
// 				require.Equal(s.T(), tc.expOutput, res.TokenMap)
// 			}
// 		})
// 	}
// }

func (s *IntegrationTestSuite) TestTxCmdDeposit() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			name:      "missing arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0"},
			expErr:    true,
			expErrMsg: "Error: accepts 7 arg(s), received 6",
		},
		{
			name:      "too many arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0", "0", s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 7 arg(s), received 8",
		},
		{
			name:     "valid",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0", "0"},
			errInRes: false,
		},
		{
			name:     "valid: multiple case",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "0,0", "10,10", "25,25", "1,1"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdDeposit()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestTx2CmdWithdraw() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	//Deposit Funds
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0", "0"}, commonFlags...)
	cmd := dexClient.CmdDeposit()
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "withdrawl [receiver] [token-a] [token-b] [list of shares-to-remove] [list of tick-index] [list of fee indexes] ",
			name:      "missing arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "0"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 5",
		},
		{
			name:      "too many arguments",
			args:      []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "0", "1", s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 7",
		},
		{
			name:     "valid",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "0", "1"},
			errInRes: false,
		},
		{
			name:     "valid: multiple case",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "2,2", "0,0", "0,1"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdWithdrawl()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestTx3CmdSwap() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "swap [receiver] [amount-in] [tokenA] [tokenB] [token-in] [minOut] [priceLimit]",
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "5", "TokenA", "TokenB", "TokenA", "4"},
			expErr:    true,
			expErrMsg: "Error: accepts 7 arg(s), received 6",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "5", "TokenA", "TokenB", "TokenA", "0", "2", s.addr1.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 7 arg(s), received 8",
		},
		{
			name:     "valid",
			args:     []string{s.addr1.String(), "2", "TokenA", "TokenB", "TokenA", "0", "0.0"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdSwap()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestTx4Cmd4laceLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			// "place-limit-order [receiver] [token-a] [token-b] [tick-index] [token-in] [amount-in]",,
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 5",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB", "10", "1"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 7",
		},
		{
			name:     "valid",
			args:     []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB", "10"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdPlaceLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestTx5CmdCancelLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	// Place Limit Order
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "0", "TokenB", "10"}, commonFlags...)
	cmd := dexClient.CmdPlaceLimitOrder()
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			//  "cancel-limit-order [receiver] [token-a] [token-b] [tick-index] [key-token] [key]"
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 5",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB", "0", "1"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 7",
		},
		{
			name:     "valid",
			args:     []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB", "0"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdCancelLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}

func (s *IntegrationTestSuite) TestTx6CmdWithdrawFilledLimitOrder() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	// Place Limit Order
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "0", "TokenB", "10"}, commonFlags...)
	cmd := dexClient.CmdPlaceLimitOrder()
	_, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	argsSwap := append([]string{s.network.Validators[0].Address.String(), "30", "TokenA", "TokenB", "TokenA", "0", "0.0"}, commonFlags...)
	cmd = dexClient.CmdSwap()
	_, err = cli.ExecTestCLICmd(clientCtx, cmd, argsSwap)
	require.NoError(s.T(), err)

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		errInRes  bool
	}{
		{
			//  "withdraw-filled-limit-order [receiver] [token-a] [token-b] [tick-index] [key-token] [key]"
			name:      "missing arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 5",
		},
		{
			name:      "too many arguments",
			args:      []string{s.addr1.String(), "TokenA", "TokenB", "0", "TokenB", "0", "1"},
			expErr:    true,
			expErrMsg: "Error: accepts 6 arg(s), received 7",
		},
		{
			name:     "valid",
			args:     []string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "0", "TokenB", "1"},
			errInRes: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdWithdrawFilledLimitOrder()
			args := append(tc.args, commonFlags...)
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				if tc.errInRes {
					require.Contains(s.T(), out.String(), tc.expErrMsg)
				} else {
					require.NoError(s.T(), err)
					var res sdk.TxResponse
					require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
					require.Zero(s.T(), res.Code, res.RawLog)
				}

			}
		})
	}
}
