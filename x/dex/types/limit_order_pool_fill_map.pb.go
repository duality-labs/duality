// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/limit_order_pool_fill_map.proto

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

type LimitOrderTrancheFillMap struct {
	PairId         string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token          string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex      int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Count          uint64                                 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	FilledReserves github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=filledReserves,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"filledReserves" yaml:"fill"`
}

func (m *LimitOrderTrancheFillMap) Reset()         { *m = LimitOrderTrancheFillMap{} }
func (m *LimitOrderTrancheFillMap) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheFillMap) ProtoMessage()    {}
func (*LimitOrderTrancheFillMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_20d7a0c99149927b, []int{0}
}
func (m *LimitOrderTrancheFillMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheFillMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheFillMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheFillMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheFillMap.Merge(m, src)
}
func (m *LimitOrderTrancheFillMap) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheFillMap) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheFillMap.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheFillMap proto.InternalMessageInfo

func (m *LimitOrderTrancheFillMap) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *LimitOrderTrancheFillMap) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LimitOrderTrancheFillMap) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *LimitOrderTrancheFillMap) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*LimitOrderTrancheFillMap)(nil), "nicholasdotsol.duality.dex.LimitOrderTrancheFillMap")
}

func init() {
	proto.RegisterFile("dex/limit_order_pool_fill_map.proto", fileDescriptor_20d7a0c99149927b)
}

var fileDescriptor_20d7a0c99149927b = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x33, 0x5f, 0xff, 0x40, 0xf3, 0x81, 0x8b, 0x50, 0x25, 0x14, 0x49, 0x4a, 0x05, 0xe9,
	0xa6, 0x99, 0x85, 0x3b, 0x97, 0xa5, 0x28, 0x05, 0xb5, 0x12, 0x77, 0x6e, 0x42, 0x9a, 0x19, 0xdb,
	0xa1, 0x37, 0xb9, 0x21, 0x33, 0xd5, 0xf4, 0x2d, 0x7c, 0xac, 0x2e, 0xbb, 0x14, 0x17, 0x41, 0xda,
	0x9d, 0xcb, 0x3e, 0x81, 0x4c, 0x1a, 0x50, 0xba, 0x9a, 0x39, 0x67, 0xe6, 0xfe, 0xce, 0x70, 0xc6,
	0xbc, 0x60, 0x3c, 0xa7, 0x20, 0x62, 0xa1, 0x02, 0xcc, 0x18, 0xcf, 0x82, 0x14, 0x11, 0x82, 0x17,
	0x01, 0x10, 0xc4, 0x61, 0xea, 0xa5, 0x19, 0x2a, 0xb4, 0x3a, 0x89, 0x88, 0xe6, 0x08, 0xa1, 0x64,
	0xa8, 0x24, 0x82, 0xc7, 0x96, 0x21, 0x08, 0xb5, 0xf2, 0x18, 0xcf, 0x3b, 0xed, 0x19, 0xce, 0xb0,
	0xbc, 0x46, 0xf5, 0xee, 0x30, 0xd1, 0xdb, 0x13, 0xf3, 0xf4, 0x4e, 0x53, 0x27, 0x1a, 0xfa, 0x88,
	0x08, 0x37, 0x02, 0xe0, 0x3e, 0x4c, 0xad, 0x33, 0xb3, 0x99, 0x86, 0x22, 0x1b, 0x33, 0x9b, 0x74,
	0x49, 0xbf, 0xe5, 0x57, 0xca, 0x6a, 0x9b, 0x0d, 0x85, 0x0b, 0x9e, 0xd8, 0xff, 0x4a, 0xfb, 0x20,
	0xac, 0x73, 0xb3, 0xa5, 0x44, 0xb4, 0x18, 0x27, 0x8c, 0xe7, 0x76, 0xad, 0x4b, 0xfa, 0x35, 0xff,
	0xd7, 0xd0, 0x33, 0x11, 0x2e, 0x13, 0x65, 0xd7, 0xbb, 0xa4, 0x5f, 0xf7, 0x0f, 0xc2, 0x7a, 0x33,
	0x4f, 0xf4, 0xfb, 0x39, 0xf3, 0xb9, 0xe4, 0xd9, 0x2b, 0x97, 0x76, 0x43, 0x23, 0x87, 0x93, 0x75,
	0xe1, 0x1a, 0x9f, 0x85, 0x7b, 0x39, 0x13, 0x6a, 0xbe, 0x9c, 0x7a, 0x11, 0xc6, 0x34, 0x42, 0x19,
	0xa3, 0xac, 0x96, 0x81, 0x64, 0x0b, 0xaa, 0x56, 0x29, 0x97, 0xde, 0x88, 0x47, 0xdf, 0x85, 0x7b,
	0xc4, 0xd9, 0x17, 0xee, 0xff, 0x55, 0x18, 0xc3, 0x75, 0x4f, 0xfb, 0x3d, 0xff, 0xe8, 0x78, 0x78,
	0xbb, 0xde, 0x3a, 0x64, 0xb3, 0x75, 0xc8, 0xd7, 0xd6, 0x21, 0xef, 0x3b, 0xc7, 0xd8, 0xec, 0x1c,
	0xe3, 0x63, 0xe7, 0x18, 0xcf, 0x83, 0x3f, 0x91, 0x0f, 0x55, 0x97, 0x23, 0x54, 0x4f, 0x08, 0xb4,
	0xea, 0x92, 0xe6, 0x54, 0xff, 0x44, 0x99, 0x3e, 0x6d, 0x96, 0x25, 0x5e, 0xfd, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xda, 0x83, 0x17, 0xdc, 0x9d, 0x01, 0x00, 0x00,
}

func (m *LimitOrderTrancheFillMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheFillMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheFillMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FilledReserves.Size()
		i -= size
		if _, err := m.FilledReserves.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheFillMap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Count != 0 {
		i = encodeVarintLimitOrderTrancheFillMap(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.TickIndex != 0 {
		i = encodeVarintLimitOrderTrancheFillMap(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLimitOrderTrancheFillMap(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintLimitOrderTrancheFillMap(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTrancheFillMap(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTrancheFillMap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheFillMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheFillMap(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheFillMap(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderTrancheFillMap(uint64(m.TickIndex))
	}
	if m.Count != 0 {
		n += 1 + sovLimitOrderTrancheFillMap(uint64(m.Count))
	}
	l = m.FilledReserves.Size()
	n += 1 + l + sovLimitOrderTrancheFillMap(uint64(l))
	return n
}

func sovLimitOrderTrancheFillMap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTrancheFillMap(x uint64) (n int) {
	return sovLimitOrderTrancheFillMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheFillMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTrancheFillMap
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
			return fmt.Errorf("proto: LimitOrderTrancheFillMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheFillMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheFillMap
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
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheFillMap
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
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheFillMap
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheFillMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilledReserves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheFillMap
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
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheFillMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FilledReserves.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTrancheFillMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTrancheFillMap
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
func skipLimitOrderTrancheFillMap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTrancheFillMap
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
					return 0, ErrIntOverflowLimitOrderTrancheFillMap
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
					return 0, ErrIntOverflowLimitOrderTrancheFillMap
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
				return 0, ErrInvalidLengthLimitOrderTrancheFillMap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTrancheFillMap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTrancheFillMap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTrancheFillMap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTrancheFillMap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTrancheFillMap = fmt.Errorf("proto: unexpected end of group")
)
