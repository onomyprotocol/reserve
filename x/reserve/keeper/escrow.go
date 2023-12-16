package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetEscrow set a specific escrow in the store from its index
func (k Keeper) SetEscrow(ctx sdk.Context, base string, escrow types.Escrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKeyPrefix))
	a := k.cdc.MustMarshal(&escrow)
	store.Set(types.EscrowKey(
		base,
	), a)
}

// GetEscrow returns a escrow from its index
func (k Keeper) GetEscrow(
	ctx sdk.Context,
	base string,
) (val types.Escrow, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKeyPrefix))

	b := store.Get(types.EscrowKey(
		base,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEscrow removes a escrow from the store
func (k Keeper) RemoveEscrow(
	ctx sdk.Context,
	base string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKeyPrefix))

	b := store.Get(types.EscrowKey(
		base,
	))

	if b == nil {
		return
	}

	store.Delete(types.EscrowKey(
		base,
	))
}
