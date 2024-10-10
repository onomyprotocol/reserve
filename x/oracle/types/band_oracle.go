package types

import (
	"fmt"

	math "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/onomyprotocol/reserve/x/oracle/utils"
)

const (
	BandPriceMultiplier uint64 = 1000000000 // 1e9
	MaxDataSize                = 256        // 256B
	QuoteUSD                   = "USD"
)

var NilPriceState = BandPriceState{}

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

// GetBytes returns the bytes representation of this oracle response packet data.
func (p OracleResponsePacketData) GetBytes() []byte {
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	return sdk.MustSortJSON(cdc.MustMarshalJSON(&p))
}

// GetCalldata gets the Band IBC request call data based on the symbols and multiplier.
func (r *BandOracleRequest) GetCalldata(legacyScheme bool) []byte {
	if legacyScheme {
		return utils.MustEncode(Input{
			Symbols:    r.Symbols,
			Multiplier: BandPriceMultiplier,
		})
	}

	return utils.MustEncode(SymbolInput{
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

// CheckPriceFeedThreshold returns true if the newPrice has changed beyond 100x or less than 1% of the last price
func CheckPriceFeedThreshold(lastPrice, newPrice math.LegacyDec) bool {
	return newPrice.GT(lastPrice.Mul(math.LegacyNewDec(100))) || newPrice.LT(lastPrice.Quo(math.LegacyNewDec(100)))
}

func DecodeOracleInput(data []byte) (OracleInput, error) {
	var (
		legacyInput LegacyBandInput
		newInput    BandInput
		err         error
	)

	if err = utils.Decode(data, &legacyInput); err == nil {
		return legacyInput, nil
	}

	if err = utils.Decode(data, &newInput); err == nil {
		return newInput, nil
	}

	return nil, fmt.Errorf("failed to decode oracle input: %w", err)
}

func DecodeOracleOutput(data []byte) (OracleOutput, error) {
	var (
		legacyOutput LegacyBandOutput
		newOutput    BandOutput
		err          error
	)

	if err = utils.Decode(data, &legacyOutput); err == nil {
		return legacyOutput, nil
	}

	if err = utils.Decode(data, &newOutput); err == nil {
		return newOutput, nil
	}

	return nil, fmt.Errorf("failed to decode oracle output: %w", err)
}

// it is assumed that the id of a symbol
// within OracleInput exists within OracleOutput

type OracleInput interface {
	PriceSymbols() []string
	PriceMultiplier() uint64
}

type (
	LegacyBandInput Input
	BandInput       SymbolInput
)

func (in LegacyBandInput) PriceSymbols() []string {
	return in.Symbols
}

func (in LegacyBandInput) PriceMultiplier() uint64 {
	return in.Multiplier
}

func (in BandInput) PriceSymbols() []string {
	return in.Symbols
}

func (in BandInput) PriceMultiplier() uint64 {
	return BandPriceMultiplier
}

type OracleOutput interface {
	Rate(id int) uint64
	Valid(id int) bool
}

type (
	LegacyBandOutput Output
	BandOutput       SymbolOutput
)

func (out LegacyBandOutput) Rate(id int) uint64 {
	return out.Pxs[id]
}

func (out LegacyBandOutput) Valid(id int) bool {
	return true
}

func (out BandOutput) Rate(id int) uint64 {
	return out.Responses[id].Rate
}

func (out BandOutput) Valid(id int) bool {
	return out.Responses[id].ResponseCode == 0
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

func NewPriceState(price math.LegacyDec, timestamp int64) *PriceState {
	return &PriceState{
		Price:     price,
		Timestamp: timestamp,
	}
}

func (p *PriceState) UpdatePrice(price math.LegacyDec, timestamp int64) {
	p.Timestamp = timestamp
	p.Price = price
}
