package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	"github.com/onomyprotocol/reserve/x/oracle/utils"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k Keeper) RequestBandRates(goCtx context.Context, msg *types.MsgRequestBandRates) (*types.MsgRequestBandRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bandOracleRequest := k.GetBandOracleRequest(ctx, msg.RequestId)
	if bandOracleRequest == nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidBandRequest, "Band oracle request not found!")
	}

	if err := k.RequestBandOraclePrices(ctx, bandOracleRequest); err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	return &types.MsgRequestBandRatesResponse{}, nil
}

func (k Keeper) UpdateBandParams(goCtx context.Context, msg *types.MsgUpdateBandParams) (*types.MsgUpdateBandParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.validateAuthority(msg.Authority); err != nil {
		return nil, err
	}

	if err := k.validateUpdateBandParams(msg); err != nil {
		return nil, err
	}
	k.SetPort(ctx, msg.BandParams.IbcPortId)
	// Only try to bind to port if it is not already bound, since we may already own port capability
	if !k.IsBound(ctx, msg.BandParams.IbcPortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, msg.BandParams.IbcPortId)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrBandPortBind, err.Error())
		}
	}

	k.SetBandParams(ctx, msg.BandParams)
	return nil, nil
}

func (k Keeper) UpdateBandOracleRequest(goCtx context.Context, msg *types.MsgUpdateBandOracleRequest) (*types.MsgUpdateBandOracleRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	if err := k.validateAuthority(msg.Authority); err != nil {
		return nil, err
	}

	if err := k.validateUpdateBandOracleRequest(msg); err != nil {
		return nil, err
	}

	request := k.GetBandOracleRequest(ctx, msg.UpdateOracleRequest.RequestId)
	if request == nil {
		return nil, errorsmod.Wrapf(types.ErrBandRequestNotFound, "cannot update requestID %T", msg.UpdateOracleRequest.RequestId)
	}

	if msg.UpdateOracleRequest.OracleScriptId > 0 {
		request.OracleScriptId = msg.UpdateOracleRequest.OracleScriptId
	}

	if len(msg.UpdateOracleRequest.Symbols) > 0 {
		request.Symbols = msg.UpdateOracleRequest.Symbols
	}

	if msg.UpdateOracleRequest.MinCount > 0 {
		request.MinCount = msg.UpdateOracleRequest.MinCount
	}

	if msg.UpdateOracleRequest.AskCount > 0 {
		request.AskCount = msg.UpdateOracleRequest.AskCount
	}

	if msg.UpdateOracleRequest.FeeLimit != nil {
		request.FeeLimit = msg.UpdateOracleRequest.FeeLimit
	}

	if msg.UpdateOracleRequest.PrepareGas > 0 {
		request.PrepareGas = msg.UpdateOracleRequest.PrepareGas
	}

	if msg.UpdateOracleRequest.ExecuteGas > 0 {
		request.ExecuteGas = msg.UpdateOracleRequest.ExecuteGas
	}

	if msg.UpdateOracleRequest.MinSourceCount > 0 {
		request.MinSourceCount = msg.UpdateOracleRequest.MinSourceCount
	}

	k.SetBandOracleRequest(ctx, *request)

	return nil, nil
}

func (k Keeper) DeleteBandOracleRequests(goCtx context.Context, msg *types.MsgDeleteBandOracleRequests) (*types.MsgDeleteBandOracleRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	if err := k.validateAuthority(msg.Authority); err != nil {
		return nil, err
	}

	if len(msg.DeleteRequestIds) == 0 {
		return nil, types.ErrInvalidBandDeleteRequest
	}

	for _, requestID := range msg.DeleteRequestIds {
		k.DeleteBandOracleRequest(ctx, requestID)
	}

	return nil, nil
}

// validateUpdateBandParams returns validate for update band params.
func (k *Keeper) validateUpdateBandParams(msg *types.MsgUpdateBandParams) error {
	if msg.BandParams.IbcRequestInterval == 0 {
		return types.ErrBadRequestInterval
	}

	if msg.BandParams.IbcSourceChannel == "" {
		return errorsmod.Wrap(types.ErrInvalidSourceChannel, "UpdateBandParamsProposal: IBC Source Channel must not be empty.")
	}
	if msg.BandParams.IbcVersion == "" {
		return errorsmod.Wrap(types.ErrInvalidVersion, "UpdateBandParamsProposal: IBC Version must not be empty.")
	}

	return nil
}

// validateUpdateBandOracleRequest returns validate for update band oracle request.
func (k *Keeper) validateUpdateBandOracleRequest(msg *types.MsgUpdateBandOracleRequest) error {
	if msg.UpdateOracleRequest == nil {
		return types.ErrInvalidBandUpdateRequest
	}

	if msg.UpdateOracleRequest != nil && len(msg.UpdateOracleRequest.Symbols) > 0 {
		callData, err := utils.Encode(types.SymbolInput{
			Symbols:            msg.UpdateOracleRequest.Symbols,
			MinimumSourceCount: uint8(msg.UpdateOracleRequest.MinCount),
		})

		if err != nil {
			return err
		}

		if len(callData) > types.MaxDataSize {
			return errorsmod.Wrapf(types.ErrTooLargeCalldata, "got: %d, maximum: %d", len(callData), types.MaxDataSize)
		}
	}

	if msg.UpdateOracleRequest != nil && msg.UpdateOracleRequest.AskCount > 0 && msg.UpdateOracleRequest.MinCount > 0 && msg.UpdateOracleRequest.AskCount < msg.UpdateOracleRequest.MinCount {
		return errorsmod.Wrapf(types.ErrInvalidAskCount, "UpdateBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", msg.UpdateOracleRequest.AskCount, msg.UpdateOracleRequest.MinCount)
	}

	if msg.UpdateOracleRequest != nil && msg.UpdateOracleRequest.FeeLimit != nil && !msg.UpdateOracleRequest.FeeLimit.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "UpdateBandOracleRequestProposal: Invalid Fee Limit (%s)", msg.UpdateOracleRequest.GetFeeLimit().String())
	}

	if msg.UpdateOracleRequest != nil && msg.UpdateOracleRequest.PrepareGas <= 0 && msg.UpdateOracleRequest.ExecuteGas > 0 {
		return errorsmod.Wrapf(types.ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Invalid Prepare Gas (%d)", msg.UpdateOracleRequest.PrepareGas)
	}

	if msg.UpdateOracleRequest != nil && msg.UpdateOracleRequest.ExecuteGas <= 0 {
		return errorsmod.Wrapf(types.ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Invalid Execute Gas (%d)", msg.UpdateOracleRequest.ExecuteGas)
	}

	return nil
}

func (k *Keeper) validateAuthority(authority string) error {
	if _, err := k.authKeeper.AddressCodec().StringToBytes(authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid authority address: %s", err)
	}

	if k.authority != authority {
		return errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, authority)
	}

	return nil
}