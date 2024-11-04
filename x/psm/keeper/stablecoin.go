package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
)

func (k Keeper) GetTotalLimitWithDenomStablecoin(ctx context.Context, denom string) (math.Int, error) {
	s, err := k.Stablecoins.Get(ctx, denom)
	if err != nil {
		return math.Int{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.LimitTotal, nil
}

func (k Keeper) GetNomType(ctx context.Context, denom string) (string, error) {
	s, err := k.Stablecoins.Get(ctx, denom)
	if err != nil {
		return "", fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.NomType, nil
}

func (k Keeper) GetFeeIn(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, err := k.Stablecoins.Get(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeIn, nil
}

func (k Keeper) GetFeeOut(ctx context.Context, denom string) (math.LegacyDec, error) {
	s, err := k.Stablecoins.Get(ctx, denom)
	if err != nil {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeOut, nil
}
