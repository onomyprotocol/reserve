package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	// KeyMCR is byte key for Minimum Collateral Ratio param.
	KeyMCR = []byte("MCR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyLR is byte key for Liquidiation Ratio param.
	KeyLR = []byte("LR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeyIR is byte key for Interest Rate param.
	KeyIR = []byte("IR") //nolint:gochecknoglobals // cosmos-sdk style
	// KeySR is byte key for Savings Rate param.
	KeySR = []byte("SR") //nolint:gochecknoglobals // cosmos-sdk style
)

var (
	// DefaultEarnRate is default value for the DefaultEarnRate param.
	DefaultMCR = "25000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnRate is default value for the DefaultBurnRate param.
	DefaultLR = "15000" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultBurnCoin is default value for the DefaultBurnCoin param.
	DefaultIR = "0600" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
	// DefaultMarketFee is default value for the MarketFee param.
	DefaultSR = "0060" //nolint:gomnd,gochecknoglobals // cosmos-sdk style
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	MCR string,
	LR string,
	IR string,
	SR string,
) Params {
	return Params{
		MCR: MCR,
		LR:  LR,
		IR:  IR,
		SR:  SR,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMCR, DefaultLR, DefaultIR, DefaultSR)
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
