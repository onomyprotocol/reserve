package keeper

import (
	"github.com/onomyprotocol/reserve/x/interquery/types"
)

var _ types.QueryServer = Keeper{}
