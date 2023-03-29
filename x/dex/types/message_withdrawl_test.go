package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgWithdrawl_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawl
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgWithdrawl{
				Creator:         "invalid_address",
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        "invalid_address",
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid fee indexes length",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid tick indexes length",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid shares to remove length",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "no withdraw specs",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				SharesToRemove:  []sdk.Int{},
			},
			err: ErrZeroWithdraw,
		},
		{
			name: "no withdraw specs",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.ZeroInt()},
			},
			err: ErrZeroWithdraw,
		},
		{
			name: "valid msg",
			msg: MsgWithdrawl{
				Creator:         sample.AccAddress(),
				Receiver:        sample.AccAddress(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
