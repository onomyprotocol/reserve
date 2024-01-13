package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMarketChannel = []byte("MarketChannel")
	// TODO: Determine the default value
	DefaultMarketChannel string = "market_channel"
)

var (
	KeyProviderChannel = []byte("ProviderChannel")
	// TODO: Determine the default value
	DefaultProviderChannel string = "provider_channel"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	marketChannel string,
	providerChannel string,
) Params {
	return Params{
		MarketChannel:   marketChannel,
		ProviderChannel: providerChannel,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMarketChannel,
		DefaultProviderChannel,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMarketChannel, &p.MarketChannel, validateMarketChannel),
		paramtypes.NewParamSetPair(KeyProviderChannel, &p.ProviderChannel, validateProviderChannel),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMarketChannel(p.MarketChannel); err != nil {
		return err
	}

	if err := validateProviderChannel(p.ProviderChannel); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMarketChannel validates the MarketChannel param
func validateMarketChannel(v interface{}) error {
	marketChannel, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = marketChannel

	return nil
}

// validateProviderChannel validates the ProviderChannel param
func validateProviderChannel(v interface{}) error {
	providerChannel, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = providerChannel

	return nil
}
