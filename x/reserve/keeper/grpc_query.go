package keeper

import (
	"reserve/x/reserve/types"
)

var _ types.QueryServer = Keeper{}
