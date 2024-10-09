package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// Each value below is the default value for each parameter when generating the default
	// genesis file.
	DefaultBandRequestInterval = int64(1) // every 1 block
	DefaultBandSourceChannel   = "channel-0"
	DefaultBandVersion         = "bandchain-1"
	DefaultBandPortID          = "oracle"
	DefaultOracleIds		   = []int64{42}

	// DefaultBandOracleRequestParams
	// TODO: Check these params
	DefaultAskCount       = uint64(16)
	DefaultMinCount       = uint64(10)
	DefaultFeeLimit       = sdk.Coins{sdk.NewCoin("uband", sdkmath.NewInt(100))}
	DefaultPrepareGas     = uint64(20000)
	DefaultExecuteGas     = uint64(100000)
	DefaultMinSourceCount = uint64(3)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// DefaultBandParams returns the default BandParams
func DefaultBandParams() BandParams {
	return BandParams{
		IbcRequestInterval: DefaultBandRequestInterval,
		IbcSourceChannel:   DefaultBandSourceChannel,
		IbcVersion:         DefaultBandVersion,
		IbcPortId:          DefaultBandPortID,
		LegacyOracleIds: 	DefaultOracleIds,
	}
}

// DefaultBandOracelRequestParams return the default BandOracelRequestParams
func DefaultBandOracelRequestParams() BandOracleRequestParams {
	return BandOracleRequestParams{
		AskCount:       DefaultAskCount,
		MinCount:       DefaultMinCount,
		FeeLimit:       DefaultFeeLimit,
		PrepareGas:     DefaultPrepareGas,
		ExecuteGas:     DefaultExecuteGas,
		MinSourceCount: DefaultMinSourceCount,
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func DefaultTestBandIbcParams() *BandParams {
	return &BandParams{
		// block request interval to send Band IBC prices
		IbcRequestInterval: 10,
		// band IBC source channel
		IbcSourceChannel: "channel-0",
		// band IBC version
		IbcVersion: "bandchain-1",
		// band IBC portID
		IbcPortId: "oracle",
	}
}
