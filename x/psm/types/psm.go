package types

import (
	"cosmossdk.io/math"
)

func GetMsgStablecoin(msg getStablecoinFromMsg) Stablecoin {
	return Stablecoin{
		Denom:               msg.GetDenom(),
		LimitTotal:          msg.GetLimitTotal(),
		FeeIn:               msg.GetFeeIn(),
		FeeOut:              msg.GetFeeOut(),
		TotalStablecoinLock: math.ZeroInt(),
		FeeMaxStablecoin:    msg.GetFeeIn().Add(msg.GetFeeOut()),
		NomType:             msg.GetNomType(),
	}
}

type getStablecoinFromMsg interface {
	GetDenom() string
	GetLimitTotal() math.Int
	// GetPrice() math.LegacyDec
	GetFeeIn() math.LegacyDec
	GetFeeOut() math.LegacyDec
	GetNomType() string
}
