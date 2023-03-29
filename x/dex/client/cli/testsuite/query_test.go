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

var testAddress = sdk.AccAddress([]byte("testAddr"))

var limitOrderTrancheList = []types.TickLiquidity{
	{
		Liquidity: &types.TickLiquidity_LimitOrderTranche{
			LimitOrderTranche: &types.LimitOrderTranche{
				PairID: &types.PairID{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:          "TokenB",
				TickIndex:        1,
				TrancheKey:       "0",
				ReservesTokenIn:  sdk.NewInt(10),
				ReservesTokenOut: sdk.ZeroInt(),
				TotalTokenIn:     sdk.NewInt(10),
				TotalTokenOut:    sdk.ZeroInt(),
			},
		},
	},
	{
		Liquidity: &types.TickLiquidity_LimitOrderTranche{
			LimitOrderTranche: &types.LimitOrderTranche{
				PairID: &types.PairID{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:          "TokenB",
				TickIndex:        2,
				TrancheKey:       "1",
				ReservesTokenIn:  sdk.NewInt(10),
				ReservesTokenOut: sdk.ZeroInt(),
				TotalTokenIn:     sdk.NewInt(10),
				TotalTokenOut:    sdk.ZeroInt(),
			},
		},
	},
}

var inactiveLimitOrderTrancheList = []types.LimitOrderTranche{
	{
		PairID:           &types.PairID{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:          "TokenB",
		TickIndex:        0,
		TrancheKey:       "0",
		TotalTokenIn:     sdk.NewInt(10),
		TotalTokenOut:    sdk.NewInt(10),
		ReservesTokenOut: sdk.NewInt(10),
		ReservesTokenIn:  sdk.NewInt(0),
	},
	{
		PairID:           &types.PairID{Token0: "TokenA", Token1: "TokenB"},
		TokenIn:          "TokenB",
		TickIndex:        0,
		TrancheKey:       "1",
		TotalTokenIn:     sdk.NewInt(10),
		TotalTokenOut:    sdk.NewInt(10),
		ReservesTokenOut: sdk.NewInt(10),
		ReservesTokenIn:  sdk.NewInt(0),
	},
}

var poolReservesList = []types.TickLiquidity{
	{
		Liquidity: &types.TickLiquidity_PoolReserves{
			PoolReserves: &types.PoolReserves{
				PairID: &types.PairID{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:   "TokenB",
				TickIndex: 0,
				Fee:       1,
				Reserves:  sdk.NewInt(10),
			},
		},
	},
	{
		Liquidity: &types.TickLiquidity_PoolReserves{
			PoolReserves: &types.PoolReserves{
				PairID: &types.PairID{
					Token0: "TokenA",
					Token1: "TokenB",
				},
				TokenIn:   "TokenB",
				TickIndex: 0,
				Fee:       3,
				Reserves:  sdk.NewInt(10),
			},
		},
	},
}

var limitOrderTrancheUserList = []types.LimitOrderTrancheUser{
	{
		PairID:          &types.PairID{Token0: "TokenA", Token1: "TokenB"},
		Token:           "TokenA",
		TickIndex:       1,
		TrancheKey:      "0",
		Address:         testAddress.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
		TakerReserves:   sdk.ZeroInt(),
	},
	{
		PairID:          &types.PairID{Token0: "TokenA", Token1: "TokenB"},
		Token:           "TokenB",
		TickIndex:       20,
		TrancheKey:      "1",
		Address:         testAddress.String(),
		SharesOwned:     sdk.NewInt(10),
		SharesWithdrawn: sdk.NewInt(0),
		SharesCancelled: sdk.NewInt(0),
		TakerReserves:   sdk.ZeroInt(),
	},
}

var genesisState types.GenesisState = types.GenesisState{
	TickLiquidityList:             append(poolReservesList, limitOrderTrancheList...),
	LimitOrderTrancheUserList:     limitOrderTrancheUserList,
	InactiveLimitOrderTrancheList: inactiveLimitOrderTrancheList,
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

func (s *QueryTestSuite) TestQueryCmdListTickLiquidity() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.TickLiquidity
	}{
		{
			name:      "valid",
			args:      []string{"TokenA<>TokenB", "TokenB"},
			expOutput: append(poolReservesList, limitOrderTrancheList...),
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
		// show-limit-order-tranche [pair-id] [tick-index] [token-in] [tranche-key]
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
		// "show-limit-order-tranche-user [address] [tranche-key]"
		{
			name:      "valid",
			args:      []string{testAddress.String(), "0"},
			expOutput: limitOrderTrancheUserList[0],
		},
		{
			name:      "invalid pair",
			args:      []string{testAddress.String(), "BADKEY"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many parameters",
			args:      []string{testAddress.String(), "0", "EXTRAARG"},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 3",
		},
		{
			name:      "no parameters",
			args:      []string{},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 0",
		},
		{
			name:      "too few parameters",
			args:      []string{testAddress.String()},
			expErr:    true,
			expErrMsg: "Error: accepts 2 arg(s), received 1",
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

func (s *QueryTestSuite) TestQueryCmdListInactiveLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput []types.LimitOrderTranche
	}{
		{
			name:      "valid",
			args:      []string{},
			expOutput: inactiveLimitOrderTrancheList,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdListInactiveLimitOrderTranche()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)

				var res types.QueryAllInactiveLimitOrderTrancheResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.InactiveLimitOrderTranche)
			}
		})
	}
}

func (s *QueryTestSuite) TestQueryCmdShowInactiveLimitOrderTranche() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	testCases := []struct {
		name      string
		args      []string
		expErr    bool
		expErrMsg string
		expOutput types.LimitOrderTranche
	}{
		// show-filled-limit-order-tranche [pair-id] [token-in] [tick-index] [tranche-index]",
		{
			name:      "valid",
			args:      []string{"TokenA<>TokenB", "TokenB", "0", "0"},
			expOutput: inactiveLimitOrderTrancheList[0],
		},
		{
			name:      "invalid pair",
			args:      []string{"TokenC<>TokenB", "TokenB", "0", "0"},
			expErr:    true,
			expErrMsg: "key not found",
		},
		{
			name:      "too many parameters",
			args:      []string{"TokenC<>TokenB", "TokenB", "0", "0", "Extra arg"},
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
			args:      []string{"TokenC<>TokenB", "TokenB", "0"},
			expErr:    true,
			expErrMsg: "Error: accepts 4 arg(s), received 3",
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := dexClient.CmdShowInactiveLimitOrderTranche()
			out, err := cli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				require.Error(s.T(), err)
				require.Contains(s.T(), out.String(), tc.expErrMsg)
			} else {
				require.NoError(s.T(), err)
				var res types.QueryGetInactiveLimitOrderTrancheResponse
				require.NoError(s.T(), clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res))
				require.NotEmpty(s.T(), res)
				require.Equal(s.T(), tc.expOutput, res.InactiveLimitOrderTranche)
			}
		})
	}
}
