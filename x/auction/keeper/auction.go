package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

func (k Keeper) NewAuction(ctx context.Context,
	startTime, currentTime time.Time,
	initialPrice, item,
	targetGoal sdk.Coin,
	startingRate string,
) (*types.Auction, error) {
	auctionId, err := k.AuctionIdSeq.Next(ctx)
	if err != nil {
		return nil, err
	}

	return &types.Auction{
		StartTime:        startTime,
		AuctionId:        auctionId,
		InitialPrice:     initialPrice,
		Item:             item,
		CurrentRate:      startingRate,
		LastDiscountTime: currentTime,
		Status:           types.AuctionStatus_AUCTION_STATUS_ACTIVE,
		TargetGoal:       targetGoal,
	}, nil
}
