package simulation

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/x/faucet/keeper"
	"github.com/NicholasDotSol/duality/x/faucet/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEmit(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEmit{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Emit simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Emit simulation not implemented"), nil, nil
	}
}
