// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/vaults/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/onomyprotocol/reserve/x/oracle/types"
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

// VaultStatus is the status of a vault.
type VaultStatus int32

const (
	// ACTIVE - vault is in use and can be changed
	active VaultStatus = 0
	// LIQUIDATING - vault is being liquidated by the vault manager, and cannot be
	// changed by the user. If liquidation fails, vaults may remain in this state.
	// An upgrade might be able to recover them.
	liquidating VaultStatus = 1
	// TRANSFER - vault is able to be transferred (payments and debits frozen until
	// it has a new owner)
	transfer VaultStatus = 2
	// CLOSED - vault was closed by the user and all assets have been paid out
	closed VaultStatus = 3
	// LIQUIDATED - vault was closed by the manager, with remaining assets paid to owner
	liquidated VaultStatus = 4
)

var VaultStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "LIQUIDATING",
	2: "TRANSFER",
	3: "CLOSED",
	4: "LIQUIDATED",
}

var VaultStatus_value = map[string]int32{
	"ACTIVE":      0,
	"LIQUIDATING": 1,
	"TRANSFER":    2,
	"CLOSED":      3,
	"LIQUIDATED":  4,
}

func (x VaultStatus) String() string {
	return proto.EnumName(VaultStatus_name, int32(x))
}

func (VaultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{0}
}

// Params defines the parameters for the module.
type Params struct {
	MintingFee         cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=minting_fee,json=mintingFee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"minting_fee"`
	StabilityFee       cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=stability_fee,json=stabilityFee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"stability_fee"`
	LiquidationPenalty cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidation_penalty"`
	MinInitialDebt     cosmossdk_io_math.Int       `protobuf:"bytes,4,opt,name=min_initial_debt,json=minInitialDebt,proto3,customtype=cosmossdk.io/math.Int" json:"min_initial_debt"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// VaultParams defines the parameters for each collateral vault type.
type VaultMamagerParams struct {
	MinCollateralRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=min_collateral_ratio,json=minCollateralRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_collateral_ratio"`
	LiquidationRatio   cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=liquidation_ratio,json=liquidationRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidation_ratio"`
	MaxDebt            cosmossdk_io_math.Int       `protobuf:"bytes,3,opt,name=max_debt,json=maxDebt,proto3,customtype=cosmossdk.io/math.Int" json:"max_debt"`
}

func (m *VaultMamagerParams) Reset()         { *m = VaultMamagerParams{} }
func (m *VaultMamagerParams) String() string { return proto.CompactTextString(m) }
func (*VaultMamagerParams) ProtoMessage()    {}
func (*VaultMamagerParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{1}
}
func (m *VaultMamagerParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultMamagerParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultMamagerParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultMamagerParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultMamagerParams.Merge(m, src)
}
func (m *VaultMamagerParams) XXX_Size() int {
	return m.Size()
}
func (m *VaultMamagerParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultMamagerParams.DiscardUnknown(m)
}

var xxx_messageInfo_VaultMamagerParams proto.InternalMessageInfo

// VaultMamager defines the manager of each collateral vault type.
type VaultMamager struct {
	Params        VaultMamagerParams    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Denom         string                `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	MintAvailable cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=mint_available,json=mintAvailable,proto3,customtype=cosmossdk.io/math.Int" json:"mint_available"`
}

func (m *VaultMamager) Reset()         { *m = VaultMamager{} }
func (m *VaultMamager) String() string { return proto.CompactTextString(m) }
func (*VaultMamager) ProtoMessage()    {}
func (*VaultMamager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{2}
}
func (m *VaultMamager) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultMamager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultMamager.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultMamager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultMamager.Merge(m, src)
}
func (m *VaultMamager) XXX_Size() int {
	return m.Size()
}
func (m *VaultMamager) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultMamager.DiscardUnknown(m)
}

var xxx_messageInfo_VaultMamager proto.InternalMessageInfo

func (m *VaultMamager) GetParams() VaultMamagerParams {
	if m != nil {
		return m.Params
	}
	return VaultMamagerParams{}
}

func (m *VaultMamager) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type Vault struct {
	Owner            string      `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Debt             types.Coin  `protobuf:"bytes,2,opt,name=debt,proto3" json:"debt"`
	CollateralLocked types.Coin  `protobuf:"bytes,3,opt,name=collateral_locked,json=collateralLocked,proto3" json:"collateral_locked"`
	Status           VaultStatus `protobuf:"varint,4,opt,name=status,proto3,enum=reserve.vaults.VaultStatus" json:"status,omitempty"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f12ab0d072f9f7c, []int{3}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

func (m *Vault) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Vault) GetDebt() types.Coin {
	if m != nil {
		return m.Debt
	}
	return types.Coin{}
}

