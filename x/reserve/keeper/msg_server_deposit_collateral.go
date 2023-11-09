package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

func (k msgServer) DepositCollateral(goCtx context.Context, msg *types.MsgDepositCollateral) (*types.MsgDepositCollateralResponse, error) {
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

	vault, found := k.GetVault(ctx, msg.Uid)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrVaultNotFound, "Vault: %s", strconv.FormatUint(msg.Uid, 10))
	}

	k.SetVault(ctx, types.Vault{
		Uid:        msg.Uid,
		Collateral: collateral.Add(vault.Collateral),
	})

	return &types.MsgDepositCollateralResponse{}, nil
}
