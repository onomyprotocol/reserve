package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reserve/x/portal/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MarketChannel(ctx),
		k.ProviderChannel(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MarketChannel returns the MarketChannel param
func (k Keeper) MarketChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMarketChannel, &res)
	return
}

// ProviderChannel returns the ProviderChannel param
func (k Keeper) ProviderChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyProviderChannel, &res)
	return
}
