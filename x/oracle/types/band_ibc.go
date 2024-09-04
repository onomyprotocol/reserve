package types

import (
	// "fmt"

	// bandobi "github.com/bandprotocol/bandchain-packet/obi"

	// bandprice "github.com/InjectiveLabs/injective-core/injective-chain/modules/oracle/bandchain/hooks/price"
)

func NewOracleRequestPacketData(clientID string, calldata []byte, r *BandOracleRequest) OracleRequestPacketData {
	return OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: uint64(r.OracleScriptId),
		Calldata:       calldata,
		AskCount:       r.AskCount,
		MinCount:       r.MinCount,
		FeeLimit:       r.FeeLimit,
		PrepareGas:     r.PrepareGas,
		ExecuteGas:     r.ExecuteGas,
	}
}