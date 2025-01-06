// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/auction/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the auction module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// list of auctions
	Auctions []Auction `protobuf:"bytes,2,rep,name=auctions,proto3" json:"auctions"`
	// list of all bid by address
	BidByAddress          []BidByAddress `protobuf:"bytes,3,rep,name=bid_by_address,json=bidByAddress,proto3" json:"bid_by_address"`
	AuctionSequence       uint64         `protobuf:"varint,4,opt,name=auction_sequence,json=auctionSequence,proto3" json:"auction_sequence,omitempty"`
	LastestAuctionPeriods int64          `protobuf:"varint,5,opt,name=lastest_auction_periods,json=lastestAuctionPeriods,proto3" json:"lastest_auction_periods,omitempty"`
	BidSequences          []BidSequence  `protobuf:"bytes,6,rep,name=bid_sequences,json=bidSequences,proto3" json:"bid_sequences"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e716c21f756a4f6, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAuctions() []Auction {
	if m != nil {
		return m.Auctions
	}
	return nil
}

func (m *GenesisState) GetBidByAddress() []BidByAddress {
	if m != nil {
		return m.BidByAddress
	}
	return nil
}

func (m *GenesisState) GetAuctionSequence() uint64 {
	if m != nil {
		return m.AuctionSequence
	}
	return 0
}

func (m *GenesisState) GetLastestAuctionPeriods() int64 {
	if m != nil {
		return m.LastestAuctionPeriods
	}
	return 0
}

func (m *GenesisState) GetBidSequences() []BidSequence {
	if m != nil {
		return m.BidSequences
	}
	return nil
}

type BidSequence struct {
	AuctionId uint64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BidSequence) Reset()         { *m = BidSequence{} }
func (m *BidSequence) String() string { return proto.CompactTextString(m) }
func (*BidSequence) ProtoMessage()    {}
func (*BidSequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e716c21f756a4f6, []int{1}
}
func (m *BidSequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidSequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidSequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidSequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidSequence.Merge(m, src)
}
func (m *BidSequence) XXX_Size() int {
	return m.Size()
}
func (m *BidSequence) XXX_DiscardUnknown() {
	xxx_messageInfo_BidSequence.DiscardUnknown(m)
}

var xxx_messageInfo_BidSequence proto.InternalMessageInfo

func (m *BidSequence) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *BidSequence) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "reserve.auction.v1.GenesisState")
	proto.RegisterType((*BidSequence)(nil), "reserve.auction.v1.BidSequence")
}

func init() { proto.RegisterFile("reserve/auction/v1/genesis.proto", fileDescriptor_3e716c21f756a4f6) }

var fileDescriptor_3e716c21f756a4f6 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbb, 0x8e, 0x9b, 0x40,
	0x14, 0x86, 0x19, 0xe3, 0x58, 0xf6, 0xd8, 0xb9, 0x8d, 0x12, 0x05, 0x11, 0x05, 0x23, 0x57, 0x24,
	0x05, 0xc4, 0x8e, 0x94, 0x2e, 0x85, 0x69, 0x72, 0x69, 0xe2, 0xe0, 0x2e, 0x0d, 0xe2, 0x32, 0x22,
	0x23, 0x19, 0x86, 0x30, 0x63, 0x2b, 0xbc, 0x45, 0xca, 0x3c, 0x42, 0xca, 0x7d, 0x0c, 0x97, 0x2e,
	0xb7, 0x5a, 0xad, 0xec, 0x62, 0x5f, 0x63, 0xe5, 0x61, 0x40, 0x68, 0x97, 0x6d, 0xac, 0xe3, 0xff,
	0xfc, 0xe7, 0x3b, 0xff, 0x88, 0x03, 0xcd, 0x02, 0x33, 0x5c, 0xec, 0xb0, 0x13, 0x6c, 0x23, 0x4e,
	0x68, 0xe6, 0xec, 0xe6, 0x4e, 0x82, 0x33, 0xcc, 0x08, 0xb3, 0xf3, 0x82, 0x72, 0x8a, 0x90, 0x74,
	0xd8, 0xd2, 0x61, 0xef, 0xe6, 0xfa, 0xf3, 0x20, 0x25, 0x19, 0x75, 0xc4, 0x6f, 0x65, 0xd3, 0x5f,
	0x24, 0x34, 0xa1, 0xa2, 0x74, 0xce, 0x95, 0x54, 0xa7, 0x1d, 0xf8, 0x3c, 0x28, 0x82, 0x54, 0xd2,
	0xf5, 0xae, 0xfd, 0xf5, 0x22, 0xe1, 0x98, 0xfd, 0x53, 0xe1, 0xe4, 0x73, 0x95, 0x68, 0xcd, 0x03,
	0x8e, 0xd1, 0x27, 0x38, 0xa8, 0x10, 0x1a, 0x30, 0x81, 0x35, 0x5e, 0xe8, 0xf6, 0xfd, 0x84, 0xf6,
	0x4a, 0x38, 0xdc, 0xd1, 0xfe, 0x6a, 0xaa, 0xfc, 0xbf, 0xb9, 0x78, 0x07, 0x3c, 0x39, 0x84, 0x5c,
	0x38, 0x94, 0x3e, 0xa6, 0xf5, 0x4c, 0xd5, 0x1a, 0x2f, 0x5e, 0x77, 0x01, 0x96, 0x55, 0xd9, 0x26,
	0x34, 0x73, 0xe8, 0x07, 0x7c, 0x12, 0x92, 0xd8, 0x0f, 0x4b, 0x3f, 0x88, 0xe3, 0x02, 0x33, 0xa6,
	0xa9, 0x82, 0x64, 0x76, 0x91, 0x5c, 0x12, 0xbb, 0xe5, 0xb2, 0xf2, 0xb5, 0x71, 0x93, 0xb0, 0xd5,
	0x40, 0x6f, 0xe1, 0x33, 0x39, 0xe3, 0x33, 0xfc, 0x7b, 0x8b, 0xb3, 0x08, 0x6b, 0x7d, 0x13, 0x58,
	0x7d, 0xef, 0xa9, 0xd4, 0xd7, 0x52, 0x46, 0x1f, 0xe1, 0xab, 0x4d, 0xc0, 0x38, 0x66, 0xdc, 0xaf,
	0x47, 0x72, 0x5c, 0x10, 0x1a, 0x33, 0xed, 0x91, 0x09, 0x2c, 0xd5, 0x7b, 0x29, 0xdb, 0xf2, 0x0d,
	0xab, 0xaa, 0x89, 0xbe, 0xc3, 0xc7, 0xe7, 0xd4, 0x35, 0x9e, 0x69, 0x03, 0x11, 0x7a, 0xfa, 0x40,
	0xe8, 0x7a, 0xdf, 0xdd, 0xcc, 0xb5, 0xce, 0x66, 0x5f, 0xe0, 0xb8, 0xe5, 0x43, 0x6f, 0x20, 0xac,
	0xf3, 0x90, 0x58, 0x7c, 0x9c, 0xbe, 0x37, 0x92, 0xca, 0xd7, 0x18, 0xe9, 0x70, 0xd8, 0xbc, 0xac,
	0x27, 0x9a, 0xcd, 0x7f, 0xf7, 0xdb, 0xfe, 0x68, 0x80, 0xc3, 0xd1, 0x00, 0xd7, 0x47, 0x03, 0xfc,
	0x3d, 0x19, 0xca, 0xe1, 0x64, 0x28, 0x97, 0x27, 0x43, 0xf9, 0xf9, 0x3e, 0x21, 0xfc, 0xd7, 0x36,
	0xb4, 0x23, 0x9a, 0x3a, 0x34, 0xa3, 0x69, 0x29, 0xae, 0x22, 0xa2, 0x1b, 0xa7, 0xbe, 0x9c, 0x3f,
	0xcd, 0xed, 0xf0, 0x32, 0xc7, 0x2c, 0x1c, 0x08, 0xc7, 0x87, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x39, 0x85, 0x50, 0xfa, 0xdb, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidSequences) > 0 {
		for iNdEx := len(m.BidSequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BidSequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.LastestAuctionPeriods != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastestAuctionPeriods))
		i--
		dAtA[i] = 0x28
	}
	if m.AuctionSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionSequence))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BidByAddress) > 0 {
		for iNdEx := len(m.BidByAddress) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BidByAddress[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BidSequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidSequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidSequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if m.AuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BidByAddress) > 0 {
		for _, e := range m.BidByAddress {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.AuctionSequence != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionSequence))
	}
	if m.LastestAuctionPeriods != 0 {
		n += 1 + sovGenesis(uint64(m.LastestAuctionPeriods))
	}
	if len(m.BidSequences) > 0 {
		for _, e := range m.BidSequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *BidSequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionId))
	}
	if m.Sequence != 0 {
		n += 1 + sovGenesis(uint64(m.Sequence))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, Auction{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidByAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidByAddress = append(m.BidByAddress, BidByAddress{})
			if err := m.BidByAddress[len(m.BidByAddress)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionSequence", wireType)
			}
			m.AuctionSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastestAuctionPeriods", wireType)
			}
			m.LastestAuctionPeriods = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastestAuctionPeriods |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidSequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidSequences = append(m.BidSequences, BidSequence{})
			if err := m.BidSequences[len(m.BidSequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *BidSequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: BidSequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidSequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
