// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dualitylabs/duality/incentives/stake.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// Stake records what coins are staked when by who for the purpose of
// calculating gauge reward distributions.
type Stake struct {
	// ID is the "autoincrementing" id of the stake, assigned at creation.
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// owner is the account originating the stake. Only the owner can withdraw
	// coins from the stake.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// start_time is the time at which the coins in the lock were staked.
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty" yaml:"start_time"`
	// coins are the tokens staked, and managed by the module account.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// start_dist_epoch is the dist epoch (defaulting to the day) at which the
	// coins in the lock were staked. This is used by distribution logic to filter
	// on stakes that have existed for longer than the distribution period (you
	// can only qualify for today's rewards if you staked your LP tokens
	// yesterday). We use int64 instead of uint64 to make testing easier.
	StartDistEpoch int64 `protobuf:"varint,5,opt,name=start_dist_epoch,json=startDistEpoch,proto3" json:"start_dist_epoch,omitempty"`
}

func (m *Stake) Reset()         { *m = Stake{} }
func (m *Stake) String() string { return proto.CompactTextString(m) }
func (*Stake) ProtoMessage()    {}
func (*Stake) Descriptor() ([]byte, []int) {
	return fileDescriptor_5820df422cd2215d, []int{0}
}
func (m *Stake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stake.Merge(m, src)
}
func (m *Stake) XXX_Size() int {
	return m.Size()
}
func (m *Stake) XXX_DiscardUnknown() {
	xxx_messageInfo_Stake.DiscardUnknown(m)
}

var xxx_messageInfo_Stake proto.InternalMessageInfo

func (m *Stake) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Stake) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Stake) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Stake) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Stake) GetStartDistEpoch() int64 {
	if m != nil {
		return m.StartDistEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Stake)(nil), "dualitylabs.duality.incentives.Stake")
}

func init() {
	proto.RegisterFile("dualitylabs/duality/incentives/stake.proto", fileDescriptor_5820df422cd2215d)
}

