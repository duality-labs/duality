// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/ticks.proto

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

type Ticks struct {
	Price       string         `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Fee         string         `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Direction   string         `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction,omitempty"`
	OrderType   string         `protobuf:"bytes,4,opt,name=orderType,proto3" json:"orderType,omitempty"`
	Reserve     string         `protobuf:"bytes,5,opt,name=reserve,proto3" json:"reserve,omitempty"`
	Token       string         `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	PairPrice   string         `protobuf:"bytes,7,opt,name=pairPrice,proto3" json:"pairPrice,omitempty"`
	PairFee     string         `protobuf:"bytes,8,opt,name=pairFee,proto3" json:"pairFee,omitempty"`
	TotalShares string         `protobuf:"bytes,9,opt,name=totalShares,proto3" json:"totalShares,omitempty"`
	Orderparams []*OrderParams `protobuf:"bytes,10,rep,name=orderparams,proto3" json:"orderparams,omitempty"`
}

func (m *Ticks) Reset()         { *m = Ticks{} }
func (m *Ticks) String() string { return proto.CompactTextString(m) }
func (*Ticks) ProtoMessage()    {}
func (*Ticks) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2fc2226c4509899, []int{0}
}
func (m *Ticks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ticks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ticks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ticks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticks.Merge(m, src)
}
func (m *Ticks) XXX_Size() int {
	return m.Size()
}
func (m *Ticks) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticks.DiscardUnknown(m)
}

var xxx_messageInfo_Ticks proto.InternalMessageInfo

func (m *Ticks) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Ticks) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Ticks) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *Ticks) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *Ticks) GetReserve() string {
	if m != nil {
		return m.Reserve
	}
	return ""
}

func (m *Ticks) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Ticks) GetPairPrice() string {
	if m != nil {
		return m.PairPrice
	}
	return ""
}

func (m *Ticks) GetPairFee() string {
	if m != nil {
		return m.PairFee
	}
	return ""
}

func (m *Ticks) GetTotalShares() string {
	if m != nil {
		return m.TotalShares
	}
	return ""
}

func (m *Ticks) GetOrderparams() []*OrderParams {
	if m != nil {
		return m.Orderparams
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticks)(nil), "nicholasdotsol.duality.dex.Ticks")
}

func init() { proto.RegisterFile("dex/ticks.proto", fileDescriptor_f2fc2226c4509899) }

var fileDescriptor_f2fc2226c4509899 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0x59, 0xf8, 0x03, 0x7f, 0xca, 0x41, 0xd3, 0x18, 0xd3, 0x10, 0xb3, 0x21, 0x5e, 0xe4,
	0x62, 0x37, 0xd1, 0x37, 0x30, 0x46, 0xe3, 0x45, 0x09, 0x70, 0xf2, 0x62, 0xca, 0x76, 0x94, 0x86,
	0x85, 0x69, 0xda, 0x62, 0xe0, 0x2d, 0x7c, 0x24, 0x8f, 0x1e, 0x39, 0x7a, 0x34, 0xf0, 0x22, 0xa6,
	0xed, 0x8a, 0x5c, 0xbc, 0xf5, 0xfb, 0xbe, 0xe9, 0x6f, 0x26, 0x33, 0xe4, 0x40, 0xc2, 0x32, 0x73,
	0x2a, 0x9f, 0x5a, 0xae, 0x0d, 0x3a, 0xa4, 0x9d, 0xb9, 0xca, 0x27, 0x58, 0x08, 0x2b, 0xd1, 0x59,
	0x2c, 0xb8, 0x5c, 0x88, 0x42, 0xb9, 0x15, 0x97, 0xb0, 0xec, 0x1c, 0xfb, 0x62, 0x34, 0x12, 0xcc,
	0x93, 0x16, 0x46, 0xcc, 0xca, 0x3f, 0xa7, 0xef, 0x55, 0x52, 0x1f, 0x79, 0x06, 0x3d, 0x22, 0x75,
	0x6d, 0x54, 0x0e, 0x2c, 0xe9, 0x26, 0xbd, 0xd6, 0x20, 0x0a, 0x7a, 0x48, 0x6a, 0xcf, 0x00, 0xac,
	0x1a, 0x3c, 0xff, 0xa4, 0x27, 0xa4, 0x25, 0x95, 0x81, 0xdc, 0x29, 0x9c, 0xb3, 0x5a, 0xf0, 0x7f,
	0x0d, 0x9f, 0x86, 0x2e, 0xa3, 0x95, 0x06, 0xf6, 0x2f, 0xa6, 0x3b, 0x83, 0x32, 0xd2, 0x34, 0x60,
	0xc1, 0xbc, 0x02, 0xab, 0x87, 0xec, 0x47, 0xfa, 0xee, 0x0e, 0xa7, 0x30, 0x67, 0x8d, 0xd8, 0x3d,
	0x08, 0x4f, 0xd3, 0x42, 0x99, 0x7e, 0x98, 0xab, 0x19, 0x69, 0x3b, 0xc3, 0xd3, 0xbc, 0xb8, 0x01,
	0x60, 0xff, 0x23, 0xad, 0x94, 0xb4, 0x4b, 0xda, 0x0e, 0x9d, 0x28, 0x86, 0x13, 0x61, 0xc0, 0xb2,
	0x56, 0x48, 0xf7, 0x2d, 0x7a, 0x47, 0xda, 0x61, 0xac, 0xb8, 0x0c, 0x46, 0xba, 0xb5, 0x5e, 0xfb,
	0xe2, 0x8c, 0xff, 0xbd, 0x41, 0xfe, 0xe0, 0xcb, 0xfb, 0xa1, 0x7c, 0xb0, 0xff, 0xf7, 0xea, 0xf6,
	0x63, 0x93, 0x26, 0xeb, 0x4d, 0x9a, 0x7c, 0x6d, 0xd2, 0xe4, 0x6d, 0x9b, 0x56, 0xd6, 0xdb, 0xb4,
	0xf2, 0xb9, 0x4d, 0x2b, 0x8f, 0xe7, 0x2f, 0xca, 0x4d, 0x16, 0x63, 0x9e, 0xe3, 0x2c, 0xbb, 0x2f,
	0xc9, 0xd7, 0xe8, 0x86, 0x58, 0x64, 0x25, 0x39, 0x5b, 0x66, 0xe1, 0x8a, 0x2b, 0x0d, 0x76, 0xdc,
	0x08, 0x27, 0xb9, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xe9, 0x3c, 0x46, 0xd9, 0x01, 0x00,
	0x00,
}

func (m *Ticks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ticks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ticks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Orderparams) > 0 {
		for iNdEx := len(m.Orderparams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orderparams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.TotalShares) > 0 {
		i -= len(m.TotalShares)
		copy(dAtA[i:], m.TotalShares)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.TotalShares)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.PairFee) > 0 {
		i -= len(m.PairFee)
		copy(dAtA[i:], m.PairFee)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.PairFee)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PairPrice) > 0 {
		i -= len(m.PairPrice)
		copy(dAtA[i:], m.PairPrice)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.PairPrice)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Reserve) > 0 {
		i -= len(m.Reserve)
		copy(dAtA[i:], m.Reserve)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Reserve)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Direction) > 0 {
		i -= len(m.Direction)
		copy(dAtA[i:], m.Direction)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Direction)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintTicks(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicks(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ticks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.Direction)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.Reserve)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.PairPrice)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.PairFee)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	l = len(m.TotalShares)
	if l > 0 {
		n += 1 + l + sovTicks(uint64(l))
	}
	if len(m.Orderparams) > 0 {
		for _, e := range m.Orderparams {
			l = e.Size()
			n += 1 + l + sovTicks(uint64(l))
		}
	}
	return n
}

func sovTicks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicks(x uint64) (n int) {
	return sovTicks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ticks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicks
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
			return fmt.Errorf("proto: Ticks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ticks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Direction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserve = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalShares = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderparams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicks
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
				return ErrInvalidLengthTicks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orderparams = append(m.Orderparams, &OrderParams{})
			if err := m.Orderparams[len(m.Orderparams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicks
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
func skipTicks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicks
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
					return 0, ErrIntOverflowTicks
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
					return 0, ErrIntOverflowTicks
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
				return 0, ErrInvalidLengthTicks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicks = fmt.Errorf("proto: unexpected end of group")
)
