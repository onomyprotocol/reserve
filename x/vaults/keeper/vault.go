package keeper

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	errors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/onomyprotocol/reserve/x/vaults/types"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
)

func (k *Keeper) CreateNewVault(
	ctx context.Context,
	owner sdk.AccAddress,
	collateral sdk.Coin,
	mint sdk.Coin,
) error {
	denom := collateral.Denom
	vm, err := k.GetVaultManager(ctx, denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", denom)
	}

	if mint.Denom != types.DefaultMintDenom {
		return fmt.Errorf("minted denom must be %s", types.DefaultMintDenom)
	}

	params := k.GetParams(ctx)
	vmParams := vm.Params

	// Check if expect min less than MinInitialDebt
	if mint.Amount.LT(params.MinInitialDebt) {
		return fmt.Errorf("initial mint should be greater than min. Got %v, expected %v", mint.Amount, params.MinInitialDebt)
	}

	// Calculate collateral ratio
	price := k.OracleKeeper.GetPrice(ctx, denom, types.DefaultMintDenom)
	if price == nil || price.IsNil() {
		return errors.Wrapf(oracletypes.ErrInvalidOracle, "CreateNewVault: can not get price with base %s quote %s", denom, types.DefaultMintDenom)
	}
	// TODO: recalculate with denom decimal?
	collateralValue := math.LegacyNewDecFromInt(collateral.Amount).Mul(*price)
	ratio := collateralValue.QuoInt(mint.Amount)

	if ratio.LT(vmParams.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vmParams.MinCollateralRatio)
	}

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(vmParams.MintingFee).TruncateInt()
	feeCoin := sdk.NewCoin(mint.Denom, feeAmount)
	mintedCoin := feeCoin.Add(mint)

	if vm.MintAvailable.LT(mintedCoin.Amount) {
		return fmt.Errorf("exeed max debt")
	}

	vaultId, vaultAddress := k.GetVaultIdAndAddress(ctx)

	// Lock collateral asset
	err = k.BankKeeper.SendCoins(ctx, owner, vaultAddress, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Mint and transfer to user and reserve
	err = k.BankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintedCoin))
	if err != nil {
		return err
	}

	err = k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(feeCoin))
	if err != nil {
		return err
	}

	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Set vault
	vault := types.Vault{
		Id:               vaultId,
		Owner:            owner.String(),
		Debt:             mintedCoin,
		CollateralLocked: collateral,
		Status:           types.ACTIVE,
		Address:          vaultAddress.String(),
	}
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}
	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoin.Amount)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtCreateVault,
			sdk.NewAttribute(types.AttributeKeyVaultId, fmt.Sprintf("%d", vaultId)),
			sdk.NewAttribute(types.AttributeKeyOwner, vault.Owner),
			sdk.NewAttribute(types.AttributeKeyMintAmount, mint.String()),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateral.String()),
			sdk.NewAttribute(types.AttributeKeyDebt, vault.Debt.String()),
			sdk.NewAttribute(types.AttributeKeyVaultAddress, vault.Address),
		),
	)

	return k.VaultsManager.Set(ctx, denom, vm)
}

func (k *Keeper) CloseVault(
	ctx context.Context,
	sender string,
	vault types.Vault,
) error {
	if sender != vault.Owner {
		return fmt.Errorf("sender is not the vault owner")
	}
	
	// User have to pay all the debt to close the vault
	if vault.Debt.Amount.GT(math.ZeroInt()) {
		err := k.BankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Owner), sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(vault.Debt))
		if err != nil {
			return err
		}
	}

	// transfer all collateral locked to owner
	lockedCoins := sdk.NewCoins(vault.CollateralLocked)
	err := k.BankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), sdk.MustAccAddressFromBech32(vault.Owner), lockedCoins)
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked.Amount = math.ZeroInt()
	vault.Status = types.CLOSED
	return k.SetVault(ctx, vault)
}

