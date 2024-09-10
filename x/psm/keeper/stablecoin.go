package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	// "github.com/cosmos/cosmos-sdk/codec"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) SetStablecoin(ctx context.Context, s types.Stablecoin) {
	store := k.storeService.OpenKVStore(ctx)

	key := types.GetKeyStableCoin(s.Denom)
	bz := k.cdc.MustMarshal(&s)

	store.Set(key, bz)
}

func (k Keeper) GetStablecoin(ctx context.Context, denom string) (types.Stablecoin, bool) {
	store := k.storeService.OpenKVStore(ctx)

	key := types.GetKeyStableCoin(denom)

	bz, err := store.Get(key)
	if bz == nil || err != nil {
		return types.Stablecoin{}, false
	}

	var token types.Stablecoin
	k.cdc.MustUnmarshal(bz, &token)

	return token, true
}

func (k Keeper) IterateStablecoin(ctx context.Context, cb func(red types.Stablecoin) (stop bool)) error {
	store := k.storeService.OpenKVStore(ctx)

	iterator, err := store.Iterator(types.KeyStableCoin, storetypes.PrefixEndBytes(types.KeyStableCoin))
	if err != nil {
		return err
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var token types.Stablecoin
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		if cb(token) {
			break
		}
	}
	return nil
}

func (k Keeper) SetLockCoin(ctx context.Context, lockCoin types.LockCoin) {
	store := k.storeService.OpenKVStore(ctx)

	key := types.GetKeyLockCoin(lockCoin.Address)
	bz := k.cdc.MustMarshal(&lockCoin)

	store.Set(key, bz)
}

func (k Keeper) GetLockCoin(ctx context.Context, addr string) (types.LockCoin, bool) {
	store := k.storeService.OpenKVStore(ctx)

	key := types.GetKeyLockCoin(addr)

	bz, err := store.Get(key)
	if bz == nil || err != nil {
		return types.LockCoin{}, false
	}

	var lockCoin types.LockCoin
	k.cdc.MustUnmarshal(bz, &lockCoin)

	return lockCoin, true
}

func (k Keeper) IterateLockCoin(ctx context.Context, cb func(red types.LockCoin) (stop bool)) error {
	store := k.storeService.OpenKVStore(ctx)

	iterator, err := store.Iterator(types.KeyLockStableCoin, storetypes.PrefixEndBytes(types.KeyLockStableCoin))
	if err != nil {
		return err
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var lockCoin types.LockCoin
		k.cdc.MustUnmarshal(iterator.Value(), &lockCoin)
		if cb(lockCoin) {
			break
		}
	}
	return nil
}

func (k Keeper) TotalStablecoinLock(ctx context.Context, denom string) math.Int {
	total := math.ZeroInt()

	k.IterateLockCoin(ctx, func(red types.LockCoin) (stop bool) {
		if red.Coin.Denom == denom {
			total.Add(red.Coin.Amount)
		}
		return false
	})
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	totalStablecoinLock := k.bankKeeper.GetBalance(ctx, moduleAddr, denom).Amount

	if !total.Equal(totalStablecoinLock) {
		panic(fmt.Sprintf("no equal %v and %v", total, totalStablecoinLock))
	}

	return total
}
