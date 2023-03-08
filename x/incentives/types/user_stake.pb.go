// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incentives/user_stake.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type UserStake struct {
	Amount    []types.Coin `protobuf:"bytes,1,rep,name=amount,proto3" json:"amount"`
	StartDate uint64       `protobuf:"varint,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   uint64       `protobuf:"varint,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Creator   string       `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *UserStake) Reset()         { *m = UserStake{} }
func (m *UserStake) String() string { return proto.CompactTextString(m) }
func (*UserStake) ProtoMessage()    {}
func (*UserStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_94b9dd2a8efddf90, []int{0}
}
func (m *UserStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStake.Merge(m, src)
}
func (m *UserStake) XXX_Size() int {
	return m.Size()
}
func (m *UserStake) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStake.DiscardUnknown(m)
}

var xxx_messageInfo_UserStake proto.InternalMessageInfo

func (m *UserStake) GetAmount() []types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *UserStake) GetStartDate() uint64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *UserStake) GetEndDate() uint64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

func (m *UserStake) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*UserStake)(nil), "dualitylabs.duality.incentives.UserStake")
}

func init() { proto.RegisterFile("incentives/user_stake.proto", fileDescriptor_94b9dd2a8efddf90) }

var fileDescriptor_94b9dd2a8efddf90 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0xe3, 0xb7, 0x55, 0x5f, 0x35, 0x6c, 0x11, 0x43, 0x28, 0xc8, 0x44, 0x4c, 0x59, 0xb0,
	0xd5, 0x32, 0xb0, 0x17, 0x46, 0xa6, 0x20, 0x16, 0x16, 0xe4, 0xa4, 0xa7, 0x60, 0xd1, 0xf8, 0x2a,
	0xfb, 0x52, 0xd1, 0x6f, 0xc1, 0xc4, 0x67, 0xea, 0xd8, 0x91, 0x09, 0xa1, 0xe4, 0x8b, 0xa0, 0xfc,
	0x53, 0xd9, 0xee, 0xf1, 0xef, 0x7e, 0xb2, 0xee, 0xf1, 0xcf, 0xb5, 0xc9, 0xc0, 0x90, 0xde, 0x82,
	0x93, 0xa5, 0x03, 0xfb, 0xe2, 0x48, 0xbd, 0x81, 0xd8, 0x58, 0x24, 0x0c, 0xf8, 0xaa, 0x54, 0x6b,
	0x4d, 0xbb, 0xb5, 0x4a, 0x9d, 0xe8, 0x67, 0x71, 0x14, 0x66, 0xa7, 0x39, 0xe6, 0xd8, 0xae, 0xca,
	0x66, 0xea, 0xac, 0x19, 0xcf, 0xd0, 0x15, 0xe8, 0x64, 0xaa, 0x1c, 0xc8, 0xed, 0x3c, 0x05, 0x52,
	0x73, 0x99, 0xa1, 0x36, 0x1d, 0xbf, 0xfa, 0x64, 0xfe, 0xf4, 0xc9, 0x81, 0x7d, 0x6c, 0x7e, 0x0a,
	0x6e, 0xfd, 0x89, 0x2a, 0xb0, 0x34, 0x14, 0xb2, 0x68, 0x14, 0x9f, 0x2c, 0xce, 0x44, 0xa7, 0x8b,
	0x46, 0x17, 0xbd, 0x2e, 0xee, 0x50, 0x9b, 0xe5, 0x78, 0xff, 0x7d, 0xe9, 0x25, 0xfd, 0x7a, 0x70,
	0xe1, 0x4f, 0x1d, 0x29, 0x4b, 0xf7, 0x8a, 0x20, 0xfc, 0x17, 0xb1, 0x78, 0x9c, 0x1c, 0x1f, 0x82,
	0xd0, 0xff, 0x0f, 0x66, 0xd5, 0xb2, 0x51, 0xcb, 0x86, 0xd8, 0x90, 0xcc, 0x82, 0x22, 0xb4, 0xe1,
	0x38, 0x62, 0xf1, 0x34, 0x19, 0xe2, 0xf2, 0x61, 0x5f, 0x71, 0x76, 0xa8, 0x38, 0xfb, 0xa9, 0x38,
	0xfb, 0xa8, 0xb9, 0x77, 0xa8, 0xb9, 0xf7, 0x55, 0x73, 0xef, 0x79, 0x91, 0x6b, 0x7a, 0x2d, 0x53,
	0x91, 0x61, 0x21, 0xfb, 0x1e, 0xae, 0x9b, 0x52, 0x86, 0x20, 0xdf, 0xe5, 0x9f, 0x1e, 0x69, 0xb7,
	0x01, 0x97, 0x4e, 0xda, 0x6b, 0x6f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0xb9, 0x8b, 0x2d,
	0x62, 0x01, 0x00, 0x00,
}

func (m *UserStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintUserStake(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if m.EndDate != 0 {
		i = encodeVarintUserStake(dAtA, i, uint64(m.EndDate))
		i--
		dAtA[i] = 0x18
	}
	if m.StartDate != 0 {
		i = encodeVarintUserStake(dAtA, i, uint64(m.StartDate))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUserStake(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserStake(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserStake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovUserStake(uint64(l))
		}
	}
	if m.StartDate != 0 {
		n += 1 + sovUserStake(uint64(m.StartDate))
	}
	if m.EndDate != 0 {
		n += 1 + sovUserStake(uint64(m.EndDate))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovUserStake(uint64(l))
	}
	return n
}

func sovUserStake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserStake(x uint64) (n int) {
	return sovUserStake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserStake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserStake
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
			return fmt.Errorf("proto: UserStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserStake
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
				return ErrInvalidLengthUserStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			m.StartDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			m.EndDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserStake
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
				return ErrInvalidLengthUserStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserStake
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
func skipUserStake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserStake
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
					return 0, ErrIntOverflowUserStake
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
					return 0, ErrIntOverflowUserStake
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
				return 0, ErrInvalidLengthUserStake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserStake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserStake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserStake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserStake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserStake = fmt.Errorf("proto: unexpected end of group")
)