var fileDescriptor_5820df422cd2215d = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xb1, 0x6e, 0xdb, 0x30,
	0x14, 0x14, 0xe5, 0xa8, 0x40, 0x98, 0x22, 0x48, 0x85, 0x0c, 0xaa, 0x07, 0x4a, 0xd0, 0x50, 0x08,
	0x45, 0x43, 0x36, 0x2e, 0xba, 0x74, 0x54, 0xdd, 0x21, 0x40, 0x27, 0xb5, 0x53, 0x97, 0x80, 0x92,
	0x59, 0x85, 0x88, 0x24, 0x0a, 0x26, 0x9d, 0x56, 0x7f, 0x91, 0xef, 0xe8, 0x47, 0x74, 0xf6, 0xe8,
	0xb1, 0x93, 0x5c, 0xd8, 0x5b, 0x47, 0x7f, 0x41, 0x41, 0x52, 0x86, 0x9d, 0x49, 0x4f, 0xf7, 0xee,
	0xf1, 0xee, 0xc8, 0x07, 0x5f, 0xcf, 0x16, 0xb4, 0xe2, 0xaa, 0xab, 0x68, 0x2e, 0xc9, 0x50, 0x13,
	0xde, 0x14, 0xac, 0x51, 0xfc, 0x81, 0x49, 0x22, 0x15, 0xbd, 0x67, 0xb8, 0x9d, 0x0b, 0x25, 0x7c,
	0x74, 0xc4, 0xc5, 0x43, 0x8d, 0x0f, 0xdc, 0xf1, 0x65, 0x29, 0x4a, 0x61, 0xa8, 0x44, 0x57, 0x76,
	0x6a, 0x1c, 0x96, 0x42, 0x94, 0x15, 0x23, 0xe6, 0x2f, 0x5f, 0x7c, 0x27, 0x8a, 0xd7, 0x4c, 0x2a,
	0x5a, 0xb7, 0x03, 0x01, 0x15, 0x42, 0xd6, 0x42, 0x92, 0x9c, 0x4a, 0x46, 0x1e, 0xae, 0x73, 0xa6,
	0xe8, 0x35, 0x29, 0x04, 0x6f, 0x6c, 0x3f, 0xfe, 0xed, 0x42, 0xef, 0x8b, 0xb6, 0xe1, 0x9f, 0x43,
	0xf7, 0x66, 0x1a, 0x80, 0x08, 0x24, 0x27, 0x99, 0x7b, 0x33, 0xf5, 0x5f, 0x41, 0x4f, 0xfc, 0x68,
	0xd8, 0x3c, 0x70, 0x23, 0x90, 0x9c, 0xa6, 0x17, 0xbb, 0x3e, 0x7c, 0xde, 0xd1, 0xba, 0xfa, 0x10,
	0x1b, 0x38, 0xce, 0x6c, 0xdb, 0x6f, 0x21, 0x94, 0x8a, 0xce, 0xd5, 0xad, 0x96, 0x0e, 0x46, 0x11,
	0x48, 0xce, 0x26, 0x63, 0x6c, 0x7d, 0xe1, 0xbd, 0x2f, 0xfc, 0x75, 0xef, 0x2b, 0x7d, 0xbf, 0xec,
	0x43, 0xe7, 0x5f, 0x1f, 0x5e, 0x1e, 0xa6, 0xde, 0x88, 0x9a, 0x2b, 0x56, 0xb7, 0xaa, 0xdb, 0xf5,
	0xe1, 0x0b, 0x2b, 0x72, 0xe8, 0xc6, 0x8f, 0xeb, 0x10, 0x64, 0xa7, 0x06, 0xd0, 0xc7, 0xf8, 0x14,
	0x7a, 0x3a, 0x81, 0x0c, 0x4e, 0xa2, 0x51, 0x72, 0x36, 0x79, 0x89, 0x6d, 0x46, 0xac, 0x33, 0xe2,
	0x21, 0x23, 0xfe, 0x28, 0x78, 0x93, 0xbe, 0xd5, 0x5a, 0xbf, 0xd6, 0x61, 0x52, 0x72, 0x75, 0xb7,
	0xc8, 0x71, 0x21, 0x6a, 0x32, 0x5c, 0x88, 0xfd, 0x5c, 0xc9, 0xd9, 0x3d, 0x51, 0x5d, 0xcb, 0xa4,
	0x19, 0x90, 0x99, 0x3d, 0xd9, 0x4f, 0xe0, 0x85, 0x35, 0x30, 0xe3, 0x52, 0xdd, 0xb2, 0x56, 0x14,
	0x77, 0x81, 0x17, 0x81, 0x64, 0x94, 0x9d, 0x1b, 0x7c, 0xca, 0xa5, 0xfa, 0xa4, 0xd1, 0xf4, 0xf3,
	0x72, 0x83, 0xc0, 0x6a, 0x83, 0xc0, 0xdf, 0x0d, 0x02, 0x8f, 0x5b, 0xe4, 0xac, 0xb6, 0xc8, 0xf9,
	0xb3, 0x45, 0xce, 0xb7, 0xc9, 0x91, 0xe8, 0xf0, 0xa0, 0x57, 0x4f, 0x36, 0xe1, 0xe7, 0xf1, 0x2e,
	0x18, 0x13, 0xf9, 0x33, 0x73, 0x61, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xae, 0xc8,
	0xb5, 0x3a, 0x02, 0x00, 0x00,
}

func (m *Stake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartDistEpoch != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.StartDistEpoch))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStake(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStake(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintStake(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintStake(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStake(dAtA []byte, offset int, v uint64) int {
	offset -= sovStake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovStake(uint64(m.ID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovStake(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovStake(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovStake(uint64(l))
		}
	}
	if m.StartDistEpoch != 0 {
		n += 1 + sovStake(uint64(m.StartDistEpoch))
	}
	return n
}

func sovStake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStake(x uint64) (n int) {
	return sovStake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStake
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
			return fmt.Errorf("proto: Stake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
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
				return ErrInvalidLengthStake
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDistEpoch", wireType)
			}
			m.StartDistEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStake
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartDistEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStake
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
func skipStake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStake
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
					return 0, ErrIntOverflowStake
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
					return 0, ErrIntOverflowStake
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
				return 0, ErrInvalidLengthStake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStake = fmt.Errorf("proto: unexpected end of group")
)
