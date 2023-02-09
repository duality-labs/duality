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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/testutil/network"
	"github.com/duality-labs/duality/testutil/nullify"
	"github.com/duality-labs/duality/x/dex/client/cli"
	"github.com/duality-labs/duality/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithFilledLimitOrderTrancheObjects(t *testing.T, n int) (*network.Network, []types.FilledLimitOrderTranche) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		filledLimitOrderTranche := types.FilledLimitOrderTranche{
			PairId:           &types.PairId{Token0: "TokenA", Token1: "TokenB"},
			TokenIn:          strconv.Itoa(i),
			TickIndex:        int64(i),
			TrancheKey:     uint64(i),
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
		}
		nullify.Fill(&filledLimitOrderTranche)
		state.FilledLimitOrderTrancheList = append(state.FilledLimitOrderTrancheList, filledLimitOrderTranche)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.FilledLimitOrderTrancheList
}

func TestShowFilledLimitOrderTranche(t *testing.T) {
	net, objs := networkWithFilledLimitOrderTrancheObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc           string
		idPairId       *types.PairId
		idTokenIn      string
		idTickIndex    int64
		idTrancheKey uint64

		args []string
		err  error
		obj  types.FilledLimitOrderTranche
	}{
		{
			desc:           "found",
			idPairId:       objs[0].PairId,
			idTokenIn:      objs[0].TokenIn,
			idTickIndex:    objs[0].TickIndex,
			idTrancheKey: objs[0].TrancheKey,

			args: common,
			obj:  objs[0],
		},
		{
			desc:           "not found",
			idPairId:       &types.PairId{Token0: "TokenA", Token1: "TokenB"},
			idTokenIn:      strconv.Itoa(100000),
			idTickIndex:    100000,
			idTrancheKey: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idPairId.Stringify(),
				tc.idTokenIn,
				strconv.Itoa(int(tc.idTickIndex)),
				strconv.Itoa(int(tc.idTrancheKey)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowFilledLimitOrderTranche(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetFilledLimitOrderTrancheResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.FilledLimitOrderTranche)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.FilledLimitOrderTranche),
				)
			}
		})
	}
}

func TestListFilledLimitOrderTranche(t *testing.T) {
	net, objs := networkWithFilledLimitOrderTrancheObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFilledLimitOrderTranche(), args)
			require.NoError(t, err)
			var resp types.QueryAllFilledLimitOrderTrancheResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.FilledLimitOrderTranche), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.FilledLimitOrderTranche),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFilledLimitOrderTranche(), args)
			require.NoError(t, err)
			var resp types.QueryAllFilledLimitOrderTrancheResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.FilledLimitOrderTranche), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.FilledLimitOrderTranche),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListFilledLimitOrderTranche(), args)
		require.NoError(t, err)
		var resp types.QueryAllFilledLimitOrderTrancheResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.FilledLimitOrderTranche),
		)
	})
}
