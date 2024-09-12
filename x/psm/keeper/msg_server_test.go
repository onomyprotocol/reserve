package keeper_test

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestMsgServerSwapToIST() {
	s.SetupTest()

	tests := []struct {
		name  string
		addr  sdk.AccAddress
		setup func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToIST

		expectPass      bool
		expectedReceive math.Int
	}{
		{
			name: "success",
			addr: s.TestAccs[0],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToIST {
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

				amountSwap := sdk.NewCoin(usdt, math.NewInt(1000))
				return &types.MsgSwapToIST{
					Address: s.TestAccs[0].String(),
					Coin:    &amountSwap,
				}
			},

			expectPass:      true,
			expectedReceive: math.NewInt(999),
		},
		{
			name: "insufficient balance",
			addr: s.TestAccs[1],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToIST {
				sc := types.Stablecoin{
					Denom:      usdt,
					LimitTotal: limitUSDT,
					Price:      math.LegacyMustNewDecFromStr("1"),
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
				}
				s.k.SetStablecoin(s.Ctx, sc)

				amountSwap := sdk.NewCoin(usdt, math.NewInt(1000))
				return &types.MsgSwapToIST{
					Address: s.TestAccs[1].String(),
					Coin:    &amountSwap,
				}
			},

			expectPass:      false,
			expectedReceive: math.NewInt(999),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			msg := t.setup(s.Ctx, s.k)

			_, err := s.msgServer.SwapToIST(s.Ctx, msg)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, types.InterStableToken)
				s.Require().Equal(t.expectedReceive, balance.Amount)

			} else {
				s.Require().Error(err)
			}

		})
	}
}

func (s *KeeperTestSuite) TestMsgSwapToStablecoin() {
	s.SetupTest()

	tests := []struct {
		name  string
		addr  sdk.AccAddress
		setup func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToStablecoin

		expectPass         bool
		expectedBalanceIST math.Int
	}{
		{
			name: "success",
			addr: s.TestAccs[0],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToStablecoin {
				// swaptoIST
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

				amountSwap := sdk.NewCoin(usdt, math.NewInt(1001))
				msg := &types.MsgSwapToIST{
					Address: s.TestAccs[0].String(),
					Coin:    &amountSwap,
				}
				_, err = s.msgServer.SwapToIST(s.Ctx, msg)
				s.Require().NoError(err)

				return &types.MsgSwapToStablecoin{
					Address: s.TestAccs[0].String(),
					ToDenom: usdt,
					Amount:  math.NewInt(1000),
				}
			},

			expectPass:         true,
			expectedBalanceIST: math.NewInt(0),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			msg := t.setup(s.Ctx, s.k)

			_, err := s.msgServer.SwapToStablecoin(s.Ctx, msg)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, types.InterStableToken)
				s.Require().Equal(t.expectedBalanceIST.String(), balance.Amount.String())

			} else {
				s.Require().Error(err)
			}

		})
	}
}
