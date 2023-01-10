// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/limit_order_tranche_user.proto

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

type LimitOrderTrancheUser struct {
	PairId          string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token           string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex       int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Count           uint64                                 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Address         string                                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	SharesOwned     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesOwned" yaml:"sharesOwned"`
	SharesWithdrawn github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=sharesWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesWithdrawn" yaml:"sharesWithdrawn"`
	SharesCancelled github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=sharesCancelled,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesCancelled" yaml:"sharesCancelled"`
}

func (m *LimitOrderTrancheUser) Reset()         { *m = LimitOrderTrancheUser{} }
func (m *LimitOrderTrancheUser) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheUser) ProtoMessage()    {}
func (*LimitOrderTrancheUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_896ef6766d8ff3c4, []int{0}
}
func (m *LimitOrderTrancheUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheUser.Merge(m, src)
}
func (m *LimitOrderTrancheUser) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheUser.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheUser proto.InternalMessageInfo

func (m *LimitOrderTrancheUser) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *LimitOrderTrancheUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LimitOrderTrancheUser) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *LimitOrderTrancheUser) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LimitOrderTrancheUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*LimitOrderTrancheUser)(nil), "nicholasdotsol.duality.dex.LimitOrderTrancheUser")
}

func init() {
	proto.RegisterFile("duality/dex/limit_order_tranche_user.proto", fileDescriptor_896ef6766d8ff3c4)
}

