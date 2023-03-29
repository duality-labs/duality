// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/deposit_record.proto

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

type DepositRecord struct {
	PairID          *PairID                                `protobuf:"bytes,1,opt,name=pairID,proto3" json:"pairID,omitempty"`
	SharesOwned     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalShares" yaml:"totalShares"`
	CenterTickIndex int64                                  `protobuf:"varint,3,opt,name=centerTickIndex,proto3" json:"centerTickIndex,omitempty"`
	LowerTickIndex  int64                                  `protobuf:"varint,4,opt,name=lowerTickIndex,proto3" json:"lowerTickIndex,omitempty"`
	UpperTickIndex  int64                                  `protobuf:"varint,5,opt,name=upperTickIndex,proto3" json:"upperTickIndex,omitempty"`
	Fee             uint64                                 `protobuf:"varint,6,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *DepositRecord) Reset()         { *m = DepositRecord{} }
func (m *DepositRecord) String() string { return proto.CompactTextString(m) }
func (*DepositRecord) ProtoMessage()    {}
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7c76d21a5add72, []int{0}
}
func (m *DepositRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositRecord.Merge(m, src)
}
func (m *DepositRecord) XXX_Size() int {
	return m.Size()
}
func (m *DepositRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DepositRecord proto.InternalMessageInfo

func (m *DepositRecord) GetPairID() *PairID {
	if m != nil {
		return m.PairID
	}
	return nil
}

func (m *DepositRecord) GetCenterTickIndex() int64 {
	if m != nil {
		return m.CenterTickIndex
	}
	return 0
}

func (m *DepositRecord) GetLowerTickIndex() int64 {
	if m != nil {
		return m.LowerTickIndex
	}
	return 0
}

func (m *DepositRecord) GetUpperTickIndex() int64 {
	if m != nil {
		return m.UpperTickIndex
	}
	return 0
}

func (m *DepositRecord) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*DepositRecord)(nil), "dualitylabs.duality.dex.DepositRecord")
}

func init() { proto.RegisterFile("duality/dex/deposit_record.proto", fileDescriptor_9d7c76d21a5add72) }

var fileDescriptor_9d7c76d21a5add72 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x73, 0x6d, 0x2d, 0x78, 0xc5, 0x1f, 0x1c, 0x82, 0xb1, 0x43, 0x12, 0x3a, 0x94, 0x80,
	0xf4, 0x02, 0x3a, 0x08, 0x8e, 0x45, 0x87, 0x4e, 0x4a, 0xea, 0xe4, 0x52, 0xd2, 0xdc, 0xd7, 0x36,
	0x34, 0xcd, 0x85, 0xdc, 0x85, 0xa6, 0xff, 0x85, 0x7f, 0x56, 0xc1, 0xa5, 0xa3, 0x38, 0x04, 0x69,
	0x37, 0x47, 0xff, 0x02, 0xc9, 0x35, 0x42, 0x2c, 0x38, 0xe5, 0xe5, 0xf1, 0xe1, 0x3d, 0xee, 0x7d,
	0xb1, 0xc5, 0x52, 0x2f, 0x0c, 0xe4, 0xd2, 0x61, 0x90, 0x39, 0x0c, 0x62, 0x2e, 0x02, 0x39, 0x4a,
	0xc0, 0xe7, 0x09, 0xa3, 0x71, 0xc2, 0x25, 0x27, 0xe7, 0x25, 0x11, 0x7a, 0x63, 0x41, 0x4b, 0x4d,
	0x19, 0x64, 0xed, 0xb3, 0x09, 0x9f, 0x70, 0xc5, 0x38, 0x85, 0xda, 0xe1, 0xed, 0x8b, 0x6a, 0x60,
	0xec, 0x05, 0xc9, 0x28, 0x28, 0x93, 0x3a, 0x6f, 0x35, 0x7c, 0x74, 0xb7, 0xab, 0x70, 0x55, 0x03,
	0xb9, 0xc1, 0xcd, 0x02, 0x19, 0x30, 0x1d, 0x59, 0xc8, 0x6e, 0x5d, 0x99, 0xf4, 0x9f, 0x32, 0xfa,
	0xa8, 0x30, 0xb7, 0xc4, 0x49, 0x8a, 0x5b, 0x62, 0xea, 0x25, 0x20, 0x1e, 0x16, 0x11, 0x30, 0xbd,
	0x66, 0x21, 0xfb, 0xb0, 0x3f, 0x5c, 0xe5, 0xa6, 0xf6, 0x91, 0x9b, 0xdd, 0x49, 0x20, 0xa7, 0xe9,
	0x98, 0xfa, 0x7c, 0xee, 0xf8, 0x5c, 0xcc, 0xb9, 0x28, 0x3f, 0x3d, 0xc1, 0x66, 0x8e, 0x5c, 0xc6,
	0x20, 0xe8, 0x20, 0x92, 0x5f, 0xb9, 0xd9, 0x92, 0x5c, 0x7a, 0xe1, 0x50, 0x25, 0x7d, 0xe7, 0x26,
	0x59, 0x7a, 0xf3, 0xf0, 0xb6, 0x53, 0x31, 0x3b, 0x6e, 0xb5, 0x87, 0xd8, 0xf8, 0xc4, 0x87, 0x48,
	0x42, 0xf2, 0x14, 0xf8, 0xb3, 0x41, 0xc4, 0x20, 0xd3, 0xeb, 0x16, 0xb2, 0xeb, 0xee, 0xbe, 0x4d,
	0xba, 0xf8, 0x38, 0xe4, 0x8b, 0x2a, 0xd8, 0x50, 0xe0, 0x9e, 0x5b, 0x70, 0x69, 0x1c, 0x57, 0xb9,
	0x83, 0x1d, 0xf7, 0xd7, 0x25, 0xa7, 0xb8, 0xfe, 0x02, 0xa0, 0x37, 0x2d, 0x64, 0x37, 0xdc, 0x42,
	0xf6, 0xef, 0x57, 0x1b, 0x03, 0xad, 0x37, 0x06, 0xfa, 0xdc, 0x18, 0xe8, 0x75, 0x6b, 0x68, 0xeb,
	0xad, 0xa1, 0xbd, 0x6f, 0x0d, 0xed, 0xf9, 0xb2, 0xf2, 0xfe, 0x72, 0xc3, 0x5e, 0x31, 0xe8, 0xef,
	0x8f, 0x93, 0xa9, 0xe3, 0xa8, 0x21, 0xc6, 0x4d, 0x75, 0x9b, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0xf9, 0xf2, 0x4c, 0x09, 0x02, 0x00, 0x00,
}

func (m *DepositRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x30
	}
	if m.UpperTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.UpperTickIndex))
		i--
		dAtA[i] = 0x28
	}
	if m.LowerTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.LowerTickIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.CenterTickIndex != 0 {
		i = encodeVarintDepositRecord(dAtA, i, uint64(m.CenterTickIndex))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDepositRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PairID != nil {
		{
			size, err := m.PairID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDepositRecord(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairID != nil {
		l = m.PairID.Size()
		n += 1 + l + sovDepositRecord(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovDepositRecord(uint64(l))
	if m.CenterTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.CenterTickIndex))
	}
	if m.LowerTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.LowerTickIndex))
	}
	if m.UpperTickIndex != 0 {
		n += 1 + sovDepositRecord(uint64(m.UpperTickIndex))
	}
	if m.Fee != 0 {
		n += 1 + sovDepositRecord(uint64(m.Fee))
	}
	return n
}

func sovDepositRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositRecord(x uint64) (n int) {
	return sovDepositRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositRecord
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
			return fmt.Errorf("proto: DepositRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
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
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
				return ErrInvalidLengthDepositRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CenterTickIndex", wireType)
			}
			m.CenterTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CenterTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerTickIndex", wireType)
			}
			m.LowerTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LowerTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperTickIndex", wireType)
			}
			m.UpperTickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpperTickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositRecord
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
			skippy, err := skipDepositRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositRecord
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
func skipDepositRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
					return 0, ErrIntOverflowDepositRecord
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
				return 0, ErrInvalidLengthDepositRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositRecord = fmt.Errorf("proto: unexpected end of group")
)
