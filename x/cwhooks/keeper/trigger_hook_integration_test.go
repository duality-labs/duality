package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/duality-labs/duality/testutil/keeper"
	"github.com/duality-labs/duality/x/cwhooks/types"
	"github.com/stretchr/testify/require"
)

type ContractExecCall struct {
	contractAddr string
	callerAddr   string
	args         []byte
	coins        sdk.Coins
}

type DummyWasmKeeper struct{}

var execCalls []ContractExecCall

func (k DummyWasmKeeper) Execute(_ sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error) {
	fmt.Printf("EXEC: %v\n", contractAddress)

	call := ContractExecCall{contractAddr: contractAddress.String(), callerAddr: caller.String(), args: msg, coins: coins}
	execCalls = append(execCalls, call)

	return []byte{}, nil
}

func NewDummyWasmKeeper() DummyWasmKeeper {
	execCalls = []ContractExecCall{}
	return DummyWasmKeeper{}
}

func TestHookLifecycle(t *testing.T) {
	wasmKeeper := NewDummyWasmKeeper()
	keeper, ctx := keepertest.CWHooksKeeper(t, wasmKeeper)
	triggerKey := "TriggerKey"
	triggerValue := "TriggerValue"
	hook := types.Hook{
		ContractAddress: sdk.AccAddress([]byte("contractAddr")).String(),
		Args:            "ARGLIST",
		Persistent:      false,
		TriggerKey:      triggerKey,
		TriggerValue:    triggerValue,
		Creator:         sdk.AccAddress([]byte("creatorAddr")).String(),
	}
	keeper.AppendHook(ctx, hook)
	keeper.EmitTrigger(ctx, triggerKey, triggerValue)
	keeper.InvokeAllTriggeredHooks(ctx)

	expectedExecCall := ContractExecCall{contractAddr: hook.ContractAddress, callerAddr: hook.Creator, args: []byte(hook.Args), coins: sdk.Coins{}}
	require.ElementsMatch(t, []ContractExecCall{expectedExecCall}, execCalls)
}
