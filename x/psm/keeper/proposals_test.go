package keeper_test

import (
	"cosmossdk.io/math"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestAddStableCoinProposal() {
	s.SetupTest()

	proAdd := types.MsgAddStableCoin{
		Denom:      usdt,
		LimitTotal: limitUSDT,
		Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}

	_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, err := s.k.Stablecoins.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)
}

func (s *KeeperTestSuite) TestUpdateStableCoinProposal() {
	s.SetupTest()

	proAdd := types.MsgAddStableCoin{
		Denom:      usdt,
		LimitTotal: limitUSDT,
		Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}

	_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, err := s.k.Stablecoins.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)

	// update stablecoin
	limitTotalUpdates := math.NewInt(2000000)

	proUpdates := types.MsgUpdatesStableCoin{
		Denom:      usdt,
		LimitTotal: limitTotalUpdates,
		Authority:  authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}

	_, err = s.msgServer.UpdatesStableCoinProposal(s.Ctx, &proUpdates)
	s.Require().NoError(err)

	stablecoin, err = s.k.Stablecoins.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitTotalUpdates)

}
