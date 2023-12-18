package keeper

import (
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// ProviderChannel - The Provider Chain IBC channel
func (k Keeper) ProviderChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyProviderChannel, &res)
	return
}

// MarketChannel - The Market Chain IBC channel
func (k Keeper) MarketChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMarketChannel, &res)
	return
}

// MarketCollateral - The Collateral IBC address on Market Chain
func (k Keeper) MarketCollateral(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMarketCollateral, &res)
	return
}

// ReserveCollateral - The Collateral IBC address on Reserve Chain
func (k Keeper) ReserveCollateral(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyReserveCollateral, &res)
	return
}

// CollateralDeposit - The Collateral amount needed to create a denom
func (k Keeper) CollateralDeposit(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyCollateralDeposit, &res)
	return
}
