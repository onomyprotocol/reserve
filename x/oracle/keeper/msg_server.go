package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k Keeper) RequestBandRates(goCtx context.Context, msg *types.MsgRequestBandRates) (*types.MsgRequestBandRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bandOracleRequest := k.GetBandOracleRequest(ctx, msg.RequestId)
	if bandOracleRequest == nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidBandRequest, "Band oracle request not found!")
	}

	if err := k.RequestBandOraclePrices(ctx, bandOracleRequest); err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	return &types.MsgRequestBandRatesResponse{}, nil
}
