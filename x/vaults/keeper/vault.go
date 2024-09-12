package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k *Keeper) CreateVault(
	ctx context.Context,
	denom string,
	owner sdk.AccAddress,
	collateral sdk.Coin,
	mint sdk.Coin,
) error {
	vm, err := k.GetVaultManager(ctx, denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", denom)
	}

	params := k.GetParams(ctx)
	vmParams := vm.Params

	// Check if expect min less than MinInitialDebt
	if mint.Amount.LT(params.MinInitialDebt) {
		return fmt.Errorf("initial mint should be greater than min. Got %d, expected %d", mint, params.MinInitialDebt)
	}

	// Calculate collateral ratio
	price := k.GetPrice(ctx, denom)
	// TODO: recalculate with denom decimal?
	collateralValue := math.LegacyNewDecFromInt(collateral.Amount).Mul(price)
	ratio := collateralValue.QuoInt(mint.Amount)

	if ratio.LT(vmParams.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vmParams.MinCollateralRatio)
	}

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(params.MintingFee).TruncateInt()
	feeCoins := sdk.NewCoins(sdk.NewCoin(mint.Denom, feeAmount))
	mintedCoins := feeCoins.Add(mint)

	// Lock collateral asset
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Mint and transfer to user and reserve
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, feeCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Set vault
	vault := types.Vault{
		Owner:            owner.String(),
		Debt:             mintedCoins[0],
		CollateralLocked: collateral,
		Status:           0,
	}
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoins[0].Amount)
	return k.VaultsManager.Set(ctx, denom, vm)
}

func (k *Keeper) MintCoin(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	mint sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	params := k.GetParams(ctx)

	lockedCoin := vault.CollateralLocked
	price := k.GetPrice(ctx, lockedCoin.Denom)
	lockedValue := math.LegacyNewDecFromInt(lockedCoin.Amount).Mul(price)

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(params.MintingFee).TruncateInt()
	feeCoins := sdk.NewCoins(sdk.NewCoin(mint.Denom, feeAmount))
	mintedAmount := feeAmount.Add(mint.Amount)
	mintedCoins := feeCoins.Add(mint)

	// calculate ratio
	ratio := lockedValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount.Add(mintedAmount)))
	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vm.Params.MinCollateralRatio)
	}

	// Mint and transfer to user and reserve
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, feeCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Owner), sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Add(sdk.NewCoin(vault.Debt.Denom, mintedAmount))
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoins[0].Amount)
	return k.VaultsManager.Set(ctx, vault.CollateralLocked.Denom, vm)

}

func (k *Keeper) Repay(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	mint sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	burnAmount := mint
	if vault.Debt.IsLT(burnAmount) {
		burnAmount = vault.Debt
	}

	err = k.bankKeeper.BurnCoins(ctx, sender, sdk.NewCoins(burnAmount))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Sub(burnAmount)
	k.SetVault(ctx, vault)

	vm.MintAvailable = vm.MintAvailable.Add(burnAmount.Amount)
	return k.VaultsManager.Set(ctx, vm.Denom, vm)
}

func (k *Keeper) Deposit(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}

	// Lock collateral asset
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Add(collateral)
	return k.SetVault(ctx, vault)
}

func (k *Keeper) Withdraw(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}

	if vault.CollateralLocked.Amount.LT(collateral.Amount) {
		fmt.Errorf("%d exeed locked amount: %d", collateral.Amount, vault.CollateralLocked.Amount)
	}

	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	newLock := vault.CollateralLocked.Sub(collateral)
	price := k.GetPrice(ctx, collateral.Denom)
	newLockValue := math.LegacyNewDecFromInt(newLock.Amount).Mul(price)
	ratio := newLockValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))

	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("ratio less than min ratio. Got: %d, min: %d", ratio, vm.Params.MinCollateralRatio)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Sub(collateral)
	return k.SetVault(ctx, vault)
}

func (k *Keeper) UpdateVaultsDebt(
	ctx context.Context,
) error {
	params := k.GetParams(ctx)
	fee := params.StabilityFee

	k.Vaults.Walk(ctx, nil, func(key uint64, vault types.Vault) (bool, error) {
		if vault.Status == 0 {
			debtAmount := vault.Debt.Amount
			newDebtAmount := math.LegacyNewDecFromInt(debtAmount).Add(math.LegacyNewDecFromInt(debtAmount).Mul(fee)).TruncateInt()
			vault.Debt.Amount = newDebtAmount
		}

		return false, nil
	})
	return nil
}

func (k *Keeper) GetVault(
	ctx context.Context,
	id uint64,
) (types.Vault, error) {
	vault, err := k.Vaults.Get(ctx, id)
	if err != nil {
		return types.Vault{}, err
	}
	return vault, nil
}

func (k *Keeper) SetVault(
	ctx context.Context,
	vault types.Vault,
) error {
	id, err := k.VaultsSequence.Next(ctx)
	if err != nil {
		return err
	}

	return k.Vaults.Set(ctx, id, vault)
}
