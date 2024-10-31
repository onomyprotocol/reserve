package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k Keeper) handleLiquidation(ctx context.Context, mintDenom string) error {
	// Implement the logic for handling liquidation
	// This might include checking conditions, updating states, and logging events
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

		liquidations, err := k.vaultKeeper.GetLiquidations(ctx, mintDenom)
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
	liquidationMap, err := k.newLiquidateMap(ctx, mintDenom, params)
	if err != nil {
		return err
	}

	// Loop through liquidationMap and liquidate
	for _, liq := range liquidationMap {
		err := k.vaultKeeper.Liquidate(ctx, *liq, mintDenom)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// TODO: testing
// TODO: implement the logic for handling liquidation per mint denom
func (k Keeper) newLiquidateMap(ctx context.Context, mintDenom string, params types.Params) (map[string]*vaultstypes.Liquidation, error) {
	// loop through all auctions
	// get liquidations data then distribute debt & collateral remain
	liquidationMap := make(map[string]*vaultstypes.Liquidation)
	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time

	err := k.Auctions.Walk(ctx, nil, func(auctionId uint64, auction types.Auction) (bool, error) {
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
	return liquidationMap, err
}
