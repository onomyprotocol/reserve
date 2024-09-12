package keeper_test

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestSwapToIST() {
	s.SetupTest()

	tests := []struct {
		name       string
		setup      func(ctx context.Context, keeper keeper.Keeper)
		addr       sdk.AccAddress
		stablecoin sdk.Coin

		expectPass      bool
		expectedReceive math.Int
		expectedFee     math.LegacyDec
	}{
		{
			name: "success",
			setup: func(ctx context.Context, keeper keeper.Keeper) {
				coinsMint := sdk.NewCoins(sdk.NewCoin(usdt, math.NewInt(1000000)))
				err := keeper.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				sc := types.Stablecoin{
					Denom:      usdt,
					LimitTotal: limitUSDT,
					Price:      math.LegacyMustNewDecFromStr("1"),
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
				}
				s.k.SetStablecoin(s.Ctx, sc)

			},
			addr:            s.TestAccs[0],
			stablecoin:      sdk.NewCoin(usdt, math.NewInt(1000)),
			expectPass:      true,
			expectedReceive: math.NewInt(999),
			expectedFee:     math.LegacyMustNewDecFromStr("1"),
		},
		{
			name: "insufficient balance",
			setup: func(ctx context.Context, keeper keeper.Keeper) {
				sc := types.Stablecoin{
					Denom:      usdt,
					LimitTotal: limitUSDT,
					Price:      math.LegacyMustNewDecFromStr("1"),
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
				}
				s.k.SetStablecoin(s.Ctx, sc)

			},
			addr:            s.TestAccs[1],
			stablecoin:      sdk.NewCoin(usdt, math.NewInt(1000)),
			expectPass:      false,
			expectedReceive: math.NewInt(999),
			expectedFee:     math.LegacyMustNewDecFromStr("1"),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup(s.Ctx, s.k)

			receiveAmount, fee, err := s.k.SwapToIST(s.Ctx, t.addr, t.stablecoin)
			if t.expectPass {
				s.Require().NoError(err)
				s.Require().Equal(t.expectedReceive, receiveAmount)
				s.Require().Equal(t.expectedFee, fee.Amount)
			} else {
				s.Require().Error(err)
			}

		})
	}
}

func (s *KeeperTestSuite) TestSwapToStablecoin() {
	s.SetupTest()

	tests := []struct {
		name    string
		setup   func(ctx context.Context, keeper keeper.Keeper)
		addr    sdk.AccAddress
		amount  math.Int
		toDenom string

		expectPass      bool
		expectedReceive math.Int
		expectedFee     math.LegacyDec
	}{
		{
			name: "success",
			setup: func(ctx context.Context, keeper keeper.Keeper) {
				coinsMint := sdk.NewCoins(sdk.NewCoin(types.InterStableToken, math.NewInt(1000000)))
				err := keeper.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				sc := types.Stablecoin{
					Denom:      usdt,
					LimitTotal: limitUSDT,
					Price:      math.LegacyMustNewDecFromStr("1"),
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
				}
				s.k.SetStablecoin(s.Ctx, sc)

			},
			addr:            s.TestAccs[0],
			amount:          math.NewInt(1000),
			toDenom:         usdt,
			expectPass:      true,
			expectedReceive: math.NewInt(999),
			expectedFee:     math.LegacyMustNewDecFromStr("1"),
		},
		{
			name: "insufficient balance",
			setup: func(ctx context.Context, keeper keeper.Keeper) {
				sc := types.Stablecoin{
					Denom:      usdt,
					LimitTotal: limitUSDT,
					Price:      math.LegacyMustNewDecFromStr("1"),
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
				}
				s.k.SetStablecoin(s.Ctx, sc)

			},
			addr:            s.TestAccs[1],
			amount:          math.NewInt(1000),
			toDenom:         usdt,
			expectPass:      false,
			expectedReceive: math.NewInt(999),
			expectedFee:     math.LegacyMustNewDecFromStr("1"),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup(s.Ctx, s.k)

			receiveAmount, fee, err := s.k.SwapToStablecoin(s.Ctx, t.addr, t.amount, t.toDenom)
			if t.expectPass {
				s.Require().NoError(err)
				s.Require().Equal(t.expectedReceive, receiveAmount)
				s.Require().Equal(t.expectedFee, fee.Amount)
			} else {
				s.Require().Error(err)
			}

		})
	}
}
