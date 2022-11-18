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

	"github.com/NicholasDotSol/duality/testutil/network"
	"github.com/NicholasDotSol/duality/testutil/nullify"
	"github.com/NicholasDotSol/duality/x/dex/client/cli"
	"github.com/NicholasDotSol/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithSharesObjects(t *testing.T, n int) (*network.Network, []types.Shares) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		shares := types.Shares{
			Address:   strconv.Itoa(i),
			PairId:    strconv.Itoa(i),
			TickIndex: int64(i),
			FeeIndex:  uint64(i),
		}
		nullify.Fill(&shares)
		state.SharesList = append(state.SharesList, shares)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.SharesList
}

func TestShowShares(t *testing.T) {
	net, objs := networkWithSharesObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc        string
		idAddress   string
		idPairId    string
		idTickIndex int64
		idFee       uint64

		args []string
		err  error
		obj  types.Shares
	}{
		{
			desc:        "found",
			idAddress:   objs[0].Address,
			idPairId:    objs[0].PairId,
			idTickIndex: objs[0].TickIndex,
			idFee:       objs[0].FeeIndex,

			args: common,
			obj:  objs[0],
		},
		{
			desc:        "not found",
			idAddress:   strconv.Itoa(100000),
			idPairId:    strconv.Itoa(100000),
			idTickIndex: int64(100000),
			idFee:       uint64(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idAddress,
				tc.idPairId,
				fmt.Sprint(tc.idTickIndex),
				fmt.Sprint(tc.idFee),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowShares(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetSharesResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Shares)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Shares),
				)
			}
		})
	}
}

func TestListShares(t *testing.T) {
	net, objs := networkWithSharesObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShares(), args)
			require.NoError(t, err)
			var resp types.QueryAllSharesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Shares), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Shares),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShares(), args)
			require.NoError(t, err)
			var resp types.QueryAllSharesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Shares), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Shares),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShares(), args)
		require.NoError(t, err)
		var resp types.QueryAllSharesResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Shares),
		)
	})
}
