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
		return &types.MsgBidResponse{
			Response: "Failed to submit bid",
		}, err
	}
	bidIdSeq, err := k.BidIdSeq.Get(ctx, msg.AuctionId)
	if err != nil {
		return nil, err
	}

	newBidId := bidIdSeq + 1
	err = k.BidIdSeq.Set(ctx, msg.AuctionId, newBidId)
	if err != nil {
		return nil, err
	}

	bid := types.Bid{
		BidId:      newBidId,
		Bidder:     msg.Bidder,
		Amount:     msg.Amount,
		ReciveRate: msg.ReciveRate,
		IsHandle:   false,
	}
	err = k.AddBidEntry(ctx, msg.AuctionId, bidderAddr, bid)
	if err != nil {
		return nil, err
	}

	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(sdk.NewEvent(
		types.EventAddBid,
		sdk.NewAttribute(types.AttributeKeyBidEntry, fmt.Sprintf("bidder %s has submit an entry with amount: %s", msg.Bidder, msg.Amount.String())),
	))

	return &types.MsgBidResponse{
		Response: "Bid Accepted",
	}, nil
}

func (k msgServer) CancelBid(ctx context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {

	err := k.CancelBidEntry(ctx, msg.AuctionId, msg.BidId)
	if err != nil {
		return nil, err
	}

	sdk.UnwrapSDKContext(ctx).EventManager().EmitEvent(sdk.NewEvent(
		types.EventUpdateBid,
		sdk.NewAttribute(types.AttributeKeyBidEntry, fmt.Sprintf("cancel bid id %v for auction %v", msg.BidId, msg.AuctionId)),
	))

	return &types.MsgCancelBidResponse{}, nil
}