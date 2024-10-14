package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called at every block, update validator set
func (k *Keeper) BeginBlocker(ctx sdk.Context) error {
	currentTime := ctx.BlockTime()
	params := k.GetParams(ctx)
	// TODO: Recalculate debt
	lastUpdate, err := k.LastUpdateTime.Get(ctx)
	if err != nil {
		return err
	}
	if currentTime.Sub(lastUpdate.Time) >= params.ChargingPeriod {
		return k.UpdateVaultsDebt(ctx, lastUpdate.Time, currentTime)
	}

	return nil
}
