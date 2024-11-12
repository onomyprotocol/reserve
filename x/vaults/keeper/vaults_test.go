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
		LiquidationPrice: math.LegacyMustNewDecFromStr("1.0"),
		Address:          "addr1_______________",
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
		denom      = "atom"
		mintDenom  = types.DefaultMintDenoms[0]
		collateral = sdk.NewCoin(denom, math.NewInt(10_000_000)) // 10 atom = 80$
		maxDebt    = math.NewInt(100_000_000)
	)
	err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("1.6"), math.LegacyMustNewDecFromStr("1.5"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
	s.Require().NoError(err)

	tests := []struct {
		name       string
		setup      func()
		denom      string
		owner      sdk.AccAddress
		collateral sdk.Coin
		mint       sdk.Coin
		expErr     bool
	}{
		{
			name: "mint less than min initial debt",
			setup: func() {
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(collateral))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(collateral))
				s.Require().NoError(err)
			},
			denom:      "atom",
			owner:      s.TestAccs[0],
			collateral: collateral,
			mint:       sdk.NewCoin(mintDenom, math.NewInt(10_000_000)),
			expErr:     true,
		},
		{
			name: "exeed max debt",
			setup: func() {
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(collateral))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(collateral))
				s.Require().NoError(err)
			},
			denom:      "atom",
			owner:      s.TestAccs[0],
			collateral: collateral,
			mint:       sdk.NewCoin(mintDenom, math.NewInt(110_000_000)),
			expErr:     true,
		},
		{
			name: "invalid ratio",
			setup: func() {
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(collateral))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(collateral))
				s.Require().NoError(err)
			},
			denom:      "atom",
			owner:      s.TestAccs[0],
			collateral: collateral,
			mint:       sdk.NewCoin(mintDenom, math.NewInt(60_000_000)),
			expErr:     true,
		},
		{
			name: "mint invalid denom",
			setup: func() {
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(collateral))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(collateral))
				s.Require().NoError(err)
			},
			denom:      "atom",
			owner:      s.TestAccs[0],
			collateral: collateral,
			mint:       sdk.NewCoin("atom", math.NewInt(60_000_000)),
			expErr:     true,
		},
		{
			name: "not the vault owner",
			setup: func() {
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(collateral))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[1], sdk.NewCoins(collateral))
				s.Require().NoError(err)
			},
			denom:      "atom",
			owner:      s.TestAccs[1],
			collateral: collateral,
			mint:       sdk.NewCoin("atom", math.NewInt(60_000_000)),
			expErr:     true,
		},
		{
			name: "success",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(collateral))
			},
			denom:      "atom",
			owner:      s.TestAccs[0],
			collateral: collateral,
			mint:       sdk.NewCoin(mintDenom, math.NewInt(50_000_000)),
			expErr:     false,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.CreateNewVault(s.Ctx, t.owner, t.collateral, t.mint)
			if t.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				vm, err := s.k.GetVaultManager(s.Ctx, denom, "nomUSD")
				s.Require().NoError(err)
				s.Require().NotEqual(maxDebt, vm.MintAvailable)

				createdVault, err := s.k.GetVault(s.Ctx, 0)
				s.Require().NoError(err)
				s.Require().Equal(createdVault.Status, types.ACTIVE)

				expectDebt := t.mint.Add(sdk.NewCoin(t.mint.Denom, math.LegacyNewDecFromInt(t.mint.Amount).Mul(types.DefaultStabilityFee).TruncateInt()))
				s.Require().Equal(createdVault.Debt, expectDebt)
				s.Require().Equal(createdVault.CollateralLocked, t.collateral)

				vaultAddr := sdk.MustAccAddressFromBech32(createdVault.Address)
				vaultBalance := s.k.BankKeeper.GetAllBalances(s.Ctx, vaultAddr)
				s.Require().Equal(vaultBalance, sdk.NewCoins(createdVault.CollateralLocked))
			}
		})
	}
}

