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
	PairId           *PairId                                `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token            string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex        int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	TrancheKey       string                                 `protobuf:"bytes,4,opt,name=trancheKey,proto3" json:"trancheKey,omitempty"`
	Address          string                                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	SharesOwned      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesOwned" yaml:"sharesOwned"`
	SharesWithdrawn  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=sharesWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesWithdrawn" yaml:"sharesWithdrawn"`
	SharesCancelled  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=sharesCancelled,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"sharesCancelled" yaml:"sharesCancelled"`
	TakerReserves github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=reservesFromSwap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reservesFromSwap" yaml:"reservesFromSwap"`
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

func (m *LimitOrderTrancheUser) GetPairId() *PairId {
	if m != nil {
		return m.PairId
	}
	return nil
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

func (m *LimitOrderTrancheUser) GetTrancheKey() string {
	if m != nil {
		return m.TrancheKey
	}
	return ""
}

func (m *LimitOrderTrancheUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*LimitOrderTrancheUser)(nil), "dualitylabs.duality.dex.LimitOrderTrancheUser")
}

func init() {
	proto.RegisterFile("duality/dex/limit_order_tranche_user.proto", fileDescriptor_896ef6766d8ff3c4)
}

var fileDescriptor_896ef6766d8ff3c4 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x1c, 0xc5, 0x3b, 0xae, 0xed, 0xda, 0xd9, 0x83, 0x32, 0xac, 0xee, 0xb8, 0x48, 0x52, 0x72, 0x90,
	0xa2, 0x6c, 0x02, 0x7a, 0x10, 0x3c, 0xae, 0x28, 0x14, 0x85, 0x95, 0xac, 0x22, 0x28, 0x52, 0xa6,
	0x99, 0x3f, 0x6d, 0x68, 0x92, 0x29, 0x33, 0x53, 0xdb, 0x7e, 0x00, 0x0f, 0xde, 0xfc, 0x58, 0x7b,
	0x5c, 0x6f, 0xe2, 0x21, 0x48, 0x7b, 0xf3, 0xd8, 0x4f, 0x20, 0x99, 0x64, 0xe3, 0x90, 0xc5, 0xc3,
	0xb2, 0xa7, 0xce, 0xff, 0xff, 0x5e, 0xdf, 0xfb, 0x05, 0x66, 0xf0, 0x23, 0x3e, 0x67, 0x49, 0xac,
	0x57, 0x01, 0x87, 0x65, 0x90, 0xc4, 0x69, 0xac, 0x87, 0x42, 0x72, 0x90, 0x43, 0x2d, 0x59, 0x16,
	0x4d, 0x60, 0x38, 0x57, 0x20, 0xfd, 0x99, 0x14, 0x5a, 0x90, 0x83, 0xca, 0x9b, 0xb0, 0x91, 0xf2,
	0xab, 0xb3, 0xcf, 0x61, 0x79, 0xb8, 0x3f, 0x16, 0x63, 0x61, 0x3c, 0x41, 0x71, 0x2a, 0xed, 0x87,
	0xf7, 0xed, 0xe8, 0x19, 0x8b, 0xe5, 0x30, 0xe6, 0xa5, 0xe4, 0xfd, 0x68, 0xe3, 0xbb, 0x6f, 0x8a,
	0xb2, 0x93, 0xa2, 0xeb, 0x5d, 0x59, 0xf5, 0x5e, 0x81, 0x24, 0xcf, 0x70, 0xa7, 0xb0, 0x0e, 0x38,
	0x45, 0x3d, 0xd4, 0xdf, 0x7b, 0xe2, 0xfa, 0xff, 0x29, 0xf5, 0xdf, 0x1a, 0x5b, 0x58, 0xd9, 0xc9,
	0x3e, 0x6e, 0x6b, 0x31, 0x85, 0x8c, 0xde, 0xe8, 0xa1, 0x7e, 0x37, 0x2c, 0x07, 0xf2, 0x00, 0x77,
	0x75, 0x1c, 0x4d, 0x07, 0x19, 0x87, 0x25, 0xdd, 0xe9, 0xa1, 0xfe, 0x4e, 0xf8, 0x6f, 0x41, 0x1c,
	0x8c, 0xab, 0xcf, 0x7c, 0x0d, 0x2b, 0x7a, 0xd3, 0xfc, 0xd1, 0xda, 0x10, 0x8a, 0x77, 0x19, 0xe7,
	0x12, 0x94, 0xa2, 0x6d, 0x23, 0x5e, 0x8c, 0x64, 0x8e, 0xf7, 0xd4, 0x84, 0x49, 0x50, 0x27, 0x8b,
	0x0c, 0x38, 0xed, 0x14, 0xea, 0xf1, 0xe9, 0x59, 0xee, 0xb6, 0x7e, 0xe5, 0xee, 0xc3, 0x71, 0xac,
	0x27, 0xf3, 0x91, 0x1f, 0x89, 0x34, 0x88, 0x84, 0x4a, 0x85, 0xaa, 0x7e, 0x8e, 0x14, 0x9f, 0x06,
	0x7a, 0x35, 0x03, 0xe5, 0x0f, 0x32, 0xfd, 0x27, 0x77, 0xed, 0x90, 0x6d, 0xee, 0x92, 0x15, 0x4b,
	0x93, 0xe7, 0x9e, 0xb5, 0xf4, 0x42, 0xdb, 0x42, 0xbe, 0x22, 0x7c, 0xbb, 0x9c, 0x3f, 0xc4, 0x7a,
	0xc2, 0x25, 0x5b, 0x64, 0x74, 0xd7, 0x74, 0x7f, 0xba, 0x72, 0x77, 0x33, 0x68, 0x9b, 0xbb, 0xf7,
	0xec, 0xfe, 0x5a, 0xf0, 0xc2, 0xa6, 0xd5, 0xe2, 0x78, 0xc1, 0xb2, 0x08, 0x92, 0x04, 0x38, 0xbd,
	0x75, 0x3d, 0x8e, 0x3a, 0xa8, 0xc9, 0x51, 0x0b, 0x35, 0x47, 0xbd, 0x21, 0xdf, 0x10, 0xbe, 0x23,
	0x41, 0x81, 0xfc, 0x02, 0xea, 0x95, 0x14, 0xe9, 0xe9, 0x82, 0xcd, 0x68, 0xd7, 0x80, 0x7c, 0xbe,
	0x32, 0xc8, 0xa5, 0xa4, 0x6d, 0xee, 0x1e, 0x94, 0x24, 0x4d, 0xc5, 0x0b, 0x2f, 0x99, 0x8f, 0x5f,
	0x9e, 0xad, 0x1d, 0x74, 0xbe, 0x76, 0xd0, 0xef, 0xb5, 0x83, 0xbe, 0x6f, 0x9c, 0xd6, 0xf9, 0xc6,
	0x69, 0xfd, 0xdc, 0x38, 0xad, 0x8f, 0x8f, 0x2d, 0x84, 0xea, 0x06, 0x1f, 0x15, 0xd7, 0xf9, 0x62,
	0x08, 0x96, 0xe6, 0x89, 0x18, 0x96, 0x51, 0xc7, 0xbc, 0x90, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x14, 0xbb, 0x84, 0xaa, 0x99, 0x03, 0x00, 0x00,
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
		size := m.TakerReserves.Size()
		i -= size
		if _, err := m.TakerReserves.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
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
	if len(m.TrancheKey) > 0 {
		i -= len(m.TrancheKey)
		copy(dAtA[i:], m.TrancheKey)
		i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(len(m.TrancheKey)))
		i--
		dAtA[i] = 0x22
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
	if m.PairId != nil {
		{
			size, err := m.PairId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLimitOrderTrancheUser(dAtA, i, uint64(size))
		}
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
	if m.PairId != nil {
		l = m.PairId.Size()
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderTrancheUser(uint64(m.TickIndex))
	}
	l = len(m.TrancheKey)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUser(uint64(l))
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
	l = m.TakerReserves.Size()
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
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUser
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
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairId == nil {
				m.PairId = &PairId{}
			}
			if err := m.PairId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrancheKey", wireType)
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
			m.TrancheKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesFromSwap", wireType)
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
			if err := m.TakerReserves.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
