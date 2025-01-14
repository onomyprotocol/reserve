package keeper

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/auction/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) error {
	if err := k.SetParams(ctx, genState.Params); err != nil {
		return err
	}
	err := k.LastestAuctionPeriods.Set(ctx, ctx.BlockTime().Unix())
	if err != nil {
		return err
	}

	for _, auction := range genState.Auctions {
		err := k.Auctions.Set(ctx, auction.AuctionId, auction)
		if err != nil {
			return err
		}
	}

	allBids := map[uint64][]types.Bid{}
	for _, bidByAddress := range genState.BidByAddress {
		err := k.BidByAddress.Set(ctx, collections.Join(bidByAddress.AuctionId, sdk.AccAddress(bidByAddress.Bidder)), bidByAddress.Bids)
		if err != nil {
			return err
		}
		allBids[bidByAddress.AuctionId] = append(allBids[bidByAddress.AuctionId], bidByAddress.Bids.Bids...)
	}

	for auctionId, bids:= range allBids {
		bidQueue := types.BidQueue{
			AuctionId: auctionId,
			Bids: bids,
		}
		err := k.Bids.Set(ctx, auctionId, bidQueue)
		if err != nil {
			return err
		}
	}

	err = k.AuctionIdSeq.Set(ctx, genState.AuctionSequence)
	if err != nil {
		return err
	}

	for _, bidSequence := range genState.BidSequences {
		err = k.BidIdSeq.Set(ctx, bidSequence.AuctionId, bidSequence.Sequence)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	lastestAuctionPeriod, err := k.LastestAuctionPeriods.Get(ctx)
	if err != nil {
		panic(err)
	}

	auctions, err := k.GetAllAuctions(ctx)
	if err != nil {
		panic(err)
	}

	bidsByAddress, err := k.GetAllBidByAddress(ctx)
	if err != nil {
		panic(err)
	}

	auctionIdSequence, err := k.AuctionIdSeq.Peek(ctx)
	if err != nil {
		panic(err)
	}

	bidsIdSequence, err := k.GetAllBidsSequence(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params: params,
		Auctions: auctions,
		BidByAddress: bidsByAddress,
		AuctionSequence: auctionIdSequence,
		LastestAuctionPeriods: lastestAuctionPeriod,
		BidSequences: bidsIdSequence,
	}
}
