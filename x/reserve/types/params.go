package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// What to do with this?
var (
	// KeyMCR is byte key for Minimum Collateralization Ratio param.
	KeyProviderChannel = []byte("ProviderChannel") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyLR is byte key for Liquidiation Ratio param.
	KeyMarketChannel = []byte("MarketChannel") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyIR is byte key for Interest Rate param.
	KeyMarketCollateral = []byte("MarketCollateral") //nolint:gochecknoglobals // cosmos-sdk style
	// KeySR is byte key for Savings Rate param.
	KeyReserveCollateral = []byte("ReserveCollateral") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyONEX is byte key for ONEX chain channel.
	KeyCollateralDeposit = []byte("CollateralDeposit") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
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
	// DefaultCollateralDeposit is default value for the Denom Collateral Deposit.
	DefaultCollateralDeposit = "100000000000000000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	providerChannel string,
	marketChannel string,
	marketCollateral string,
	reserveCollateral string,
	collateralDeposit string,
) Params {
	return Params{
		ProviderChannel:   providerChannel,
		MarketChannel:     marketChannel,
		MarketCollateral:  marketCollateral,
		ReserveCollateral: reserveCollateral,
		CollateralDeposit: collateralDeposit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultProviderChannel,
		DefaultMarketChannel,
		DefaultMarketCollateral,
		DefaultReserveCollateral,
		DefaultCollateralDeposit,
	)
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
