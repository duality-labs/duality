package simulation

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/x/dex/keeper"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgWithdrawal(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// simAccount, _ := simtypes.RandomAcc(r, accs)
		// msg := &types.MsgWithdrawal{
		// 	Creator: simAccount.Address.String(),
		// }

		// TODO: Handling the Withdrawal simulation

		return simtypes.NoOpMsg(types.ModuleName, types.TypeMsgWithdrawal, "Withdrawal simulation not implemented"), nil, nil
	}
}
