package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/onomyprotocol/reserve/x/auction/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// keepers
		authKeeper types.AccountKeeper
		bankKeeper types.BankKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		// timestamp of lastest auction period
		lastestAuctionPeriod collections.Item[time.Time]

		// Auctions maps auction id with auction struct
		Auctions collections.Map[uint64, types.Auction]

		// Bids maps auction id with bids queue
		Bids collections.Map[uint64, types.BidQueue]

		// BidByAddress maps bidder address + auction id to a bid entry
		BidByAddress collections.Map[collections.Pair[uint64, sdk.AccAddress], types.Bid]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)
	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		Auctions:     collections.NewMap(sb, types.AuctionsPrefix, "auctions", collections.Uint64Key, codec.CollValue[types.Auction](cdc)),
		Bids:         collections.NewMap(sb, types.BidsPrefix, "bids", collections.Uint64Key, codec.CollValue[types.BidQueue](cdc)),
		BidByAddress: collections.NewMap(sb, types.BidByAddressPrefix, "bids_by_address", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.Bid](cdc)),
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// clear all state of the given auction id
func (k Keeper) DeleteAuction(ctx context.Context, auctionId uint64) error {
	err := k.Auctions.Remove(ctx, auctionId)
	if err != nil {
		return fmt.Errorf("failed to remove auction: %s", err)
	}

	// clear the bid queue
	err = k.Bids.Remove(ctx, auctionId)
	if err != nil {
		return fmt.Errorf("failed to remove bid queue: %s", err)
	}

	// clear all bids for that auction id
	rng := collections.NewPrefixedPairRange[uint64, sdk.AccAddress](auctionId)
	return k.BidByAddress.Clear(ctx, rng)
}

// AddBidEntry adds new bid entry for the given auction id
func (k Keeper) AddBidEntry(ctx context.Context, auctionId uint64, bidderAddr sdk.AccAddress, bid types.Bid) error {
	has, err := k.Auctions.Has(ctx, auctionId)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("cannot bid for non-existing/expired auction with id: %v", auctionId)
	}

	has = k.authKeeper.HasAccount(ctx, bidderAddr)
	if !has {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proposer address %s: account does not exist", bid.Bidder)
	}

	bidQueue, err := k.Bids.Get(ctx, auctionId)
	if err != nil {
		return err
	}

	_, has = bidQueue.Bids[bid.Bidder]
	if has {
		return fmt.Errorf("bid entry already exist for address %s and auction %v", bid.Bidder, auctionId)
	}

	bidQueue.Bids[bid.Bidder] = &bid
	err = k.Bids.Set(ctx, auctionId, bidQueue)
	if err != nil {
		return err
	}

	err = k.BidByAddress.Set(ctx, collections.Join(auctionId, bidderAddr), bid)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBidEntry udpdate existing bid entry for the given auction id
func (k Keeper) UpdateBidEntry(ctx context.Context, auctionId uint64, bidderAddr sdk.AccAddress, updatedBid types.Bid) error {
	has, err := k.Auctions.Has(ctx, auctionId)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("cannot bid for non-existing/expired auction with id: %v", auctionId)
	}

	has = k.authKeeper.HasAccount(ctx, bidderAddr)
	if !has {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proposer address %s: account does not exist", updatedBid.Bidder)
	}

	bidQueue, err := k.Bids.Get(ctx, auctionId)
	if err != nil {
		return err
	}

	currBid, has := bidQueue.Bids[updatedBid.Bidder]
	if !has {
		return fmt.Errorf("bid entry does not exist for address %s and auction %v", updatedBid.Bidder, auctionId)
	}

	// locked additional amount when bidder raise the amount
	// or refund amount when bidder lower the amount
	if currBid.Amount.Amount.Equal(updatedBid.Amount.Amount) {
		return errors.New("updated bidding amount must be different from the current bidding amount")
	}

	// update the entry
	bidQueue.Bids[currBid.Bidder] = currBid
	err = k.Bids.Set(ctx, auctionId, bidQueue)
	if err != nil {
		return err
	}

	err = k.BidByAddress.Set(ctx, collections.Join(auctionId, bidderAddr), *currBid)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) lockedToken(ctx context.Context, amt sdk.Coins, bidderAdrr sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAdrr, types.ModuleName, amt)
}

func (k Keeper) refundToken(ctx context.Context, amt sdk.Coins, bidderAdrr sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidderAdrr, amt)
}
