package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
)

func (k Keeper) GetTotalLimitWithDenomStablecoin(ctx context.Context, denom string) (math.Int, error) {
	s, err := k.StablecoinInfo.Get(ctx, denom)
	if err != nil {
		return math.Int{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.LimitTotal, nil
}

func (k Keeper) GetFeeIn(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, err := k.StablecoinInfo.Get(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeIn, nil
}

func (k Keeper) GetFeeOut(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, err := k.StablecoinInfo.Get(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeOut, nil
}
