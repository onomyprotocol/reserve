package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDenom set a specific denom in the store from its index
func (k Keeper) SetDenom(ctx sdk.Context, base string, denom types.Denom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))
	a := k.cdc.MustMarshal(&denom)
	store.Set(types.DenomKey(
		base,
	), a)
}

// GetDenom returns a denom from its index
func (k Keeper) GetDenom(
	ctx sdk.Context,
	base string,
) (val types.Denom, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		base,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDenom removes a denom from the store
func (k Keeper) RemoveDenom(
	ctx sdk.Context,
	base string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomKeyPrefix))

	b := store.Get(types.DenomKey(
		base,
	))

	if b == nil {
		return
	}

	store.Delete(types.DenomKey(
		base,
	))
}
