package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestAddStableCoinProposal() {
	s.SetupTest()

	proAdd := types.AddStableCoinProposal{
		Title:       "title",
		Description: "description",
		Denom:       usdt,
		LimitTotal:  limitUSDT,
		Price:       math.LegacyMustNewDecFromStr("1"),
		FeeIn:       math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:      math.LegacyMustNewDecFromStr("0.001"),
	}

	err := s.k.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, found := s.k.GetStablecoin(s.Ctx, proAdd.Denom)
	s.Require().True(found)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)
}

func (s *KeeperTestSuite) TestUpdateStableCoinProposal() {
	s.SetupTest()

	proAdd := types.AddStableCoinProposal{
		Title:       "title",
		Description: "description",
		Denom:       usdt,
		LimitTotal:  limitUSDT,
		Price:       math.LegacyMustNewDecFromStr("1"),
		FeeIn:       math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:      math.LegacyMustNewDecFromStr("0.001"),
	}

	err := s.k.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, found := s.k.GetStablecoin(s.Ctx, proAdd.Denom)
	s.Require().True(found)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)

	// update stablecoin
	limitTotalUpdates := math.NewInt(2000000)

	proUpdates := types.UpdatesStableCoinProposal{
		Title:             "title",
		Description:       "description",
		Denom:             usdt,
		UpdatesLimitTotal: limitTotalUpdates,
		Price:             math.LegacyMustNewDecFromStr("1"),
		FeeIn:             math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:            math.LegacyMustNewDecFromStr("0.001"),
	}

	err = s.k.UpdatesStableCoinProposal(s.Ctx, &proUpdates)
	s.Require().NoError(err)

	stablecoin, found = s.k.GetStablecoin(s.Ctx, proAdd.Denom)
	s.Require().True(found)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitTotalUpdates)

}
