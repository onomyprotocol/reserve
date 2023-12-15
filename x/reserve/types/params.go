package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyMCR is byte key for Minimum Collateralization Ratio param.
	KeyMCR = []byte("MCR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyLR is byte key for Liquidiation Ratio param.
	KeyLR = []byte("LR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyIR is byte key for Interest Rate param.
	KeyIR = []byte("IR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeySR is byte key for Savings Rate param.
	KeySR = []byte("SR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyONEX is byte key for ONEX chain channel.
	KeyONEX = []byte("ONEX") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultMinCollateralizationRatio is default value for the Minimum Collateralization Ratio.
	DefaultMinCollateralizationRatio = "25000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultLiquidationRatio is default value for the Liquidation Ratio.
	DefaultLiquidationRatio = "15000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultInterestRate is default value for the Interest Rate.
	DefaultInterestRate = "0600" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultSavingsRate is default value for the Savings Rate.
	DefaultSavingsRate = "0060" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultProviderChannel is default value for the Provider channel.
	DefaultProviderChannel = "0" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultMarketChannel is default value for the Market channel.
	DefaultMarketChannel = "0" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultReserveChannel is default value for the Market channel.
	DefaultReserveChannel = "0" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultMarketCollateral is default value for the Market channel.
	DefaultMarketCollateral = "0" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultReserveCollateral is default value for the Market channel.
	DefaultReserveCollateral = "0" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minCollateralizationRatio string,
	liquidationRatio string,
	interestRate string,
	savingsRate string,
	providerChannel string,
	marketChannel string,
	marketCollateral string,
	reserveCollateral string,
) Params {
	return Params{
		MinCollateralizationRatio: minCollateralizationRatio,
		LiquidationRatio:          liquidationRatio,
		InterestRate:              interestRate,
		SavingsRate:               savingsRate,
		ProviderChannel:           providerChannel,
		MarketChannel:             marketChannel,
		MarketCollateral:          marketCollateral,
		ReserveCollateral:         reserveCollateral,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMinCollateralizationRatio, DefaultLiquidationRatio, DefaultInterestRate, DefaultSavingsRate, DefaultProviderChannel, DefaultMarketChannel, DefaultMarketCollateral, DefaultReserveCollateral)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
