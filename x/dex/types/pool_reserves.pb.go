// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/pool_reserves.proto

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

type PoolReserves struct {
	PairID    *PairID                                `protobuf:"bytes,1,opt,name=pairID,proto3" json:"pairID,omitempty"`
	TokenIn   string                                 `protobuf:"bytes,2,opt,name=tokenIn,proto3" json:"tokenIn,omitempty"`
	TickIndex int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Reserves  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=reserves,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserves" yaml:"reserves"`
	Fee       uint64                                 `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *PoolReserves) Reset()         { *m = PoolReserves{} }
func (m *PoolReserves) String() string { return proto.CompactTextString(m) }
func (*PoolReserves) ProtoMessage()    {}
func (*PoolReserves) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37077b416662cb1, []int{0}
}
func (m *PoolReserves) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolReserves) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolReserves.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolReserves) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolReserves.Merge(m, src)
}
func (m *PoolReserves) XXX_Size() int {
	return m.Size()
}
func (m *PoolReserves) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolReserves.DiscardUnknown(m)
}

var xxx_messageInfo_PoolReserves proto.InternalMessageInfo

func (m *PoolReserves) GetPairID() *PairID {
	if m != nil {
		return m.PairID
	}
	return nil
}

func (m *PoolReserves) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *PoolReserves) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *PoolReserves) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*PoolReserves)(nil), "dualitylabs.duality.dex.PoolReserves")
}

func init() { proto.RegisterFile("duality/dex/pool_reserves.proto", fileDescriptor_d37077b416662cb1) }

var fileDescriptor_d37077b416662cb1 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x4f, 0x49, 0xad, 0xd0, 0x2f, 0xc8, 0xcf, 0xcf, 0x89, 0x2f, 0x4a, 0x2d,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x2a, 0xc8,
	0x49, 0x4c, 0x2a, 0xd6, 0x83, 0xb2, 0xf5, 0x52, 0x52, 0x2b, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3,
	0xc1, 0x6a, 0xf4, 0x41, 0x2c, 0x88, 0x72, 0x29, 0x49, 0x14, 0xf3, 0x12, 0x33, 0x8b, 0xe2, 0x33,
	0x53, 0x20, 0x52, 0x4a, 0x7f, 0x18, 0xb9, 0x78, 0x02, 0xf2, 0xf3, 0x73, 0x82, 0xa0, 0x16, 0x08,
	0x99, 0x73, 0xb1, 0x81, 0x54, 0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xc9, 0xeb,
	0xe1, 0xb0, 0x4b, 0x2f, 0x00, 0xac, 0x2c, 0x08, 0xaa, 0x5c, 0x48, 0x82, 0x8b, 0xbd, 0x24, 0x3f,
	0x3b, 0x35, 0xcf, 0x33, 0x4f, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x92, 0xe1,
	0xe2, 0x2c, 0xc9, 0x4c, 0xce, 0xf6, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x56, 0x60, 0xd4, 0x60,
	0x0e, 0x42, 0x08, 0x08, 0x65, 0x72, 0x71, 0xc0, 0x7c, 0x27, 0xc1, 0x02, 0xd2, 0xe8, 0xe4, 0x7b,
	0xe2, 0x9e, 0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf,
	0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x33, 0xaf, 0xe4, 0xd5, 0x3d, 0x79, 0xb8, 0x09, 0x9f, 0xee,
	0xc9, 0xf3, 0x57, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0xc1, 0x44, 0x94, 0x82, 0xe0, 0x92, 0x42, 0x02,
	0x5c, 0xcc, 0x69, 0xa9, 0xa9, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x20, 0xa6, 0x93, 0xeb,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0x23, 0x59, 0x0e, 0xf5, 0xb5,
	0x2e, 0x28, 0x08, 0x60, 0x1c, 0xfd, 0x0a, 0x70, 0x68, 0x82, 0x5d, 0x91, 0xc4, 0x06, 0x0e, 0x4c,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xd7, 0x5d, 0x86, 0xb9, 0x01, 0x00, 0x00,
}

func (m *PoolReserves) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolReserves) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolReserves) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Reserves.Size()
		i -= size
		if _, err := m.Reserves.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolReserves(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.TickIndex != 0 {
		i = encodeVarintPoolReserves(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintPoolReserves(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if m.PairID != nil {
		{
			size, err := m.PairID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPoolReserves(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolReserves(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolReserves(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolReserves) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairID != nil {
		l = m.PairID.Size()
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovPoolReserves(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovPoolReserves(uint64(m.TickIndex))
	}
	l = m.Reserves.Size()
	n += 1 + l + sovPoolReserves(uint64(l))
	if m.Fee != 0 {
		n += 1 + sovPoolReserves(uint64(m.Fee))
	}
	return n
}

func sovPoolReserves(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolReserves(x uint64) (n int) {
	return sovPoolReserves(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolReserves) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolReserves
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
			return fmt.Errorf("proto: PoolReserves: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolReserves: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairID == nil {
				m.PairID = &PairID{}
			}
			if err := m.PairID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
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
				return ErrInvalidLengthPoolReserves
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolReserves
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserves.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolReserves
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoolReserves(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolReserves
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
func skipPoolReserves(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
					return 0, ErrIntOverflowPoolReserves
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
				return 0, ErrInvalidLengthPoolReserves
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolReserves
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolReserves
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolReserves        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolReserves          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolReserves = fmt.Errorf("proto: unexpected end of group")
)
