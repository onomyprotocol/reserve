package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) SetStablecoin(ctx context.Context, s types.Stablecoin) error {
	store := k.storeService.OpenKVStore(ctx)

	key := types.GetKeyStableCoin(s.Denom)
	bz := k.cdc.MustMarshal(&s)

	return store.Set(key, bz)
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

func (k Keeper) GetTotalLimitWithDenomStablecoin(ctx context.Context, denom string) (math.Int, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.Int{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.LimitTotal, nil
}

// func (k Keeper) GetPrice(ctx context.Context, denom string) (math.LegacyDec, error) {
// 	s, found := k.GetStablecoin(ctx, denom)
// 	if !found {
// 		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
// 	}
// 	return s.Price, nil
// }

func (k Keeper) GetFeeIn(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeIn, nil
}

func (k Keeper) GetFeeOut(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeOut, nil
}