func (k *Keeper) MintCoin(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	mint sdk.Coin,
) error {
	params := k.GetParams(ctx)
	if mint.Denom != params.MintDenom {
		return fmt.Errorf("minted denom must be %s", types.DefaultMintDenom)
	}
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	if sender.String() != vault.Owner {
		return fmt.Errorf("sender is not the vault owner")
	}
	if vault.Status != types.ACTIVE {
		return fmt.Errorf("vault is not actived")
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	lockedCoin := vault.CollateralLocked
	price := k.OracleKeeper.GetPrice(ctx, lockedCoin.Denom, types.DefaultMintDenom)
	if price == nil || price.IsNil() {
		return errors.Wrapf(oracletypes.ErrInvalidOracle, "MintCoin: can not get price with base %s quote %s", lockedCoin.Denom, types.DefaultMintDenom)
	}
	lockedValue := math.LegacyNewDecFromInt(lockedCoin.Amount).Mul(*price)

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(vm.Params.MintingFee).TruncateInt()
	feeCoin := sdk.NewCoin(mint.Denom, feeAmount)
	mintedCoin := feeCoin.Add(mint)

	// calculate ratio
	ratio := lockedValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount.Add(mintedCoin.Amount)))
	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vm.Params.MinCollateralRatio)
	}

	if vm.MintAvailable.LT(mintedCoin.Amount) {
		return fmt.Errorf("exeed max debt")
	}

	// Mint and transfer to user and reserve
	err = k.BankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintedCoin))
	if err != nil {
		return err
	}

	err = k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(feeCoin))
	if err != nil {
		return err
	}

	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Owner), sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Add(sdk.NewCoin(vault.Debt.Denom, mintedCoin.Amount))
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoin.Amount)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtMint,
			sdk.NewAttribute(types.AttributeKeyVaultId, fmt.Sprintf("%d", vaultId)),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyMintAmount, mint.String()),
		),
	)

	return k.VaultsManager.Set(ctx, vault.CollateralLocked.Denom, vm)
}

func (k *Keeper) RepayDebt(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	repay sdk.Coin,
) error {
	params := k.GetParams(ctx)
	if repay.Denom != params.MintDenom {
		return fmt.Errorf("minted denom must be %s", types.DefaultMintDenom)
	}

	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	if sender.String() != vault.Owner {
		return fmt.Errorf("sender is not the vault owner")
	}
	if vault.Status != types.ACTIVE {
		return fmt.Errorf("vault is not actived")
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	burnAmount := repay
	if vault.Debt.IsLT(burnAmount) {
		burnAmount = vault.Debt
	}

	err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return err
	}

	err = k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Sub(burnAmount)
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	vm.MintAvailable = vm.MintAvailable.Add(burnAmount.Amount)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtRepay,
			sdk.NewAttribute(types.AttributeKeyVaultId, fmt.Sprintf("%d", vaultId)),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyRepayAmount, repay.String()),
		),
	)

	return k.VaultsManager.Set(ctx, vm.Denom, vm)
}

func (k *Keeper) DepositToVault(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}

	if sender.String() != vault.Owner {
		return fmt.Errorf("sender is not the vault owner")
	}

	if collateral.Denom != vault.CollateralLocked.Denom {
		return fmt.Errorf("vaultId %d does not accept denom %s", vaultId, collateral.Denom)
	}

	if vault.Status != types.ACTIVE {
		return fmt.Errorf("vault is not actived")
	}

	// Lock collateral asset
	err = k.BankKeeper.SendCoins(ctx, sender, sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Add(collateral)

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtDeposit,
			sdk.NewAttribute(types.AttributeKeyVaultId, fmt.Sprintf("%d", vaultId)),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateral.String()),
		),
	)

	return k.SetVault(ctx, vault)
}

