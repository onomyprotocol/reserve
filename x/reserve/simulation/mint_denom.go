package simulation

import (
	"math/rand"

	"reserve/x/reserve/keeper"
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintDenom(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintDenom{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MintDenom simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MintDenom simulation not implemented"), nil, nil
	}
}
