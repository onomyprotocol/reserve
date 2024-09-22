package types

import (
	"cosmossdk.io/math"
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultMintingFee            = math.LegacyMustNewDecFromStr("1")
	DefaultStabilityFee          = math.LegacyMustNewDecFromStr("1")
	DefaultLiquidationPenalty    = math.LegacyMustNewDecFromStr("1")
	DefaultMinInitialDebt        = math.NewInt(1)
	DefaultRecalculateDebtPeriod = uint64(1)
	DefaultLiquidatePeriod       = uint64(1)

	KeyMintingFee            = []byte("MintingFee")
	KeyStabilityFee          = []byte("StabilityFee")
	KeyLiquidationPenalty    = []byte("LiquidationPenalty")
	KeyMinInitialDebt        = []byte("MinInitialDebt")
	KeyRecalculateDebtPeriod = []byte("RecalculateDebtPeriod")
	KeyLiquidatePeriod       = []byte("LiquidatePeriod")
)

// NewParams creates a new Params instance.
func NewParams(
	mintingFee math.LegacyDec,
	stabilityFee math.LegacyDec,
	liquidationPenalty math.LegacyDec,
	minInitialDebt math.Int,
	recalculateDebtPeriod uint64,
	liquidatePeriod uint64,
) Params {
	return Params{
		MintingFee:            mintingFee,
		StabilityFee:          stabilityFee,
		LiquidationPenalty:    liquidationPenalty,
		MinInitialDebt:        minInitialDebt,
		RecalculateDebtPeriod: recalculateDebtPeriod,
		LiquidatePeriod:       liquidatePeriod,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMintingFee,
		DefaultStabilityFee,
		DefaultLiquidationPenalty,
		DefaultMinInitialDebt,
		DefaultRecalculateDebtPeriod,
		DefaultLiquidatePeriod,
	)
}

// Validate validates the set of params.
func (m Params) Validate() error {
	if err := validateMintingFee(m.MintingFee); err != nil {
		return err
	}
	if err := validateStabilityFee(m.StabilityFee); err != nil {
		return err
	}
	if err := validateLiquidationPenalty(m.LiquidationPenalty); err != nil {
		return err
	}
	if err := validateMinInitialDebt(m.MinInitialDebt); err != nil {
		return err
	}
	if err := validateRecalculateDebtPeriod(m.RecalculateDebtPeriod); err != nil {
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

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintingFee, &p.MintingFee, validateMintingFee),
		paramtypes.NewParamSetPair(KeyStabilityFee, &p.StabilityFee, validateStabilityFee),
		paramtypes.NewParamSetPair(KeyLiquidationPenalty, &p.LiquidationPenalty, validateLiquidationPenalty),
		paramtypes.NewParamSetPair(KeyMinInitialDebt, &p.MinInitialDebt, validateMinInitialDebt),
		paramtypes.NewParamSetPair(KeyRecalculateDebtPeriod, &p.RecalculateDebtPeriod, validateRecalculateDebtPeriod),
		paramtypes.NewParamSetPair(KeyLiquidatePeriod, &p.LiquidatePeriod, validateLiquidatePeriod),
	}
}
