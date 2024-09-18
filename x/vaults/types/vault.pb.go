// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/vaults/vault.proto

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
	return fileDescriptor_1f64e7967fdcb058, []int{0}
}

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
	return fileDescriptor_1f64e7967fdcb058, []int{0}
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
	return fileDescriptor_1f64e7967fdcb058, []int{1}
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
	proto.RegisterType((*VaultMamager)(nil), "reserve.vaults.VaultMamager")
	proto.RegisterType((*Vault)(nil), "reserve.vaults.Vault")
}

func init() { proto.RegisterFile("reserve/vaults/vault.proto", fileDescriptor_1f64e7967fdcb058) }

var fileDescriptor_1f64e7967fdcb058 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0xef, 0x40, 0x69, 0x60, 0x8a, 0xb5, 0x4c, 0xd0, 0x94, 0x25, 0x59, 0x36, 0x9c, 0x08, 0x09,
	0xbb, 0x02, 0x17, 0xaf, 0x2d, 0x54, 0xb3, 0x09, 0xa2, 0x6c, 0x11, 0x13, 0x2f, 0x64, 0x76, 0x77,
	0x5c, 0x26, 0xec, 0xce, 0xe0, 0xcc, 0xb4, 0xca, 0x37, 0x30, 0x3d, 0xf9, 0x05, 0x7a, 0x30, 0x5e,
	0x3c, 0x7a, 0xe0, 0x23, 0x78, 0xe0, 0x48, 0x38, 0x19, 0x0f, 0xc4, 0xd0, 0x83, 0x9f, 0xc1, 0x9b,
	0xd9, 0x99, 0xdd, 0x58, 0x8d, 0x07, 0x2f, 0xed, 0xce, 0xfb, 0xfd, 0x79, 0xbf, 0xf7, 0xf2, 0xa0,
	0x25, 0x88, 0x24, 0x62, 0x40, 0xbc, 0x01, 0xee, 0xa7, 0x4a, 0x9a, 0x3f, 0xf7, 0x4c, 0x70, 0xc5,
	0x51, 0xa3, 0xc0, 0x5c, 0x83, 0x59, 0x8b, 0x09, 0x4f, 0xb8, 0x86, 0xbc, 0xfc, 0xcb, 0xb0, 0xac,
	0xe5, 0xd2, 0x81, 0x0b, 0x1c, 0xa5, 0xc4, 0x3b, 0xc3, 0x02, 0x67, 0xb2, 0x00, 0x17, 0x70, 0x46,
	0x19, 0xf7, 0xf4, 0x6f, 0x51, 0xb2, 0x23, 0x2e, 0x33, 0x2e, 0xbd, 0x10, 0x4b, 0xe2, 0x0d, 0x36,
	0x43, 0xa2, 0xf0, 0xa6, 0x17, 0x71, 0xca, 0x0a, 0x7c, 0xc9, 0xe0, 0xc7, 0xa6, 0x91, 0x79, 0xfc,
	0xdd, 0xaa, 0x08, 0x3b, 0xd9, 0x6a, 0xf5, 0x0b, 0x80, 0xf3, 0x47, 0x79, 0xfd, 0x09, 0xce, 0x70,
	0x42, 0x04, 0xea, 0xc2, 0x9a, 0x21, 0xb4, 0x80, 0x03, 0xd6, 0xea, 0x5b, 0xab, 0xee, 0x9f, 0xf3,
	0xb8, 0x93, 0xec, 0x67, 0x9a, 0xd9, 0x99, 0xbb, 0xbc, 0x59, 0xa9, 0x7c, 0xfa, 0xf1, 0x79, 0x1d,
	0x04, 0x85, 0x18, 0x2d, 0xc2, 0x99, 0x98, 0x30, 0x9e, 0xb5, 0xa6, 0x1c, 0xb0, 0x36, 0x17, 0x98,
	0x07, 0x7a, 0x01, 0x1b, 0x19, 0x65, 0xea, 0x18, 0x0f, 0x30, 0x4d, 0x71, 0x98, 0x92, 0xd6, 0x74,
	0x0e, 0x77, 0x1e, 0xe4, 0x06, 0xdf, 0x6e, 0x56, 0xee, 0x99, 0xe0, 0x32, 0x3e, 0x75, 0x29, 0xf7,
	0x32, 0xac, 0x4e, 0x5c, 0x9f, 0xa9, 0xeb, 0x8b, 0x0d, 0x58, 0x4c, 0xe4, 0x33, 0x65, 0xfa, 0xdc,
	0xc9, 0x7d, 0xda, 0xa5, 0xcd, 0xea, 0x4f, 0x00, 0x67, 0x74, 0x30, 0xe4, 0xc2, 0x19, 0xfe, 0x86,
	0x11, 0xa1, 0xe3, 0xcf, 0x75, 0x5a, 0xd7, 0x17, 0x1b, 0x8b, 0x85, 0xb8, 0x1d, 0xc7, 0x82, 0x48,
	0xd9, 0x53, 0x82, 0xb2, 0x24, 0x30, 0x34, 0xf4, 0x10, 0x56, 0x63, 0x12, 0x2a, 0x9d, 0xb3, 0xbe,
	0xb5, 0xe4, 0x16, 0xdc, 0x7c, 0xcf, 0x6e, 0xb1, 0x67, 0x77, 0x87, 0x53, 0x36, 0x39, 0xa4, 0x56,
	0xa0, 0x03, 0xb8, 0x10, 0xf1, 0x34, 0xc5, 0x8a, 0x08, 0x9c, 0x1e, 0xa7, 0x3c, 0x3a, 0x25, 0xb1,
	0x9e, 0xe7, 0x7f, 0x6d, 0x9a, 0xbf, 0xe5, 0x7b, 0x5a, 0x8d, 0xb6, 0x61, 0x4d, 0x2a, 0xac, 0xfa,
	0xb2, 0x55, 0x75, 0xc0, 0x5a, 0x63, 0x6b, 0xf9, 0x9f, 0xcb, 0xef, 0x69, 0x4a, 0x50, 0x50, 0xd7,
	0x3f, 0x00, 0x58, 0x9f, 0xa8, 0xa3, 0xfb, 0xb0, 0xd6, 0xde, 0x39, 0xf4, 0x8f, 0xba, 0xcd, 0x8a,
	0x05, 0x87, 0x23, 0xa7, 0x86, 0x23, 0x45, 0x07, 0x04, 0x39, 0xb0, 0xbe, 0xe7, 0x1f, 0x3c, 0xf7,
	0x77, 0xdb, 0x87, 0xfe, 0xfe, 0xe3, 0x26, 0xb0, 0xee, 0x0e, 0x47, 0x4e, 0x3d, 0xa5, 0xaf, 0xfb,
	0x34, 0xc6, 0x8a, 0xb2, 0x04, 0x59, 0x70, 0xf6, 0x30, 0x68, 0xef, 0xf7, 0x1e, 0x75, 0x83, 0xe6,
	0x94, 0x35, 0x3f, 0x1c, 0x39, 0xb3, 0x4a, 0x60, 0x26, 0x5f, 0x11, 0x91, 0xbb, 0xee, 0xec, 0x3d,
	0xed, 0x75, 0x77, 0x9b, 0xd3, 0xc6, 0x35, 0x4a, 0xb9, 0x24, 0x31, 0xb2, 0x21, 0x2c, 0x5d, 0xbb,
	0xbb, 0xcd, 0xaa, 0xd5, 0x18, 0x8e, 0x1c, 0x58, 0x9a, 0x92, 0xd8, 0xaa, 0xbe, 0xfb, 0x68, 0x57,
	0x3a, 0xfe, 0xe5, 0xad, 0x0d, 0xae, 0x6e, 0x6d, 0xf0, 0xfd, 0xd6, 0x06, 0xef, 0xc7, 0x76, 0xe5,
	0x6a, 0x6c, 0x57, 0xbe, 0x8e, 0xed, 0xca, 0x4b, 0x2f, 0xa1, 0xea, 0xa4, 0x1f, 0xba, 0x11, 0xcf,
	0x3c, 0xce, 0x78, 0x76, 0xae, 0xef, 0x32, 0xe2, 0xa9, 0x57, 0x9e, 0xed, 0xdb, 0xf2, 0x70, 0xd5,
	0xf9, 0x19, 0x91, 0x61, 0x4d, 0x13, 0xb6, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x55, 0x3d,
	0x74, 0x84, 0x03, 0x00, 0x00,
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
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
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
		i = encodeVarintVault(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.CollateralLocked.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Debt.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VaultMamager) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovVault(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.MintAvailable.Size()
	n += 1 + l + sovVault(uint64(l))
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
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.Debt.Size()
	n += 1 + l + sovVault(uint64(l))
	l = m.CollateralLocked.Size()
	n += 1 + l + sovVault(uint64(l))
	if m.Status != 0 {
		n += 1 + sovVault(uint64(m.Status))
	}
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VaultMamager) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
				return ErrIntOverflowVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
					return ErrIntOverflowVault
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
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
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
					return ErrIntOverflowVault
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
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func skipVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
				return 0, ErrInvalidLengthVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVault = fmt.Errorf("proto: unexpected end of group")
)