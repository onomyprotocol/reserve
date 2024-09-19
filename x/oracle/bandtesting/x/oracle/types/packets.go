package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	bandoracletypes "github.com/onomyprotocol/reserve/x/oracle/bandchain/oracle/types"

)

// NewOracleRequestPacketData contructs a new OracleRequestPacketData instance
func NewOracleRequestPacketData(
	clientID string, oracleScriptID OracleScriptID, calldata []byte, askCount uint64, minCount uint64, feeLimit sdk.Coins, prepareGas uint64, executeGas uint64,
) OracleRequestPacketData {
	return OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: oracleScriptID,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// ValidateBasic is used for validating the request.
func (p OracleRequestPacketData) ValidateBasic() error {
	if p.MinCount <= 0 {
		return errors.Wrapf(bandoracletypes.ErrInvalidMinCount, "got: %d", p.MinCount)
	}
	if p.AskCount < p.MinCount {
		return errors.Wrapf(bandoracletypes.ErrInvalidAskCount, "got: %d, min count: %d", p.AskCount, p.MinCount)
	}
	if len(p.ClientID) > MaxClientIDLength {
		return bandoracletypes.WrapMaxError(bandoracletypes.ErrTooLongClientID, len(p.ClientID), MaxClientIDLength)
	}
	if p.PrepareGas <= 0 {
		return errors.Wrapf(bandoracletypes.ErrInvalidOwasmGas, "invalid prepare gas: %d", p.PrepareGas)
	}
	if p.ExecuteGas <= 0 {
		return errors.Wrapf(bandoracletypes.ErrInvalidOwasmGas, "invalid execute gas: %d", p.ExecuteGas)
	}
	if p.PrepareGas+p.ExecuteGas > MaximumOwasmGas {
		return errors.Wrapf(bandoracletypes.ErrInvalidOwasmGas, "sum of prepare gas and execute gas (%d) exceed %d", p.PrepareGas+p.ExecuteGas, MaximumOwasmGas)
	}
	if !p.FeeLimit.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, p.FeeLimit.String())
	}
	return nil
}

// GetBytes is a helper for serialising
func (p OracleRequestPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

func NewOracleRequestPacketAcknowledgement(requestID RequestID) *OracleRequestPacketAcknowledgement {
	return &OracleRequestPacketAcknowledgement{
		RequestID: requestID,
	}
}

// NewOracleResponsePacketData contructs a new OracleResponsePacketData instance
func NewOracleResponsePacketData(
	clientID string, requestID RequestID, ansCount uint64, requestTime int64,
	resolveTime int64, resolveStatus ResolveStatus, result []byte,
) OracleResponsePacketData {
	return OracleResponsePacketData{
		ClientID:      clientID,
		RequestID:     requestID,
		AnsCount:      ansCount,
		RequestTime:   requestTime,
		ResolveTime:   resolveTime,
		ResolveStatus: resolveStatus,
		Result:        result,
	}
}

// GetBytes returns the bytes representation of this oracle response packet data.
func (p OracleResponsePacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}