func (s *KeeperTestSuite) TestRepayDebt() {
	s.SetupTest()
	var (
		denom           = "atom"
		repayAsset      = sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(2000000))
		collateralAsset = sdk.NewCoin(denom, math.NewInt(100000000000))
		fund            = sdk.NewCoin(denom, math.NewInt(1000000000000))
		maxDebt         = math.NewInt(2000000000)
		mintedCoin      = sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(300000000))
	)
	err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
	s.Require().NoError(err)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		repayAsset sdk.Coin
		expErr     bool
		isZeroDebt bool
	}{
		{
			name: "success",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			repayAsset: repayAsset,
			expErr:     false,
		},
		{
			name: "repay more than debt",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(mintedCoin))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:    1,
			sender:     s.TestAccs[0],
			repayAsset: sdk.NewCoin("nomUSD", math.NewInt(500000000)),
			expErr:     false,
			isZeroDebt: true,
		},
		{
			name: "not the vault owner",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				s.FundAccount(s.TestAccs[1], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:    0,
			sender:     s.TestAccs[1],
			repayAsset: repayAsset,
			expErr:     true,
		},
		{
			name: "repay incorrect denom",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
			},
			vaultId:    0,
			sender:     s.TestAccs[1],
			repayAsset: collateralAsset,
			expErr:     true,
		},
		{
			name: "inactive vault",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.SetVault(s.Ctx, types.Vault{
					Owner:            s.TestAccs[0].String(),
					Id:               0,
					Status:           types.LIQUIDATED,
					Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(1000)),
					CollateralLocked: sdk.NewCoin(denom, math.ZeroInt()),
				})
				s.Require().NoError(err)
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			repayAsset: collateralAsset,
			expErr:     true,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()

			vaultBefore, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			debt := vaultBefore.Debt

			err = s.k.RepayDebt(s.Ctx, t.vaultId, t.sender, t.repayAsset)
			if t.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				vaultAfter, err := s.k.GetVault(s.Ctx, t.vaultId)
				s.Require().NoError(err)

				if t.isZeroDebt {
					s.Require().Equal(vaultAfter.Debt.Amount, math.ZeroInt())
				} else {
					s.Require().Equal(vaultAfter.Debt, debt.Sub(repayAsset))
				}
			}
		})
	}
}

func (s *KeeperTestSuite) TestDepositToVault() {
	s.SetupTest()
	var (
		denom           = "atom"
		invalidCoin     = sdk.NewCoin("invalid", math.NewInt(1000000))
		coin            = sdk.NewCoin(denom, math.NewInt(1000000))
		collateralAsset = sdk.NewCoin(denom, math.NewInt(100000000000))
		fund            = sdk.NewCoin(denom, math.NewInt(1000000000000))
		maxDebt         = math.NewInt(2000000000)
		mintedCoin      = sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(200000000))
	)
	err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
	s.Require().NoError(err)

	tests := []struct {
		name         string
		setup        func()
		vaultId      uint64
		sender       sdk.AccAddress
		depositAsset sdk.Coin
		expErr       bool
	}{
		{
			name: "success",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:      0,
			sender:       s.TestAccs[0],
			depositAsset: coin,
			expErr:       false,
		},
		{
			name: "deposit wrong token",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:      0,
			sender:       s.TestAccs[0],
			depositAsset: invalidCoin,
			expErr:       true,
		},
		{
			name: "deposit to inactive vault",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.SetVault(s.Ctx, types.Vault{
					Id:               0,
					Status:           types.LIQUIDATED,
					Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(1000)),
					CollateralLocked: sdk.NewCoin(denom, math.ZeroInt()),
				})
				s.Require().NoError(err)
			},
			vaultId:      0,
			sender:       s.TestAccs[0],
			depositAsset: invalidCoin,
			expErr:       true,
		},
		{
			name: "insufficient balance",
			setup: func() {
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:      0,
			sender:       s.TestAccs[0],
			depositAsset: coin,
			expErr:       true,
		},
		{
			name: "not vault owner",
			setup: func() {
				s.FundAccount(s.TestAccs[1], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
			},
			vaultId:      0,
			sender:       s.TestAccs[1],
			depositAsset: coin,
			expErr:       true,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.DepositToVault(s.Ctx, t.vaultId, t.sender, t.depositAsset)
			if t.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				vault, err := s.k.GetVault(s.Ctx, t.vaultId)
				s.Require().NoError(err)
				s.Require().Equal(collateralAsset.Add(t.depositAsset), vault.CollateralLocked)

				vaultAddr := sdk.MustAccAddressFromBech32(vault.Address)
				vaultBalance := s.k.BankKeeper.GetAllBalances(s.Ctx, vaultAddr)
				s.Require().Equal(vaultBalance, sdk.NewCoins(vault.CollateralLocked))
			}

		})
	}
}

