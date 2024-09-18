package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// EndBlocker called at every block, update validator set
func (k *Keeper) BeginBlocker(ctx sdk.Context) error {
	height := ctx.BlockHeight()
	params := k.GetParams(ctx)
	// TODO: Recalculate debt
	if height%int64(params.RecalculateDebtPeriod) == 0 {
		k.UpdateVaultsDebt(ctx)
	}

	k.Vaults.Walk(ctx, nil, func(key uint64, vault types.Vault) (bool, error) {
		liquidated, err := k.ShouldLiquidate(ctx, vault)
		if err != nil && liquidated {

		}

	})

	// TODO: Check liquidate

	return nil
}
