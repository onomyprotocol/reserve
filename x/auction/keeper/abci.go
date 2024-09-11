package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) error {
	params := k.GetParams(ctx)

	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time
	lastAuctionPeriods, err := k.lastestAuctionPeriod.Get(ctx)
	if err != nil {
		return err
	}

	// check if has reached the next auction periods
	if lastAuctionPeriods.Add(params.AuctionPeriods).Before(currentTime) {
		return nil
	}

	k.lastestAuctionPeriod.Set(ctx, lastAuctionPeriods.Add(params.AuctionDurations))

	// TODO: check vault module for liquidate vault

	// loop through all auctions
	err = k.Auctions.Walk(ctx, nil, func(auctionId uint64, auction types.Auction) (bool, error) {
		// check if auction is ended or a bidder won
		if auction.EndTime.After(currentTime) ||
			auction.Status == types.AuctionStatus_AUCTION_STATUS_EXPIRED ||
			auction.Status == types.AuctionStatus_AUCTION_STATUS_FINISHED {
			if auction.FinalBid == nil ||
				auction.FinalBid.Bidder == "" ||
				auction.FinalBid.Amount.IsZero() {
				// TODO: notify vault module about auction without winner
			}

			bidderAddr, err := k.authKeeper.AddressCodec().StringToBytes(auction.FinalBid.Bidder)
			if err != nil {
				err := k.revertFinishedStatus(ctx, auction, currentTime)
				return err == nil, err
			}

			spendable := k.bankKeeper.SpendableCoins(ctx, bidderAddr)
			if spendable.AmountOf(auction.FinalBid.Amount.Denom).LT(auction.FinalBid.Amount.Amount) {
				// if bidder does not have enough token to pay, revert the status of auction
				err := k.revertFinishedStatus(ctx, auction, currentTime)
				return err == nil, err
			}

			// send the bid amount to auction module
			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAddr, types.ModuleName, sdk.NewCoins(auction.FinalBid.Amount))
			if err != nil {
				err := k.revertFinishedStatus(ctx, auction, currentTime)
				return err == nil, err
			}

			// send the liquidate assets to auction winner
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidderAddr, auction.Items)
			if err != nil {
				err := k.revertFinishedStatus(ctx, auction, currentTime)
				return err == nil, err
			}

			// TODO: notify vault module about the winner and return raised token from the auction

			// clear the auction afterward
			err = k.DeleteAuction(ctx, auction.AuctionId)
			if err != nil {
				return true, err
			}

			// skip other logic
			return false, nil
		}

		// check if reach next reduce step
		if auction.LastDiscountTime.Add(params.ReduceStep).Before(currentTime) {
			// get new discount rate
			newRate, err := k.discountRate(auction, params)
			if err != nil {
				return true, err
			}

			// apply new changes
			auction.CurrentRate = newRate
			auction.LastDiscountTime = auction.LastDiscountTime.Add(params.ReduceStep)

			// update new rate and last discount time
			err = k.Auctions.Set(ctx, auctionId, auction)
			if err != nil {
				return true, err
			}
		}

		highestBid, amt, err := k.checkBidEntry(ctx, auction)
		if err != nil {
			return true, err
		}
		if highestBid == "" || amt.Amount.IsZero() {
			return false, nil
		}

		// update status and final bid
		auction.Status = types.AuctionStatus_AUCTION_STATUS_FINISHED
		auction.FinalBid = &types.Bid{
			Bidder: highestBid,
			Amount: amt,
		}
		err = k.Auctions.Set(ctx, auctionId, auction)
		if err != nil {
			return true, err
		}

		return false, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) revertFinishedStatus(ctx context.Context, auction types.Auction, currTime time.Time) error {
	auction.FinalBid = nil
	if currTime.After(auction.EndTime) {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_EXPIRED
	} else {
		auction.Status = types.AuctionStatus_AUCTION_STATUS_ACTIVE
	}

	return k.Auctions.Set(ctx, auction.AuctionId, auction)
}

func (k Keeper) checkBidEntry(ctx context.Context, auction types.Auction) (highestBidder string, amt sdk.Coin, err error) {
	denom := auction.InitialPrice.Denom

	bidQueue, err := k.Bids.Get(ctx, auction.AuctionId)
	if err != nil {
		return "", sdk.NewCoin(denom, sdkmath.ZeroInt()), err
	}

	currentRate, err := sdkmath.LegacyNewDecFromStr(auction.CurrentRate)
	if err != nil {
		return "", sdk.NewCoin(denom, sdkmath.ZeroInt()), err
	}

	currentPriceAmt := sdkmath.LegacyNewDecFromInt(auction.InitialPrice.Amount).Mul(currentRate).RoundInt()

	maxBidder := struct {
		addr string
		amt  sdkmath.Int
	}{
		addr: "",
		amt:  currentPriceAmt,
	}
	for addr, bid := range bidQueue.Bids {
		// get the highest bid that greater or equal the current price
		if bid.Amount.Amount.GT(maxBidder.amt) {
			maxBidder.addr = addr
			maxBidder.amt = bid.Amount.Amount
		}
	}

	if maxBidder.addr == "" {
		return "", sdk.NewCoin(denom, sdkmath.ZeroInt()), err
	}

	return maxBidder.addr, sdk.NewCoin(denom, maxBidder.amt), nil

}

func (k Keeper) discountRate(auction types.Auction, params types.Params) (string, error) {
	lowestRate, err := sdkmath.LegacyNewDecFromStr(params.LowestRate)
	if err != nil {
		return sdkmath.LegacyZeroDec().String(), err
	}

	discountRate, err := sdkmath.LegacyNewDecFromStr(params.DiscountRate)
	if err != nil {
		return sdkmath.LegacyZeroDec().String(), err
	}

	currentRate, err := sdkmath.LegacyNewDecFromStr(auction.CurrentRate)
	if err != nil {
		return sdkmath.LegacyZeroDec().String(), err
	}

	if currentRate.LT(lowestRate) || currentRate.Sub(discountRate).LT(lowestRate) {
		return currentRate.String(), nil
	}

	newCurrentRate := currentRate.Sub(discountRate)

	return newCurrentRate.String(), nil
}
