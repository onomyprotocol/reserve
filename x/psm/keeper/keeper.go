package keeper

import (
	"fmt"

	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"

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

		BankKeeper    types.BankKeeper
		AccountKeeper types.AccountKeeper
		OracleKeeper  types.OracleKeeper

		Stablecoins collections.Map[string, types.Stablecoin]
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
	oracleKeeper types.OracleKeeper,
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

		BankKeeper:    bankKeeper,
		AccountKeeper: accountKeeper,
		OracleKeeper:  oracleKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Stablecoins:   collections.NewMap(sb, types.KeyStableCoin, "stablecoins", collections.StringKey, codec.CollValue[types.Stablecoin](cdc)),
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

func (k Keeper) TotalStablecoinLock(ctx context.Context, nameStablecoin string) (math.Int, error) {
	sc, err := k.Stablecoins.Get(ctx, nameStablecoin)
	if err != nil {
		return math.Int{}, fmt.Errorf("canot found stablecoin name %s", nameStablecoin)
	}

	return sc.TotalStablecoinLock, nil
}

func (k Keeper) AddTotalStablecoinLock(ctx context.Context, amountAdd sdk.Coin) error {
	sc, err := k.Stablecoins.Get(ctx, amountAdd.Denom)
	if err != nil {
		return fmt.Errorf("canot found stablecoin name %s", amountAdd.Denom)
	}

	sc.TotalStablecoinLock = sc.TotalStablecoinLock.Add(amountAdd.Amount)
	if sc.TotalStablecoinLock.GT(sc.LimitTotal) {
		return fmt.Errorf("exceed limitTotal stablecoin %s", amountAdd.Denom)
	}

	return k.Stablecoins.Set(ctx, sc.Denom, sc)
}

func (k Keeper) SubTotalStablecoinLock(ctx context.Context, amountSub sdk.Coin) error {
	sc, err := k.Stablecoins.Get(ctx, amountSub.Denom)
	if err != nil {
		return fmt.Errorf("canot found stablecoin name %s", amountSub.Denom)
	}

	sc.TotalStablecoinLock = sc.TotalStablecoinLock.Sub(amountSub.Amount)
	if sc.TotalStablecoinLock.LT(math.ZeroInt()) {
		return fmt.Errorf("not enough stablecoins %s to pay", amountSub.Denom)
	}
	return k.Stablecoins.Set(ctx, sc.Denom, sc)
}
