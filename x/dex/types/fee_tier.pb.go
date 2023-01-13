// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/fee_tier.proto

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

type FeeTier struct {
	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fee int64  `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *FeeTier) Reset()         { *m = FeeTier{} }
func (m *FeeTier) String() string { return proto.CompactTextString(m) }
func (*FeeTier) ProtoMessage()    {}
func (*FeeTier) Descriptor() ([]byte, []int) {
	return fileDescriptor_a606db6eb639052c, []int{0}
}
func (m *FeeTier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeTier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeTier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeTier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeTier.Merge(m, src)
}
func (m *FeeTier) XXX_Size() int {
	return m.Size()
}
func (m *FeeTier) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeTier.DiscardUnknown(m)
}

var xxx_messageInfo_FeeTier proto.InternalMessageInfo

func (m *FeeTier) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FeeTier) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*FeeTier)(nil), "nicholasdotsol.duality.dex.FeeTier")
}

func init() { proto.RegisterFile("duality/dex/fee_tier.proto", fileDescriptor_a606db6eb639052c) }

var fileDescriptor_a606db6eb639052c = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x4f, 0x49, 0xad, 0xd0, 0x4f, 0x4b, 0x4d, 0x8d, 0x2f, 0xc9, 0x4c, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xca, 0xcb, 0x4c, 0xce, 0xc8, 0xcf, 0x49, 0x2c,
	0x4e, 0xc9, 0x2f, 0x29, 0xce, 0xcf, 0xd1, 0x83, 0x2a, 0xd5, 0x4b, 0x49, 0xad, 0x50, 0xd2, 0xe6,
	0x62, 0x77, 0x4b, 0x4d, 0x0d, 0xc9, 0x4c, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0x4e, 0x4b, 0x4d, 0x95,
	0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x9d, 0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x37, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0xdf, 0x0f, 0x6a, 0x9b, 0x4b, 0x7e, 0x49, 0x70, 0x7e, 0x8e, 0x3e, 0xcc, 0x61, 0x15, 0x60, 0xa7,
	0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x5a, 0xb9, 0x6b, 0xe8, 0xb6, 0x00, 0x00, 0x00,
}

func (m *FeeTier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeTier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeTier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != 0 {
		i = encodeVarintFeeTier(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintFeeTier(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeeTier(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeeTier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeTier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFeeTier(uint64(m.Id))
	}
	if m.Fee != 0 {
		n += 1 + sovFeeTier(uint64(m.Fee))
	}
	return n
}

func sovFeeTier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeeTier(x uint64) (n int) {
	return sovFeeTier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeTier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeeTier
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
			return fmt.Errorf("proto: FeeTier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeTier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeTier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeTier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFeeTier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeeTier
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
func skipFeeTier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeeTier
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
					return 0, ErrIntOverflowFeeTier
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
					return 0, ErrIntOverflowFeeTier
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
				return 0, ErrInvalidLengthFeeTier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeeTier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeeTier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeeTier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeeTier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeeTier = fmt.Errorf("proto: unexpected end of group")
)
