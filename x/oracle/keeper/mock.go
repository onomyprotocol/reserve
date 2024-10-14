package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	psmtypes "github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	_, ok := k.price[symbol]

	if !ok {
		k.SetPrice(ctx, symbol, math.LegacyZeroDec())
	}
	return nil
}

func (k Keeper) GetPrice(ctx context.Context, base, quote string) *math.LegacyDec {
	base_price, ok := k.price[base]

	if !ok {
		panic("call SetPrice " + base)
	}

	quote_price, ok := k.price[quote]

	if !ok {
		panic("call SetPrice " + quote)
	}
	multiplier := base_price.Quo(quote_price)
	return &multiplier
}

func (k Keeper) SetPrice(ctx context.Context, denom string, price math.LegacyDec) {
	k.price[denom] = price
}

func (k msgServer) SetPrice(ctx context.Context, msg *types.MsgSetPrice) (*types.MsgSetPriceResponse, error) {
	k.k.SetPrice(ctx, msg.Denom, msg.Price)
	return &types.MsgSetPriceResponse{}, nil
}

func (k Keeper) QueryPrice(ctx context.Context, msg *types.MsgGetPrice) (*types.MsgGetPriceResponse, error) {
	price := k.GetPrice(ctx, msg.Denom, psmtypes.DenomStable)
	return &types.MsgGetPriceResponse{Price: *price}, nil
}
