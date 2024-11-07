package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker called at every block, update validator set
func (k *Keeper) BeginBlocker(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
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
