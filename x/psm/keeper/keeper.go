package keeper

import (
	"context"
	"cosmossdk.io/math"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		addressCodec address.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message.
		// Typically, this should be the x/gov module account.
		authority string

		Schema collections.Schema
		Params collections.Item[types.Params]
		// this line is used by starport scaffolding # collection/type

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	// addressCodec address.Codec,
	storeService store.KVStoreService,
	// logger log.Logger,
	authority string,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) Keeper {
	// if _, err := addressCodec.StringToBytes(authority); err != nil {
	// 	panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	// }

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc: cdc,
		// addressCodec: addressCodec,
		storeService: storeService,
		authority:    authority,
		// logger:       logger,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		// this line is used by starport scaffolding # collection/instantiate
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

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

func (k Keeper) GetTotalLimitWithDenomStablecoin(ctx sdk.Context, denom string) (math.Int, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.Int{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.LimitTotal, nil
}

func (k Keeper) SwapToStablecoin(ctx sdk.Context, addr sdk.AccAddress, amount math.Int, toDenom string) (math.Int, sdk.Coin, error) {
	asset := k.bankKeeper.GetBalance(ctx, addr, types.InterStableToken)

	if asset.Amount.LT(amount) {
		return math.ZeroInt(), sdk.Coin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, toDenom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.Coin{}, err
	}
	amountStablecoin := amount.ToLegacyDec().Quo(multiplier).RoundInt()

	fee, err := k.PayFeesOut(ctx, amountStablecoin, toDenom)
	if err != nil {
		return math.Int{}, sdk.Coin{}, err
	}

	receiveAmount := amountStablecoin.Sub(fee)
	return receiveAmount, sdk.NewCoin(toDenom, fee), nil
}

func (k Keeper) SwaptoIST(ctx sdk.Context, addr sdk.AccAddress, stablecoin sdk.Coin) (math.Int, sdk.Coin, error) {
	asset := k.bankKeeper.GetBalance(ctx, addr, stablecoin.Denom)

	if asset.Amount.LT(stablecoin.Amount) {
		return math.ZeroInt(), sdk.Coin{}, fmt.Errorf("insufficient balance")
	}

	multiplier, err := k.GetPrice(ctx, stablecoin.Denom)
	if err != nil || multiplier.IsZero() {
		return math.Int{}, sdk.Coin{}, err
	}

	amountIST := multiplier.Mul(stablecoin.Amount.ToLegacyDec()).RoundInt()

	fee, err := k.PayFeesIn(ctx, amountIST, stablecoin.Denom)
	if err != nil {
		return math.Int{}, sdk.Coin{}, err
	}

	receiveAmountIST := amountIST.Sub(fee)
	return receiveAmountIST, sdk.NewCoin(types.InterStableToken, fee), nil
}

func (k Keeper) GetPrice(ctx sdk.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.Price, nil
}

func (k Keeper) GetFeeIn(ctx sdk.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeIn, nil
}

func (k Keeper) GetFeeOut(ctx sdk.Context, denom string) (math.LegacyDec, error) {
	s, found := k.GetStablecoin(ctx, denom)
	if !found {
		return math.LegacyDec{}, fmt.Errorf("not found Stable coin %s", denom)
	}
	return s.FeeOut, nil
}

func (k Keeper) PayFeesOut(ctx sdk.Context, amount math.Int, denom string) (math.Int, error) {
	ratioSwapOutFees, err := k.GetFeeOut(ctx, denom)
	if err != nil {
		return math.Int{}, err
	}

	fee := ratioSwapOutFees.MulInt(amount).RoundInt()
	return fee, nil
}

func (k Keeper) PayFeesIn(ctx sdk.Context, amount math.Int, denom string) (math.Int, error) {
	ratioSwapInFees, err := k.GetFeeIn(ctx, denom)
	if err != nil {
		return math.Int{}, err
	}
	fee := ratioSwapInFees.MulInt(amount).RoundInt()
	return fee, nil
}
