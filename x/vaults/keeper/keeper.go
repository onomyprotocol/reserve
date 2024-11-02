package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/vaults/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeService  storetypes.KVStoreService
	BankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	// Temporarily leave it public to easily replace it with mocks.
	// TODO: Make it private
	OracleKeeper types.OracleKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string

	Schema          collections.Schema
	Params          collections.Item[types.Params]
	VaultsManager   collections.Map[string, types.VaultMamager]
	Vaults          collections.Map[uint64, types.Vault]
	VaultsSequence  collections.Sequence
	LastUpdateTime  collections.Item[types.LastUpdate]
	ShortfallAmount collections.Item[math.Int]
}

// NewKeeper returns a new keeper by codec and storeKey inputs.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ok types.OracleKeeper,
	authority string,
) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		authority:       authority,
		cdc:             cdc,
		storeService:    storeService,
		accountKeeper:   ak,
		OracleKeeper:    ok,
		BankKeeper:      bk,
		Params:          collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		VaultsManager:   collections.NewMap(sb, types.VaultManagerKeyPrefix, "vaultmanagers", collections.StringKey, codec.CollValue[types.VaultMamager](cdc)),
		Vaults:          collections.NewMap(sb, types.VaultKeyPrefix, "vaults", collections.Uint64Key, codec.CollValue[types.Vault](cdc)),
		VaultsSequence:  collections.NewSequence(sb, types.VaultSequenceKeyPrefix, "sequence"),
		LastUpdateTime:  collections.NewItem(sb, types.LastUpdateKeyPrefix, "last_update", codec.CollValue[types.LastUpdate](cdc)),
		ShortfallAmount: collections.NewItem(sb, types.ShortfallKeyPrefix, "shortfall", sdk.IntValue),
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
	stabilityFee math.LegacyDec,
	mintingFee math.LegacyDec,
	liquidationPenalty math.LegacyDec,
	oracleScript int64,
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
			StabilityFee:       stabilityFee,
			LiquidationPenalty: liquidationPenalty,
			MintingFee:         mintingFee,
		},
		MintAvailable: maxDebt,
	}
	err := k.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, denom, oracleScript)
	if err != nil {
		return err
	}

	return k.VaultsManager.Set(ctx, denom, vm)
}

func (k *Keeper) UpdatesCollateralAsset(
	ctx context.Context,
	denom string,
	minCollateralRatio math.LegacyDec,
	liquidationRatio math.LegacyDec,
	maxDebt math.Int,
	stabilityFee math.LegacyDec,
	mintingFee math.LegacyDec,
	liquidationPenalty math.LegacyDec,
) error {
	// Check if asset alreay be actived
	vm, err := k.GetVaultManager(ctx, denom)
	if err != nil {
		return fmt.Errorf("denom %s not activated", denom)
	}
	amountMinted := vm.Params.MaxDebt.Sub(vm.MintAvailable)

	vm.Params.MinCollateralRatio = minCollateralRatio
	vm.Params.LiquidationRatio = liquidationRatio
	vm.Params.MaxDebt = maxDebt
	vm.Params.StabilityFee = stabilityFee
	vm.Params.MintingFee = mintingFee
	vm.Params.LiquidationPenalty = liquidationPenalty

	vm.MintAvailable = maxDebt.Sub(amountMinted)

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
