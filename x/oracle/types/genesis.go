package types

func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		BandParams: DefaultBandParams(),
		BandOracleRequestParams: DefaultBandOracelRequestParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// TODO: validate stuff in genesis
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}
