package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// InitGenesis - Init store state from genesis data
//
// CONTRACT: old coins from the FeeCollectionKeeper need to be transferred through
// a genesis port script to the new fee collector account
func (k *Keeper) InitGenesis(ctx context.Context, data types.GenesisState) error {
	err := k.SetParams(ctx, data.Params)
	if err != nil {
		return err
	}
	for _, vm := range data.VaultManagers {
		err := k.VaultsManager.Set(ctx, vm.Denom, vm)
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
		k.LastUpdateTime.Set(ctx, *data.LastUpdate)
	} else {
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		k.LastUpdateTime.Set(ctx, types.LastUpdate{Time: sdkCtx.BlockTime()})
	}

	k.ShortfallAmount.Set(ctx, data.ShortfallAmount)

	return nil
}
