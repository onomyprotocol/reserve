package keeper_test

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (s *KeeperTestSuite) TestBeginBlock() {
	var (
		stabilityFee    = math.LegacyMustNewDecFromStr("0.1") //10%
		denom           = "atom"
		maxDebt         = math.NewInt(100_000_000_000_000_000)
		fund            = sdk.NewCoin(denom, math.NewInt(1000000000000))
		collateralAsset = sdk.NewCoin(denom, math.NewInt(100000000000))
		mintedCoin      = sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(200000000))
	)

	tests := []struct {
		name              string
		setup             func()
		expErr            bool
		expDebt           sdk.Coin
		expLastUpdateTime time.Time
		hasVaults         bool
	}{
		{
			name: "success: one vault",
			setup: func() { // 100000000000atom debt 210000000nomUSD(get 200000000nomUSD + 10000000nomUSD MintingFee)
				err := s.k.ActiveCollateralAsset(s.Ctx,
					denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("1.6"),
					math.LegacyMustNewDecFromStr("1.5"),
					maxDebt, stabilityFee,
					types.DefaultMintingFee,
					types.DefaultLiquidationPenalty, 1, 1,
				)
				s.Require().NoError(err)

				s.FundAccount(s.TestAccs[0], types.ModuleName, sdk.NewCoins(fund))
				err = s.k.CreateNewVault(s.Ctx, s.TestAccs[0], collateralAsset, mintedCoin)
				s.Require().NoError(err)
				vault, err := s.k.GetVault(s.Ctx, 0)
				s.Require().NoError(err)
				fmt.Println(vault.Debt)

				p := s.k.GetParams(s.Ctx)
				p.ChargingPeriod = time.Second * 15
				err = s.k.SetParams(s.Ctx, p)
				s.Require().NoError(err)
				err = s.k.LastUpdateTime.Set(s.Ctx, types.LastUpdate{Time: time.Now().Add(-time.Second * 16)})
				s.Require().NoError(err)
			},
			expErr:            false,
			expDebt:           sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(210000010)), // 10% * 15/(60 * 60 * 24 * 365) * 210000000 + 210000000 = 210000010
			expLastUpdateTime: s.Ctx.BlockTime(),
			hasVaults:         false,
		},
		{
			name: "success: no vault, LastUpdateTime updates",
			setup: func() { // 100000000000atom debt 210000000nomUSD(get 200000000nomUSD + 10000000nomUSD MintingFee)
				err := s.k.ActiveCollateralAsset(s.Ctx,
					denom, denom, "nomUSD", "USD", math.LegacyMustNewDecFromStr("1.6"),
					math.LegacyMustNewDecFromStr("1.5"),
					maxDebt, stabilityFee,
					types.DefaultMintingFee,
					types.DefaultLiquidationPenalty, 1, 1,
				)
				s.Require().NoError(err)

				p := s.k.GetParams(s.Ctx)
				p.ChargingPeriod = time.Second * 15
				err = s.k.SetParams(s.Ctx, p)
				s.Require().NoError(err)
				err = s.k.LastUpdateTime.Set(s.Ctx, types.LastUpdate{Time: time.Now().Add(-time.Second * 16)})
				s.Require().NoError(err)
			},
			expErr:            false,
			expLastUpdateTime: s.Ctx.BlockTime(),
			hasVaults:         false,
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()
			t.setup()
			err := s.k.BeginBlocker(s.Ctx)
			if t.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				LastUpdateTime, err := s.k.LastUpdateTime.Get(s.Ctx)
				s.Require().NoError(err)
				s.Require().Equal(t.expLastUpdateTime, LastUpdateTime.Time)

				if t.hasVaults {
					vault, err := s.k.GetVault(s.Ctx, 0)
					s.Require().NoError(err)
					fmt.Println(vault.Debt)
					s.Require().True(vault.Debt.Equal(t.expDebt))
				}
			}
		})
	}
}
