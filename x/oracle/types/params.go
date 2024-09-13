package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
	// Each value below is the default value for each parameter when generating the default
	// genesis file.
	DefaultBandRequestInterval = int64(1) // every 1 block
	DefaultBandSourceChannel   = "channel-0"
	DefaultBandVersion         = "bandchain-1"
	DefaultBandPortID          = "oracle"
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
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
