package keeper

import (
	"github.com/onomyprotocol/reserve/x/portal/types"
)

var _ types.QueryServer = Keeper{}
