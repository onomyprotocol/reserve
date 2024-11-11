## Concept

- `auction module` is implementing the *Dutch Auction* method which mean the price are decrease until a user bid for that price or the price reach floor price. 

### Initialization of an auction
- At each begin block `auction module` check for auction periods, if reach a new auction period, auction module calls over `vault module` to check for any vault that are in debt and need to liquidate the assets. `vault module` returns information needed to create an auction.


- An auction is defined as
```
type Auction struct {
    // for simplicity, will use vault id that start the auction as auction id
	AuctionId uint64 
    
	// starting price (currently only support usd stable token)
	InitialPrice string 
    
	// items defines liquidate assets
	Item types.Coin
    
	// current_rate defines the rate compare with the initial price
	CurrentRate string 
    
	// last_discount_time defines the last time a discount has been apply
	LastDiscountTime time.Time  
    
    // token amount raised from the auction
	TokenRaised      types.Coin 
    
	// status defines auction current status
	Status AuctionStatus 
    
	// target_goal defines the debt the auction is trying to recover
	TargetGoal types.Coin 
    
	// vault_id defines id of auction vault
	VaultId uint64 
}
```

### Auction update per block
- For each begin block, we update the auctions information accordingly, This includes:
    - Check for auction start time to start the process and open for bidding.
    - Check if auction reach next `reduce_step` to update the price based on `discount_rate`. If the price dips to `lowest_rate`, handle any bids left, after that update auction status `toAuctionStatus_AUCTION_STATUS_FINISHED` and closed for bidding.
    - Check if there is any bid entry matches the current price. If there is, send the collatheral amount accordingly with the bid amount and the price its accepts. If auction out of collatheral then update auction status to `AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL` and closed for bidding.
    - Check if auction status is either `toAuctionStatus_AUCTION_STATUS_FINISHED` or `AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL` to end the auction.

:::info
In the case where the bid amount cover a portion of the remain collatheral, we will send those remain to the bidder and refund the rest back to the user so that auction won't end with auction has some small amount that no one bid for.
:::

### Bidding process
- Bidder can submit a transaction to participate in the auction.

eg:
```
message MsgBid {
    string auction_id = 1
    string bidder = 2
    sdk.Coin amount = 3
}
```

A bid is defined as

```
type Bid struct {
    // id of bid
  uint64 bid_id = 1;
  
  // bidder address
  string bidder = 2;

  // bidding amount
  sdk.Coin amount = 3;
  
  // recive_price defines the price that the bid is willing to pay
  string recive_price = 4 

  bool is_handle = 5;

  // index in auction bid_queue
  uint64 index = 6;
}
```

- The bid amount denom must matches that of the vault debt denom.

- The bidding process occur through out the auction process, users can submit an entry indicates the amount they are willing to pay for a certain price.

- Auction keeper use a queue to store all bid entries for an auction. Bid will be handle by the order of time it got added in the queue.

- To update the entry, user must cancel the current bid entry and resubmit

eg:
```
message MsgCancelBid {
    uint64 auction_id = 1
    string bidder = 2
    uint64 bid_id = 3
}
```

- When the price is dropped to an amount that less than or equal the bid `recive_price` the amount of collatheral the bidder recieve = bid_amount / recieve_price

### Auction end process
- By the end of the auction ( either auction status updated to `AuctionStatus_AUCTION_STATUS_FINISHED` or `AuctionStatus_AUCTION_STATUS_OUT_OF_COLLATHERAL`), send token raised from auctions to `vault` module. 
- Appends the liquidation result to a `liquidationMap`. If `liquidationMap` not empty notified `vault` module with the liquidation result through `vault keeper` interface `Liquidate` func. 
- Refunds all unhandled bids to the corresponding bidder. Afterward, clear all state of the auction.

Liquidation is defined as
```
type Liquidation struct {
	Denom                  string                             
	LiquidatingVaults      []*Vault                           
	VaultLiquidationStatus map[uint64]*VaultLiquidationStatus 
}

```