var fileDescriptor_896ef6766d8ff3c4 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x3f, 0x8b, 0xdb, 0x30,
	0x1c, 0xb5, 0x7a, 0xb9, 0x5c, 0x4f, 0x1d, 0x0a, 0xe2, 0x7a, 0x88, 0xa3, 0xd8, 0xc1, 0x43, 0x09,
	0x85, 0xb3, 0x87, 0x6e, 0x1d, 0xd3, 0x42, 0x09, 0x94, 0x06, 0x9c, 0x96, 0x42, 0x3b, 0x04, 0xc5,
	0x12, 0xb1, 0x88, 0x2c, 0x05, 0x49, 0x26, 0xce, 0x07, 0xe8, 0xde, 0xb9, 0x9f, 0x28, 0x63, 0xc6,
	0xd2, 0xc1, 0x94, 0x64, 0xeb, 0x98, 0x4f, 0x50, 0xfc, 0x27, 0xae, 0xc9, 0x56, 0x6e, 0x92, 0xde,
	0xfb, 0xfd, 0x7e, 0xef, 0xbd, 0xe1, 0xc1, 0x97, 0x34, 0x23, 0x82, 0xdb, 0x4d, 0x48, 0x59, 0x1e,
	0x0a, 0x9e, 0x72, 0x3b, 0x53, 0x9a, 0x32, 0x3d, 0xb3, 0x9a, 0xc8, 0x38, 0x61, 0xb3, 0xcc, 0x30,
	0x1d, 0xac, 0xb4, 0xb2, 0x0a, 0xdd, 0x49, 0x1e, 0x27, 0x4a, 0x10, 0x43, 0x95, 0x35, 0x4a, 0x04,
	0xcd, 0x69, 0x40, 0x59, 0x7e, 0x77, 0xb3, 0x50, 0x0b, 0x55, 0xad, 0x85, 0xe5, 0xaf, 0xbe, 0xf0,
	0x7f, 0xf4, 0xe0, 0xb3, 0xf7, 0xa5, 0xe8, 0xa4, 0xd4, 0xfc, 0x58, 0x4b, 0x7e, 0x32, 0x4c, 0xa3,
	0x5b, 0xd8, 0x5f, 0x11, 0xae, 0xc7, 0x14, 0x83, 0x01, 0x18, 0x5e, 0x47, 0x0d, 0x42, 0x37, 0xf0,
	0xd2, 0xaa, 0x25, 0x93, 0xf8, 0x51, 0x45, 0xd7, 0x00, 0x3d, 0x87, 0xd7, 0x96, 0xc7, 0xcb, 0xb1,
	0xa4, 0x2c, 0xc7, 0x17, 0x03, 0x30, 0xbc, 0x88, 0xfe, 0x11, 0xe5, 0x4d, 0xac, 0x32, 0x69, 0x71,
	0x6f, 0x00, 0x86, 0xbd, 0xa8, 0x06, 0x08, 0xc3, 0x2b, 0x42, 0xa9, 0x66, 0xc6, 0xe0, 0xcb, 0x4a,
	0xeb, 0x04, 0x51, 0x06, 0x9f, 0x98, 0x84, 0x68, 0x66, 0x26, 0x6b, 0xc9, 0x28, 0xee, 0x97, 0xd3,
	0xd1, 0x74, 0x5b, 0x78, 0xce, 0xaf, 0xc2, 0x7b, 0xb1, 0xe0, 0x36, 0xc9, 0xe6, 0x41, 0xac, 0xd2,
	0x30, 0x56, 0x26, 0x55, 0xa6, 0x79, 0xee, 0x0d, 0x5d, 0x86, 0x76, 0xb3, 0x62, 0x26, 0x18, 0x4b,
	0xfb, 0xa7, 0xf0, 0xba, 0x22, 0xc7, 0xc2, 0x43, 0x1b, 0x92, 0x8a, 0xd7, 0x7e, 0x87, 0xf4, 0xa3,
	0xee, 0x0a, 0xfa, 0x06, 0xe0, 0xd3, 0x1a, 0x7f, 0xe6, 0x36, 0xa1, 0x9a, 0xac, 0x25, 0xbe, 0xaa,
	0xbc, 0xbf, 0xfe, 0xb7, 0xf7, 0xb9, 0xd0, 0xb1, 0xf0, 0x6e, 0xbb, 0xfe, 0xed, 0xc0, 0x8f, 0xce,
	0x57, 0x3b, 0x39, 0xde, 0x10, 0x19, 0x33, 0x21, 0x18, 0xc5, 0x8f, 0x1f, 0x96, 0xa3, 0x15, 0x3a,
	0xcf, 0xd1, 0x0e, 0xda, 0x1c, 0x2d, 0x33, 0x7a, 0xb7, 0xdd, 0xbb, 0x60, 0xb7, 0x77, 0xc1, 0xef,
	0xbd, 0x0b, 0xbe, 0x1f, 0x5c, 0x67, 0x77, 0x70, 0x9d, 0x9f, 0x07, 0xd7, 0xf9, 0x72, 0xdf, 0xf1,
	0xff, 0xd0, 0x74, 0xee, 0xad, 0xb2, 0x53, 0x25, 0xc2, 0x53, 0x5d, 0xf3, 0xaa, 0xb0, 0x55, 0x94,
	0x79, 0xbf, 0x2a, 0xdb, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x02, 0xae, 0x3b, 0xcc,
	0x02, 0x00, 0x00,
}

func (m *LimitOrderTrancheUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesCancelled.Size()
		i -= size
		if _, err := m.SharesCancelled.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.SharesWithdrawn.Size()
		i -= size
		if _, err := m.SharesWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Count != 0 {
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.TickIndex != 0 {
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTrancheUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTrancheUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderTrancheUser(uint64(m.TickIndex))
	}
	if m.Count != 0 {
		n += 1 + sovLimitOrderTrancheUser(uint64(m.Count))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	l = m.SharesWithdrawn.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	l = m.SharesCancelled.Size()
	n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	return n
}

func sovLimitOrderTrancheUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTrancheUser(x uint64) (n int) {
	return sovLimitOrderTrancheUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTrancheUser
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
			return fmt.Errorf("proto: LimitOrderTrancheUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
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
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
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
					return ErrIntOverflowLimitOrderTrancheUser
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
					return ErrIntOverflowLimitOrderTrancheUser
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesCancelled", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesCancelled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTrancheUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
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
func skipLimitOrderTrancheUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTrancheUser
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
					return 0, ErrIntOverflowLimitOrderTrancheUser
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
					return 0, ErrIntOverflowLimitOrderTrancheUser
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
				return 0, ErrInvalidLengthLimitOrderTrancheUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTrancheUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTrancheUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTrancheUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTrancheUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTrancheUser = fmt.Errorf("proto: unexpected end of group")
)
