package keeper

import (
	"time"
	"strconv"
	sdk "github.com/cosmos/cosmos-sdk/types"
	prefix "cosmossdk.io/store/prefix"
	runtime "github.com/cosmos/cosmos-sdk/runtime"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	errorsmod "cosmossdk.io/errors"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// SetBandParams sets the Band params in the state
func (k Keeper) SetBandParams(ctx sdk.Context, bandParams types.BandParams)  error{
	bz := k.cdc.MustMarshal(&bandParams)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.BandParamsKey, bz)
}

// GetBandParams gets the Band params stored in the state
func (k Keeper) GetBandParams(ctx sdk.Context) types.BandParams {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.BandParamsKey)

	if err != nil {
		return types.DefaultGenesis().BandParams
	}

	if bz == nil {
		return types.DefaultGenesis().BandParams
	}

	var bandParams types.BandParams
	k.cdc.MustUnmarshal(bz, &bandParams)
	return bandParams
}

// SetBandCallData sets the Band IBC oracle request call data
func (k Keeper) SetBandCallDataRecord(ctx sdk.Context, record *types.CalldataRecord) error {
	bz := k.cdc.MustMarshal(record)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.GetBandCallDataRecordKey(record.ClientId), bz)
}

// DeleteBandCallDataRecord deletes the Band IBC oracle request call data
func (k Keeper) DeleteBandCallDataRecord(ctx sdk.Context, clientID uint64) error{
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(types.GetBandCallDataRecordKey(clientID))
}

// GetAllBandCalldataRecords gets all Band oracle request CallData for each clientID
func (k Keeper) GetAllBandCalldataRecords(ctx sdk.Context) []*types.CalldataRecord {
	calldataRecords := make([]*types.CalldataRecord, 0)
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandIBCCalldataStore := prefix.NewStore(kvStore, types.BandCallDataRecordKey)

	iterator := bandIBCCalldataStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var calldataRecord types.CalldataRecord
		k.cdc.MustUnmarshal(iterator.Value(), &calldataRecord)
		calldataRecords = append(calldataRecords, &calldataRecord)
	}

	return calldataRecords
}

// GetBandCallDataRecord gets the Band oracle request CallDataRecord for a given clientID
func (k Keeper) GetBandCallDataRecord(ctx sdk.Context, clientID uint64) *types.CalldataRecord {
	var callDataRecord types.CalldataRecord
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.GetBandCallDataRecordKey(clientID))
	if err != nil {
		return nil
	}
	if bz == nil {
		return nil
	}
	k.cdc.MustUnmarshal(bz, &callDataRecord)
	return &callDataRecord
}

// GetBandLatestClientID returns the latest clientID of Band oracle request packet data.
func (k Keeper) GetBandLatestClientID(ctx sdk.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.LatestClientIDKey)
	if err != nil {
		return 0
	}
	if bz == nil {
		return 0
	}
	clientID := sdk.BigEndianToUint64(bz)
	return clientID
}

// SetBandLatestClientID sets the latest clientID of Band oracle request packet data.
func (k Keeper) SetBandLatestClientID(ctx sdk.Context, clientID uint64) error {
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.LatestClientIDKey, sdk.Uint64ToBigEndian(clientID))
}

// SetBandOracleRequest sets the Band IBC oracle request data
func (k Keeper) SetBandOracleRequest(ctx sdk.Context, req types.BandOracleRequest) error {
	bz := k.cdc.MustMarshal(&req)
	store := k.storeService.OpenKVStore(ctx)
	return store.Set(types.GetBandOracleRequestIDKey(req.RequestId), bz)
}

// GetBandOracleRequest gets the Band oracle request data
func (k Keeper) GetBandOracleRequest(ctx sdk.Context, requestID uint64) *types.BandOracleRequest {
	var bandOracleRequest types.BandOracleRequest
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.GetBandOracleRequestIDKey(requestID))
	if err != nil {
		return nil
	}
	if bz == nil {
		return nil
	}

	k.cdc.MustUnmarshal(bz, &bandOracleRequest)
	return &bandOracleRequest
}

// DeleteBandOracleRequest deletes the Band oracle request call data
func (k Keeper) DeleteBandOracleRequest(ctx sdk.Context, requestID uint64) error{
	store := k.storeService.OpenKVStore(ctx)
	return store.Delete(types.GetBandOracleRequestIDKey(requestID))
}

// GetAllBandOracleRequests gets all Band oracle requests for each requestID
func (k Keeper) GetAllBandOracleRequests(ctx sdk.Context) []*types.BandOracleRequest {
	bandIBCOracleRequests := make([]*types.BandOracleRequest, 0)
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bandIBCOracleRequestStore := prefix.NewStore(kvStore, types.BandOracleRequestIDKey)

	iterator := bandIBCOracleRequestStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bandIBCOracleRequest types.BandOracleRequest
		k.cdc.MustUnmarshal(iterator.Value(), &bandIBCOracleRequest)
		bandIBCOracleRequests = append(bandIBCOracleRequests, &bandIBCOracleRequest)
	}

	return bandIBCOracleRequests
}

// RequestBandOraclePrices creates and sends an IBC packet to fetch band oracle price feed data through IBC.
func (k *Keeper) RequestBandOraclePrices(
	ctx sdk.Context,
	req *types.BandOracleRequest,
) (err error) {
	bandIBCParams := k.GetBandParams(ctx)
	sourcePortID := bandIBCParams.IbcPortId
	sourceChannel := bandIBCParams.IbcSourceChannel

	calldata := req.GetCalldata(types.IsLegacySchemeOracleScript(req.OracleScriptId, bandIBCParams))

	sourceChannelEnd, found := k.ibcKeeperFn().ChannelKeeper.GetChannel(ctx, sourcePortID, sourceChannel)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unknown channel %s port %s", sourceChannel, sourcePortID)
	}

	// retrieve the dynamic capability for this channel
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePortID, sourceChannel))
	if !ok {
		return errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	destinationPort := sourceChannelEnd.Counterparty.PortId
	destinationChannel := sourceChannelEnd.Counterparty.ChannelId
	sequence, found := k.ibcKeeperFn().ChannelKeeper.GetNextSequenceSend(ctx, sourcePortID, sourceChannel)

	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unknown sequence number for channel %s port %s", sourceChannel, sourcePortID)
	}

	clientID := k.GetBandLatestClientID(ctx) + 1
	packetData := types.NewOracleRequestPacketData(strconv.Itoa(int(clientID)), calldata, req)

	// Creating custom oracle packet data
	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePortID,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 0),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)), // Arbitrarily high timeout for now
	)

	// Send packet to IBC, authenticating with channelCap
	_, err = k.ibcKeeperFn().ChannelKeeper.SendPacket(
		ctx,
		channelCap,
		packet.SourcePort,
		packet.SourceChannel,
		packet.TimeoutHeight,
		packet.TimeoutTimestamp,
		packet.Data,
	)
	if err != nil {
		return err
	}

	// Persist the sequence number and OracleRequest CallData. CallData contains list of symbols.
	// This is used to map the prices/rates with the symbols upon receiving oracle response from Band IBC.
	k.SetBandCallDataRecord(ctx, &types.CalldataRecord{
		ClientId: clientID,
		Calldata: calldata,
	})

	k.SetBandLatestClientID(ctx, clientID)

	return
}
