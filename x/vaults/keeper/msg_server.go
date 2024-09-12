package keeper

import (
	"context"

	"github.com/onomyprotocol/reserve/x/vaults/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) ActiveCollateral(ctx context.Context, msg *types.MsgActiveCollateral) (*types.MsgActiveCollateralResponse, error) {
	err := k.ActiveCollateralAsset(ctx, msg.Denom, msg.MinCollateralRatio, msg.LiquidationRatio, msg.MaxDebt)
	if err != nil {
		return nil, err
	}

	return &types.MsgActiveCollateralResponse{}, nil
}

func (k msgServer) CreateVault(ctx context.Context, msg *types.MsgCreateVault) (*types.MsgCreateVaultResponse, error) {
	err := k.CreateNewVault(ctx, msg.Denom, sdk.MustAccAddressFromBech32(msg.Owner), msg.Collateral, msg.Minted)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateVaultResponse{}, nil
}

func (k msgServer) Deposit(ctx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	err := k.DepositToVault(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgDepositResponse{}, nil
}

func (k msgServer) Withdraw(ctx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	err := k.WithdrawFromVault(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawResponse{}, nil
}

func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	err := k.MintCoin(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintResponse{}, nil
}

func (k msgServer) Repay(ctx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	err := k.RepayDebt(ctx, msg.VaultId, sdk.MustAccAddressFromBech32(msg.Sender), msg.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgRepayResponse{}, nil
}
