package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reserve/x/reserve/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MarketCollateral(ctx),
		k.ReserveCollateral(ctx),
		k.CollateralDeposit(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MarketCollateral returns the MarketCollateral param
func (k Keeper) MarketCollateral(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMarketCollateral, &res)
	return
}

// ReserveCollateral returns the ReserveCollateral param
func (k Keeper) ReserveCollateral(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyReserveCollateral, &res)
	return
}

// CollateralDeposit returns the CollateralDeposit param
func (k Keeper) CollateralDeposit(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyCollateralDeposit, &res)
	return
}
