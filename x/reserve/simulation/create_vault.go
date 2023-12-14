package simulation

import (
	"math/rand"

	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateVault(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateVault{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateVault simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateVault simulation not implemented"), nil, nil
	}
}
