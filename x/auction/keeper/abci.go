package keeper

import (
	"context"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) error {
	params := k.GetParams(ctx)

	currentTime := sdk.UnwrapSDKContext(ctx).BlockHeader().Time
	lastAuctionPeriods_unix, err := k.LastestAuctionPeriod.Get(ctx)
	if err != nil {
		return err
	}
	lastAuctionPeriods := time.Unix(lastAuctionPeriods_unix, 0)
	// check if has reached the next auction periods
	fmt.Println("sumtime", lastAuctionPeriods.Add(params.AuctionPeriods).Unix())
	fmt.Println("currentTime", currentTime.Unix())
	// lastAuctionPeriods.Add(params.AuctionPeriods) > currentTime
	if lastAuctionPeriods.Add(params.AuctionPeriods).Before(currentTime) {
		fmt.Println("dem 123")
		// update latest auction period
		err := k.LastestAuctionPeriod.Set(ctx, lastAuctionPeriods.Add(params.AuctionPeriods).Unix())
		if err != nil {
			return err
		}

		liquidations, err := k.vaultKeeper.GetLiquidations(ctx)
		fmt.Println("liquidations", liquidations)
		if err != nil {
			return err
		}
		if len(liquidations) > 0 {
			fmt.Println()
			fmt.Println(liquidations[0].LiquidatingVaults[0].LiquidationPrice.String())
		}

		liquidatedVaults := make([]*vaultstypes.Vault, 0)
		for _, liq := range liquidations {
			liquidatedVaults = append(liquidatedVaults, liq.LiquidatingVaults...)
		}

		// create new auction for this vault
		for _, vault := range liquidatedVaults {
			//calcualte initial price and target price
			auction, err := k.GetNewAuction(ctx, currentTime, k.calculateInitAuctionPrice(ctx, vault.CollateralLocked), vault.CollateralLocked, vault.Debt, vault.Id)
			if err != nil {
				return err
			}
			fmt.Println("+++++1", auction.AuctionId)

			found, err := k.Auctions.Has(ctx, auction.AuctionId)
			if !found {
				err = k.Auctions.Set(ctx, auction.AuctionId, *auction)
				if err != nil {
					return err
				}
				err = k.Bids.Set(ctx, auction.AuctionId, types.BidQueue{})
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

		// true
		// 1728641755
		// 1728641705
		fmt.Println()
		fmt.Println("mmmmmmm")
		fmt.Println(auction.Status)
		fmt.Println(auction.EndTime.Unix())
		fmt.Println(currentTime.Unix())
		if auction.Status == types.AuctionStatus_AUCTION_STATUS_FINISHED ||
			auction.Status == types.AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL ||
			auction.EndTime.Before(currentTime) {

			var liquidation_tmp vaultstypes.Liquidation
			liquidation_tmp.Denom = auction.Item.Denom
			liquidation_tmp.LiquidatingVaults = append(liquidation_tmp.LiquidatingVaults, &vault)

			liquidation_tmp.VaultLiquidationStatus = make(map[uint64]*vaultstypes.VaultLiquidationStatus)

			var vaultLiquidationStatus_tmp vaultstypes.VaultLiquidationStatus

			vaultLiquidationStatus_tmp.Sold = auction.TokenRaised
			vaultLiquidationStatus_tmp.RemainCollateral = auction.Item

			liquidation_tmp.VaultLiquidationStatus[vault.Id] = &vaultLiquidationStatus_tmp

			liquidationMap[auction.Item.Denom] = &liquidation_tmp

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
			fmt.Println("xoa heetttttttt", auction.AuctionId)
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
		fmt.Println("liquidateeeeeeeeeeeeeeeeee")
		_, _, err := k.vaultKeeper.Liquidate(ctx, *liq)
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

	currentRate, err := sdkmath.LegacyNewDecFromStr(auction.CurrentRate)
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

		if currentRate.Mul(auction.InitialPrice.Amount.ToLegacyDec()).TruncateInt().LTE(bid.Amount.Amount) {
			bidderAddr, err := k.authKeeper.AddressCodec().StringToBytes(bid.Bidder)
			if err != nil {
				continue
			}

			receiveRate, err := sdkmath.LegacyNewDecFromStr(bid.ReciveRate)
			if err != nil {
				continue
			}

			receivePrice := receiveRate.Mul(auction.InitialPrice.Amount.ToLegacyDec()).TruncateInt()
			receiveAmt := bid.Amount.Amount.Quo(receivePrice)
			receiveCoin := sdk.NewCoin(itemDenom, receiveAmt)
			// if out of collatheral
			if auction.Item.Amount.LT(receiveAmt) {
				auction.Status = types.AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL
				continue
			}

			err = k.bankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), bidderAddr, sdk.NewCoins(receiveCoin))
			if err != nil {
				continue
			}

			// update auction collatheral
			auction.Item = auction.Item.Sub(receiveCoin)

			auction.TokenRaised = auction.TokenRaised.Add(bid.Amount)

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
