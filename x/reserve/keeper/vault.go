package keeper

import (
	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVault set a specific vault in the store from its index
func (k Keeper) SetVault(ctx sdk.Context, vault types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))
	a := k.cdc.MustMarshal(&vault)
	store.Set(types.VaultKey(
		vault.Uid,
	), a)
}

// GetVault returns a vault from its index
func (k Keeper) GetVault(
	ctx sdk.Context,
	uid uint64,
) (val types.Vault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))

	b := store.Get(types.VaultKey(
		uid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVault removes a vault from the store
func (k Keeper) RemoveVault(
	ctx sdk.Context,
	uid uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKeyPrefix))

	b := store.Get(types.VaultKey(
		uid,
	))

	if b == nil {
		return
	}

	store.Delete(types.VaultKey(
		uid,
	))
}
