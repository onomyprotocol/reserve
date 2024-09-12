// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/psm/v1/tx.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the module parameters to update.
	//
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgSwapToIST struct {
	Address string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    *types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (m *MsgSwapToIST) Reset()         { *m = MsgSwapToIST{} }
func (m *MsgSwapToIST) String() string { return proto.CompactTextString(m) }
func (*MsgSwapToIST) ProtoMessage()    {}
func (*MsgSwapToIST) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{2}
}
func (m *MsgSwapToIST) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapToIST) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapToIST.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapToIST) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapToIST.Merge(m, src)
}
func (m *MsgSwapToIST) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapToIST) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapToIST.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapToIST proto.InternalMessageInfo

func (m *MsgSwapToIST) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSwapToIST) GetCoin() *types.Coin {
	if m != nil {
		return m.Coin
	}
	return nil
}

type MsgSwapToISTResponse struct {
}

func (m *MsgSwapToISTResponse) Reset()         { *m = MsgSwapToISTResponse{} }
func (m *MsgSwapToISTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapToISTResponse) ProtoMessage()    {}
func (*MsgSwapToISTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{3}
}
func (m *MsgSwapToISTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapToISTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapToISTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapToISTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapToISTResponse.Merge(m, src)
}
func (m *MsgSwapToISTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapToISTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapToISTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapToISTResponse proto.InternalMessageInfo

type MsgSwapToStablecoin struct {
	Address string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ToDenom string                `protobuf:"bytes,2,opt,name=to_denom,json=toDenom,proto3" json:"to_denom,omitempty"`
	Amount  cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *MsgSwapToStablecoin) Reset()         { *m = MsgSwapToStablecoin{} }
func (m *MsgSwapToStablecoin) String() string { return proto.CompactTextString(m) }
func (*MsgSwapToStablecoin) ProtoMessage()    {}
func (*MsgSwapToStablecoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{4}
}
func (m *MsgSwapToStablecoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapToStablecoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapToStablecoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapToStablecoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapToStablecoin.Merge(m, src)
}
func (m *MsgSwapToStablecoin) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapToStablecoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapToStablecoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapToStablecoin proto.InternalMessageInfo

func (m *MsgSwapToStablecoin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgSwapToStablecoin) GetToDenom() string {
	if m != nil {
		return m.ToDenom
	}
	return ""
}

type MsgSwapToStablecoinResponse struct {
}

func (m *MsgSwapToStablecoinResponse) Reset()         { *m = MsgSwapToStablecoinResponse{} }
func (m *MsgSwapToStablecoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapToStablecoinResponse) ProtoMessage()    {}
func (*MsgSwapToStablecoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ff2d5421e71e2a, []int{5}
}
func (m *MsgSwapToStablecoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapToStablecoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapToStablecoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapToStablecoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapToStablecoinResponse.Merge(m, src)
}
func (m *MsgSwapToStablecoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapToStablecoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapToStablecoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapToStablecoinResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "reserve.psm.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "reserve.psm.v1.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgSwapToIST)(nil), "reserve.psm.v1.MsgSwapToIST")
	proto.RegisterType((*MsgSwapToISTResponse)(nil), "reserve.psm.v1.MsgSwapToISTResponse")
	proto.RegisterType((*MsgSwapToStablecoin)(nil), "reserve.psm.v1.MsgSwapToStablecoin")
	proto.RegisterType((*MsgSwapToStablecoinResponse)(nil), "reserve.psm.v1.MsgSwapToStablecoinResponse")
}

func init() { proto.RegisterFile("reserve/psm/v1/tx.proto", fileDescriptor_d0ff2d5421e71e2a) }

