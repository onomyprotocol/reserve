package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/vaults module sentinel errors
var (
	ErrInvalidSigner                   = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidActiveCollateralProposal = sdkerrors.Register(ModuleName, 2, "invalid active collateral proposal")
)
