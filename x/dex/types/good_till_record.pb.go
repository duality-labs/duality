// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/good_till_record.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GoodTillRecord struct {
	// see limitOrderTranche.proto for details on goodTillDate
	GoodTillDate time.Time `protobuf:"bytes,1,opt,name=GoodTillDate,proto3,stdtime" json:"GoodTillDate"`
	TrancheRef   []byte    `protobuf:"bytes,2,opt,name=trancheRef,proto3" json:"trancheRef,omitempty"`
}

func (m *GoodTillRecord) Reset()         { *m = GoodTillRecord{} }
func (m *GoodTillRecord) String() string { return proto.CompactTextString(m) }
func (*GoodTillRecord) ProtoMessage()    {}
func (*GoodTillRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f0c5e2cff15201, []int{0}
}
func (m *GoodTillRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodTillRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodTillRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodTillRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodTillRecord.Merge(m, src)
}
func (m *GoodTillRecord) XXX_Size() int {
	return m.Size()
}
func (m *GoodTillRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodTillRecord.DiscardUnknown(m)
}

var xxx_messageInfo_GoodTillRecord proto.InternalMessageInfo

func (m *GoodTillRecord) GetGoodTillDate() time.Time {
	if m != nil {
		return m.GoodTillDate
	}
	return time.Time{}
}

func (m *GoodTillRecord) GetTrancheRef() []byte {
	if m != nil {
		return m.TrancheRef
	}
	return nil
}

func init() {
	proto.RegisterType((*GoodTillRecord)(nil), "dualitylabs.duality.dex.GoodTillRecord")
}

func init() {
	proto.RegisterFile("duality/dex/good_till_record.proto", fileDescriptor_33f0c5e2cff15201)
}

var fileDescriptor_33f0c5e2cff15201 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xcd, 0x7a, 0x10, 0x59, 0x8b, 0x87, 0x20, 0x58, 0x72, 0xd8, 0x94, 0x9e, 0x0a, 0xe2, 0x2e,
	0xe8, 0x1f, 0x14, 0x45, 0xcf, 0xa1, 0x27, 0x2f, 0x65, 0x93, 0xdd, 0x6e, 0x17, 0x36, 0x4e, 0x48,
	0x26, 0x90, 0xfa, 0x15, 0xfd, 0xac, 0x1e, 0x7b, 0xf4, 0xa4, 0x92, 0xfc, 0x88, 0x74, 0x93, 0x80,
	0xbd, 0xbd, 0x79, 0xbc, 0xf7, 0x66, 0xde, 0xd0, 0xb9, 0xaa, 0xa5, 0xb3, 0xb8, 0x13, 0x4a, 0x37,
	0xc2, 0x00, 0xa8, 0x35, 0x5a, 0xe7, 0xd6, 0xa5, 0xce, 0xa0, 0x54, 0xbc, 0x28, 0x01, 0x21, 0xbc,
	0x1b, 0x34, 0x4e, 0xa6, 0x15, 0x1f, 0x30, 0x57, 0xba, 0x89, 0x62, 0x03, 0x60, 0x9c, 0x16, 0x5e,
	0x96, 0xd6, 0x1b, 0x81, 0x36, 0xd7, 0x15, 0xca, 0xbc, 0xe8, 0x9d, 0xd1, 0xad, 0x01, 0x03, 0x1e,
	0x8a, 0x13, 0xea, 0xd9, 0xf9, 0x27, 0xbd, 0x79, 0x05, 0x50, 0x2b, 0xeb, 0x5c, 0xe2, 0xf7, 0x84,
	0x6f, 0x74, 0x32, 0x32, 0xcf, 0x12, 0xf5, 0x94, 0xcc, 0xc8, 0xe2, 0xfa, 0x31, 0xe2, 0x7d, 0x3e,
	0x1f, 0xf3, 0xf9, 0x6a, 0xcc, 0x5f, 0x5e, 0x1d, 0xbe, 0xe3, 0x60, 0xff, 0x13, 0x93, 0xe4, 0xcc,
	0x19, 0x32, 0x4a, 0xb1, 0x94, 0x1f, 0xd9, 0x56, 0x27, 0x7a, 0x33, 0xbd, 0x98, 0x91, 0xc5, 0x24,
	0xf9, 0xc7, 0x2c, 0x5f, 0x0e, 0x2d, 0x23, 0xc7, 0x96, 0x91, 0xdf, 0x96, 0x91, 0x7d, 0xc7, 0x82,
	0x63, 0xc7, 0x82, 0xaf, 0x8e, 0x05, 0xef, 0xf7, 0xc6, 0xe2, 0xb6, 0x4e, 0x79, 0x06, 0xb9, 0x18,
	0x4a, 0x3e, 0x9c, 0x1a, 0x8f, 0x83, 0x68, 0xfc, 0x8f, 0x70, 0x57, 0xe8, 0x2a, 0xbd, 0xf4, 0x27,
	0x3d, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xdb, 0xbb, 0x84, 0x3f, 0x01, 0x00, 0x00,
}

func (m *GoodTillRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodTillRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodTillRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TrancheRef) > 0 {
		i -= len(m.TrancheRef)
		copy(dAtA[i:], m.TrancheRef)
		i = encodeVarintGoodTillRecord(dAtA, i, uint64(len(m.TrancheRef)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.GoodTillDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.GoodTillDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGoodTillRecord(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGoodTillRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovGoodTillRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoodTillRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.GoodTillDate)
	n += 1 + l + sovGoodTillRecord(uint64(l))
	l = len(m.TrancheRef)
	if l > 0 {
		n += 1 + l + sovGoodTillRecord(uint64(l))
	}
	return n
}

func sovGoodTillRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGoodTillRecord(x uint64) (n int) {
	return sovGoodTillRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoodTillRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoodTillRecord
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
			return fmt.Errorf("proto: GoodTillRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodTillRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodTillDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoodTillRecord
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
				return ErrInvalidLengthGoodTillRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGoodTillRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.GoodTillDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheRef", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoodTillRecord
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
				return ErrInvalidLengthGoodTillRecord
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGoodTillRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrancheRef = append(m.TrancheRef[:0], dAtA[iNdEx:postIndex]...)
			if m.TrancheRef == nil {
				m.TrancheRef = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoodTillRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGoodTillRecord
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
func skipGoodTillRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoodTillRecord
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
					return 0, ErrIntOverflowGoodTillRecord
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
					return 0, ErrIntOverflowGoodTillRecord
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
				return 0, ErrInvalidLengthGoodTillRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGoodTillRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGoodTillRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGoodTillRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoodTillRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGoodTillRecord = fmt.Errorf("proto: unexpected end of group")
)
