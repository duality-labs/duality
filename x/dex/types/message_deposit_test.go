package types

import (
	"testing"

	"github.com/NicholasDotSol/duality/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeposit
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgDeposit{
				Creator:     "invalid_address",
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{},
				TickIndexes: []int64{},
				AmountsA:    []sdk.Int{},
				AmountsB:    []sdk.Int{},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    "invalid address",
				FeeIndexes:  []uint64{},
				TickIndexes: []int64{},
				AmountsA:    []sdk.Int{},
				AmountsB:    []sdk.Int{},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid fee indexes length",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{0},
				TickIndexes: []int64{},
				AmountsA:    []sdk.Int{},
				AmountsB:    []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "invalid tick indexes length",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{},
				TickIndexes: []int64{0},
				AmountsA:    []sdk.Int{},
				AmountsB:    []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "invalid amounts A length",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{},
				TickIndexes: []int64{},
				AmountsA:    []sdk.Int{sdk.OneInt()},
				AmountsB:    []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "invalid amounts B length",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{},
				TickIndexes: []int64{},
				AmountsA:    []sdk.Int{},
				AmountsB:    []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		}, {
			name: "valid msg",
			msg: MsgDeposit{
				Creator:     sample.AccAddress(),
				Receiver:    sample.AccAddress(),
				FeeIndexes:  []uint64{0},
				TickIndexes: []int64{0},
				AmountsA:    []sdk.Int{sdk.OneInt()},
				AmountsB:    []sdk.Int{sdk.OneInt()},
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
