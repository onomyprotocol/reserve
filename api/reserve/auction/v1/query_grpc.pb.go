// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: reserve/auction/v1/query.proto

package auctionv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName                = "/reserve.auction.v1.Query/Params"
	Query_QueryAllAuction_FullMethodName       = "/reserve.auction.v1.Query/QueryAllAuction"
	Query_QueryAllBids_FullMethodName          = "/reserve.auction.v1.Query/QueryAllBids"
	Query_QueryAllBidderBids_FullMethodName    = "/reserve.auction.v1.Query/QueryAllBidderBids"
	Query_QueryAllBidsByAddress_FullMethodName = "/reserve.auction.v1.Query/QueryAllBidsByAddress"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	QueryAllAuction(ctx context.Context, in *QueryAllAuctionRequest, opts ...grpc.CallOption) (*QueryAllAuctionResponse, error)
	QueryAllBids(ctx context.Context, in *QueryAllBidsRequest, opts ...grpc.CallOption) (*QueryAllBidsResponse, error)
	QueryAllBidderBids(ctx context.Context, in *QueryAllBidderBidsRequest, opts ...grpc.CallOption) (*QueryAllBidderBidsResponse, error)
	QueryAllBidsByAddress(ctx context.Context, in *QueryAllBidsByAddressRequest, opts ...grpc.CallOption) (*QueryAllBidsByAddressResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryAllAuction(ctx context.Context, in *QueryAllAuctionRequest, opts ...grpc.CallOption) (*QueryAllAuctionResponse, error) {
	out := new(QueryAllAuctionResponse)
	err := c.cc.Invoke(ctx, Query_QueryAllAuction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryAllBids(ctx context.Context, in *QueryAllBidsRequest, opts ...grpc.CallOption) (*QueryAllBidsResponse, error) {
	out := new(QueryAllBidsResponse)
	err := c.cc.Invoke(ctx, Query_QueryAllBids_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryAllBidderBids(ctx context.Context, in *QueryAllBidderBidsRequest, opts ...grpc.CallOption) (*QueryAllBidderBidsResponse, error) {
	out := new(QueryAllBidderBidsResponse)
	err := c.cc.Invoke(ctx, Query_QueryAllBidderBids_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryAllBidsByAddress(ctx context.Context, in *QueryAllBidsByAddressRequest, opts ...grpc.CallOption) (*QueryAllBidsByAddressResponse, error) {
	out := new(QueryAllBidsByAddressResponse)
	err := c.cc.Invoke(ctx, Query_QueryAllBidsByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	QueryAllAuction(context.Context, *QueryAllAuctionRequest) (*QueryAllAuctionResponse, error)
	QueryAllBids(context.Context, *QueryAllBidsRequest) (*QueryAllBidsResponse, error)
	QueryAllBidderBids(context.Context, *QueryAllBidderBidsRequest) (*QueryAllBidderBidsResponse, error)
	QueryAllBidsByAddress(context.Context, *QueryAllBidsByAddressRequest) (*QueryAllBidsByAddressResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) QueryAllAuction(context.Context, *QueryAllAuctionRequest) (*QueryAllAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllAuction not implemented")
}
func (UnimplementedQueryServer) QueryAllBids(context.Context, *QueryAllBidsRequest) (*QueryAllBidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllBids not implemented")
}
func (UnimplementedQueryServer) QueryAllBidderBids(context.Context, *QueryAllBidderBidsRequest) (*QueryAllBidderBidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllBidderBids not implemented")
}
func (UnimplementedQueryServer) QueryAllBidsByAddress(context.Context, *QueryAllBidsByAddressRequest) (*QueryAllBidsByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllBidsByAddress not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryAllAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryAllAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryAllAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryAllAuction(ctx, req.(*QueryAllAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryAllBids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryAllBids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryAllBids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryAllBids(ctx, req.(*QueryAllBidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryAllBidderBids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBidderBidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryAllBidderBids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryAllBidderBids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryAllBidderBids(ctx, req.(*QueryAllBidderBidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryAllBidsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBidsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryAllBidsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QueryAllBidsByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryAllBidsByAddress(ctx, req.(*QueryAllBidsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reserve.auction.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryAllAuction",
			Handler:    _Query_QueryAllAuction_Handler,
		},
		{
			MethodName: "QueryAllBids",
			Handler:    _Query_QueryAllBids_Handler,
		},
		{
			MethodName: "QueryAllBidderBids",
			Handler:    _Query_QueryAllBidderBids_Handler,
		},
		{
			MethodName: "QueryAllBidsByAddress",
			Handler:    _Query_QueryAllBidsByAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reserve/auction/v1/query.proto",
}
