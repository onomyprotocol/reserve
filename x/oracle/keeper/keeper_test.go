package keeper_test

import (
	"testing"
	testifysuite "github.com/stretchr/testify/suite"
	apptesting "github.com/onomyprotocol/reserve/app/apptesting"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestKeeperTestSuite(t *testing.T) {
	testifysuite.Run(t, new(KeeperTestSuite))
}