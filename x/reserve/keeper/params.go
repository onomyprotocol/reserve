package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return k.getParams(ctx)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) getParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// MCR - the Minimum Collateralization Ratio
func (k Keeper) MCR(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMCR, &res)
	return
}

// LR - the Liquidation Ratio
func (k Keeper) LR(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyLR, &res)
	return
}

// IR - the Interest Rate
func (k Keeper) IR(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyIR, &res)
	return
}

// SR - the Saving Rate
func (k Keeper) SR(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeySR, &res)
	return
}