func (k *Keeper) WithdrawFromVault(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	if vault.Status != types.ACTIVE {
		return fmt.Errorf("vault is not actived")
	}

	if sender.String() != vault.Owner {
		return fmt.Errorf("sender is not the vault owner")
	}

	if collateral.Denom != vault.CollateralLocked.Denom {
		return fmt.Errorf("vaultId %d does not accept denom %s", vaultId, collateral.Denom)
	}

	if vault.CollateralLocked.Amount.LT(collateral.Amount) {
		return fmt.Errorf("%d exeed locked amount: %d", collateral.Amount, vault.CollateralLocked.Amount)
	}

	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	newLock := vault.CollateralLocked.Sub(collateral)
	price := k.OracleKeeper.GetPrice(ctx, collateral.Denom, types.DefaultMintDenom)
	// defensive programming: should never happen since when withdraw should always have a valid oracle price
	if price == nil || price.IsNil() {
		return errors.Wrapf(oracletypes.ErrInvalidOracle, "WithdrawFromVault: can not get price with base %s quote %s", collateral.Denom, types.DefaultMintDenom)
	}

	newLockValue := math.LegacyNewDecFromInt(newLock.Amount).Mul(*price)
	ratio := newLockValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))

	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("ratio less than min ratio. Got: %d, min: %d", ratio, vm.Params.MinCollateralRatio)
	}

	err = k.BankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), sender, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Sub(collateral)
	if vault.CollateralLocked.Amount.Equal(math.ZeroInt()) {
		vault.Status = types.CLOSED
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtWithdraw,
			sdk.NewAttribute(types.AttributeKeyVaultId, fmt.Sprintf("%d", vaultId)),
			sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
			sdk.NewAttribute(types.AttributeKeyCollateral, collateral.String()),
		),
	)

	return k.SetVault(ctx, vault)
}

func (k *Keeper) UpdateVaultsDebt(
	ctx context.Context,
	lastUpdateTime time.Time,
	currentTime time.Time,
) error {
	deltaDur := currentTime.Sub(lastUpdateTime)
	rate := math.LegacyNewDec(deltaDur.Milliseconds()).Quo(math.LegacyNewDec((time.Hour * 24 * 365).Milliseconds())) // divice 365 days
	// Get stability fee of all denoms
	fees := make(map[string]math.LegacyDec, 0)
	err := k.VaultsManager.Walk(ctx, nil, func(denom string, vm types.VaultMamager) (bool, error) {
		fees[denom] = vm.Params.StabilityFee.Mul(rate)
		return false, nil
	})
	if err != nil {
		return err
	}

	err = k.Vaults.Walk(ctx, nil, func(id uint64, vault types.Vault) (bool, error) {
		var err error
		if vault.Status == types.ACTIVE {
			debtAmount := vault.Debt.Amount
			newDebtAmount := math.LegacyNewDecFromInt(debtAmount).Add(math.LegacyNewDecFromInt(debtAmount).Mul(fees[vault.CollateralLocked.Denom])).TruncateInt()
			vault.Debt.Amount = newDebtAmount
			err = k.Vaults.Set(ctx, id, vault)
		}

		return false, err
	})
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return k.LastUpdateTime.Set(ctx, types.LastUpdate{Time: sdkCtx.BlockTime()})
}

func (k *Keeper) shouldLiquidate(
	vault types.Vault,
	price math.LegacyDec,
	liquidationRatio math.LegacyDec,
) (bool, error) {
	// Only liquidate ACTIVE vault
	if vault.Status != types.ACTIVE {
		return false, nil
	}

	collateralValue := math.LegacyNewDecFromInt(vault.CollateralLocked.Amount).Mul(price)
	if math.LegacyNewDecFromInt(vault.Debt.Amount).Equal(math.LegacyZeroDec()) {
		return false, nil
	}
	ratio := collateralValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))

	if ratio.LTE(liquidationRatio) {
		return true, nil
	}
	return false, nil
}

