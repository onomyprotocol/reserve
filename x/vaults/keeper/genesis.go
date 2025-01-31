package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// InitGenesis - Init store state from genesis data
func (k *Keeper) InitGenesis(ctx context.Context, data types.GenesisState) error {
	err := k.SetParams(ctx, data.Params)
	if err != nil {
		return err
	}
	for _, vm := range data.VaultManagers {
		key := getVMKey(vm.Denom, vm.Params.MintDenom)
		err := k.VaultsManager.Set(ctx, key, vm)
		if err != nil {
			return err
		}
	}
	for _, vault := range data.Vaults {
		err := k.SetVault(ctx, vault)
		if err != nil {
			return err
		}
	}

	if data.LastUpdate != nil {
		err = k.LastUpdateTime.Set(ctx, *data.LastUpdate)
		if err != nil {
			return err
		}
	} else {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err = k.LastUpdateTime.Set(ctx, types.LastUpdate{Time: sdkCtx.BlockTime()})
		if err != nil {
			return err
		}
	}

	for _, shortfall := range data.ShortfallAmounts {
		err = k.ShortfallAmount.Set(ctx, shortfall.Denom, shortfall.Amount)
		if err != nil {
			return err
		}
	}

	return k.VaultsSequence.Set(ctx, data.VaultSequence)
}

func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState{
	vaults, err := k.GetAllVault(ctx)
	if err != nil {
		panic(err)
	}

	vms, err := k.GetAllVaultManagers(ctx)
	if err != nil {
		panic(err)
	}

	shortfalls, err := k.GetAllShortfall(ctx)
	if err != nil {
		panic(err)
	}

	lastUpdate, err := k.LastUpdateTime.Get(ctx)
	if err != nil {
		panic(err)
	}

	vaultSequence, err := k.VaultsSequence.Peek(ctx)
	if err != nil {
		panic(err)
	}

	genState := types.NewGenesisState(
		k.GetParams(ctx),
		vms,
		vaults,
		&lastUpdate,
		shortfalls,
		vaultSequence,
	)

	return genState
}
