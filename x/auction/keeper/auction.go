package keeper

import (
	"context"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

// return aution, is create, error
func (k Keeper) GetNewAuction(ctx context.Context,
	startTime time.Time,
	initialPrice math.LegacyDec,
	item, targetGoal sdk.Coin,
	vaultId uint64,
) (*types.Auction, bool, error) {
	var newAuction *types.Auction
	k.Auctions.Walk(ctx, nil, func(key uint64, value types.Auction) (stop bool, err error) {
		if value.VaultId == vaultId {
			newAuction = &value
			return true, nil
		}
		return false, nil
	})
	if newAuction != nil {
		return newAuction, false, nil
	}
	newAuction, err := k.NewAuction(ctx, startTime, initialPrice, item, targetGoal, vaultId)
	if err != nil {
		return newAuction, true, err
	}
	return newAuction, true, nil
}

func (k Keeper) NewAuction(ctx context.Context,
	startTime time.Time,
	initialPrice math.LegacyDec,
	item, targetGoal sdk.Coin,
	vaultId uint64,
) (*types.Auction, error) {
	auctionId, err := k.AuctionIdSeq.Next(ctx)
	if err != nil {
		return nil, err
	}
	params := k.GetParams(ctx)

	return &types.Auction{
		AuctionId:        auctionId,
		InitialPrice:     initialPrice.String(),
		Item:             item,
		CurrentRate:      params.StartingRate,
		LastDiscountTime: startTime,
		Status:           types.AuctionStatus_AUCTION_STATUS_ACTIVE,
		TargetGoal:       targetGoal,
		TokenRaised:      sdk.NewCoin(vaultstypes.DefaultMintDenom, math.ZeroInt()),
		VaultId:          vaultId,
	}, nil
}
