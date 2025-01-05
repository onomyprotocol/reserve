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
	VaultsManager   collections.Map[string, types.VaultManager]
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
		VaultsManager:   collections.NewMap(sb, types.VaultManagerKeyPrefix, "vaultmanagers", collections.StringKey, codec.CollValue[types.VaultManager](cdc)),
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
	CollateralDenom string,
	CollateralSymbol string,
	MintDenom string,
	MintSymbol string,
	minCollateralRatio math.LegacyDec,
	liquidationRatio math.LegacyDec,
	maxDebt math.Int,
	stabilityFee math.LegacyDec,
	mintingFee math.LegacyDec,
	liquidationPenalty math.LegacyDec,
	collateralOracleScript int64,
	mintOracleScript int64,
	collateralDecimals, mintDecimals uint64,
) error {
	// Check if asset alreay be actived
	actived, vmKey := k.IsActived(ctx, CollateralDenom, MintDenom)
	if actived {
		return fmt.Errorf("denom %s already be actived", CollateralDenom)
	}
	vm := types.VaultManager{
		Denom:  CollateralDenom,
		Symbol: CollateralSymbol,
		Params: types.VaultManagerParams{
			MintDenom:          MintDenom,
			MintSymbol:         MintSymbol,
			MinCollateralRatio: minCollateralRatio,
			LiquidationRatio:   liquidationRatio,
			MaxDebt:            maxDebt,
			StabilityFee:       stabilityFee,
			LiquidationPenalty: liquidationPenalty,
			MintingFee:         mintingFee,
		},
		MintAvailable: maxDebt,
	}
	err := k.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, CollateralSymbol, collateralOracleScript)
	if err != nil {
		return err
	}

	err = k.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, MintSymbol, mintOracleScript)
	if err != nil {
		return err
	}

	err = k.OracleKeeper.SetPairDecimalsRate(ctx, CollateralSymbol, MintSymbol, collateralDecimals, mintDecimals)
	if err != nil {
		return err
	}
	return k.VaultsManager.Set(ctx, vmKey, vm)
}

func (k *Keeper) UpdatesCollateralAsset(
	ctx context.Context,
	denom string,
	CollateralSymBol string,
	mintDenom string,
	minCollateralRatio math.LegacyDec,
	liquidationRatio math.LegacyDec,
	maxDebt math.Int,
	stabilityFee math.LegacyDec,
	mintingFee math.LegacyDec,
	liquidationPenalty math.LegacyDec,
	CollateralOracleScript int64,
) error {
	// Check if asset alreay be actived
	key := getVMKey(denom, mintDenom)
	vm, err := k.GetVaultManager(ctx, denom, mintDenom)
	if err != nil {
		return fmt.Errorf("pair %s not activated", key)
	}
	amountMinted := vm.Params.MaxDebt.Sub(vm.MintAvailable)

	vm.Params.MinCollateralRatio = minCollateralRatio
	vm.Params.LiquidationRatio = liquidationRatio
	vm.Params.MaxDebt = maxDebt
	vm.Params.StabilityFee = stabilityFee
	vm.Params.MintingFee = mintingFee
	vm.Params.LiquidationPenalty = liquidationPenalty

	vm.MintAvailable, err = maxDebt.SafeSub(amountMinted)
	if err != nil {
		return err
	}

	err = k.OracleKeeper.AddNewSymbolToBandOracleRequest(ctx, CollateralSymBol, CollateralOracleScript)
	if err != nil {
		return err
	}

	return k.VaultsManager.Set(ctx, key, vm)
}

func (k *Keeper) GetVaultManager(
	ctx context.Context,
	collateralDenom string,
	mintDenom string,
) (types.VaultManager, error) {
	key := getVMKey(collateralDenom, mintDenom)
	vm, err := k.VaultsManager.Get(ctx, key)
	if err != nil {
		return types.VaultManager{}, err
	}
	return vm, nil
}

func (k *Keeper) IsActived(
	ctx context.Context,
	collateralDenom string,
	mintDenom string,
) (bool, string) {
	keyStr := getVMKey(collateralDenom, mintDenom)
	has, _ := k.VaultsManager.Has(ctx, keyStr)
	return has, keyStr
}

func getVMKey(
	collateralDenom string,
	mintDenom string,
) string {
	return fmt.Sprintf("%s-%s", collateralDenom, mintDenom)
}

func (k *Keeper) GetAllowedMintDenoms(ctx context.Context) []string {
	return k.GetParams(ctx).AllowedMintDenom
}

func (k *Keeper) mintDebt(ctx context.Context, vmKey string, vm types.VaultManager, coin sdk.Coin) error {
	err := k.BankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	vm.MintAvailable, err = vm.MintAvailable.SafeSub(coin.Amount)
	if err != nil {
		return err
	}
	return k.VaultsManager.Set(ctx, vmKey, vm)
}

func (k *Keeper) burnDebt(ctx context.Context, vmKey string, vm types.VaultManager, coin sdk.Coin) error {
	err := k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	vm.MintAvailable = vm.MintAvailable.Add(coin.Amount)
	return k.VaultsManager.Set(ctx, vmKey, vm)
}