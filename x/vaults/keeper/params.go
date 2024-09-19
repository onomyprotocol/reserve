package keeper

import (
	"context"

	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return types.Params{}
	}
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	return k.Params.Set(ctx, params)
}
