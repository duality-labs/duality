package cli_test

// JCP TODO: fix me
// import (
// 	"fmt"
// 	"strconv"
// 	"testing"

// 	sdkmath "cosmossdk.io/math"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// 	"github.com/stretchr/testify/require"

// 	"github.com/duality-labs/duality/testutil/network"
// 	"github.com/duality-labs/duality/x/cwhooks/client/cli"
// )

// func TestCreateHook(t *testing.T) {
// 	net := network.New(t)
// 	val := net.Validators[0]
// 	ctx := val.ClientCtx

// 	fields := []string{"xyz", "xyz", "false", "xyz", "xyz"}
// 	tests := []struct {
// 		desc string
// 		args []string
// 		err  error
// 		code uint32
// 	}{
// 		{
// 			desc: "valid",
// 			args: []string{
// 				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
// 				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
// 				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
// 				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
// 			},
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			require.NoError(t, net.WaitForNextBlock())

// 			args := []string{}
// 			args = append(args, fields...)
// 			args = append(args, tc.args...)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateHook(), args)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			var resp sdk.TxResponse
// 			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
// 		})
// 	}
// }

// func TestDeleteHook(t *testing.T) {
// 	net := network.New(t)

// 	val := net.Validators[0]
// 	ctx := val.ClientCtx

// 	fields := []string{"xyz", "xyz", "false", "xyz", "xyz"}
// 	common := []string{
// 		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
// 		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
// 		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
// 		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
// 	}
// 	args := []string{}
// 	args = append(args, fields...)
// 	args = append(args, common...)
// 	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateHook(), args)
// 	require.NoError(t, err)

// 	tests := []struct {
// 		desc string
// 		id   string
// 		args []string
// 		code uint32
// 		err  error
// 	}{
// 		{
// 			desc: "valid",
// 			id:   "0",
// 			args: common,
// 		},
// 		{
// 			desc: "key not found",
// 			id:   "1",
// 			args: common,
// 			code: sdkerrors.ErrKeyNotFound.ABCICode(),
// 		},
// 		{
// 			desc: "invalid key",
// 			id:   "invalid",
// 			err:  strconv.ErrSyntax,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			require.NoError(t, net.WaitForNextBlock())

// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteHook(), append([]string{tc.id}, tc.args...))
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			var resp sdk.TxResponse
// 			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
// 		})
// 	}
// }
