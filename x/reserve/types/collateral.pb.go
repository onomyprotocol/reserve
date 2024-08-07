// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/collateral.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Collateral struct {
	Base             string                                  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Display          string                                  `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	MinimumDeposit   github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=minimum_deposit,json=minimumDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"minimum_deposit"`
	LendingRatio     uint64                                  `protobuf:"varint,4,opt,name=lending_ratio,json=lendingRatio,proto3" json:"lending_ratio,omitempty"`
	LiquidationRatio uint64                                  `protobuf:"varint,5,opt,name=liquidation_ratio,json=liquidationRatio,proto3" json:"liquidation_ratio,omitempty"`
	InitTime         int64                                   `protobuf:"varint,6,opt,name=init_time,json=initTime,proto3" json:"init_time,omitempty"`
}

func (m *Collateral) Reset()         { *m = Collateral{} }
func (m *Collateral) String() string { return proto.CompactTextString(m) }
func (*Collateral) ProtoMessage()    {}
func (*Collateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_813e0280592e4b6c, []int{0}
}
func (m *Collateral) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collateral.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collateral.Merge(m, src)
}
func (m *Collateral) XXX_Size() int {
	return m.Size()
}
func (m *Collateral) XXX_DiscardUnknown() {
	xxx_messageInfo_Collateral.DiscardUnknown(m)
}

var xxx_messageInfo_Collateral proto.InternalMessageInfo

func (m *Collateral) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Collateral) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Collateral) GetLendingRatio() uint64 {
	if m != nil {
		return m.LendingRatio
	}
	return 0
}

func (m *Collateral) GetLiquidationRatio() uint64 {
	if m != nil {
		return m.LiquidationRatio
	}
	return 0
}

func (m *Collateral) GetInitTime() int64 {
	if m != nil {
		return m.InitTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Collateral)(nil), "reserve.Collateral")
}

func init() { proto.RegisterFile("reserve/collateral.proto", fileDescriptor_813e0280592e4b6c) }

var fileDescriptor_813e0280592e4b6c = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x33, 0x6d, 0xff, 0xed, 0xbf, 0x83, 0x9f, 0x83, 0xe0, 0xa0, 0x30, 0x2d, 0xba, 0xb0,
	0x20, 0x76, 0x16, 0xe2, 0x0b, 0xb4, 0x3e, 0x41, 0x70, 0x21, 0x6e, 0x4a, 0xda, 0x0c, 0xf1, 0xe2,
	0x7c, 0xc4, 0xcc, 0x54, 0xec, 0xda, 0x17, 0xf0, 0x31, 0x7c, 0x94, 0x2e, 0xbb, 0x14, 0x17, 0x45,
	0x92, 0x17, 0x91, 0x4c, 0x12, 0x71, 0x35, 0xe7, 0xfe, 0xce, 0x99, 0xcb, 0xe5, 0x60, 0x9a, 0x09,
	0x2b, 0xb2, 0x17, 0xc1, 0x17, 0x46, 0xca, 0xc8, 0x89, 0x2c, 0x92, 0xe3, 0x34, 0x33, 0xce, 0x90,
	0x5e, 0xed, 0x9c, 0x1c, 0x25, 0x26, 0x31, 0x9e, 0xf1, 0x52, 0x55, 0xf6, 0xd9, 0x5b, 0x0b, 0xe3,
	0xe9, 0xef, 0x1f, 0x42, 0x70, 0x67, 0x1e, 0x59, 0x41, 0xd1, 0x10, 0x8d, 0xfa, 0xa1, 0xd7, 0x84,
	0xe2, 0x5e, 0x0c, 0x36, 0x95, 0xd1, 0x8a, 0xb6, 0x3c, 0x6e, 0x46, 0x72, 0x8f, 0xf7, 0x15, 0x68,
	0x50, 0x4b, 0x35, 0x8b, 0x45, 0x6a, 0x2c, 0x38, 0xda, 0x2e, 0x13, 0x13, 0xbe, 0xde, 0x0e, 0x82,
	0xaf, 0xed, 0xe0, 0x22, 0x01, 0xf7, 0xb8, 0x9c, 0x8f, 0x17, 0x46, 0xf1, 0x85, 0xb1, 0xca, 0xd8,
	0xfa, 0xb9, 0xb2, 0xf1, 0x13, 0x77, 0xab, 0x54, 0xd8, 0xf1, 0xd4, 0x80, 0x0e, 0xf7, 0xea, 0x3d,
	0xb7, 0xd5, 0x1a, 0x72, 0x8e, 0x77, 0xa5, 0xd0, 0x31, 0xe8, 0x64, 0x96, 0x45, 0x0e, 0x0c, 0xed,
	0x0c, 0xd1, 0xa8, 0x13, 0xee, 0xd4, 0x30, 0x2c, 0x19, 0xb9, 0xc4, 0x87, 0x12, 0x9e, 0x97, 0x10,
	0x97, 0x93, 0xae, 0x83, 0xff, 0x7c, 0xf0, 0xe0, 0x8f, 0x51, 0x85, 0x4f, 0x71, 0x1f, 0x34, 0xb8,
	0x99, 0x03, 0x25, 0x68, 0x77, 0x88, 0x46, 0xed, 0xf0, 0x7f, 0x09, 0xee, 0x40, 0x89, 0xc9, 0xcd,
	0x47, 0xce, 0xd0, 0x3a, 0x67, 0x68, 0x93, 0x33, 0xf4, 0x9d, 0x33, 0xf4, 0x5e, 0xb0, 0x60, 0x53,
	0xb0, 0xe0, 0xb3, 0x60, 0xc1, 0xc3, 0x71, 0x53, 0xee, 0x2b, 0x6f, 0x94, 0x3f, 0x7d, 0xde, 0xf5,
	0x1d, 0x5e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x34, 0x44, 0xc3, 0x7e, 0x01, 0x00, 0x00,
}

