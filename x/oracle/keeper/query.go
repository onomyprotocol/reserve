package keeper

import (
	"reserve/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
