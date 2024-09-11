// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: cosmos/vaults/tx.proto

package vaults

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Msg_ActiveCollateral_FullMethodName = "/cosmos.vaults.Msg/ActiveCollateral"
	Msg_CreateVault_FullMethodName      = "/cosmos.vaults.Msg/CreateVault"
	Msg_Deposit_FullMethodName          = "/cosmos.vaults.Msg/Deposit"
	Msg_Withdraw_FullMethodName         = "/cosmos.vaults.Msg/Withdraw"
	Msg_Mint_FullMethodName             = "/cosmos.vaults.Msg/Mint"
	Msg_Repay_FullMethodName            = "/cosmos.vaults.Msg/Repay"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Msg defines the vaults Msg service.
type MsgClient interface {
	// ActiveCollateral defines a method for enable a collateral asset
	ActiveCollateral(ctx context.Context, in *MsgActiveCollateral, opts ...grpc.CallOption) (*MsgActiveCollateralResponse, error)
	// CreateVault defines a method for creating a new vault and mint token
	CreateVault(ctx context.Context, in *MsgCreateVault, opts ...grpc.CallOption) (*MsgCreateVaultResponse, error)
	// Deposit defines a method for depositing collateral assets to vault
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing collateral assets out of the vault
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
	// Mint defines a method for minting more tokens
	Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error)
	// Repay defines a method for reducing debt by burning tokens
	Repay(ctx context.Context, in *MsgRepay, opts ...grpc.CallOption) (*MsgMintResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ActiveCollateral(ctx context.Context, in *MsgActiveCollateral, opts ...grpc.CallOption) (*MsgActiveCollateralResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgActiveCollateralResponse)
	err := c.cc.Invoke(ctx, Msg_ActiveCollateral_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateVault(ctx context.Context, in *MsgCreateVault, opts ...grpc.CallOption) (*MsgCreateVaultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgCreateVaultResponse)
	err := c.cc.Invoke(ctx, Msg_CreateVault_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, Msg_Deposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, Msg_Withdraw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Mint(ctx context.Context, in *MsgMint, opts ...grpc.CallOption) (*MsgMintResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgMintResponse)
	err := c.cc.Invoke(ctx, Msg_Mint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Repay(ctx context.Context, in *MsgRepay, opts ...grpc.CallOption) (*MsgMintResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgMintResponse)
	err := c.cc.Invoke(ctx, Msg_Repay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
//
// Msg defines the vaults Msg service.
type MsgServer interface {
	// ActiveCollateral defines a method for enable a collateral asset
	ActiveCollateral(context.Context, *MsgActiveCollateral) (*MsgActiveCollateralResponse, error)
	// CreateVault defines a method for creating a new vault and mint token
	CreateVault(context.Context, *MsgCreateVault) (*MsgCreateVaultResponse, error)
	// Deposit defines a method for depositing collateral assets to vault
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	// Withdraw defines a method for withdrawing collateral assets out of the vault
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
	// Mint defines a method for minting more tokens
	Mint(context.Context, *MsgMint) (*MsgMintResponse, error)
	// Repay defines a method for reducing debt by burning tokens
	Repay(context.Context, *MsgRepay) (*MsgMintResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) ActiveCollateral(context.Context, *MsgActiveCollateral) (*MsgActiveCollateralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveCollateral not implemented")
}
func (UnimplementedMsgServer) CreateVault(context.Context, *MsgCreateVault) (*MsgCreateVaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVault not implemented")
}
func (UnimplementedMsgServer) Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedMsgServer) Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedMsgServer) Mint(context.Context, *MsgMint) (*MsgMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (UnimplementedMsgServer) Repay(context.Context, *MsgRepay) (*MsgMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Repay not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ActiveCollateral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgActiveCollateral)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ActiveCollateral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ActiveCollateral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ActiveCollateral(ctx, req.(*MsgActiveCollateral))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateVault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVault)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateVault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVault(ctx, req.(*MsgCreateVault))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Deposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Mint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Repay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRepay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Repay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Repay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Repay(ctx, req.(*MsgRepay))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vaults.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActiveCollateral",
			Handler:    _Msg_ActiveCollateral_Handler,
		},
		{
			MethodName: "CreateVault",
			Handler:    _Msg_CreateVault_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
		{
			MethodName: "Mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "Repay",
			Handler:    _Msg_Repay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/vaults/tx.proto",
}
