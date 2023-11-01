package keeper

import (
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

var _ types.QueryServer = Keeper{}
