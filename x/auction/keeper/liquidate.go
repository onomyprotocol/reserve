package keeper

import (
	"context"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k Keeper) handleLiquidation(ctx context.Context, mintDenom string) error {
	params := k.GetParams(ctx)

	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time
	lastestAuctionPeriod, err := k.LastestAuctionPeriods.Get(ctx, "LastestAuctionPeriods")
	if err != nil {
		return err
	}
	lastAuctionPeriods := time.Unix(lastestAuctionPeriod, 0)
	// check if has reached the next auction periods
	if lastAuctionPeriods.Add(params.AuctionPeriods).Before(currentTime) {
		// update latest auction period
		err = k.LastestAuctionPeriods.Set(ctx, "LastestAuctionPeriods", lastAuctionPeriods.Add(params.AuctionPeriods).Unix())
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
				err = k.Bids.Set(ctx, auction.AuctionId, types.BidQueue{AuctionId: auction.AuctionId, Bids: []types.Bid{}})
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

	return nil
}

// TODO: testing
// TODO: implement the logic for handling liquidation per mint denom
func (k Keeper) newLiquidateMap(ctx context.Context, mintDenom string, params types.Params) (map[string]*vaultstypes.Liquidation, error) {
	// loop through all auctions
	// get liquidations data then distribute debt & collateral remain
	liquidationMap := make(map[string]*vaultstypes.Liquidation)
	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time

	// loop through all auctions
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
		currentRate := math.LegacyMustNewDecFromStr(auction.CurrentRate)
		lowestRate := math.LegacyMustNewDecFromStr(params.LowestRate)
		if auction.Status == types.AuctionStatus_AUCTION_STATUS_FINISHED ||
			auction.Status == types.AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL ||
			currentRate.Equal(lowestRate) {
			liquidation_tmp, ok := liquidationMap[auction.Item.Denom]
			if ok && liquidation_tmp != nil {
				liquidation_tmp.DebtDenom = auction.Item.Denom
				liquidation_tmp.LiquidatingVaults = append(liquidation_tmp.LiquidatingVaults, &vault)
				liquidation_tmp.VaultLiquidationStatus[vault.Id] = &vaultstypes.VaultLiquidationStatus{
					Sold:             auction.TokenRaised,
					RemainCollateral: auction.Item,
				}
			} else {
				liquidation_tmp = &vaultstypes.Liquidation{
					DebtDenom:              auction.Item.Denom,
					MintDenom:              mintDenom,
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
			err = k.refundBidders(ctx, bidQueue)
			if err != nil {
				return true, err
			}

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

	if err != nil {
		return nil, err
	}
	return liquidationMap, nil
}
