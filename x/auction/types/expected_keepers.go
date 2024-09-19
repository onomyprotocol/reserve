package types

import (
	"context"

	"cosmossdk.io/math"

	addresscodec "cosmossdk.io/core/address"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vaulttypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

type LiquidateVaults struct {
	VaultId      uint64
	TargetGoal   sdk.Coin
	Collatheral  sdk.Coin
	InitialPrice sdk.Coin
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	AddressCodec() addresscodec.Codec
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

type VaultKeeper interface {
	GetLiquidatedVaults(ctx context.Context) ([]vaulttypes.Vault, error)
	NotifyVault(ctx context.Context, tokenRaised, collatheralUnsold sdk.Coin, isReachedGoal bool) error
}

type OracleKeeper interface {
	GetPrice(ctx context.Context, denom string) math.LegacyDec
}
