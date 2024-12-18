package keeper_test

import (
	"testing"

	apptesting "github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	testifysuite "github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	k         keeper.Keeper
	msgServer types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	testifysuite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.k = s.App.OracleKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.k)
}
