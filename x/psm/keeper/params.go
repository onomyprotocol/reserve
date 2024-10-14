package keeper

import (
	"context"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	return store.Set(types.ParamsKey, bz)
}

// GetParams gets the x/staking module parameters.
func (k Keeper) GetParams(ctx context.Context) (params types.Params, err error) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.ParamsKey)
	if err != nil {
		return params, err
	}

	if bz == nil {
		return params, nil
	}

	err = k.cdc.Unmarshal(bz, &params)
	return params, err
}
