## Params
- auction module will need to set these params:

```
type Params struct {
    // defines how long (either in blocktime or blockheight)
    // between each auction
    auction_periods time.Time|uint64
    
    // defines how long the auction will takes
    auction_durations time.Duration|uint64
    
    // period between each price reduction
    reduce_step time.Time|uint64
    
    // rate compared with the collaterals price from the
    // oracle at which the auction will start with
    starting_rate float64
    
    // rate compared with the initial price that the price
    // can drop to
    lowest_rate float64
    
    // rate that are decrease every reduce_step
    discount_rate float64
    
}
```
