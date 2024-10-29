package keeper

import (
	"context"
	"fmt"

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
		authKeeper   types.AccountKeeper
		bankKeeper   types.BankKeeper
		vaultKeeper  types.VaultKeeper
		OracleKeeper types.OracleKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		// timestamp of lastest auction period (Unix timestamp)
		LastestAuctionPeriod int64

		AuctionIdSeq collections.Sequence

		// bid id seq by auction id
		BidIdSeq collections.Map[uint64, uint64]

		// Auctions maps auction id with auction struct
		Auctions collections.Map[uint64, types.Auction]

		// Bids maps auction id with bids queue
		Bids collections.Map[uint64, types.BidQueue]

		// BidByAddress maps bidder auction id + address to a bid entry
		BidByAddress collections.Map[collections.Pair[uint64, sdk.AccAddress], types.Bids]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	vk types.VaultKeeper,
	ok types.OracleKeeper,
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
		authKeeper:   ak,
		bankKeeper:   bk,
		vaultKeeper:  vk,
		OracleKeeper: ok,
		AuctionIdSeq: collections.NewSequence(sb, types.AuctionIdSeqPrefix, "auction_id_sequence"),
		BidIdSeq:     collections.NewMap(sb, types.BidIdSeqPrefix, "bid_id_sequence", collections.Uint64Key, collections.Uint64Value),
		Auctions:     collections.NewMap(sb, types.AuctionsPrefix, "auctions", collections.Uint64Key, codec.CollValue[types.Auction](cdc)),
		Bids:         collections.NewMap(sb, types.BidsPrefix, "bids", collections.Uint64Key, codec.CollValue[types.BidQueue](cdc)),
		BidByAddress: collections.NewMap(sb, types.BidByAddressPrefix, "bids_by_address", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.Bids](cdc)), //nolint:staticcheck
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

	// clear the bid seq tracking
	err = k.BidIdSeq.Remove(ctx, auctionId)
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

	bid.Index = uint64(len(bidQueue.Bids))

	bidQueue.Bids = append(bidQueue.Bids, &bid)
	err = k.Bids.Set(ctx, auctionId, bidQueue)
	if err != nil {
		return err
	}

	found, err := k.BidByAddress.Has(ctx, collections.Join(auctionId, bidderAddr))
	if found {
		bids, err := k.BidByAddress.Get(ctx, collections.Join(auctionId, bidderAddr))
		if err != nil {
			return err
		}
		bids.Bids = append(bids.Bids, &bid)
		err = k.BidByAddress.Set(ctx, collections.Join(auctionId, bidderAddr), bids)
		if err != nil {
			return err
		}
	} else {
		bids := types.Bids{
			Bids: []*types.Bid{&bid},
		}
		err = k.BidByAddress.Set(ctx, collections.Join(auctionId, bidderAddr), bids)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return k.lockedToken(ctx, sdk.NewCoins(bid.Amount), bid.Bidder)
}

// CancelBidEntry cancel existing bid entry for the given auction id
func (k Keeper) CancelBidEntry(ctx context.Context, auctionId, bidId uint64) error {
	has, err := k.Auctions.Has(ctx, auctionId)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("cannot bid for non-existing/expired auction with id: %v", auctionId)
	}

	bidQueue, err := k.Bids.Get(ctx, auctionId)
	if err != nil {
		return err
	}

	var refundAddr string
	var refundAmt sdk.Coin
	for i, bid := range bidQueue.Bids {
		if bid.BidId == bidId {
			bid.IsHandle = true
			bidQueue.Bids[i] = bid
			refundAddr = bid.Bidder
			refundAmt = bid.Amount
			break
		}
	}

	err = k.Bids.Set(ctx, auctionId, bidQueue)
	if err != nil {
		return err
	}

	if refundAddr == "" || refundAmt.IsNil() {
		return fmt.Errorf("cannot find bid entry with id %v for auction %v", bidId, auctionId)
	}

	return k.refundToken(ctx, sdk.NewCoins(refundAmt), refundAddr)
}

func (k Keeper) lockedToken(ctx context.Context, amt sdk.Coins, bidderAdrr string) error {
	bidderAcc, err := k.authKeeper.AddressCodec().StringToBytes(bidderAdrr)
	if err != nil {
		return err
	}

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAcc, types.ModuleName, amt)
}

func (k Keeper) refundToken(ctx context.Context, amt sdk.Coins, bidderAdrr string) error {
	bidderAcc, err := k.authKeeper.AddressCodec().StringToBytes(bidderAdrr)
	if err != nil {
		return err
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidderAcc, amt)
}
