package cli_test

import (
	"fmt"
	"testing"

	"github.com/NicholasDotSol/duality/testutil/network"
	dexClient "github.com/NicholasDotSol/duality/x/dex/client/cli"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type QueryTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	addr1 sdk.AccAddress
	addr2 sdk.AccAddress
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

func (s *QueryTestSuite) SetupSuite() {
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

	var commonFlags = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
		fmt.Sprintf("--%s=%s", flags.FlagGas, "200000000"),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, s.network.Validators[0].Address.String()),
	}

	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	args := append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "10", "10", "0", "1", "false"}, commonFlags...)
	cmd := dexClient.CmdDeposit()
	_, err = cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	args = append([]string{s.network.Validators[0].Address.String(), "TokenA", "TokenB", "20", "TokenB", "10"}, commonFlags...)
	cmd = dexClient.CmdPlaceLimitOrder()
	_, err = cli.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(s.T(), err)

	//s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr1, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))

	//s.fundAccount(clientCtx, s.network.Validators[0].Address, s.addr2, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)), sdk.NewCoin("TokenA", sdk.NewInt(100000))))
}

func (s *QueryTestSuite) TestQueryCmdShowFeeTier() {
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

func (s *QueryTestSuite) TestQueryCmdListFeeTier() {
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

func (s *QueryTestSuite) TestQueryCmdShowTick() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	//clientCtx.OutputFormat = outputFormat

	ValidTestReserve := sdk.NewInt(10)
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.TickLiquidity
	}{
		{
			// show-tick-liquidity [pair-id] [token-in] [tick-index] [liquidity-type] [liquidity-index]
			name: "valid",
			args: []string{"TokenA<>TokenB", "TokenA", `"-3"`, types.LiquidityTypeLP, "3"},
			expOutput: types.TickLiquidity{
				PairId: &types.PairId{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:        "TokenA",
				TickIndex:      -3,
				LiquidityType:  types.LiquidityTypeLP,
				LiquidityIndex: 3,
				LPReserve:      &ValidTestReserve,
			},
		},
		{
			name:      "uninitialized tick",
			args:      []string{"TokenA<>TokenB", "TokenA", "1", types.LiquidityTypeLP, "0"},
			expErr:    true,
			expErrMsg: "key not found",
		},

		{
			name:      "pair not specified",
			args:      []string{"1"},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 1",
		},

		{
			name:      "tick not specified",
			args:      []string{"TokenA<>TokenB", "TokenA", "1", types.LiquidityTypeLP},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 4",
		},

		{
			name:      "too many ticks",
			args:      []string{"Token", "TokenA", "1", types.LiquidityTypeLP, "1"},
			expErr:    true,
			expErrMsg: "PairId does not conform to pattern",
		},

		{
			name:      "multiple pairIds",
			args:      []string{"TokenA<>TokenB", "TokenA<>stake", `"-3"`, types.LiquidityTypeLP, "3"},
			expErr:    true,
			expErrMsg: "key not found",
		},

		{
			name:      "too many arguments",
			args:      []string{"TokenA<>TokenB", "TokenA", `"-3`, types.LiquidityTypeLP, "3", "10"},
			expErr:    true,
			expErrMsg: "Error: accepts 5 arg(s), received 6",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowTickLiquidity()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryGetTickLiquidityResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res.TickLiquidity)
				require.Equal(s.T(), tc.expOutput, res.TickLiquidity)

			}
		})
	}
}

func (s *QueryTestSuite) TestQueryCmdListTick() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	ValidTestReserve := sdk.NewInt(10)
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.TickLiquidity
	}{
		{
			name: "valid",
			args: []string{},
			expOutput: []types.TickLiquidity{
				types.TickLiquidity{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:        "TokenA",
					TickIndex:      -3,
					LiquidityType:  types.LiquidityTypeLP,
					LiquidityIndex: 3,
					LPReserve:      &ValidTestReserve,
				},
				types.TickLiquidity{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:        "TokenB",
					TickIndex:      3,
					LiquidityType:  types.LiquidityTypeLP,
					LiquidityIndex: 3,
					LPReserve:      &ValidTestReserve,
				},
				types.TickLiquidity{
					PairId: &types.PairId{
						Token0: "TokenA",
						Token1: "TokenB",
					},
					TokenIn:        "TokenB",
					TickIndex:      20,
					LiquidityType:  types.LiquidityTypeLO,
					LiquidityIndex: 0,
					LimitOrderTranche: &types.LimitOrderTranche{
						PairId: &types.PairId{
							Token0: "TokenA",
							Token1: "TokenB",
						},
						TokenIn:          "TokenB",
						TickIndex:        20,
						ReservesTokenIn:  sdk.NewInt(10),
						ReservesTokenOut: sdk.ZeroInt(),
						TotalTokenIn:     sdk.NewInt(10),
						TotalTokenOut:    sdk.ZeroInt(),
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListTickLiquidity()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)
				var res types.QueryAllTickLiquidityResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.TickLiquidity)

			}
		})
	}
}

func (s *QueryTestSuite) TestQueryCmdShowUserPosition() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.UserPositions
	}{
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

func (s *QueryTestSuite) TestQueryCmdShowLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.LimitOrderTranche
	}{
		//"list-limit-order-tranche [pair-id] [token-in]"
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

func (s *QueryTestSuite) TestQueryCmdShowLimitOrderTrancheUser() {
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

func (s *QueryTestSuite) TestQueryCmdListLimitOrderTrancheUser() {
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

func (s *QueryTestSuite) TestQueryCmdListUserLimitOrders() {
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
				err := clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
				require.NoError(s.T(), err)
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.LimitOrders)

			}
		})
	}
}
func (s *QueryTestSuite) TestQueryCmdListUserDeposits() {
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
				require.Equal(s.T(), tc.expOutput, res.Deposits)

			}
		})
	}
}

// func (s *QueryTestSuite) TestQuery11CmdShowTokenMap() {
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
// 			expOutput: types.TokenMap{},
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

// func (s *QueryTestSuite) TestQuery12CmdListTokenMap() {
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
