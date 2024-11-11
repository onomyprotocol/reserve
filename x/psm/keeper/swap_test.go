package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestSwapToOnomyStableToken() {
	tests := []struct {
		name          string
		addr          sdk.AccAddress
		offerCoin     sdk.Coin
		expectedDenom string

		setup func()

		expectPass      bool
		expectedReceive math.Int
	}{
		{
			name:          "success",
			addr:          s.TestAccs[0],
			offerCoin:     sdk.NewCoin(usdt, math.NewInt(1000)),
			expectedDenom: types.ReserveStableCoinDenom,
			setup: func() {
				coinsMint := sdk.NewCoins(sdk.NewCoin(usdt, math.NewInt(1000000)))
				err := s.k.BankKeeper.MintCoins(s.Ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = s.k.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				_, err = s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:    authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:        usdt,
					SymBol:       usdt,
					LimitTotal:   limitUSDT,
					FeeIn:        math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:       math.LegacyMustNewDecFromStr("0.001"),
					OracleScript: 44,
				})
				s.Require().NoError(err)
			},

			expectPass:      true,
			expectedReceive: math.NewInt(999),
		},
		{
			name: "insufficient balance",
			addr: s.TestAccs[1],
			setup: func() {
				_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:    authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:        usdc,
					SymBol:       usdc,
					LimitTotal:   limitUSDC,
					FeeIn:        math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:       math.LegacyMustNewDecFromStr("0.001"),
					OracleScript: 44,
				})
				s.Require().NoError(err)
			},

			expectPass:      false,
			expectedReceive: math.NewInt(999),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()
			t.setup()

			err := s.k.SwapToOnomyStableToken(s.Ctx, t.addr, t.offerCoin, t.expectedDenom)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, types.ReserveStableCoinDenom)
				s.Require().Equal(t.expectedReceive, balance.Amount)

			} else {
				s.Require().Error(err)
			}

		})
	}
}

func (s *KeeperTestSuite) TestSwapToOtherStablecoin() {
	tests := []struct {
		name          string
		addr          sdk.AccAddress
		offerCoin     sdk.Coin
		expectedDenom string

		setup func()

		expectPass      bool
		expectedReceive math.Int
	}{
		{
			name:          "success",
			addr:          s.TestAccs[0],
			offerCoin:     sdk.NewCoin(types.ReserveStableCoinDenom, math.NewInt(1000)),
			expectedDenom: usdt,
			setup: func() {
				coinsMint := sdk.NewCoins(sdk.NewCoin(usdt, math.NewInt(2000000)))
				err := s.k.BankKeeper.MintCoins(s.Ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = s.k.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				_, err = s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:    authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:        usdt,
					SymBol:       usdt,
					LimitTotal:   limitUSDT,
					FeeIn:        math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:       math.LegacyMustNewDecFromStr("0.001"),
					OracleScript: 44,
				})
				s.Require().NoError(err)

				// lock
				err = s.k.SwapToOnomyStableToken(s.Ctx, s.TestAccs[0], coinsMint[0], types.ReserveStableCoinDenom)
				s.Require().NoError(err)
			},

			expectPass:      true,
			expectedReceive: math.NewInt(999),
		},
		{
			name: "insufficient balance",
			addr: s.TestAccs[1],
			setup: func() {
				_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:    authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:        usdc,
					SymBol:       usdc,
					LimitTotal:   limitUSDC,
					FeeIn:        math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:       math.LegacyMustNewDecFromStr("0.001"),
					OracleScript: 44,
				})
				s.Require().NoError(err)
			},

			expectPass:      false,
			expectedReceive: math.NewInt(999),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()
			t.setup()

			err := s.k.SwapToOtherStablecoin(s.Ctx, t.addr, t.offerCoin, t.expectedDenom)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, usdt)
				s.Require().True(t.expectedReceive.Equal(balance.Amount))
			} else {
				s.Require().Error(err)
			}

		})
	}
}
