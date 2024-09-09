package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/oracle module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1, "expected gov account as only signer for proposal message")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrInvalidBandRequest   = sdkerrors.Register(ModuleName, 4, "Invalid Band IBC Request")
)
