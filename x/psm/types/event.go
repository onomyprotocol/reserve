package types

const (
	EventAddStablecoin    = "add_stablecoin"
	EventSwapTofxUSD      = "swap_stablecoin_to_fxUSD"
	EventSwapToStablecoin = "swap_fxUSD_to_stablecoin"

	AttributeAmount         = "amount"
	AttributeStablecoinName = "stablecoin_name"
	AttributeReceive        = "receive"
	AttributeFeeIn          = "fee_in"
	AttributeFeeOut         = "fee_out"
)
