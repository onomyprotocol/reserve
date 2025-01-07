package keeper_test

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestAddStableCoinProposal() {
	s.SetupTest()

	proAdd := types.MsgAddStableCoin{
		Denom:                usdt,
		AddressPayStableInit: s.communityAddress.String(),
		AmountStableInit:     amountStableInit,
		LimitTotal:           limitUSDT,
		Symbol:               usdt,
		Authority:            authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:                math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:               math.LegacyMustNewDecFromStr("0.001"),
		OracleScript:         44,
	}

	_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, err := s.k.StablecoinInfos.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)

	proAdd2 := types.MsgAddStableCoin{
		Denom:                usdc,
		AddressPayStableInit: s.communityAddress.String(),
		AmountStableInit:     amountStableInit,
		LimitTotal:           limitUSDT,
		Symbol:               usdt,
		Authority:            authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:                math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:               math.LegacyMustNewDecFromStr("0.001"),
		OracleScript:         44,
	}

	_, err = s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd2)
	s.Require().NoError(err)

	stablecoin2, err := s.k.StablecoinInfos.Get(s.Ctx, proAdd2.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin2.Denom, proAdd2.Denom)
	s.Require().Equal(stablecoin2.LimitTotal, limitUSDT)

	coinsStable := s.App.BankKeeper.GetAllBalances(s.Ctx, s.App.AccountKeeper.GetModuleAddress(types.ModuleName))
	s.Require().Equal(coinsStable.AmountOf(usdt), amountStableInit)
	s.Require().Equal(coinsStable.AmountOf(usdc), amountStableInit)
}

func (s *KeeperTestSuite) TestAddStableCoinProposalFromAnyAddress() {
	s.SetupTest()

	s.FundAccount(s.TestAccs[2], types.ModuleName, sdk.NewCoins(sdk.NewCoin(usdt, amountStableInit)))

	proAdd := types.MsgAddStableCoin{
		Denom:                usdt,
		AddressPayStableInit: s.TestAccs[2].String(),
		AmountStableInit:     amountStableInit,
		LimitTotal:           limitUSDT,
		Symbol:               usdt,
		Authority:            authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:                math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:               math.LegacyMustNewDecFromStr("0.001"),
		OracleScript:         44,
	}

	_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, err := s.k.StablecoinInfos.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)

	coinsStable := s.App.BankKeeper.GetAllBalances(s.Ctx, s.App.AccountKeeper.GetModuleAddress(types.ModuleName))
	s.Require().Equal(coinsStable.AmountOf(usdt), amountStableInit)
}

func (s *KeeperTestSuite) TestUpdateStableCoinProposal() {
	s.SetupTest()

	proAdd := types.MsgAddStableCoin{
		Denom:                usdt,
		AddressPayStableInit: s.communityAddress.String(),
		AmountStableInit:     amountStableInit,
		LimitTotal:           limitUSDT,
		Symbol:               usdt,
		Authority:            authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:                math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:               math.LegacyMustNewDecFromStr("0.001"),
		OracleScript:         44,
	}

	_, err := s.msgServer.AddStableCoinProposal(s.Ctx, &proAdd)
	s.Require().NoError(err)

	stablecoin, err := s.k.StablecoinInfos.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitUSDT)

	// update stablecoin
	limitTotalUpdates := math.NewInt(2000000)

	proUpdates := types.MsgUpdatesStableCoin{
		Denom:        usdt,
		LimitTotal:   limitTotalUpdates,
		Authority:    authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		FeeIn:        math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:       math.LegacyMustNewDecFromStr("0.001"),
		Symbol:       usdt,
		OracleScript: 44,
	}

	_, err = s.msgServer.UpdatesStableCoinProposal(s.Ctx, &proUpdates)
	s.Require().NoError(err)

	stablecoin, err = s.k.StablecoinInfos.Get(s.Ctx, proAdd.Denom)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin.Denom, proAdd.Denom)
	s.Require().Equal(stablecoin.LimitTotal, limitTotalUpdates)

}