func (s *KeeperTestSuite) TestWithdrawFromVault() {

	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000000))
		invalidCoin   = sdk.NewCoin("invalid", math.NewInt(1000000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(100000000000))
		fund          = sdk.NewCoin(denom, math.NewInt(10000000000000))
		maxDebt       = math.NewInt(2000000000)
		mintedCoin    = sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(200000000))
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
		expErr     bool
	}{
		{
			name: "success",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))

				err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
				s.Require().NoError(err)

				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], coinMintToAcc, mintedCoin)
				s.Require().NoError(err)

			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
			expErr:     false,
		},
		{
			name: "inactive vault",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err := s.k.SetVault(s.Ctx, types.Vault{
					Owner:            s.TestAccs[0].String(),
					Id:               0,
					Status:           types.LIQUIDATED,
					Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(1000)),
					CollateralLocked: sdk.NewCoin(denom, math.ZeroInt()),
				})
				s.Require().NoError(err)

			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
			expErr:     true,
		},
		{
			name: "non owner",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))

				err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
				s.Require().NoError(err)

				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], coinMintToAcc, mintedCoin)
				s.Require().NoError(err)

			},
			vaultId:    0,
			sender:     s.TestAccs[1],
			collateral: coin,
			expErr:     true,
		},
		{
			name: "withdraw incorrect token",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))

				err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
				s.Require().NoError(err)

				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], coinMintToAcc, mintedCoin)
				s.Require().NoError(err)

			},
			vaultId:    0,
			sender:     s.TestAccs[1],
			collateral: invalidCoin,
			expErr:     true,
		},
		{
			name: "Violating the minimum collateral ratio after withdrawal",
			setup: func() {
				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))

				err := s.k.ActiveCollateralAsset(s.Ctx, denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt, types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
				s.Require().NoError(err)

				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], coinMintToAcc, mintedCoin)
				s.Require().NoError(err)

			},
			vaultId:    0,
			sender:     s.TestAccs[1],
			collateral: coinMintToAcc,
			expErr:     true,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()
			t.setup()

			vaultBefore, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)

			err = s.k.WithdrawFromVault(s.Ctx, t.vaultId, t.sender, t.collateral)
			if t.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				vaultAfter, err := s.k.GetVault(s.Ctx, t.vaultId)
				s.Require().NoError(err)
				s.Require().Equal(vaultBefore.CollateralLocked.Sub(vaultAfter.CollateralLocked), t.collateral)

				vaultAddr := sdk.MustAccAddressFromBech32(vaultBefore.Address)
				vaultBalanceAfter := s.k.BankKeeper.GetAllBalances(s.Ctx, vaultAddr)
				s.Require().Equal(vaultBalanceAfter, sdk.NewCoins(vaultAfter.CollateralLocked))
			}

		})
	}
}

