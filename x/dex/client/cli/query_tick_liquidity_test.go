package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/duality-labs/duality/testutil/network"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/client/cli"
	"github.com/duality-labs/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithTickLiquidityObjects(t *testing.T, n int) (*network.Network, []types.TickLiquidity) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		tickLiquidity := types.TickLiquidity{
			Liquidity: &types.TickLiquidity_LimitOrderTranche{
				LimitOrderTranche: &types.LimitOrderTranche{
					PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
					TokenIn:          "TokenA",
					TickIndex:        int64(i),
					TrancheIndex:     uint64(i),
					ReservesTokenIn:  sdk.NewInt(0),
					ReservesTokenOut: sdk.NewInt(0),
					TotalTokenIn:     sdk.NewInt(0),
					TotalTokenOut:    sdk.NewInt(0),
				},
			},
		}
		nullify.Fill(&tickLiquidity)
		state.TickLiquidityList = append(state.TickLiquidityList, tickLiquidity)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.TickLiquidityList
}

func TestShowTickLiquidity(t *testing.T) {
	net, objs := networkWithTickLiquidityObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc             string
		idPairId         *types.PairId
		idTokenIn        string
		idTickIndex      int64
		idLiquidityType  string
		idLiquidityIndex uint64

		args []string
		err  error
		obj  types.TickLiquidity
	}{
		{
			desc:             "found",
			idPairId:         objs[0].PairId(),
			idTokenIn:        objs[0].TokenIn(),
			idTickIndex:      objs[0].TickIndex(),
			idLiquidityType:  objs[0].LiquidityType(),
			idLiquidityIndex: objs[0].LiquidityIndex(),

			args: common,
			obj:  objs[0],
		},
		{
			desc:             "not found",
			idPairId:         &types.PairId{Token0: "TokenA", Token1: "TokenB"},
			idTokenIn:        strconv.Itoa(100000),
			idTickIndex:      100000,
			idLiquidityType:  strconv.Itoa(100000),
			idLiquidityIndex: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idPairId.Stringify(),
				tc.idTokenIn,
				strconv.Itoa(int(tc.idTickIndex)),
				tc.idLiquidityType,
				strconv.Itoa(int(tc.idLiquidityIndex)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTickLiquidity(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTickLiquidityResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.TickLiquidity)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.TickLiquidity),
				)
			}
		})
	}
}

func TestListTickLiquidity(t *testing.T) {
	net, objs := networkWithTickLiquidityObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTickLiquidity(), args)
			require.NoError(t, err)
			var resp types.QueryAllTickLiquidityResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.TickLiquidity), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.TickLiquidity),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTickLiquidity(), args)
			require.NoError(t, err)
			var resp types.QueryAllTickLiquidityResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.TickLiquidity), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.TickLiquidity),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTickLiquidity(), args)
		require.NoError(t, err)
		var resp types.QueryAllTickLiquidityResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.TickLiquidity),
		)
	})
}
