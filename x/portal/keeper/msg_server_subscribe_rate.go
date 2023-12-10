package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/onomyprotocol/reserve/x/portal/types"
)

func (k msgServer) SendSubscribeRate(goCtx context.Context, msg *types.MsgSendSubscribeRate) (*types.MsgSendSubscribeRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.SubscribeRatePacketData

	packet.Denom = msg.Denom

	// Transmit the packet
	err := k.TransmitSubscribeRatePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendSubscribeRateResponse{}, nil
}