func (s *KeeperTestSuite) TestLiquidate() {
	// s.SetupTest()

	vaultOwnerAddr := sdk.AccAddress([]byte("addr1_______________"))

	tests := []struct {
		name            string
		liquidation     types.Liquidation
		expVaultStatus  []types.VaultStatus
		shortfallAmount sdk.Coin
		moduleBalances  sdk.Coins
		reserveBalances sdk.Coins
	}{
		{
			name: "single vault - sold all, enough to cover debt",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold all at price 7,
				// Sold = 35
				// Remain collateral = 0
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(35_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.ZeroInt()),
					},
				},
			},
			expVaultStatus:  []types.VaultStatus{types.CLOSED},
			reserveBalances: sdk.NewCoins(sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(10_000_000))),
		},
		{
			name: "single vault - sold all, not enough to cover debt",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold all at price 4,
				// Sold = 20
				// Remain collateral = 0
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(20_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.ZeroInt()),
					},
				},
			},
			expVaultStatus:  []types.VaultStatus{types.LIQUIDATED},
			shortfallAmount: sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(5_000_000)),
		},
		{
			name: "single vault - remain collateral, enough to cover debt",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold 1 at 7
				// Sold 2 at 6.5
				// Sold 1 at 6
				// Sold = 26
				// Remain collateral = 1
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(26_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(1_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.LIQUIDATED},
			reserveBalances: sdk.NewCoins(
				sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(1_000_000)),
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can reconstitute vault",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold = 0
				// Remain collateral = 5
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.ZeroInt()),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(5_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.ACTIVE},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can reconstitute vault",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold = 0
				// Remain collateral = 5
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.ZeroInt()),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(5_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.ACTIVE},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can not reconstitute vault",
			liquidation: types.Liquidation{
				DebtDenom: "atom",
				MintDenom: types.DefaultMintDenoms[0],
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold 1 at 7
				// Sold = 7
				// Remain collateral = 4
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(7_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(4_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.LIQUIDATED},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.NewInt(4_000_000)), // (25/7)*0.05
			),
			shortfallAmount: sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(18_000_000)),
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()
			err := s.k.ActiveCollateralAsset(s.Ctx, "atom", "atom", "nomUSD", "USD", math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), math.NewInt(1000_000_000), types.DefaultStabilityFee, types.DefaultMintingFee, types.DefaultLiquidationPenalty, 1, 1)
			s.Require().NoError(err)

			for _, vault := range t.liquidation.LiquidatingVaults {
				vaultId, vaultAddr := s.App.VaultsKeeper.GetVaultIdAndAddress(s.Ctx)
				vault.Id = vaultId
				vault.Address = vaultAddr.String()

				err := s.App.VaultsKeeper.SetVault(s.Ctx, *vault)
				s.Require().NoError(err)

				// Fund collateral locked for vault
				lockCoins := sdk.NewCoins(t.liquidation.VaultLiquidationStatus[vaultId].RemainCollateral)
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, lockCoins)
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, vaultAddr, lockCoins)
				s.Require().NoError(err)

				// Fund sold coins to vault Module
				soldCoins := sdk.NewCoins(t.liquidation.VaultLiquidationStatus[vaultId].Sold)
				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, soldCoins)
				s.Require().NoError(err)
			}

			err = s.App.VaultsKeeper.Liquidate(s.Ctx, t.liquidation, types.DefaultMintDenoms[0])
			s.Require().NoError(err)

			if t.reserveBalances != nil {
				reserveModuleAddr := s.App.AccountKeeper.GetModuleAddress(types.ReserveModuleName)
				reserveBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, reserveModuleAddr)
				// t.reserveBalances.Sort()
				s.Require().Equal(reserveBalance, t.reserveBalances)
			}

			if !t.shortfallAmount.IsNil() {
				shortfallAmount, err := s.App.VaultsKeeper.ShortfallAmount.Get(s.Ctx)
				s.Require().NoError(err)
				s.Require().Equal(t.shortfallAmount.Amount, shortfallAmount)
			}

			for i, vault := range t.liquidation.LiquidatingVaults {
				updatedVault, err := s.App.VaultsKeeper.GetVault(s.Ctx, vault.Id)
				s.Require().NoError(err)
				s.Require().Equal(updatedVault.Status, t.expVaultStatus[i])
			}

		})
	}
}
