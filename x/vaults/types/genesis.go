package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

func NewGenesisState(
	params Params,
	vms []VaultManager,
	vaults []Vault,
	lastUpdate *LastUpdate,
	shortfall sdk.Coins,
	vaultsSequence uint64,
) *GenesisState {
	return &GenesisState{
		Params:           params,
		VaultManagers:    vms,
		Vaults:           vaults,
		LastUpdate:       lastUpdate,
		ShortfallAmounts: shortfall,
		VaultSequence:    vaultsSequence,
	}
}
