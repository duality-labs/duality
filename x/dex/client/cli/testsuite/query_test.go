package cli_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/testutil/network"
	dexClient "github.com/duality-labs/duality/x/dex/client/cli"
	"github.com/duality-labs/duality/x/dex/types"
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

var TEN = sdk.NewInt(10)

var testAddress = sdk.AccAddress([]byte("testAddr"))

var feeTierList = []types.FeeTier{
	{Id: 0, Fee: 1},
	{Id: 1, Fee: 3},
	{Id: 2, Fee: 5},
	{Id: 3, Fee: 10},
}

var limitOrderTrancheList = []types.TickLiquidity{
	{PairId: &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:        "TokenB",
		TickIndex:      1,
		LiquidityIndex: 0,
		LiquidityType:  types.LiquidityTypeLO,
		LimitOrderTranche: &types.LimitOrderTranche{
			ReservesTokenIn:  sdk.NewInt(10),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.NewInt(10),
			TotalTokenOut:    sdk.ZeroInt(),
		},
	},
	{PairId: &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:        "TokenB",
		TickIndex:      2,
		LiquidityIndex: 0,
		LiquidityType:  types.LiquidityTypeLO,
		LimitOrderTranche: &types.LimitOrderTranche{
			ReservesTokenIn:  sdk.NewInt(10),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.NewInt(10),
			TotalTokenOut:    sdk.ZeroInt(),
		},
	},
}

var filledLimitOrderTrancheList = []types.FilledLimitOrderTranche{
	{PairId: &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:          "TokenB",
		TickIndex:        0,
		TrancheIndex:     0,
		TotalTokenIn:     sdk.NewInt(10),
		TotalTokenOut:    sdk.NewInt(10),
		ReservesTokenOut: sdk.NewInt(10),
	},
	{PairId: &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:          "TokenB",
		TickIndex:        0,
		TrancheIndex:     1,
		TotalTokenIn:     sdk.NewInt(10),
		TotalTokenOut:    sdk.NewInt(10),
		ReservesTokenOut: sdk.NewInt(10),
	},
}

var poolReservesList = []types.TickLiquidity{
	{
		PairId:         &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:        "TokenA",
		TickIndex:      0,
		LiquidityIndex: 0,
		LiquidityType:  types.LiquidityTypeLP,
		LPReserve:      &TEN},
	{
		PairId:         &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:        "TokenA",
		TickIndex:      0,
		LiquidityIndex: 1,
		LiquidityType:  types.LiquidityTypeLP,
		LPReserve:      &TEN},
}

var limitOrderTrancheUserList = []types.LimitOrderTrancheUser{
	{
		PairId:          &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		Token:           "TokenB",
		TickIndex:       1,
		Count:           0,
		Address:         testAddress.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
	},
	{
		PairId:          &types.PairId{Token0: "TokenA", Token1: "TokenB"},
		Token:           "TokenA",
		TickIndex:       20,
		Count:           0,
		Address:         testAddress.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
	},
}

var genesisState types.GenesisState = types.GenesisState{
	TickLiquidityList:           append(poolReservesList, limitOrderTrancheList...),
	LimitOrderTrancheUserList:   limitOrderTrancheUserList,
	FilledLimitOrderTrancheList: filledLimitOrderTrancheList,
	FeeTierList:                 feeTierList,
}

func (s *QueryTestSuite) SetupSuite() {

	s.T().Log("setting up integration test suite")

	config := network.DefaultConfig()
	json, err := config.Codec.MarshalJSON(&genesisState)
	config.GenesisState["dex"] = json
	require.NoError(s.T(), err)

	nw := network.New(s.T(), config)
	s.network = nw

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *QueryTestSuite) TestQueryCmdShowFeeTier() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
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
			name:      "valid",
			args:      []string{"0"},
			expOutput: feeTierList[0],
		},
		{
			name:      "valid 2",
			args:      []string{"1"},
			expOutput: feeTierList[1],
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
			name:      "valid",
			args:      []string{},
			expOutput: feeTierList,
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

func (s *QueryTestSuite) TestQueryCmdListTickLiquidity() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	targetTick := poolReservesList[0]
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.TickLiquidity
	}{
		{
			name: "valid",
			args: []string{"TokenA<>TokenB", "TokenB"},
			expOutput: []types.TickLiquidity{
				{
					Liquidity: &types.TickLiquidity_PoolReserves{
						PoolReserves: &types.PoolReserves{
							PairId: &types.PairId{
								Token0: "TokenA",
								Token1: "TokenB",
							},
							TokenIn:   "TokenB",
							TickIndex: 3,
							Fee:       3,
							Reserves:  sdk.NewInt(10),
						},
					},
				},
				{
					Liquidity: &types.TickLiquidity_LimitOrderTranche{
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
		//show-limit-order-tranche [pair-id] [tick-index] [token-in] [tranche-index]
		{
			name:      "valid",
			args:      []string{"TokenA<>TokenB", "1", "TokenB", "0"},
			expOutput: *limitOrderTrancheList[0].GetLimitOrderTranche(),
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
			name:      "valid",
			args:      []string{"TokenA<>TokenB", "1", "TokenB", "0", testAddress.String()},
			expOutput: limitOrderTrancheUserList[0],
		},
		{
			name:      "invalid pair",
			args:      []string{"TokenB<>TokenC", "20", "TokenB", "0", testAddress.String()},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many parameters",
			args:      []string{"TokenA<>TokenB", "20", "TokenB", "0", "1", testAddress.String()},
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
		{
			name:      "valid",
			args:      []string{},
			expOutput: limitOrderTrancheUserList,
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

func (s *QueryTestSuite) TestQueryListFilledLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.FilledLimitOrderTranche
	}{
		{
			name:      "valid",
			args:      []string{},
			expOutput: filledLimitOrderTrancheList,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListFilledLimitOrderTranche()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllFilledLimitOrderTrancheResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.FilledLimitOrderTranche)

			}
		})
	}
}
