package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) error {
	// get allowed mint denom

	err := k.handleLiquidation(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) fillBids(ctx context.Context, auction types.Auction, bidQueue types.BidQueue) error {
	itemDenom := auction.Item.Denom

	currentRate, err := math.LegacyNewDecFromStr(auction.CurrentRate)
	if err != nil {
		return err
	}

	vault, err := k.vaultKeeper.GetVault(ctx, auction.VaultId)
	if err != nil {
		return err
	}

	for i, bid := range bidQueue.Bids {
		if bid.IsHandle {
			continue
		}

		initPrices, err := math.LegacyNewDecFromStr(auction.InitialPrice)
		if err != nil {
			continue
		}

		receivePrice, err := math.LegacyNewDecFromStr(bid.RecivePrice)
		if err != nil {
			continue
		}

		// Only handle bid if: (rate * init price) <= receive price
		if currentRate.Mul(initPrices).LTE(receivePrice) {
			bidderAddr, err := k.authKeeper.AddressCodec().StringToBytes(bid.Bidder)
			if err != nil {
				continue
			}

			receiveAmt := bid.Amount.Amount.ToLegacyDec().Quo(receivePrice).TruncateInt()
			receiveCoin := sdk.NewCoin(itemDenom, receiveAmt)
			// if out of collatheral
			if auction.Item.Amount.LT(receiveAmt) {
				auction.Status = types.AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL

				amountBuy := auction.Item.Amount.ToLegacyDec().Mul(receivePrice).TruncateInt()

				amountRefund := bid.Amount.Amount.Sub(amountBuy)
				// send all auction item
				err = k.bankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), bidderAddr, sdk.NewCoins(auction.Item))
				if err != nil {
					continue
				}

				err = k.refundToken(ctx, sdk.NewCoins(sdk.NewCoin(bid.Amount.Denom, amountRefund)), bid.Bidder)
				if err != nil {
					continue
				}

				auction.Item = sdk.NewCoin(auction.Item.Denom, math.ZeroInt())
				auction.TokenRaised = auction.TokenRaised.Add(sdk.NewCoin(bid.Amount.Denom, amountBuy))
				bidQueue.Bids[i].IsHandle = true
				break
			} else {
				err = k.bankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), bidderAddr, sdk.NewCoins(receiveCoin))
				if err != nil {
					continue
				}

				// update auction collatheral
				auction.Item = auction.Item.Sub(receiveCoin)

				auction.TokenRaised = auction.TokenRaised.Add(bid.Amount)
			}

			if auction.TokenRaised.IsGTE(auction.TargetGoal) {
				auction.Status = types.AuctionStatus_AUCTION_STATUS_FINISHED
				bidQueue.Bids[i].IsHandle = true
				break
			}

			bidQueue.Bids[i].IsHandle = true
		}
	}

	// update auction status
	err = k.Auctions.Set(ctx, auction.AuctionId, auction)
	if err != nil {
		return err
	}

	// update bid queue
	err = k.Bids.Set(ctx, auction.AuctionId, bidQueue)
	if err != nil {
		return err
	}

	return nil

}

func (k Keeper) refundBidders(ctx context.Context, bidQueue types.BidQueue) error {
	for _, bid := range bidQueue.Bids {
		if bid.IsHandle {
			continue
		}

		err := k.refundToken(ctx, sdk.NewCoins(bid.Amount), bid.Bidder)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) discountRate(auction types.Auction, params types.Params) (string, error) {
	lowestRate, err := math.LegacyNewDecFromStr(params.LowestRate)
	if err != nil {
		return math.LegacyZeroDec().String(), err
	}

	discountRate, err := math.LegacyNewDecFromStr(params.DiscountRate)
	if err != nil {
		return math.LegacyZeroDec().String(), err
	}

	currentRate, err := math.LegacyNewDecFromStr(auction.CurrentRate)
	if err != nil {
		return math.LegacyZeroDec().String(), err
	}

	if currentRate.LT(lowestRate) || currentRate.Sub(discountRate).LT(lowestRate) {
		return currentRate.String(), nil
	}

	newCurrentRate := currentRate.Sub(discountRate)

	return newCurrentRate.String(), nil
}
