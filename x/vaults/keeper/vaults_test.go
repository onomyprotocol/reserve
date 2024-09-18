package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (s *KeeperTestSuite) TestVaultsStore() {
	s.SetupTest()

	v := types.Vault{
		Owner:            s.TestAccs[0].String(),
		Debt:             sdk.NewCoin("atom", math.NewInt(1000)),
		CollateralLocked: sdk.NewCoin("atom", math.NewInt(1000)),
		Status:           types.LIQUIDATED,
	}
	err := s.k.SetVault(s.Ctx, v)
	s.Require().NoError(err)

	vault, err := s.k.GetVault(s.Ctx, 0)
	s.Require().NoError(err)

	s.Require().Equal(v, vault)
}

func (s *KeeperTestSuite) TestCreateNewVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		denom      string
		owner      sdk.AccAddress
		collateral sdk.Coin
		mint       sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
			},
			denom:      denom,
			owner:      s.TestAccs[0],
			collateral: coin,
			mint:       coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.CreateNewVault(s.Ctx, t.denom, t.owner, t.collateral, t.mint)
			s.Require().NoError(err)

			vm, err := s.k.GetVaultManager(s.Ctx, denom)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vm.MintAvailable)
		})
	}
}

func (s *KeeperTestSuite) TestRepayDebt() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name    string
		setup   func()
		vaultID uint64
		sender  sdk.AccAddress
		mint    sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
			},
			vaultID: 0,
			sender:  s.TestAccs[0],
			mint:    coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.RepayDebt(s.Ctx, t.vaultID, t.sender, t.mint)
			s.Require().NoError(err)

			vm, err := s.k.GetVaultManager(s.Ctx, denom)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vm.MintAvailable)
		})
	}
}

func (s *KeeperTestSuite) TestDepositToVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.DepositToVault(s.Ctx, t.vaultId, t.sender, t.collateral)
			s.Require().NoError(err)

			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vault.CollateralLocked)
		})
	}
}

func (s *KeeperTestSuite) TestWithdrawFromVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.WithdrawFromVault(s.Ctx, t.vaultId, t.sender, t.collateral)
			s.Require().NoError(err)

			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vault.CollateralLocked)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateVaultsDebt() {
	s.SetupTest()
	var (
		denom              = "atom"
		maxDebt            = math.NewInt(10000)
		feeStabilityUpdate = math.LegacyMustNewDecFromStr("0.5")
	)

	tests := []struct {
		name    string
		setup   func()
		vaultId uint64
	}{
		{
			name: "success",
			setup: func() {
				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err := s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				// update params
				uP := types.DefaultParams()
				uP.StabilityFee = feeStabilityUpdate
				s.k.SetParams(s.Ctx, uP)
			},
			vaultId: 0,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.UpdateVaultsDebt(s.Ctx)
			s.Require().NoError(err)

			// expect
			expectDebtAmount := math.LegacyNewDecFromInt(maxDebt).Add(math.LegacyNewDecFromInt(maxDebt).Mul(feeStabilityUpdate)).TruncateInt()
			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().Equal(expectDebtAmount.String(), vault.Debt.Amount.String())
		})
	}
}

func (s *KeeperTestSuite) TestGetLiquidateVaults() {
	s.SetupTest()
	var (
		denom1        = "atom"
		denom2        = "osmo"
		coin          = sdk.NewCoin(denom1, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom1, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom1, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)
				err = s.k.ActiveCollateralAsset(s.Ctx, denom2, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom1, maxDebt),
					CollateralLocked: sdk.NewCoin(denom1, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			vaults, prices, err := s.k.GetLiquidateVaults(s.Ctx)
			s.Require().NoError(err)

			// current price = 1, vaults is empty,
			s.Require().Equal(2, len(prices))
			s.Require().Equal(0, len(vaults))
		})
	}
}
