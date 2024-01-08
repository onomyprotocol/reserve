package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMarketCollateral = []byte("MarketCollateral")
	// TODO: Determine the default value
	DefaultMarketCollateral string = "market_collateral"
)

var (
	KeyReserveCollateral = []byte("ReserveCollateral")
	// TODO: Determine the default value
	DefaultReserveCollateral string = "reserve_collateral"
)

var (
	KeyCollateralDeposit = []byte("CollateralDeposit")
	// TODO: Determine the default value
	DefaultCollateralDeposit string = "collateral_deposit"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	marketCollateral string,
	reserveCollateral string,
	collateralDeposit string,
) Params {
	return Params{
		MarketCollateral:  marketCollateral,
		ReserveCollateral: reserveCollateral,
		CollateralDeposit: collateralDeposit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMarketCollateral,
		DefaultReserveCollateral,
		DefaultCollateralDeposit,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMarketCollateral, &p.MarketCollateral, validateMarketCollateral),
		paramtypes.NewParamSetPair(KeyReserveCollateral, &p.ReserveCollateral, validateReserveCollateral),
		paramtypes.NewParamSetPair(KeyCollateralDeposit, &p.CollateralDeposit, validateCollateralDeposit),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMarketCollateral(p.MarketCollateral); err != nil {
		return err
	}

	if err := validateReserveCollateral(p.ReserveCollateral); err != nil {
		return err
	}

	if err := validateCollateralDeposit(p.CollateralDeposit); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMarketCollateral validates the MarketCollateral param
func validateMarketCollateral(v interface{}) error {
	marketCollateral, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = marketCollateral

	return nil
}

// validateReserveCollateral validates the ReserveCollateral param
func validateReserveCollateral(v interface{}) error {
	reserveCollateral, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = reserveCollateral

	return nil
}

// validateCollateralDeposit validates the CollateralDeposit param
func validateCollateralDeposit(v interface{}) error {
	collateralDeposit, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = collateralDeposit

	return nil
}
