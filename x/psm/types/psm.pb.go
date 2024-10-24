// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/psm/v1/psm.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Stablecoin struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// limit total stablecoin module support
	LimitTotal cosmossdk_io_math.Int       `protobuf:"bytes,2,opt,name=limit_total,json=limitTotal,proto3,customtype=cosmossdk.io/math.Int" json:"limit_total"`
	FeeIn      cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=fee_in,json=feeIn,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee_in"`
	FeeOut     cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=fee_out,json=feeOut,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee_out"`
}

func (m *Stablecoin) Reset()         { *m = Stablecoin{} }
func (m *Stablecoin) String() string { return proto.CompactTextString(m) }
func (*Stablecoin) ProtoMessage()    {}
func (*Stablecoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_59572214fa05fb2f, []int{0}
}
func (m *Stablecoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stablecoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stablecoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stablecoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stablecoin.Merge(m, src)
}
func (m *Stablecoin) XXX_Size() int {
	return m.Size()
}
func (m *Stablecoin) XXX_DiscardUnknown() {
	xxx_messageInfo_Stablecoin.DiscardUnknown(m)
}

var xxx_messageInfo_Stablecoin proto.InternalMessageInfo

func (m *Stablecoin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Stablecoin)(nil), "reserve.psm.v1.Stablecoin")
}

func init() { proto.RegisterFile("reserve/psm/v1/psm.proto", fileDescriptor_59572214fa05fb2f) }

var fileDescriptor_59572214fa05fb2f = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x93, 0x7e, 0x5f, 0x2b, 0xae, 0x22, 0x18, 0x2a, 0xc4, 0x0a, 0x69, 0xf1, 0x54, 0x44,
	0xb3, 0x06, 0xdf, 0xa0, 0x14, 0xb1, 0x22, 0x88, 0xd5, 0x93, 0x97, 0x92, 0xa4, 0xd3, 0x34, 0x98,
	0xdd, 0x29, 0xdd, 0x6d, 0xb0, 0xcf, 0xe0, 0xc5, 0xc7, 0xf0, 0xe8, 0xc1, 0x87, 0xe8, 0xb1, 0x78,
	0x12, 0x0f, 0x45, 0xda, 0x83, 0xaf, 0x21, 0xbb, 0x1b, 0x41, 0xf0, 0xe6, 0x25, 0xd9, 0xf9, 0xff,
	0x86, 0xdf, 0x0e, 0x3b, 0xc4, 0x1d, 0x83, 0x80, 0x71, 0x0e, 0x74, 0x24, 0x18, 0xcd, 0x03, 0xf5,
	0xf3, 0x47, 0x63, 0x94, 0xe8, 0x6c, 0x15, 0xc4, 0x57, 0x51, 0x1e, 0xd4, 0xaa, 0x09, 0x26, 0xa8,
	0x11, 0x55, 0x27, 0xd3, 0x55, 0xdb, 0x8d, 0x51, 0x30, 0x14, 0x3d, 0x03, 0x4c, 0x51, 0xa0, 0xed,
	0x90, 0xa5, 0x1c, 0xa9, 0xfe, 0x16, 0x91, 0x67, 0x1a, 0x68, 0x14, 0x0a, 0xa0, 0x79, 0x10, 0x81,
	0x0c, 0x03, 0x1a, 0x63, 0xca, 0x0d, 0xdf, 0x7f, 0x28, 0x11, 0x72, 0x2d, 0xc3, 0x28, 0x03, 0x15,
	0x3a, 0x55, 0x52, 0xee, 0x03, 0x47, 0xe6, 0xda, 0x0d, 0xbb, 0xb9, 0xde, 0x35, 0x85, 0x73, 0x45,
	0x36, 0xb2, 0x94, 0xa5, 0xb2, 0x27, 0x51, 0x86, 0x99, 0x5b, 0x6a, 0xd8, 0xcd, 0xcd, 0xd6, 0xf1,
	0x6c, 0x51, 0xb7, 0xde, 0x17, 0xf5, 0x1d, 0x73, 0x83, 0xe8, 0xdf, 0xf9, 0x29, 0x52, 0x16, 0xca,
	0xa1, 0xdf, 0xe1, 0xf2, 0xf5, 0xe5, 0x88, 0x14, 0xb3, 0x75, 0xb8, 0x7c, 0xfa, 0x7c, 0x3e, 0xb0,
	0xbb, 0x44, 0x4b, 0x6e, 0x94, 0xc3, 0x39, 0x23, 0x95, 0x01, 0x40, 0x2f, 0xe5, 0xee, 0x3f, 0x6d,
	0x0b, 0x0a, 0xdb, 0xde, 0x6f, 0xdb, 0x05, 0x24, 0x61, 0x3c, 0x6d, 0x43, 0xfc, 0xc3, 0xd9, 0x86,
	0xb8, 0x5b, 0x1e, 0x00, 0x74, 0xb8, 0x73, 0x4e, 0xd6, 0x94, 0x09, 0x27, 0xd2, 0xfd, 0xff, 0x57,
	0x95, 0x9a, 0xe5, 0x72, 0x22, 0x5b, 0xa7, 0xb3, 0xa5, 0x67, 0xcf, 0x97, 0x9e, 0xfd, 0xb1, 0xf4,
	0xec, 0xc7, 0x95, 0x67, 0xcd, 0x57, 0x9e, 0xf5, 0xb6, 0xf2, 0xac, 0xdb, 0xc3, 0x24, 0x95, 0xc3,
	0x49, 0xe4, 0xc7, 0xc8, 0x28, 0x72, 0x64, 0x53, 0xfd, 0x7c, 0x31, 0x66, 0xf4, 0x7b, 0x9d, 0xf7,
	0x7a, 0xa1, 0x72, 0x3a, 0x02, 0x11, 0x55, 0x34, 0x3d, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0xf0, 0x97, 0x09, 0xec, 0x01, 0x00, 0x00,
}

func (m *Stablecoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stablecoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stablecoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeOut.Size()
		i -= size
		if _, err := m.FeeOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPsm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.FeeIn.Size()
		i -= size
		if _, err := m.FeeIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPsm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.LimitTotal.Size()
		i -= size
		if _, err := m.LimitTotal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPsm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPsm(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPsm(dAtA []byte, offset int, v uint64) int {
	offset -= sovPsm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stablecoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPsm(uint64(l))
	}
	l = m.LimitTotal.Size()
	n += 1 + l + sovPsm(uint64(l))
	l = m.FeeIn.Size()
	n += 1 + l + sovPsm(uint64(l))
	l = m.FeeOut.Size()
	n += 1 + l + sovPsm(uint64(l))
	return n
}

func sovPsm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPsm(x uint64) (n int) {
	return sovPsm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stablecoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPsm
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
			return fmt.Errorf("proto: Stablecoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stablecoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPsm
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
				return ErrInvalidLengthPsm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPsm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitTotal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPsm
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
				return ErrInvalidLengthPsm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPsm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LimitTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeIn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPsm
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
				return ErrInvalidLengthPsm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPsm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeOut", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPsm
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
				return ErrInvalidLengthPsm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPsm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPsm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPsm
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
func skipPsm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPsm
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
					return 0, ErrIntOverflowPsm
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
					return 0, ErrIntOverflowPsm
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
				return 0, ErrInvalidLengthPsm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPsm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPsm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPsm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPsm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPsm = fmt.Errorf("proto: unexpected end of group")
)