func (m *Vault) GetCollateralLocked() types.Coin {
	if m != nil {
		return m.CollateralLocked
	}
	return types.Coin{}
}

func (m *Vault) GetStatus() VaultStatus {
	if m != nil {
		return m.Status
	}
	return active
}

func init() {
	proto.RegisterEnum("reserve.vaults.VaultStatus", VaultStatus_name, VaultStatus_value)
	proto.RegisterType((*Params)(nil), "reserve.vaults.Params")
	proto.RegisterType((*VaultMamagerParams)(nil), "reserve.vaults.VaultMamagerParams")
	proto.RegisterType((*VaultMamager)(nil), "reserve.vaults.VaultMamager")
	proto.RegisterType((*Vault)(nil), "reserve.vaults.Vault")
}

func init() { proto.RegisterFile("reserve/vaults/params.proto", fileDescriptor_1f12ab0d072f9f7c) }

var fileDescriptor_1f12ab0d072f9f7c = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x4f, 0x03, 0x45,
	0x14, 0xc7, 0xbb, 0x6d, 0xa9, 0x30, 0x85, 0xba, 0x8c, 0x55, 0xcb, 0x92, 0x2c, 0x4d, 0x4f, 0x86,
	0x84, 0x5d, 0x81, 0xc4, 0x18, 0x6f, 0xa5, 0x2d, 0x66, 0x63, 0x45, 0x68, 0x11, 0x12, 0x3c, 0x34,
	0xb3, 0xbb, 0xc3, 0x32, 0x61, 0x77, 0xa6, 0xee, 0x4e, 0x2b, 0xfd, 0x06, 0xa6, 0x27, 0xbf, 0x40,
	0x13, 0x8d, 0x31, 0xf1, 0xc8, 0x81, 0x8f, 0xe0, 0x81, 0x23, 0xe1, 0x64, 0x3c, 0x10, 0x03, 0x07,
	0xfc, 0x0a, 0xde, 0xcc, 0xce, 0x6c, 0xeb, 0x12, 0x3c, 0x18, 0x7a, 0x69, 0xba, 0xf3, 0xde, 0xfb,
	0xfd, 0xe7, 0xfd, 0xdf, 0xcc, 0x80, 0xf5, 0x10, 0x47, 0x38, 0x1c, 0x62, 0x73, 0x88, 0x06, 0x3e,
	0x8f, 0xcc, 0x3e, 0x0a, 0x51, 0x10, 0x19, 0xfd, 0x90, 0x71, 0x06, 0x4b, 0x49, 0xd0, 0x90, 0x41,
	0xad, 0xec, 0x31, 0x8f, 0x89, 0x90, 0x19, 0xff, 0x93, 0x59, 0xda, 0x0c, 0xc1, 0x42, 0xe4, 0xf8,
	0xf8, 0x05, 0x42, 0x5b, 0x45, 0x01, 0xa1, 0xcc, 0x14, 0xbf, 0xc9, 0x92, 0xee, 0xb0, 0x28, 0x60,
	0x91, 0x69, 0xa3, 0x08, 0x9b, 0xc3, 0x6d, 0x1b, 0x73, 0xb4, 0x6d, 0x3a, 0x8c, 0xd0, 0x24, 0xbe,
	0x26, 0xe3, 0x3d, 0x29, 0x24, 0x3f, 0x64, 0xa8, 0xf6, 0x4b, 0x0e, 0x14, 0x0e, 0x05, 0x1e, 0x9e,
	0x82, 0x62, 0x40, 0x28, 0x27, 0xd4, 0xeb, 0x9d, 0x63, 0x5c, 0x51, 0xaa, 0xca, 0x47, 0x4b, 0x7b,
	0x9f, 0xdc, 0x3e, 0x6c, 0x64, 0xfe, 0x78, 0xd8, 0x58, 0x97, 0x55, 0x91, 0x7b, 0x69, 0x10, 0x66,
	0x06, 0x88, 0x5f, 0x18, 0x6d, 0xec, 0x21, 0x67, 0xd4, 0xc4, 0xce, 0xfd, 0xcd, 0x16, 0x48, 0xa0,
	0x4d, 0xec, 0xfc, 0xfa, 0x7c, 0xbd, 0xa9, 0x74, 0x40, 0x82, 0xda, 0xc7, 0x18, 0x7e, 0x03, 0x56,
	0x22, 0x8e, 0x6c, 0xe2, 0x13, 0x3e, 0x12, 0xe8, 0xec, 0x5c, 0xe8, 0xe5, 0x19, 0x2c, 0x86, 0x7b,
	0xe0, 0x3d, 0x9f, 0x7c, 0x3b, 0x20, 0x2e, 0xe2, 0x84, 0xd1, 0x5e, 0x1f, 0x53, 0xe4, 0xf3, 0x51,
	0x25, 0x37, 0x97, 0x04, 0x4c, 0x21, 0x0f, 0x25, 0x11, 0x9e, 0x01, 0x35, 0x20, 0xb4, 0x47, 0x28,
	0xe1, 0x04, 0xf9, 0x3d, 0x17, 0xdb, 0xbc, 0x92, 0x17, 0x2a, 0x1f, 0x27, 0x2a, 0xef, 0xbf, 0x56,
	0xb1, 0x28, 0x4f, 0xf1, 0x2d, 0xca, 0x25, 0xbf, 0x14, 0x10, 0x6a, 0x49, 0x50, 0x13, 0xdb, 0xfc,
	0xb3, 0xea, 0x5f, 0x3f, 0x6e, 0x28, 0xe3, 0xe7, 0xeb, 0xcd, 0x0f, 0xa7, 0x93, 0xbf, 0x9a, 0x1e,
	0x1f, 0x39, 0x9c, 0xda, 0x75, 0x16, 0xc0, 0x93, 0x78, 0xe5, 0x4b, 0x14, 0x20, 0x0f, 0x87, 0xc9,
	0xcc, 0x2e, 0x40, 0x39, 0xde, 0x94, 0xc3, 0x7c, 0x1f, 0x71, 0x1c, 0x22, 0xbf, 0x17, 0xc6, 0x9b,
	0x9e, 0x73, 0x78, 0x30, 0x20, 0xb4, 0x31, 0x43, 0x76, 0x62, 0x22, 0x74, 0xc0, 0x6a, 0xda, 0x67,
	0x29, 0x33, 0xdf, 0x20, 0xd5, 0x14, 0x50, 0x8a, 0x7c, 0x01, 0x16, 0x03, 0x74, 0x25, 0xbd, 0xcd,
	0xbd, 0xd1, 0xdb, 0x77, 0x02, 0x74, 0x15, 0x9b, 0x5a, 0xfb, 0x4d, 0x01, 0xcb, 0x69, 0xcb, 0x60,
	0x0b, 0x14, 0xe4, 0x4d, 0x12, 0xf6, 0x14, 0x77, 0x6a, 0xc6, 0xcb, 0xdb, 0x68, 0xbc, 0x36, 0x78,
	0x6f, 0x29, 0xd6, 0x97, 0xe0, 0xa4, 0x18, 0x96, 0xc1, 0x82, 0x8b, 0x29, 0x0b, 0x64, 0xf7, 0x1d,
	0xf9, 0x01, 0x4f, 0x41, 0x3c, 0x54, 0xde, 0x43, 0x43, 0x44, 0x7c, 0x64, 0xfb, 0xf8, 0xcd, 0x0d,
	0xac, 0xc4, 0x9c, 0xfa, 0x14, 0x53, 0xfb, 0x5b, 0x01, 0x0b, 0x62, 0x63, 0xd0, 0x00, 0x0b, 0xec,
	0x3b, 0x8a, 0xc3, 0x64, 0xba, 0x95, 0xfb, 0x9b, 0xad, 0x72, 0x52, 0x5c, 0x77, 0xdd, 0x10, 0x47,
	0x51, 0x97, 0x87, 0x84, 0x7a, 0x1d, 0x99, 0x06, 0x3f, 0x05, 0x79, 0xe1, 0x64, 0x56, 0x74, 0xbb,
	0x66, 0x24, 0xb9, 0xf1, 0x2b, 0x61, 0x24, 0xaf, 0x84, 0xd1, 0x60, 0x84, 0xa6, 0x9b, 0x14, 0x15,
	0xf0, 0x08, 0xac, 0xa6, 0x8e, 0x94, 0xcf, 0x9c, 0x4b, 0xec, 0x8a, 0x7e, 0xfe, 0x2f, 0x46, 0xfd,
	0xb7, 0xbc, 0x2d, 0xaa, 0xe1, 0x2e, 0x28, 0x44, 0x1c, 0xf1, 0x41, 0x24, 0x2e, 0x4d, 0x69, 0x67,
	0xfd, 0x3f, 0xcd, 0xef, 0x8a, 0x94, 0x4e, 0x92, 0xba, 0xf9, 0x93, 0x02, 0x8a, 0xa9, 0x75, 0xf8,
	0x01, 0x28, 0xd4, 0x1b, 0xc7, 0xd6, 0x49, 0x4b, 0xcd, 0x68, 0x60, 0x3c, 0xa9, 0x16, 0x90, 0xc3,
	0xc9, 0x10, 0xc3, 0x2a, 0x28, 0xb6, 0xad, 0xa3, 0xaf, 0xad, 0x66, 0xfd, 0xd8, 0x3a, 0xf8, 0x5c,
	0x55, 0xb4, 0x77, 0xc7, 0x93, 0x6a, 0x71, 0x76, 0xbc, 0xa8, 0x07, 0x35, 0xb0, 0x78, 0xdc, 0xa9,
	0x1f, 0x74, 0xf7, 0x5b, 0x1d, 0x35, 0xab, 0x2d, 0x8f, 0x27, 0xd5, 0x45, 0x1e, 0x22, 0x1a, 0x9d,
	0xe3, 0x30, 0xa6, 0x36, 0xda, 0x5f, 0x75, 0x5b, 0x4d, 0x35, 0x27, 0xa9, 0x8e, 0xcf, 0x22, 0xec,
	0x42, 0x1d, 0x80, 0x29, 0xb5, 0xd5, 0x54, 0xf3, 0x5a, 0x69, 0x3c, 0xa9, 0x82, 0x29, 0x14, 0xbb,
	0x5a, 0xfe, 0xfb, 0x9f, 0xf5, 0xcc, 0x9e, 0x75, 0xfb, 0xa8, 0x2b, 0x77, 0x8f, 0xba, 0xf2, 0xe7,
	0xa3, 0xae, 0xfc, 0xf0, 0xa4, 0x67, 0xee, 0x9e, 0xf4, 0xcc, 0xef, 0x4f, 0x7a, 0xe6, 0xcc, 0xf4,
	0x08, 0xbf, 0x18, 0xd8, 0x86, 0xc3, 0x02, 0x93, 0x51, 0x16, 0x8c, 0xc4, 0x93, 0xeb, 0x30, 0xdf,
	0x7c, 0x75, 0xcb, 0xf9, 0xa8, 0x8f, 0x23, 0xbb, 0x20, 0x12, 0x76, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x84, 0x31, 0x3d, 0xc9, 0x43, 0x06, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MintingFee.Equal(that1.MintingFee) {
		return false
	}
	if !this.StabilityFee.Equal(that1.StabilityFee) {
		return false
	}
	if !this.LiquidationPenalty.Equal(that1.LiquidationPenalty) {
		return false
	}
	if !this.MinInitialDebt.Equal(that1.MinInitialDebt) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinInitialDebt.Size()
		i -= size
		if _, err := m.MinInitialDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.StabilityFee.Size()
		i -= size
		if _, err := m.StabilityFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MintingFee.Size()
		i -= size
		if _, err := m.MintingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultMamagerParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultMamagerParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultMamagerParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxDebt.Size()
		i -= size
		if _, err := m.MaxDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.LiquidationRatio.Size()
		i -= size
		if _, err := m.LiquidationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinCollateralRatio.Size()
		i -= size
		if _, err := m.MinCollateralRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultMamager) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultMamager) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultMamager) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MintAvailable.Size()
		i -= size
		if _, err := m.MintAvailable.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.CollateralLocked.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Debt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MintingFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.StabilityFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinInitialDebt.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *VaultMamagerParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinCollateralRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.LiquidationRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDebt.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *VaultMamager) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MintAvailable.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Debt.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CollateralLocked.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.Status != 0 {
		n += 1 + sovParams(uint64(m.Status))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StabilityFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StabilityFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitialDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *VaultMamagerParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: VaultMamagerParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultMamagerParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCollateralRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCollateralRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDebt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *VaultMamager) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: VaultMamager: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultMamager: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintAvailable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintAvailable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Vault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Debt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralLocked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralLocked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= VaultStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)