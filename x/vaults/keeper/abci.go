package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called at every block, update validator set
func (k *Keeper) BeginBlocker(ctx sdk.Context) error {
	height := ctx.BlockHeight()
	params := k.GetParams(ctx)
	// TODO: Recalculate debt
	if height%int64(params.RecalculateDebtPeriod) == 0 {
		k.UpdateVaultsDebt(ctx)
	}

	return nil
}
