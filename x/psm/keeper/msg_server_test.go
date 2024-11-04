package keeper_test

import (
	"context"
	// "fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestMsgServerSwapTonomUSD() {
	s.SetupTest()

	tests := []struct {
		name  string
		addr  sdk.AccAddress
		setup func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToNom

		expectPass      bool
		expectedReceive math.Int
	}{
		{
			name: "success",
			addr: s.TestAccs[0],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToNom {
				coinsMint := sdk.NewCoins(sdk.NewCoin(usdt, math.NewInt(1000000)))
				err := keeper.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				_, err = s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:      usdt,
					LimitTotal: limitUSDT,
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
					NomType:    "nomUSD",
				})
				s.Require().NoError(err)

				amountSwap := sdk.NewCoin(usdt, math.NewInt(1000))
				return &types.MsgSwapToNom{
					Address: s.TestAccs[0].String(),
					Coin:    amountSwap,
				}
			},

			expectPass:      true,
			expectedReceive: math.NewInt(999),
		},
		{
			name: "insufficient balance",
			addr: s.TestAccs[1],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToNom {
				_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:      usdc,
					LimitTotal: limitUSDC,
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
					NomType:    "nomUSD",
				})
				s.Require().NoError(err)

				amountSwap := sdk.NewCoin(usdc, math.NewInt(1000))
				return &types.MsgSwapToNom{
					Address: s.TestAccs[1].String(),
					Coin:    amountSwap,
				}
			},

			expectPass:      false,
			expectedReceive: math.NewInt(999),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			msg := t.setup(s.Ctx, s.k)

			_, err := s.msgServer.SwapToNom(s.Ctx, msg)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, types.DefaultMintDenoms[0])
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

		expectPass            bool
		expectedBalancenomUSD math.Int
	}{
		{
			name: "success",
			addr: s.TestAccs[0],
			setup: func(ctx context.Context, keeper keeper.Keeper) *types.MsgSwapToStablecoin {
				// swaptonomUSD
				coinsMint := sdk.NewCoins(sdk.NewCoin(usdt, math.NewInt(1000000)))
				err := keeper.BankKeeper.MintCoins(ctx, types.ModuleName, coinsMint)
				s.Require().NoError(err)
				err = keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, s.TestAccs[0], coinsMint)
				s.Require().NoError(err)

				_, err = s.msgServer.AddStableCoinProposal(s.Ctx, &types.MsgAddStableCoin{
					Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
					Denom:      usdt,
					LimitTotal: limitUSDT,
					FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
					FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
					NomType:    "nomUSD",
				})
				s.Require().NoError(err)

				amountSwap := sdk.NewCoin(usdt, math.NewInt(1001))
				msg := &types.MsgSwapToNom{
					Address: s.TestAccs[0].String(),
					Coin:    amountSwap,
				}
				_, err = s.msgServer.SwapToNom(s.Ctx, msg)
				s.Require().NoError(err)

				return &types.MsgSwapToStablecoin{
					Address: s.TestAccs[0].String(),
					ToDenom: usdt,
					Coin:    sdk.NewCoin(types.DefaultMintDenoms[0], math.NewInt(1000)),
				}
			},

			expectPass:            true,
			expectedBalancenomUSD: math.NewInt(0),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			msg := t.setup(s.Ctx, s.k)

			_, err := s.msgServer.SwapToStablecoin(s.Ctx, msg)
			if t.expectPass {
				s.Require().NoError(err)
				balance := s.k.BankKeeper.GetBalance(s.Ctx, t.addr, types.DefaultMintDenoms[0])
				s.Require().Equal(t.expectedBalancenomUSD.String(), balance.Amount.String())

			} else {
				s.Require().Error(err)
			}

		})
	}
}
