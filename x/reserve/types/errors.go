package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/reserve module sentinel errors
var (
	// ErrInsufficientBalance - the user balance is insufficient for the operation.
	ErrInsufficientBalance = sdkerrors.Register(ModuleName, 1, "insufficient balance") // nolint: gomnd
)
