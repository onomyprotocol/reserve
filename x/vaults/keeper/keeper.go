package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/vaults/types"

	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeService  storetypes.KVStoreService
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema         collections.Schema
	Params         collections.Item[types.Params]
	VaultsManager  collections.Map[string, types.VaultMamager]
	Vaults         collections.Map[uint64, types.Vault]
	VaultsSequence collections.Sequence
}

// NewKeeper returns a new keeper by codec and storeKey inputs.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	authority string,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		authority:      authority,
		cdc:            cdc,
		storeService:   storeService,
		accountKeeper:  ak,
		bankKeeper:     bk,
		Params:         collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		VaultsManager:  collections.NewMap(sb, types.VaultManagerKey, "vaultmanagers", collections.StringKey, codec.CollValue[types.VaultMamager](cdc)),
		Vaults:         collections.NewMap(sb, types.VaultKey, "vaults", collections.Uint64Key, codec.CollValue[types.Vault](cdc)),
		VaultsSequence: collections.NewSequence(sb, types.VaultSequenceKey, "sequence"),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return &k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k *Keeper) ActiveCollateralAsset(
	ctx context.Context,
	denom string,
	minCollateralRatio math.LegacyDec,
	liquidationRatio math.LegacyDec,
	maxDebt math.Int,
) error {
	// Check if asset alreay be actived
	actived := k.IsActived(ctx, denom)
	if actived {
		return fmt.Errorf("denom %s already be actived", denom)
	}
	vm := types.VaultMamager{
		Denom: denom,
		Params: types.VaultMamagerParams{
			MinCollateralRatio: minCollateralRatio,
			LiquidationRatio:   liquidationRatio,
			MaxDebt:            maxDebt,
		},
		MintAvailable: maxDebt,
	}
	return k.VaultsManager.Set(ctx, denom, vm)
}

func (k *Keeper) GetVaultManager(
	ctx context.Context,
	denom string,
) (types.VaultMamager, error) {
	vm, err := k.VaultsManager.Get(ctx, denom)
	if err != nil {
		return types.VaultMamager{}, err
	}
	return vm, nil
}

func (k *Keeper) IsActived(
	ctx context.Context,
	denom string,
) bool {
	has, _ := k.VaultsManager.Has(ctx, denom)
	return has
}

func (k *Keeper) GetPrice(
	ctx context.Context,
	denom string,
) math.LegacyDec {

	// TODO: Call price module api
	return math.LegacyNewDec(1)
}
