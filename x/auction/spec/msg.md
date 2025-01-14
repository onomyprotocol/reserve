
### CLI

A user can query and interact with the `auction` module using the CLI.

### Tx

#### Bid

The `bid` command allows users submit a bid entry to an auction.

```shell
onomyd tx auction bid [auction-id] [amount] [recive_rate] [flags]
```

Example:

```shell
onomyd tx bid 0 1000fxUSD 0.8 --from mykey
```

#### Cancel bid

The `cancel-bid` command allows users cancel their bid entry of an auction.

```shell
onomyd tx auction cancel-bid [bid-id] [auction_id] [flags]
```

Example:

```shell
onomyd tx cancel-bid 1 0 --from mykey
```
