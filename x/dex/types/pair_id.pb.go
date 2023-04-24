// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/pair_id.proto

package types

import (
	fmt "fmt"
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

type PairID struct {
	Token0 string `protobuf:"bytes,1,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1 string `protobuf:"bytes,2,opt,name=token1,proto3" json:"token1,omitempty"`
}

func (m *PairID) Reset()         { *m = PairID{} }
func (m *PairID) String() string { return proto.CompactTextString(m) }
func (*PairID) ProtoMessage()    {}
func (*PairID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1919813da3dc14c8, []int{0}
}
func (m *PairID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairID.Merge(m, src)
}
func (m *PairID) XXX_Size() int {
	return m.Size()
}
func (m *PairID) XXX_DiscardUnknown() {
	xxx_messageInfo_PairID.DiscardUnknown(m)
}

var xxx_messageInfo_PairID proto.InternalMessageInfo

func (m *PairID) GetToken0() string {
	if m != nil {
		return m.Token0
	}
	return ""
}

func (m *PairID) GetToken1() string {
	if m != nil {
		return m.Token1
	}
	return ""
}

func init() {
	proto.RegisterType((*PairID)(nil), "dualitylabs.duality.dex.PairID")
}

func init() { proto.RegisterFile("duality/dex/pair_id.proto", fileDescriptor_1919813da3dc14c8) }

var fileDescriptor_1919813da3dc14c8 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x4f, 0x49, 0xad, 0xd0, 0x2f, 0x48, 0xcc, 0x2c, 0x8a, 0xcf, 0x4c, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x4a, 0xe5, 0x24, 0x26, 0x15, 0xeb, 0x41, 0xd9,
	0x7a, 0x29, 0xa9, 0x15, 0x4a, 0x16, 0x5c, 0x6c, 0x01, 0x89, 0x99, 0x45, 0x9e, 0x2e, 0x42, 0x62,
	0x5c, 0x6c, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x06, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50,
	0x1e, 0x5c, 0xdc, 0x50, 0x82, 0x09, 0x49, 0xdc, 0xd0, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0xa1, 0x76, 0xe9, 0x82, 0x2c, 0x86, 0x71, 0xf4, 0x2b, 0xc0, 0x2e, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x31, 0xc9, 0x3d, 0x8b,
	0xbd, 0x00, 0x00, 0x00,
}

func (m *PairID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintPairId(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintPairId(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPairId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPairId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PairID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovPairId(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovPairId(uint64(l))
	}
	return n
}

func sovPairId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPairId(x uint64) (n int) {
	return sovPairId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PairID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPairId
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
			return fmt.Errorf("proto: PairID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairId
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
				return ErrInvalidLengthPairId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairId
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
				return ErrInvalidLengthPairId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPairId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPairId
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
func skipPairId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPairId
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
					return 0, ErrIntOverflowPairId
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
					return 0, ErrIntOverflowPairId
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
				return 0, ErrInvalidLengthPairId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPairId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPairId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPairId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPairId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPairId = fmt.Errorf("proto: unexpected end of group")
)
