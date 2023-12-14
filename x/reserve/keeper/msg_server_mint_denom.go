package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintDenom(goCtx context.Context, msg *types.MsgMintDenom) (*types.MsgMintDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintDenomResponse{}, nil
}
