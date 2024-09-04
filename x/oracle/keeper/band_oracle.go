package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// SetBandParams sets the Band params in the state
func (k Keeper) SetBandParams(ctx sdk.Context, bandParams types.BandParams) {
	bz := k.cdc.MustMarshal(&bandParams)
	store := k.storeService.OpenKVStore(ctx)
	store.Set(types.BandParamsKey, bz)
}

// GetBandParams gets the Band params stored in the state
func (k Keeper) GetBandParams(ctx sdk.Context) types.BandParams {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.BandParamsKey)

	if err != nil {
		return types.DefaultGenesis().BandParams
	}

	if bz == nil {
		return types.DefaultGenesis().BandParams
	}

	var bandParams types.BandParams
	k.cdc.MustUnmarshal(bz, &bandParams)
	return bandParams
}
