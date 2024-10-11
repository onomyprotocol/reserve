package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k Keeper) GetNewAuction(ctx context.Context,
	startTime time.Time,
	initialPrice, item, targetGoal sdk.Coin,
	vaultId uint64,
) (*types.Auction, error) {
	var newAuction types.Auction
	k.Auctions.Walk(ctx, nil, func(key uint64, value types.Auction) (stop bool, err error) {
		if value.VaultId == vaultId {
			newAuction = value
			return true, nil
		}
		return false, nil
	})
	if newAuction.Status == 1 {
		return &newAuction, nil
	}

	return k.NewAuction(ctx, startTime, initialPrice, item, targetGoal, vaultId)
}

func (k Keeper) NewAuction(ctx context.Context,
	startTime time.Time,
	initialPrice, item, targetGoal sdk.Coin,
	vaultId uint64,
) (*types.Auction, error) {
	auctionId, err := k.AuctionIdSeq.Next(ctx)
	if err != nil {
		return nil, err
	}
	params := k.GetParams(ctx)

	startingRate, err := math.LegacyNewDecFromStr(params.StartingRate)
	if err != nil {
		return nil, fmt.Errorf("invalid starting rate params: %v", err)
	}
	lowestRate, err := math.LegacyNewDecFromStr(params.LowestRate)
	if err != nil {
		return nil, fmt.Errorf("invalid lowest rate params: %v", err)
	}
	discountRate, err := math.LegacyNewDecFromStr(params.DiscountRate)
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
		TokenRaised:      sdk.NewCoin(vaultstypes.DefaultMintDenom, math.ZeroInt()),
		VaultId:          vaultId,
	}, nil
}
