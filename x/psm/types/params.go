package types

import (
	"cosmossdk.io/math"
	"fmt"
)

var (
	DefaultLimitTotal           = math.NewInt(100_000_000)
	DefaultAcceptablePriceRatio = math.LegacyMustNewDecFromStr("0.001")
)

// NewParams creates a new Params instance.
func NewParams(
	limitTotal math.Int,
	AcceptablePriceRatio math.LegacyDec,
) Params {
	return Params{
		LimitTotal:           limitTotal,
		AcceptablePriceRatio: AcceptablePriceRatio,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultLimitTotal, DefaultAcceptablePriceRatio,
	)
}

// Validate validates the set of params.
func (m Params) Validate() error {
	if err := validateLimitTotal(m.LimitTotal); err != nil {
		return err
	}
	if m.AcceptablePriceRatio.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("AcceptablePriceRatio must be positive")
	}
	return nil
}

func validateLimitTotal(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("total limit rate cannot be negative or nil: %s", v)
	}

	return nil
}
