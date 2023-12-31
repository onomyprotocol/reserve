package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

func (k msgServer) CreateVault(goCtx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// CoinAmsg and CoinBmsg pre-sort from raw msg
	collateral, err := sdk.ParseCoinNormalized(msg.Collateral)
	if err != nil {
		panic(err)
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.validateSenderBalance(ctx, creator, sdk.NewCoins(collateral)); err != nil {
		return nil, err
	}

	uid := k.GetUidCount(ctx)

	k.SetVault(ctx, types.Vault{
		Uid:        uid,
		Collateral: collateral,
	})

	_ = ctx

	return &types.MsgCreateVaultResponse{}, nil
}
