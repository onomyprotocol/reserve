package types

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
)

var (
	DefaultMintingFee            = math.LegacyMustNewDecFromStr("0.05")
	DefaultStabilityFee          = math.LegacyMustNewDecFromStr("0.05")
	DefaultLiquidationPenalty    = math.LegacyMustNewDecFromStr("0.05")
	DefaultMinInitialDebt        = math.NewInt(50_000_000)
	DefaultRecalculateDebtPeriod = time.Hour
	DefaultLiquidatePeriod       = time.Hour
	DefaultMintDenom             = "nomUSD"

	KeyMintingFee            = []byte("MintingFee")
	KeyStabilityFee          = []byte("StabilityFee")
	KeyLiquidationPenalty    = []byte("LiquidationPenalty")
	KeyMinInitialDebt        = []byte("MinInitialDebt")
	KeyRecalculateDebtPeriod = []byte("RecalculateDebtPeriod")
	KeyLiquidatePeriod       = []byte("LiquidatePeriod")
)

// NewParams creates a new Params instance.
func NewParams(
	minInitialDebt math.Int,
	mintDenom string,
	chargingPeriod time.Duration,
	liquidatePeriod time.Duration,
) Params {
	return Params{
		MinInitialDebt:  minInitialDebt,
		LiquidatePeriod: liquidatePeriod,
		ChargingPeriod:  chargingPeriod,
		MintDenom:       mintDenom,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMinInitialDebt,
		DefaultMintDenom,
		DefaultRecalculateDebtPeriod,
		DefaultLiquidatePeriod,
	)
}

// Validate validates the set of params.
func (m Params) Validate() error {
	if err := validateMinInitialDebt(m.MinInitialDebt); err != nil {
		return err
	}
	if err := validateRecalculateDebtPeriod(m.ChargingPeriod); err != nil {
		return err
	}
	if err := validateLiquidatePeriod(m.LiquidatePeriod); err != nil {
		return err
	}
	return nil
}

func validateRecalculateDebtPeriod(i interface{}) error {
	return nil
}
func validateLiquidatePeriod(i interface{}) error {
	return nil
}

func validateStabilityFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("total limit rate cannot be negative or nil: %s", v)
	}

	return nil
}

func validateLiquidationPenalty(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("total limit rate cannot be negative or nil: %s", v)
	}

	return nil
}

func validateMinInitialDebt(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("total limit rate cannot be negative or nil: %s", v)
	}

	return nil
}

func validateMintingFee(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("total limit rate cannot be negative or nil: %s", v)
	}

	return nil
}
