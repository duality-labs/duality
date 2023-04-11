// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/incentives/lock.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/duality-labs/duality/x/dex/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Lock is a single lock unit by period defined by the x/lockup module.
// It's a record of a locked coin at a specific time. It stores owner, duration,
// unlock time and the number of coins locked. A state of a period lock is
// created upon lock creation, and deleted once the lock has been matured after
// the `duration` has passed since unbonding started.
type Lock struct {
	// ID is the unique id of the lock.
	// The ID of the lock is decided upon lock creation, incrementing by 1 for
	// every lock.
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Owner is the account address of the lock owner.
	// Only the owner can modify the state of the lock.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// Duration is the time needed for a lock to mature after unlocking has
	// started.
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// EndTime refers to the time at which the lock would mature and get deleted.
	// This value is first initialized when an unlock has started for the lock,
	// end time being block time + duration.
	EndTime time.Time `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	// Coins are the tokens locked within the lock, kept in the module account.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6be01928b83ef65, []int{0}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Lock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Lock) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Lock) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *Lock) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// QueryCondition is a struct used for querying locks upon different conditions.
type QueryCondition struct {
	// PairID represents the token pair we are looking to lock up
	PairID *types1.PairID `protobuf:"bytes,1,opt,name=pairID,proto3" json:"pairID,omitempty"`
	// Start tick inclusive
	StartTick int64 `protobuf:"varint,2,opt,name=startTick,proto3" json:"startTick,omitempty"`
	// End tick exclusive
	EndTick int64 `protobuf:"varint,3,opt,name=endTick,proto3" json:"endTick,omitempty"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6be01928b83ef65, []int{1}
}
func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return m.Size()
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetPairID() *types1.PairID {
	if m != nil {
		return m.PairID
	}
	return nil
}

func (m *QueryCondition) GetStartTick() int64 {
	if m != nil {
		return m.StartTick
	}
	return 0
}

func (m *QueryCondition) GetEndTick() int64 {
	if m != nil {
		return m.EndTick
	}
	return 0
}

func init() {
	proto.RegisterType((*Lock)(nil), "duality.incentives.Lock")
	proto.RegisterType((*QueryCondition)(nil), "duality.incentives.QueryCondition")
}

func init() { proto.RegisterFile("duality/incentives/lock.proto", fileDescriptor_a6be01928b83ef65) }

