package keeper

import (
	"context"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

func (k Keeper) NewAuction(ctx context.Context,
	startTime time.Time,
	initialPrice, item,
	targetGoal sdk.Coin,
	vaultId uint64,
) (*types.Auction, error) {
	auctionId, err := k.AuctionIdSeq.Next(ctx)
	if err != nil {
		return nil, err
	}
	params := k.GetParams(ctx)

	startingRate, err := sdkmath.LegacyNewDecFromStr(params.StartingRate)
	if err != nil {
		return nil, fmt.Errorf("invalid starting rate params: %v", err)
	}
	lowestRate, err := sdkmath.LegacyNewDecFromStr(params.LowestRate)
	if err != nil {
		return nil, fmt.Errorf("invalid lowest rate params: %v", err)
	}
	discountRate, err := sdkmath.LegacyNewDecFromStr(params.DiscountRate)
	if err != nil {
		return nil, fmt.Errorf("invalid discount rate params: %v", err)
	}
	endTime := startTime.Add(time.Duration(startingRate.Sub(lowestRate).Quo(discountRate).Ceil().RoundInt64() * int64(params.ReduceStep)))

	return &types.Auction{
		StartTime:        startTime,
		EndTime:          endTime,
		AuctionId:        auctionId,
		InitialPrice:     initialPrice,
		Item:             item,
		CurrentRate:      params.StartingRate,
		LastDiscountTime: startTime,
		Status:           types.AuctionStatus_AUCTION_STATUS_ACTIVE,
		TargetGoal:       targetGoal,
		VaultId:          vaultId,
	}, nil
}
