package keeper

import (
	"context"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) error {
	params := k.GetParams(ctx)

	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time
	lastAuctionPeriods_unix, err := k.lastestAuctionPeriod.Get(ctx)
	if err != nil {
		return err
	}
	lastAuctionPeriods := time.Unix(lastAuctionPeriods_unix, 0)
	// check if has reached the next auction periods
	if lastAuctionPeriods.Add(params.AuctionPeriods).Before(currentTime) {
		// update latest auction period
		err := k.lastestAuctionPeriod.Set(ctx, lastAuctionPeriods.Add(params.AuctionPeriods).Unix())
		if err != nil {
			return err
		}

		liquidations, err := k.vaultKeeper.GetLiquidations(ctx)
		if err != nil {
			return err
		}

		liquidatedVaults := make([]*vaultstypes.Vault, 0)
		for _, liq := range liquidations {
			liquidatedVaults = append(liquidatedVaults, liq.LiquidatingVaults...)
		}

		// create new auction for this vault
		for _, vault := range liquidatedVaults {
			//calcualte initial price and target price
			auction, isCreate, err := k.GetNewAuction(ctx, currentTime, vault.LiquidationPrice, vault.CollateralLocked, vault.Debt, vault.Id)
			if err != nil {
				return err
			}

			if isCreate {
				err = k.Auctions.Set(ctx, auction.AuctionId, *auction)
				if err != nil {
					return err
				}
				err = k.Bids.Set(ctx, auction.AuctionId, types.BidQueue{AuctionId: auction.AuctionId, Bids: []*types.Bid{}})
				if err != nil {
					return err
				}
				err = k.BidIdSeq.Set(ctx, auction.AuctionId, 0)
				if err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}
		}
	}

	// loop through all auctions
	// get liquidations data then distribute debt & collateral remain
	liquidationMap := make(map[string]*vaultstypes.Liquidation)
	err = k.Auctions.Walk(ctx, nil, func(auctionId uint64, auction types.Auction) (bool, error) {
		bidQueue, err := k.Bids.Get(ctx, auction.AuctionId)
		if err != nil {
			return true, err
		}
		vault, err := k.vaultKeeper.GetVault(ctx, auction.VaultId)
		if err != nil {
			return true, err
		}

		needCleanup := false
		if auction.Status == types.AuctionStatus_AUCTION_STATUS_FINISHED ||
			auction.Status == types.AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL ||
			auction.EndTime.Before(currentTime) {
			liquidation_tmp, ok := liquidationMap[auction.Item.Denom]
			if ok && liquidation_tmp != nil {
				liquidation_tmp.Denom = auction.Item.Denom
				liquidation_tmp.LiquidatingVaults = append(liquidation_tmp.LiquidatingVaults, &vault)
				liquidation_tmp.VaultLiquidationStatus[vault.Id].Sold = liquidation_tmp.VaultLiquidationStatus[vault.Id].Sold.Add(auction.TokenRaised)
				liquidation_tmp.VaultLiquidationStatus[vault.Id].RemainCollateral = liquidation_tmp.VaultLiquidationStatus[vault.Id].RemainCollateral.Add(auction.Item)
			} else {
				liquidation_tmp = &vaultstypes.Liquidation{
					Denom:                  auction.Item.Denom,
					LiquidatingVaults:      []*vaultstypes.Vault{&vault},
					VaultLiquidationStatus: make(map[uint64]*vaultstypes.VaultLiquidationStatus),
				}

				liquidation_tmp.VaultLiquidationStatus[vault.Id] = &vaultstypes.VaultLiquidationStatus{
					Sold:             auction.TokenRaised,
					RemainCollateral: auction.Item,
				}
				liquidationMap[auction.Item.Denom] = liquidation_tmp
			}
			err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, vaultstypes.ModuleName, sdk.NewCoins(liquidationMap[auction.Item.Denom].VaultLiquidationStatus[vault.Id].Sold))
			if err != nil {
				return true, err
			}

			needCleanup = true
			// skip other logic
		}

		if needCleanup {
			k.refundBidders(ctx, bidQueue)

			// clear the auction afterward
			err = k.DeleteAuction(ctx, auction.AuctionId)
			if err != nil {
				return true, err
			}

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

		err = k.fillBids(ctx, auction, bidQueue)
		if err != nil {
			return true, err
		}

		return false, nil
	})

	// Loop through liquidationMap and liquidate
	for _, liq := range liquidationMap {
		err := k.vaultKeeper.Liquidate(ctx, *liq)
		if err != nil {
			return err
		}
	}
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
			}

			bidQueue.Bids[i].IsHandle = true
		}

		// update auction status
		err = k.Auctions.Set(ctx, auction.AuctionId, auction)
		if err != nil {
			return err
		}
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