var fileDescriptor_a6be01928b83ef65 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xbb, 0x6e, 0xdb, 0x30,
	0x14, 0xb5, 0xfc, 0x48, 0x13, 0xa6, 0x48, 0x0b, 0xa2, 0x83, 0xe2, 0xb6, 0x92, 0xa1, 0xa1, 0xf0,
	0xd0, 0x90, 0xb5, 0x3b, 0x14, 0xe8, 0xa8, 0x78, 0x09, 0x90, 0xa1, 0x15, 0x32, 0x75, 0x09, 0x28,
	0x89, 0x55, 0x08, 0x3d, 0x28, 0x88, 0x54, 0x6a, 0x8d, 0xfd, 0x83, 0x8c, 0xfd, 0x86, 0x7e, 0x49,
	0xc6, 0x8c, 0x9d, 0xec, 0xc2, 0xde, 0x3a, 0xe6, 0x0b, 0x0a, 0x52, 0x62, 0x1d, 0x34, 0x93, 0x48,
	0x9e, 0x7b, 0xcf, 0x3d, 0xe7, 0x5c, 0x81, 0xd7, 0x71, 0x4d, 0x32, 0x26, 0x1b, 0xcc, 0x8a, 0x88,
	0x16, 0x92, 0x5d, 0x53, 0x81, 0x33, 0x1e, 0xa5, 0xa8, 0xac, 0xb8, 0xe4, 0x10, 0x76, 0x30, 0xda,
	0xc1, 0xe3, 0x17, 0x09, 0x4f, 0xb8, 0x86, 0xb1, 0x3a, 0xb5, 0x95, 0x63, 0x27, 0xe1, 0x3c, 0xc9,
	0x28, 0xd6, 0xb7, 0xb0, 0xfe, 0x8a, 0xe3, 0xba, 0x22, 0x92, 0xf1, 0xa2, 0xc3, 0xdd, 0xff, 0x71,
	0xc9, 0x72, 0x2a, 0x24, 0xc9, 0x4b, 0x43, 0x10, 0x71, 0x91, 0x73, 0x81, 0x43, 0x22, 0x28, 0xbe,
	0x9e, 0x85, 0x54, 0x92, 0x19, 0x8e, 0x38, 0x33, 0x04, 0xc7, 0x46, 0x69, 0x4c, 0x97, 0xb8, 0x24,
	0xac, 0xba, 0x64, 0x71, 0x0b, 0x79, 0xeb, 0x3e, 0x18, 0x9e, 0xf3, 0x28, 0x85, 0x47, 0xa0, 0x7f,
	0xb6, 0xb0, 0xad, 0x89, 0x35, 0x1d, 0x06, 0xfd, 0xb3, 0x05, 0x7c, 0x03, 0x46, 0xfc, 0x5b, 0x41,
	0x2b, 0xbb, 0x3f, 0xb1, 0xa6, 0x07, 0xfe, 0xf3, 0xfb, 0x95, 0xfb, 0xb4, 0x21, 0x79, 0xf6, 0xd1,
	0xd3, 0xcf, 0x5e, 0xd0, 0xc2, 0xf0, 0x0a, 0xec, 0x1b, 0xb9, 0xf6, 0x60, 0x62, 0x4d, 0x0f, 0xe7,
	0xc7, 0xa8, 0xd5, 0x8b, 0x8c, 0x5e, 0xb4, 0xe8, 0x0a, 0xfc, 0xd9, 0xed, 0xca, 0xed, 0xfd, 0x59,
	0xb9, 0xd0, 0xb4, 0xbc, 0xe5, 0x39, 0x93, 0x34, 0x2f, 0x65, 0x73, 0xbf, 0x72, 0x9f, 0xb5, 0xfc,
	0x06, 0xf3, 0x7e, 0xac, 0x5d, 0x2b, 0xf8, 0xc7, 0x0e, 0x03, 0xb0, 0x4f, 0x8b, 0xf8, 0x52, 0x99,
	0xb7, 0x87, 0x7a, 0xd2, 0xf8, 0xd1, 0xa4, 0x0b, 0x93, 0x8c, 0xff, 0x52, 0x8d, 0xda, 0x91, 0x9a,
	0x4e, 0xef, 0x46, 0x91, 0x3e, 0xa1, 0x45, 0xac, 0x4a, 0x21, 0x01, 0x23, 0x95, 0x93, 0xb0, 0x47,
	0x93, 0x81, 0x96, 0xde, 0x26, 0x89, 0x54, 0x92, 0xa8, 0x4b, 0x12, 0x9d, 0x72, 0x56, 0xf8, 0xef,
	0x14, 0xdf, 0xcf, 0xb5, 0x3b, 0x4d, 0x98, 0xbc, 0xaa, 0x43, 0x14, 0xf1, 0x1c, 0x77, 0xb1, 0xb7,
	0x9f, 0x13, 0x11, 0xa7, 0x58, 0x36, 0x25, 0x15, 0xba, 0x41, 0x04, 0x2d, 0xb3, 0xf7, 0xdd, 0x02,
	0x47, 0x9f, 0x6b, 0x5a, 0x35, 0xa7, 0xbc, 0x88, 0x99, 0x76, 0xf2, 0x01, 0xec, 0xa9, 0x2d, 0x74,
	0x79, 0x1f, 0xce, 0x5d, 0xd4, 0x2d, 0x28, 0x23, 0xa1, 0x30, 0x67, 0x14, 0xd3, 0x25, 0xfa, 0xa4,
	0xcb, 0x82, 0xae, 0x1c, 0xbe, 0x02, 0x07, 0x42, 0x92, 0x4a, 0x5e, 0xb0, 0x28, 0xd5, 0x8b, 0x19,
	0x04, 0xbb, 0x07, 0x68, 0x83, 0xd6, 0x57, 0x94, 0xea, 0x4d, 0x0c, 0x02, 0x73, 0xf5, 0xcf, 0x6f,
	0x37, 0x8e, 0x75, 0xb7, 0x71, 0xac, 0xdf, 0x1b, 0xc7, 0xba, 0xd9, 0x3a, 0xbd, 0xbb, 0xad, 0xd3,
	0xfb, 0xb5, 0x75, 0x7a, 0x5f, 0xe6, 0x0f, 0xec, 0x74, 0x83, 0x4f, 0x94, 0x0a, 0x73, 0xc1, 0xcb,
	0x87, 0xbf, 0xb7, 0xb6, 0x17, 0xee, 0xe9, 0xb8, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xa6, 0x43, 0xaf, 0x01, 0x03, 0x00, 0x00,
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLock(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLock(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLock(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTick != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.EndTick))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTick != 0 {
		i = encodeVarintLock(dAtA, i, uint64(m.StartTick))
		i--
		dAtA[i] = 0x10
	}
	if m.PairID != nil {
		{
			size, err := m.PairID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLock(dAtA []byte, offset int, v uint64) int {
	offset -= sovLock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLock(uint64(m.ID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLock(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovLock(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLock(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovLock(uint64(l))
		}
	}
	return n
}

func (m *QueryCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairID != nil {
		l = m.PairID.Size()
		n += 1 + l + sovLock(uint64(l))
	}
	if m.StartTick != 0 {
		n += 1 + sovLock(uint64(m.StartTick))
	}
	if m.EndTick != 0 {
		n += 1 + sovLock(uint64(m.EndTick))
	}
	return n
}

func sovLock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLock(x uint64) (n int) {
	return sovLock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func (m *QueryCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLock
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
			return fmt.Errorf("proto: QueryCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
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
				return ErrInvalidLengthLock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairID == nil {
				m.PairID = &types1.PairID{}
			}
			if err := m.PairID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTick", wireType)
			}
			m.StartTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTick", wireType)
			}
			m.EndTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLock
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
func skipLock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
					return 0, ErrIntOverflowLock
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
				return 0, ErrInvalidLengthLock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLock = fmt.Errorf("proto: unexpected end of group")
)