var fileDescriptor_d0ff2d5421e71e2a = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0x12, 0x41,
	0x18, 0x65, 0x5b, 0xa5, 0x32, 0x12, 0x7f, 0xac, 0x58, 0x60, 0x6b, 0x17, 0x82, 0x26, 0x12, 0x94,
	0x99, 0x52, 0x13, 0x13, 0x7b, 0x13, 0x8d, 0x91, 0x03, 0xd1, 0x2c, 0x35, 0x31, 0x5e, 0x9a, 0x81,
	0x9d, 0x2c, 0x1b, 0x3b, 0x3b, 0x9b, 0x9d, 0x01, 0xcb, 0xcd, 0x78, 0xf4, 0xe4, 0x3f, 0xe0, 0xc5,
	0x93, 0x47, 0x0e, 0x8d, 0x7f, 0x43, 0x8f, 0x4d, 0x4f, 0xc6, 0x43, 0x63, 0xe0, 0xc0, 0xbf, 0x61,
	0x76, 0x67, 0x16, 0x64, 0xad, 0xd5, 0x0b, 0xe1, 0x9b, 0xf7, 0xbd, 0xf9, 0xde, 0x7b, 0xdf, 0x2c,
	0xc8, 0x07, 0x84, 0x93, 0x60, 0x48, 0x90, 0xcf, 0x29, 0x1a, 0x36, 0x90, 0x38, 0x80, 0x7e, 0xc0,
	0x04, 0xd3, 0xaf, 0x28, 0x00, 0xfa, 0x9c, 0xc2, 0x61, 0xc3, 0xb8, 0x8e, 0xa9, 0xeb, 0x31, 0x14,
	0xfd, 0xca, 0x16, 0x23, 0xdf, 0x63, 0x9c, 0x32, 0x8e, 0x28, 0x77, 0x42, 0x2a, 0xe5, 0x8e, 0x02,
	0x8a, 0x12, 0xd8, 0x8b, 0x2a, 0x24, 0x0b, 0x05, 0xe5, 0x1c, 0xe6, 0x30, 0x79, 0x1e, 0xfe, 0x53,
	0xa7, 0x1b, 0x09, 0x15, 0x3e, 0x0e, 0x30, 0x8d, 0x29, 0xa6, 0x1a, 0xd3, 0xc5, 0x9c, 0xa0, 0x61,
	0xa3, 0x4b, 0x04, 0x6e, 0xa0, 0x1e, 0x73, 0x3d, 0x89, 0x57, 0xbe, 0x69, 0xe0, 0x6a, 0x9b, 0x3b,
	0xaf, 0x7c, 0x1b, 0x0b, 0xf2, 0x32, 0x62, 0xea, 0x0f, 0x41, 0x06, 0x0f, 0x44, 0x9f, 0x05, 0xae,
	0x18, 0x15, 0xb4, 0xb2, 0x56, 0xcd, 0x34, 0x0b, 0x27, 0x87, 0xf5, 0x9c, 0xd2, 0xf2, 0xd8, 0xb6,
	0x03, 0xc2, 0x79, 0x47, 0x04, 0xae, 0xe7, 0x58, 0x8b, 0x56, 0xfd, 0x11, 0x48, 0xcb, 0xd9, 0x85,
	0x95, 0xb2, 0x56, 0xbd, 0xbc, 0xbd, 0x0e, 0x97, 0x63, 0x80, 0xf2, 0xfe, 0x66, 0xe6, 0xe8, 0xb4,
	0x94, 0xfa, 0x3a, 0x1b, 0xd7, 0x34, 0x4b, 0x11, 0x76, 0xb6, 0x3e, 0xcc, 0xc6, 0xb5, 0xc5, 0x55,
	0x1f, 0x67, 0xe3, 0xda, 0x66, 0x6c, 0xeb, 0x20, 0x32, 0x96, 0x10, 0x59, 0x29, 0x82, 0x7c, 0xe2,
	0xc8, 0x22, 0xdc, 0x67, 0x1e, 0x27, 0x15, 0x02, 0xb2, 0x6d, 0xee, 0x74, 0xde, 0x61, 0x7f, 0x97,
	0xb5, 0x3a, 0xbb, 0x7a, 0x01, 0xac, 0x61, 0xa9, 0x59, 0xba, 0xb1, 0xe2, 0x52, 0xaf, 0x83, 0x0b,
	0x61, 0x16, 0x4a, 0x6f, 0x11, 0x2a, 0x87, 0x61, 0x58, 0x50, 0x85, 0x05, 0x9f, 0x30, 0xd7, 0xb3,
	0xa2, 0xb6, 0x9d, 0x6c, 0xa8, 0x32, 0x26, 0x57, 0xd6, 0x41, 0xee, 0xf7, 0x31, 0xf3, 0xf1, 0x5f,
	0x34, 0x70, 0x63, 0x0e, 0x74, 0x04, 0xee, 0xee, 0x93, 0x90, 0x7d, 0x8e, 0x8c, 0x22, 0xb8, 0x24,
	0xd8, 0x9e, 0x4d, 0x3c, 0x46, 0x23, 0x29, 0x19, 0x6b, 0x4d, 0xb0, 0xa7, 0x61, 0xa9, 0x3f, 0x07,
	0x69, 0x4c, 0xd9, 0xc0, 0x13, 0x85, 0xd5, 0xb2, 0x56, 0xcd, 0x36, 0xb7, 0xc2, 0xec, 0x7e, 0x9c,
	0x96, 0x6e, 0x4a, 0xa9, 0xdc, 0x7e, 0x0b, 0x5d, 0x86, 0x28, 0x16, 0x7d, 0xd8, 0xf2, 0xc4, 0xc9,
	0x61, 0x1d, 0x28, 0x0f, 0x2d, 0x4f, 0xa8, 0x88, 0x25, 0x3f, 0x21, 0x7e, 0x13, 0x6c, 0x9c, 0xa1,
	0x31, 0xf6, 0xb0, 0xfd, 0x79, 0x05, 0xac, 0xb6, 0xb9, 0xa3, 0xbf, 0x06, 0xd9, 0xa5, 0xa7, 0x51,
	0x4a, 0xae, 0x34, 0xb1, 0x03, 0xe3, 0xee, 0x3f, 0x1a, 0xe2, 0x09, 0xfa, 0x0b, 0x90, 0x59, 0x6c,
	0xe8, 0xd6, 0x19, 0xac, 0x39, 0x6a, 0xdc, 0x39, 0x0f, 0x9d, 0x5f, 0x68, 0x83, 0x6b, 0x7f, 0x44,
	0x7e, 0xfb, 0xaf, 0xcc, 0x45, 0x93, 0x71, 0xef, 0x3f, 0x9a, 0xe2, 0x29, 0xc6, 0xc5, 0xf7, 0x61,
	0xa8, 0xcd, 0x67, 0x47, 0x13, 0x53, 0x3b, 0x9e, 0x98, 0xda, 0xcf, 0x89, 0xa9, 0x7d, 0x9a, 0x9a,
	0xa9, 0xe3, 0xa9, 0x99, 0xfa, 0x3e, 0x35, 0x53, 0x6f, 0xee, 0x3b, 0xae, 0xe8, 0x0f, 0xba, 0xb0,
	0xc7, 0x28, 0x62, 0x1e, 0xa3, 0xa3, 0xe8, 0x3b, 0xeb, 0xb1, 0x7d, 0xb4, 0xfc, 0x9e, 0xc5, 0xc8,
	0x27, 0xbc, 0x9b, 0x8e, 0xd0, 0x07, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0x96, 0x66, 0x02,
	0x4a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	SwapToIST(ctx context.Context, in *MsgSwapToIST, opts ...grpc.CallOption) (*MsgSwapToISTResponse, error)
	SwapToStablecoin(ctx context.Context, in *MsgSwapToStablecoin, opts ...grpc.CallOption) (*MsgSwapToStablecoinResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/reserve.psm.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapToIST(ctx context.Context, in *MsgSwapToIST, opts ...grpc.CallOption) (*MsgSwapToISTResponse, error) {
	out := new(MsgSwapToISTResponse)
	err := c.cc.Invoke(ctx, "/reserve.psm.v1.Msg/SwapToIST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapToStablecoin(ctx context.Context, in *MsgSwapToStablecoin, opts ...grpc.CallOption) (*MsgSwapToStablecoinResponse, error) {
	out := new(MsgSwapToStablecoinResponse)
	err := c.cc.Invoke(ctx, "/reserve.psm.v1.Msg/SwapToStablecoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	SwapToIST(context.Context, *MsgSwapToIST) (*MsgSwapToISTResponse, error)
	SwapToStablecoin(context.Context, *MsgSwapToStablecoin) (*MsgSwapToStablecoinResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) SwapToIST(ctx context.Context, req *MsgSwapToIST) (*MsgSwapToISTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapToIST not implemented")
}
func (*UnimplementedMsgServer) SwapToStablecoin(ctx context.Context, req *MsgSwapToStablecoin) (*MsgSwapToStablecoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapToStablecoin not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reserve.psm.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapToIST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapToIST)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapToIST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reserve.psm.v1.Msg/SwapToIST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapToIST(ctx, req.(*MsgSwapToIST))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapToStablecoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapToStablecoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapToStablecoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reserve.psm.v1.Msg/SwapToStablecoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapToStablecoin(ctx, req.(*MsgSwapToStablecoin))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reserve.psm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SwapToIST",
			Handler:    _Msg_SwapToIST_Handler,
		},
		{
			MethodName: "SwapToStablecoin",
			Handler:    _Msg_SwapToStablecoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reserve/psm/v1/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapToIST) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapToIST) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapToIST) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Coin != nil {
		{
			size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapToISTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapToISTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapToISTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapToStablecoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapToStablecoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapToStablecoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ToDenom) > 0 {
		i -= len(m.ToDenom)
		copy(dAtA[i:], m.ToDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapToStablecoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapToStablecoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapToStablecoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapToIST) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Coin != nil {
		l = m.Coin.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapToISTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapToStablecoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapToStablecoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapToIST) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapToIST: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapToIST: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Coin == nil {
				m.Coin = &types.Coin{}
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapToISTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapToISTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapToISTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapToStablecoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapToStablecoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapToStablecoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapToStablecoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapToStablecoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapToStablecoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)