func (this *Collateral) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Collateral)
	if !ok {
		that2, ok := that.(Collateral)
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
	if this.Base != that1.Base {
		return false
	}
	if this.Display != that1.Display {
		return false
	}
	if !this.MinimumDeposit.Equal(that1.MinimumDeposit) {
		return false
	}
	if this.LendingRatio != that1.LendingRatio {
		return false
	}
	if this.LiquidationRatio != that1.LiquidationRatio {
		return false
	}
	if this.InitTime != that1.InitTime {
		return false
	}
	return true
}
func (m *Collateral) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collateral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collateral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitTime != 0 {
		i = encodeVarintCollateral(dAtA, i, uint64(m.InitTime))
		i--
		dAtA[i] = 0x30
	}
	if m.LiquidationRatio != 0 {
		i = encodeVarintCollateral(dAtA, i, uint64(m.LiquidationRatio))
		i--
		dAtA[i] = 0x28
	}
	if m.LendingRatio != 0 {
		i = encodeVarintCollateral(dAtA, i, uint64(m.LendingRatio))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.MinimumDeposit.Size()
		i -= size
		if _, err := m.MinimumDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCollateral(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintCollateral(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintCollateral(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollateral(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollateral(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collateral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovCollateral(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovCollateral(uint64(l))
	}
	l = m.MinimumDeposit.Size()
	n += 1 + l + sovCollateral(uint64(l))
	if m.LendingRatio != 0 {
		n += 1 + sovCollateral(uint64(m.LendingRatio))
	}
	if m.LiquidationRatio != 0 {
		n += 1 + sovCollateral(uint64(m.LiquidationRatio))
	}
	if m.InitTime != 0 {
		n += 1 + sovCollateral(uint64(m.InitTime))
	}
	return n
}

func sovCollateral(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollateral(x uint64) (n int) {
	return sovCollateral(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collateral) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollateral
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
			return fmt.Errorf("proto: Collateral: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collateral: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
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
				return ErrInvalidLengthCollateral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollateral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
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
				return ErrInvalidLengthCollateral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollateral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
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
				return ErrInvalidLengthCollateral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollateral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendingRatio", wireType)
			}
			m.LendingRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LendingRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			m.LiquidationRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidationRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitTime", wireType)
			}
			m.InitTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollateral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollateral(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollateral
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
func skipCollateral(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollateral
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
					return 0, ErrIntOverflowCollateral
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
					return 0, ErrIntOverflowCollateral
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
				return 0, ErrInvalidLengthCollateral
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollateral
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollateral
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollateral        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollateral          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollateral = fmt.Errorf("proto: unexpected end of group")
)
