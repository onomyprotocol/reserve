package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
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

func (k msgServer) UpdateParams(ctx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (k msgServer) Bid(ctx context.Context, msg *types.MsgBid) (*types.MsgBidResponse, error) {
	bidderAddr, err := k.authKeeper.AddressCodec().StringToBytes(msg.Bidder)
	if err != nil {
		return nil, err
	}

	bid := types.Bid{
		Bidder: msg.Bidder,
		Amount: msg.Amount,
	}
	err = k.AddBidEntry(ctx, msg.AuctionId, bidderAddr, bid)
	if err != nil {
		return nil, err
	}

	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(sdk.NewEvent(
		types.EventAddBid,
		sdk.NewAttribute(types.AttributeKeyBidEntry, fmt.Sprintf("bidder %s has submit an entry with amount: %s", msg.Bidder, msg.Amount.String())),
	))

	return &types.MsgBidResponse{}, nil
}

func (k msgServer) UpdateBid(ctx context.Context, msg *types.MsgUpdateBid) (*types.MsgUpdateBidResponse, error) {
	bidderAddr, err := k.authKeeper.AddressCodec().StringToBytes(msg.Bidder)
	if err != nil {
		return nil, err
	}

	bid := types.Bid{
		Bidder: msg.Bidder,
		Amount: msg.Amount,
	}
	err = k.UpdateBidEntry(ctx, msg.AuctionId, bidderAddr, bid)
	if err != nil {
		return nil, err
	}

	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(sdk.NewEvent(
		types.EventUpdateBid,
		sdk.NewAttribute(types.AttributeKeyBidEntry, fmt.Sprintf("bidder %s has update their entry to amount: %s", msg.Bidder, msg.Amount.String())),
	))

	return &types.MsgUpdateBidResponse{}, nil
}
