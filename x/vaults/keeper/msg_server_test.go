package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"

	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (s *KeeperTestSuite) TestBurnShortfallByMintDenom() {
	testcases := []struct {
		name                        string
		mintDenom                   string
		setup                       func() types.MsgBurnShortfall
		expShortfallAmountAfterBurn math.Int
		expReserveBalcesAfterBurn   math.Int
		expPass                     bool
	}{
		{
			name:      "success, shortfallAmount is less than reserve balances",
			mintDenom: "fxUSD",
			setup: func() types.MsgBurnShortfall {
				// make sure reserve has money
				mintCoin := sdk.NewCoins(sdk.NewCoin("fxUSD", math.NewInt(10_000_000)))
				s.FundAccount(s.TestAccs[0], types.ModuleName, mintCoin)
				err := s.k.BankKeeper.SendCoinsFromAccountToModule(s.Ctx, s.TestAccs[0], types.ReserveModuleName, mintCoin)
				s.Require().NoError(err)

				// make sure Guaranteed Shortfall Amount
				err = s.k.ShortfallAmount.Set(s.Ctx, "fxUSD", math.NewInt(1_000_000))
				s.Require().NoError(err)
				return types.MsgBurnShortfall{
					Authority: "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
					MintDenom: "fxUSD",
				}
			},
			expShortfallAmountAfterBurn: math.ZeroInt(),
			expReserveBalcesAfterBurn:   math.NewInt(9_000_000),
			expPass:                     true,
		},
		{
			name:      "success, reserve balancess is less than shortfallAmount",
			mintDenom: "fxUSD",
			setup: func() types.MsgBurnShortfall {
				// make sure reserve has money
				mintCoin := sdk.NewCoins(sdk.NewCoin("fxUSD", math.NewInt(1_000_000)))
				s.FundAccount(s.TestAccs[0], types.ModuleName, mintCoin)
				err := s.k.BankKeeper.SendCoinsFromAccountToModule(s.Ctx, s.TestAccs[0], types.ReserveModuleName, mintCoin)
				s.Require().NoError(err)

				// make sure Guaranteed Shortfall Amount
				err = s.k.ShortfallAmount.Set(s.Ctx, "fxUSD", math.NewInt(10_000_000))
				s.Require().NoError(err)
				return types.MsgBurnShortfall{
					Authority: "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
					MintDenom: "fxUSD",
				}
			},
			expShortfallAmountAfterBurn: math.NewInt(9_000_000),
			expReserveBalcesAfterBurn:   math.ZeroInt(),
			expPass:                     true,
		},
		{
			name:      "success, reserve balancess no money",
			mintDenom: "fxUSD",
			setup: func() types.MsgBurnShortfall {
				// make sure Guaranteed Shortfall Amount
				err := s.k.ShortfallAmount.Set(s.Ctx, "fxUSD", math.NewInt(10_000_000))
				s.Require().NoError(err)
				return types.MsgBurnShortfall{
					Authority: "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
					MintDenom: "fxUSD",
				}
			},
			expShortfallAmountAfterBurn: math.NewInt(10_000_000),
			expReserveBalcesAfterBurn:   math.ZeroInt(),
			expPass:                     true,
		},
		{
			name:      "fail, government account not the signatory for the proposed message",
			mintDenom: "fxUSD",
			setup: func() types.MsgBurnShortfall {
				// make sure reserve has money
				mintCoin := sdk.NewCoins(sdk.NewCoin("fxUSD", math.NewInt(1_000_000)))
				s.FundAccount(s.TestAccs[0], types.ModuleName, mintCoin)
				err := s.k.BankKeeper.SendCoinsFromAccountToModule(s.Ctx, s.TestAccs[0], types.ReserveModuleName, mintCoin)
				s.Require().NoError(err)

				// make sure Guaranteed Shortfall Amount
				err = s.k.ShortfallAmount.Set(s.Ctx, "fxUSD", math.NewInt(10_000_000))
				s.Require().NoError(err)
				return types.MsgBurnShortfall{
					Authority: s.TestAccs[0].String(),
					MintDenom: "fxUSD",
				}
			},
			expShortfallAmountAfterBurn: math.NewInt(9_000_000),
			expReserveBalcesAfterBurn:   math.ZeroInt(),
			expPass:                     false,
		},
		{
			name:      "fail, denom is not in the allowed mint denom list",
			mintDenom: "fxUSD",
			setup: func() types.MsgBurnShortfall {
				// make sure reserve has money
				mintCoin := sdk.NewCoins(sdk.NewCoin("fxUSD", math.NewInt(1_000_000)))
				s.FundAccount(s.TestAccs[0], types.ModuleName, mintCoin)
				err := s.k.BankKeeper.SendCoinsFromAccountToModule(s.Ctx, s.TestAccs[0], types.ReserveModuleName, mintCoin)
				s.Require().NoError(err)

				// make sure Guaranteed Shortfall Amount
				err = s.k.ShortfallAmount.Set(s.Ctx, "fxUSD", math.NewInt(10_000_000))
				s.Require().NoError(err)
				return types.MsgBurnShortfall{
					Authority: s.TestAccs[0].String(),
					MintDenom: "atom",
				}
			},
			expShortfallAmountAfterBurn: math.NewInt(9_000_000),
			expReserveBalcesAfterBurn:   math.ZeroInt(),
			expPass:                     false,
		},
	}

	for _, t := range testcases {
		s.Run(t.name, func() {
			s.SetupTest()
			msg := t.setup()

			// burn Shortfall
			_, err := s.msgServer.BurnShortfall(s.Ctx, &msg)
			if t.expPass {
				s.Require().NoError(err)

				// check reserve balances after burn
				reserveBalces := s.k.BankKeeper.GetAllBalances(s.Ctx, s.App.AccountKeeper.GetModuleAddress(types.ReserveModuleName))
				s.Require().True(reserveBalces.AmountOf(t.mintDenom).Equal(t.expReserveBalcesAfterBurn))

				// check ShortfallAmount after burn

				shortfallAmountAfterBurn, err := s.k.ShortfallAmount.Get(s.Ctx, t.mintDenom)
				s.Require().NoError(err)

				s.Require().True(shortfallAmountAfterBurn.Equal(t.expShortfallAmountAfterBurn))
			} else {
				s.Require().Error(err)
			}
		})
	}
}
