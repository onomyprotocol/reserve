package types

import (
	// "context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint: staticcheck
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
)

// ChannelKeeper defines the expected IBC channel keeper.
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(
		ctx sdk.Context,
		chanCap *capabilitytypes.Capability,
		sourcePort string,
		sourceChannel string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
		data []byte,
	) (sequence uint64, err error)
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
}

// PortKeeper defines the expected IBC port keeper.
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

// ScopedKeeper defines the expected IBC scoped keeper.
// type ScopedKeeper interface {
// 	GetCapability(ctx context.Context, name string) (*capabilitytypes.Capability, bool)
// 	AuthenticateCapability(ctx context.Context, cap *capabilitytypes.Capability, name string) bool
// 	ClaimCapability(ctx context.Context, cap *capabilitytypes.Capability, name string) error
// }
