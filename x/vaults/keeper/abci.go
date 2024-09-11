package keeper

import (
	"context"
)

// EndBlocker called at every block, update validator set
func (k *Keeper) BeginBlocker(ctx context.Context) error {
	// TODO: Recalculate debt

	// TODO: Check liquidate

	return nil
}
