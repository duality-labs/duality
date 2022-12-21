package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
				Creator:        "invalid_address",
				Receiver:       sample.AccAddress(),
				FeeIndexes:     []uint64{},
				TickIndexes:    []int64{},
				SharesToRemove: []sdk.Int{},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgWithdrawl{
				Creator:        sample.AccAddress(),
				Receiver:       "invalid_address",
				FeeIndexes:     []uint64{},
				TickIndexes:    []int64{},
				SharesToRemove: []sdk.Int{},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid fee indexes length",
			msg: MsgWithdrawl{
				Creator:        sample.AccAddress(),
				Receiver:       sample.AccAddress(),
				FeeIndexes:     []uint64{},
				TickIndexes:    []int64{0},
				SharesToRemove: []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "invalid tick indexes length",
			msg: MsgWithdrawl{
				Creator:        sample.AccAddress(),
				Receiver:       sample.AccAddress(),
				FeeIndexes:     []uint64{0},
				TickIndexes:    []int64{},
				SharesToRemove: []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "invalid shares to remove length",
			msg: MsgWithdrawl{
				Creator:        sample.AccAddress(),
				Receiver:       sample.AccAddress(),
				FeeIndexes:     []uint64{0},
				TickIndexes:    []int64{0},
				SharesToRemove: []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "valid msg",
			msg: MsgWithdrawl{
				Creator:        sample.AccAddress(),
				Receiver:       sample.AccAddress(),
				FeeIndexes:     []uint64{0},
				TickIndexes:    []int64{0},
				SharesToRemove: []sdk.Int{sdk.OneInt()},
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
