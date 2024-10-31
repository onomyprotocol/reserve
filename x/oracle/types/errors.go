package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/oracle module sentinel errors
var (
	ErrInvalidSigner            = sdkerrors.Register(ModuleName, 1, "expected gov account as only signer for proposal message")
	ErrInvalidPacketTimeout     = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion           = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrInvalidBandRequest       = sdkerrors.Register(ModuleName, 4, "Invalid Band IBC Request")
	ErrBandPortBind             = sdkerrors.Register(ModuleName, 5, "could not claim port capability")
	ErrBadRequestInterval       = sdkerrors.Register(ModuleName, 6, "invalid Band IBC request interval")
	ErrInvalidSourceChannel     = sdkerrors.Register(ModuleName, 7, "invalid IBC source channel")
	ErrBadSymbolsCount          = sdkerrors.Register(ModuleName, 8, "invalid symbols count")
	ErrTooLargeCalldata         = sdkerrors.Register(ModuleName, 9, "too large calldata")
	ErrInvalidMinCount          = sdkerrors.Register(ModuleName, 10, "invalid min count")
	ErrInvalidAskCount          = sdkerrors.Register(ModuleName, 11, "invalid ask count")
	ErrInvalidOwasmGas          = sdkerrors.Register(ModuleName, 12, "invalid owasm gas")
	ErrInvalidBandUpdateRequest = sdkerrors.Register(ModuleName, 13, "Invalid Band Update Request Proposal")
	ErrBandRequestNotFound      = sdkerrors.Register(ModuleName, 14, "Band Oracle Request not found")
	ErrInvalidBandDeleteRequest = sdkerrors.Register(ModuleName, 15, "Invalid Band Delete Request Proposal")
	ErrResolveStatusNotSuccess  = sdkerrors.Register(ModuleName, 16, "Band Oracle request is not resolved successfully")
	ErrInvalidOracle            = sdkerrors.Register(ModuleName, 17, "invalid oracle")
)