func (k *Keeper) GetLiquidations(
	ctx context.Context,
) ([]*types.Liquidation, error) {

	// denom to liquidationRatios
	liquidationRatios := make(map[string]math.LegacyDec)
	// denom to price
	prices := make(map[string]math.LegacyDec)
	// denom to Liquidation
	liquidations := make(map[string]*types.Liquidation)

	err := k.VaultsManager.Walk(ctx, nil, func(key string, vm types.VaultMamager) (bool, error) {
		price := k.OracleKeeper.GetPrice(ctx, vm.Denom, types.DefaultMintDenom)
		if price == nil || price.IsNil() {
			return true, errors.Wrapf(oracletypes.ErrInvalidOracle, "GetLiquidations: can not get price with base %s quote %s", vm.Denom, types.DefaultMintDenom)
		}
		prices[vm.Denom] = *price
		liquidationRatios[vm.Denom] = vm.Params.LiquidationRatio
		liquidations[vm.Denom] = types.NewEmptyLiquidation(vm.Denom)

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	err = k.Vaults.Walk(ctx, nil, func(id uint64, vault types.Vault) (bool, error) {
		denom := vault.CollateralLocked.Denom
		shouldLiquidate, err := k.shouldLiquidate(vault, prices[denom], liquidationRatios[denom])
		if shouldLiquidate && err == nil {
			liquidations[denom].LiquidatingVaults = append(liquidations[denom].LiquidatingVaults, &vault)
			liquidations[denom].VaultLiquidationStatus[id] = &types.VaultLiquidationStatus{}

			vault.Status = types.LIQUIDATING
			vault.LiquidationPrice = prices[denom]
			err := k.SetVault(ctx, vault)
			if err != nil {
				return true, err
			}
		}

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	var result []*types.Liquidation
	for _, liquidation := range liquidations {
		if len(liquidation.LiquidatingVaults) != 0 {
			result = append(result, liquidation)
		}
	}

	return result, nil
}

// TODO: Separate this func
func (k *Keeper) Liquidate(
	ctx context.Context,
	liquidation types.Liquidation,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(ctx)

	vm, err := k.GetVaultManager(ctx, liquidation.Denom)
	if err != nil {
		return err
	}

	vaultIds := ""
	totalDebt := sdk.NewCoin(params.MintDenom, math.ZeroInt())
	sold := sdk.NewCoin(params.MintDenom, math.ZeroInt())
	totalCollateralRemain := sdk.NewCoin(liquidation.Denom, math.ZeroInt())

	for _, vault := range liquidation.LiquidatingVaults {
		vaultIds = vaultIds + fmt.Sprintf("%d, ", vault.Id)
		totalDebt = totalDebt.Add(vault.Debt)
		// transfer all remain collateral locked in vault to vaults module for distributing.
		vaultAddr := sdk.MustAccAddressFromBech32(vault.Address)
		balances := k.BankKeeper.GetAllBalances(ctx, vaultAddr)
		err := k.BankKeeper.SendCoinsFromAccountToModule(ctx, vaultAddr, types.ModuleName, balances)
		if err != nil {
			return err
		}
		vault.Status = types.LIQUIDATED
	}

	for _, status := range liquidation.VaultLiquidationStatus {
		sold = sold.Add(status.Sold)
		totalCollateralRemain = totalCollateralRemain.Add(status.RemainCollateral)
	}

	// Sold amount enough to cover debt
	if sold.Amount.GTE(totalDebt.Amount) {
		// Burn debt
		err := k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(totalDebt))
		if err != nil {
			return err
		}
		// Increase mint available
		vm.MintAvailable = vm.MintAvailable.Add(totalDebt.Amount)
		err = k.VaultsManager.Set(ctx, liquidation.Denom, vm)
		if err != nil {
			return err
		}

		// If remain sold, send to reserve
		remain := sold.Sub(totalDebt)
		if remain.Amount.GT(math.ZeroInt()) {
			err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(remain))
			if err != nil {
				return err
			}
		}

		// Take the liquidation penalty and send back to vault owner
		for _, vault := range liquidation.LiquidatingVaults {
			collateralRemain := liquidation.VaultLiquidationStatus[vault.Id].RemainCollateral

			if collateralRemain.Amount.Equal(math.ZeroInt()) {
				vault.CollateralLocked.Amount = math.ZeroInt()
				vault.Debt.Amount = math.ZeroInt()
				vault.Status = types.CLOSED
				continue
			}

			penaltyAmount := math.LegacyNewDecFromInt(vault.Debt.Amount).Quo(vault.LiquidationPrice).Mul(vm.Params.LiquidationPenalty).TruncateInt()

			vault.Debt.Amount = math.ZeroInt()
			if penaltyAmount.GTE(collateralRemain.Amount) {
				err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(collateralRemain))
				if err != nil {
					return err
				}
				vault.CollateralLocked.Amount = math.ZeroInt()
				vault.Status = types.CLOSED
			} else {
				err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(sdk.NewCoin(collateralRemain.Denom, penaltyAmount)))
				if err != nil {
					return err
				}
				vault.CollateralLocked.Amount = collateralRemain.Amount.Sub(penaltyAmount)
				err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(vault.CollateralLocked))
				if err != nil {
					return err
				}
			}
		}
	} else {
		// does not raise enough to cover nomUSD debt

		// Burn sold amount
		err := k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sold))
		if err != nil {
			return err
		}
		// Increase mint available
		vm.MintAvailable = vm.MintAvailable.Add(sold.Amount)
		err = k.VaultsManager.Set(ctx, liquidation.Denom, vm)
		if err != nil {
			return err
		}

		// No collateral remain
		if totalCollateralRemain.Amount.Equal(math.ZeroInt()) {
			// Update vaults status
			for _, vault := range liquidation.LiquidatingVaults {
				soldAmount := liquidation.VaultLiquidationStatus[vault.Id].Sold.Amount
				if soldAmount.GTE(vault.Debt.Amount) {
					vault.Debt.Amount = math.ZeroInt()
				} else {
					vault.Debt.Amount = vault.Debt.Amount.Sub(soldAmount)
				}
				vault.CollateralLocked.Amount = math.ZeroInt()
				// LIQUIDATED
				err = k.SetVault(ctx, *vault)
				if err != nil {
					return err
				}
			}
			currentShortfall, err := k.ShortfallAmount.Get(ctx)
			if err != nil {
				return err
			}
			shortfallAmount := totalDebt.Sub(sold).Amount
			newShortfall := currentShortfall.Add(shortfallAmount)

			sdkCtx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.TypeEvtLiquidate,
					sdk.NewAttribute(types.AttributeKeyLiquidateVaults, vaultIds),
					sdk.NewAttribute(types.AttributeKeyBurnAmount, sold.String()),
					sdk.NewAttribute(types.AttributeKeyShortfallAmount, shortfallAmount.String()),
				),
			)

			return k.ShortfallAmount.Set(ctx, newShortfall)
		} else {
			// If there some collateral asset remain, try to reconstitue vault
			// Priority by collateral ratio at momment
			// So that mean we need less resource for high ratio vault

			ratios := make([]math.LegacyDec, 0)
			//TODO: Sort by CR in GetLiquidations could reduce calculate here
			for _, vault := range liquidation.LiquidatingVaults {
				collateralRemain := liquidation.VaultLiquidationStatus[vault.Id].RemainCollateral.Amount
				penaltyAmount := math.LegacyNewDecFromInt(vault.Debt.Amount).Quo(vault.LiquidationPrice).Mul(vm.Params.LiquidationPenalty).TruncateInt()

				// If remain collateral not enough for penalty,
				// transfer all and mark vault CLOSED
				if penaltyAmount.GT(collateralRemain) {
					soldAmount := liquidation.VaultLiquidationStatus[vault.Id].Sold.Amount
					if soldAmount.GTE(vault.Debt.Amount) {
						vault.Debt.Amount = math.ZeroInt()
					} else {
						vault.Debt.Amount = vault.Debt.Amount.Sub(soldAmount)
					}
					penaltyAmount = collateralRemain
					vault.CollateralLocked.Amount = collateralRemain
					vault.Status = types.CLOSED
				}
				err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(sdk.NewCoin(liquidation.Denom, penaltyAmount)))
				if err != nil {
					return err
				}
				vault.CollateralLocked.Amount = vault.CollateralLocked.Amount.Sub(penaltyAmount)
				totalCollateralRemain.Amount = totalCollateralRemain.Amount.Sub(penaltyAmount)

				ratio := math.LegacyNewDecFromInt(vault.CollateralLocked.Amount).Mul(vault.LiquidationPrice).Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))
				ratios = append(ratios, ratio)
			}

			// Sort the vaults by CR in descending order
			sort.Slice(liquidation.LiquidatingVaults, func(i, j int) bool {
				return ratios[i].GT(ratios[j])
			})

			// Try to reconstitue vaults
			// list contains both LIQUIDATED & CLOSED,
			// only reconstitue LIQUIDATED vaults
			totalRemainDebt := totalDebt.Sub(sold)
			for _, vault := range liquidation.LiquidatingVaults {
				if vault.Status != types.LIQUIDATED {
					continue
				}
				// if remain debt & collateral can cover full vault
				// open again
				if vault.Debt.IsLTE(totalRemainDebt) && vault.CollateralLocked.IsLTE(totalCollateralRemain) {
					// Lock collateral to vault address
					err := k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(vault.CollateralLocked))
					if err != nil {
						return err
					}
					totalRemainDebt = totalRemainDebt.Sub(vault.Debt)
					totalCollateralRemain = totalCollateralRemain.Sub(vault.CollateralLocked)

					vault.Status = types.ACTIVE
				} else {
					// Update debt then mark liquidated
					soldAmount := liquidation.VaultLiquidationStatus[vault.Id].Sold.Amount
					if soldAmount.GTE(vault.Debt.Amount) {
						vault.Debt.Amount = math.ZeroInt()
					} else {
						vault.Debt.Amount = vault.Debt.Amount.Sub(soldAmount)
					}
					vault.CollateralLocked.Amount = math.ZeroInt()
				}
			}

			// if remain collateral, send to reserve
			if totalCollateralRemain.Amount.GT(math.ZeroInt()) {
				err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(totalCollateralRemain))
				if err != nil {
					return err
				}
			}

			// if remain debt, send shortfall
			if totalRemainDebt.Amount.GT(math.ZeroInt()) {
				// Update vaults status
				for _, vault := range liquidation.LiquidatingVaults {
					err = k.SetVault(ctx, *vault)
					if err != nil {
						return err
					}
				}
				currentShortfall, err := k.ShortfallAmount.Get(ctx)
				if err != nil {
					return err
				}
				newShortfall := currentShortfall.Add(totalRemainDebt.Amount)

				sdkCtx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.TypeEvtLiquidate,
						sdk.NewAttribute(types.AttributeKeyLiquidateVaults, vaultIds),
						sdk.NewAttribute(types.AttributeKeyBurnAmount, sold.String()),
						sdk.NewAttribute(types.AttributeKeyShortfallAmount, totalRemainDebt.String()),
					),
				)

				return k.ShortfallAmount.Set(ctx, newShortfall)
			}

		}
	}
	// Update vaults status
	for _, vault := range liquidation.LiquidatingVaults {
		err = k.SetVault(ctx, *vault)
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeEvtLiquidate,
			sdk.NewAttribute(types.AttributeKeyLiquidateVaults, vaultIds),
			sdk.NewAttribute(types.AttributeKeyBurnAmount, sold.String()),
		),
	)

	return err
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
	return k.Vaults.Set(ctx, vault.Id, vault)
}

func (k *Keeper) GetVaultIdAndAddress(
	ctx context.Context,
) (uint64, sdk.AccAddress) {
	id, err := k.VaultsSequence.Next(ctx)
	if err != nil {
		return 0, sdk.AccAddress{}
	}
	address := address.Module(types.ModuleName, []byte(strconv.Itoa(int(id))))

	return id, address
}
