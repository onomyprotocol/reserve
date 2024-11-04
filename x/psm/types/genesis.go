package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1
const DefaultMintDenom = "nomUSD"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:      DefaultParams(),
		Stablecoins: []Stablecoin{},
		Noms:        []string{DefaultMintDenom},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
