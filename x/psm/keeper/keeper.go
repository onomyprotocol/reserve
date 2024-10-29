package keeper

import (
	"fmt"

	"context"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"

	// "cosmossdk.io/math"
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

		// stablecoin / totalStablecoinLock
		totalStablecoinLock collections.Map[string, math.Int]
		FeeMaxStablecoin    collections.Map[string, string]
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

		BankKeeper:          bankKeeper,
		AccountKeeper:       accountKeeper,
		OracleKeeper:        oracleKeeper,
		Params:              collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		totalStablecoinLock: collections.NewMap(sb, types.KeyTotalStablecoinLock, "total_stablecoin_lock", collections.StringKey, sdk.IntValue),
		FeeMaxStablecoin:    collections.NewMap(sb, types.KeyFeeMax, "fee_max_stablecoin", collections.StringKey, collections.StringValue),
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

func (k Keeper) TotalStablecoinLock(ctx context.Context, nameStablecoin string) (math.Int, error) {
	total := math.ZeroInt()

	err := k.totalStablecoinLock.Walk(ctx, nil, func(key string, value math.Int) (stop bool, err error) {
		if key == nameStablecoin {
			total = total.Add(value)
		}
		return false, nil
	})

	return total, err
}
