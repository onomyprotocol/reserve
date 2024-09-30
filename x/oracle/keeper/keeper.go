package keeper

import (
	"fmt"
	"context"
	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	"github.com/onomyprotocol/reserve/x/oracle/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		ibcKeeperFn func() *ibckeeper.Keeper
		capabilityScopedFn func(string) capabilitykeeper.ScopedKeeper
		// scopedKeeper capabilitykeeper.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	ibcKeeperFn func() *ibckeeper.Keeper,
	capabilityScopedFn func(string) capabilitykeeper.ScopedKeeper,
	// scopedKeeper capabilitykeeper.ScopedKeeper,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		ibcKeeperFn:  ibcKeeperFn,
		capabilityScopedFn: capabilityScopedFn,
		// scopedKeeper: scopedKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", types.ModuleName)
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k *Keeper) ChanCloseInit(ctx context.Context, portID, channelID string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.ScopedKeeper().GetCapability(sdkCtx, capName)
	if !ok {
		return errorsmod.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.ibcKeeperFn().ChannelKeeper.ChanCloseInit(sdkCtx, portID, channelID, chanCap)
}

// ShouldBound checks if the IBC app module can be bound to the desired port
func (k *Keeper) ShouldBound(ctx context.Context, portID string) bool {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	scopedKeeper := k.ScopedKeeper()
	if scopedKeeper == nil {
		return false
	}
	_, ok := scopedKeeper.GetCapability(sdkCtx, host.PortPath(portID))
	return !ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k *Keeper) BindPort(ctx context.Context, portID string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	cap := k.ibcKeeperFn().PortKeeper.BindPort(sdkCtx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k *Keeper) GetPort(ctx context.Context) string {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k *Keeper) SetPort(ctx context.Context, portID string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k *Keeper) AuthenticateCapability(ctx context.Context, cap *capabilitytypes.Capability, name string) bool {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return k.ScopedKeeper().AuthenticateCapability(sdkCtx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k *Keeper) ClaimCapability(ctx context.Context, cap *capabilitytypes.Capability, name string) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return k.ScopedKeeper().ClaimCapability(sdkCtx, cap, name)
}

// ScopedKeeper returns the ScopedKeeper
func (k *Keeper) ScopedKeeper() exported.ScopedKeeper {
	return k.capabilityScopedFn(types.ModuleName)
}
