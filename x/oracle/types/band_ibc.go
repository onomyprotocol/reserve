package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const BandPriceMultiplier uint64 = 1000000000 // 1e9
type RequestID int64

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

// GetBytes is a helper for serialising
func (p OracleRequestPacketData) GetBytes() []byte {
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	return sdk.MustSortJSON(cdc.MustMarshalJSON(&p))
}

// GetCalldata gets the Band IBC request call data based on the symbols and multiplier.
func (r *BandOracleRequest) GetCalldata(legacyScheme bool) []byte {
	if legacyScheme {
		return MustEncode(Input{
			Symbols:    r.Symbols,
			Multiplier: BandPriceMultiplier,
		})
	}

	return MustEncode(SymbolInput{
		Symbols:            r.Symbols,
		MinimumSourceCount: uint8(r.MinSourceCount),
	})
}

func IsLegacySchemeOracleScript(scriptID int64, params BandParams) bool {
	for _, id := range params.LegacyOracleIds {
		if id == scriptID {
			return true
		}
	}

	return false
}

type SymbolInput struct {
	Symbols            []string `json:"symbols"`
	MinimumSourceCount uint8    `json:"minimum_source_count"`
}

type SymbolOutput struct {
	Responses []Response `json:"responses"`
}

type Response struct {
	Symbol       string `json:"symbol"`
	ResponseCode uint8  `json:"response_code"`
	Rate         uint64 `json:"rate"`
}

type Input struct {
	Symbols    []string `json:"symbols"`
	Multiplier uint64   `json:"multiplier"`
}

type Output struct {
	Pxs []uint64 `json:"pxs"`
}

type Price struct {
	Symbol      string    `json:"symbol"`
	Multiplier  uint64    `json:"multiplier"`
	Px          uint64    `json:"px"`
	RequestID   RequestID `json:"request_id"`
	ResolveTime int64     `json:"resolve_time"`
}

func NewPrice(symbol string, multiplier, px uint64, reqID RequestID, resolveTime int64) Price {
	return Price{
		Symbol:      symbol,
		Multiplier:  multiplier,
		Px:          px,
		RequestID:   reqID,
		ResolveTime: resolveTime,
	}
}
