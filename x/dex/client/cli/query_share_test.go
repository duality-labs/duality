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

func networkWithShareObjects(t *testing.T, n int) (*network.Network, []types.Share) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		share := types.Share{
			Owner:  strconv.Itoa(i),
			Token0: strconv.Itoa(i),
			Token1: strconv.Itoa(i),
			Price:  strconv.Itoa(i),
			Fee:    strconv.Itoa(i),
		}
		nullify.Fill(&share)
		state.ShareList = append(state.ShareList, share)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ShareList
}

func TestShowShare(t *testing.T) {
	net, objs := networkWithShareObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc     string
		idOwner  string
		idToken0 string
		idToken1 string
		idPrice  string
		idFee    string

		args []string
		err  error
		obj  types.Share
	}{
		{
			desc:     "found",
			idOwner:  objs[0].Owner,
			idToken0: objs[0].Token0,
			idToken1: objs[0].Token1,
			idPrice:  objs[0].Price,
			idFee:    objs[0].Fee,

			args: common,
			obj:  objs[0],
		},
		{
			desc:     "not found",
			idOwner:  strconv.Itoa(100000),
			idToken0: strconv.Itoa(100000),
			idToken1: strconv.Itoa(100000),
			idPrice:  strconv.Itoa(100000),
			idFee:    strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOwner,
				tc.idToken0,
				tc.idToken1,
				tc.idPrice,
				tc.idFee,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowShare(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetShareResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Share)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Share),
				)
			}
		})
	}
}

func TestListShare(t *testing.T) {
	net, objs := networkWithShareObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShare(), args)
			require.NoError(t, err)
			var resp types.QueryAllShareResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Share), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Share),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShare(), args)
			require.NoError(t, err)
			var resp types.QueryAllShareResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Share), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Share),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListShare(), args)
		require.NoError(t, err)
		var resp types.QueryAllShareResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Share),
		)
	})
}
