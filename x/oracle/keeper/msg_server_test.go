package keeper_test

import (
	"time"

	"github.com/onomyprotocol/reserve/x/oracle/types"
)

func (s *KeeperTestSuite) TestUpdateParams() {
	s.SetupTest()

	paramDefault := s.k.GetParams(s.Ctx)
	s.Require().Equal(paramDefault.AllowedPriceDelay, types.DefauAllowedPriceDelay)

	allowedPriceDelayUpdate := time.Hour * 10

	msgUpdateParams := types.MsgUpdateParams{
		Authority: s.k.GetAuthority(),
		Params:    types.NewParams(allowedPriceDelayUpdate),
	}

	_, err := s.msgServer.UpdateParams(s.Ctx, &msgUpdateParams)
	s.Require().NoError(err)

	paramsNew := s.k.GetParams(s.Ctx)
	s.Require().Equal(paramsNew.AllowedPriceDelay, allowedPriceDelayUpdate)
}
