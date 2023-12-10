package keeper

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/onomyprotocol/reserve/x/portal/types"
)

// TransmitSubscribeRatePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitSubscribeRatePacket(
	ctx sdk.Context,
	packetData types.SubscribeRatePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvSubscribeRatePacket processes packet reception
func (k Keeper) OnRecvSubscribeRatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.SubscribeRatePacketData) (packetAck types.SubscribeRatePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}

// OnAcknowledgementSubscribeRatePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementSubscribeRatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.SubscribeRatePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.SubscribeRatePacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutSubscribeRatePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutSubscribeRatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.SubscribeRatePacketData) error {

	// TODO: packet timeout logic

	return nil
}

func (k msgServer) SubscribeRate(ctx sdk.Context, denom string, port string) error {

	creator := k.accountKeeper.GetModuleAddress(types.ModuleName).String()
	srcPort := port
	srcChannel := "1"

	RelativePacketTimeoutTimestamp := uint64((time.Duration(1) * time.Minute).Nanoseconds())

	// Get the relative timeout timestamp
	timeoutTimestamp := uint64((ctx.BlockHeader().Time.UnixNano())) + RelativePacketTimeoutTimestamp

	msg := types.NewMsgSendSubscribeRate(
		creator,
		srcPort,
		srcChannel,
		timeoutTimestamp,
		denom,
	)

	k.SendSubscribeRate(ctx, msg)
}
