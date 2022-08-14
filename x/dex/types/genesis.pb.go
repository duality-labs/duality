// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params                     Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	NodesList                  []Nodes                 `protobuf:"bytes,2,rep,name=nodesList,proto3" json:"nodesList"`
	NodesCount                 uint64                  `protobuf:"varint,3,opt,name=nodesCount,proto3" json:"nodesCount,omitempty"`
	VirtualPriceTickQueueList  []VirtualPriceTickQueue `protobuf:"bytes,4,rep,name=virtualPriceTickQueueList,proto3" json:"virtualPriceTickQueueList"`
	VirtualPriceTickQueueCount uint64                  `protobuf:"varint,5,opt,name=virtualPriceTickQueueCount,proto3" json:"virtualPriceTickQueueCount,omitempty"`
	TicksList                  []Ticks                 `protobuf:"bytes,6,rep,name=ticksList,proto3" json:"ticksList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetNodesList() []Nodes {
	if m != nil {
		return m.NodesList
	}
	return nil
}

func (m *GenesisState) GetNodesCount() uint64 {
	if m != nil {
		return m.NodesCount
	}
	return 0
}

func (m *GenesisState) GetVirtualPriceTickQueueList() []VirtualPriceTickQueue {
	if m != nil {
		return m.VirtualPriceTickQueueList
	}
	return nil
}

func (m *GenesisState) GetVirtualPriceTickQueueCount() uint64 {
	if m != nil {
		return m.VirtualPriceTickQueueCount
	}
	return 0
}

func (m *GenesisState) GetTicksList() []Ticks {
	if m != nil {
		return m.TicksList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nicholasdotsol.duality.dex.GenesisState")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xdb, 0x3f, 0xfc, 0x49, 0x5c, 0x4c, 0xd4, 0xc6, 0x03, 0xf6, 0xb0, 0x22, 0x27, 0x2e,
	0xb6, 0x11, 0xef, 0xc6, 0xa0, 0x86, 0x8b, 0x21, 0x08, 0xc6, 0x83, 0x17, 0xb2, 0xb4, 0x9b, 0xb2,
	0xa1, 0x74, 0x6a, 0x77, 0xd7, 0xc0, 0x5b, 0xf8, 0x32, 0xbe, 0x03, 0x47, 0x8e, 0x9e, 0x8c, 0x81,
	0x17, 0x31, 0xbb, 0x5d, 0x91, 0x03, 0xe0, 0xad, 0xfd, 0x76, 0x7e, 0xdf, 0x7c, 0x33, 0x19, 0x74,
	0x14, 0xd2, 0x89, 0x1f, 0xd1, 0x84, 0x72, 0xc6, 0xbd, 0x34, 0x03, 0x01, 0x8e, 0x9b, 0xb0, 0x60,
	0x08, 0x31, 0xe1, 0x21, 0x08, 0x0e, 0xb1, 0x17, 0x4a, 0x12, 0x33, 0x31, 0xf5, 0x42, 0x3a, 0x71,
	0x8f, 0x23, 0x88, 0x40, 0x97, 0xf9, 0xea, 0x2b, 0x27, 0xdc, 0x43, 0x65, 0x92, 0x92, 0x8c, 0x8c,
	0x8d, 0x87, 0x7b, 0xa0, 0x94, 0x04, 0x42, 0xfa, 0x23, 0xd4, 0x94, 0xf0, 0xca, 0x32, 0x21, 0x49,
	0xdc, 0x4f, 0x33, 0x16, 0xd0, 0xbe, 0x60, 0xc1, 0xa8, 0xff, 0x22, 0xa9, 0xa4, 0xeb, 0x90, 0x52,
	0x0d, 0x54, 0x7b, 0x2f, 0xa0, 0xfd, 0x56, 0x9e, 0xad, 0x27, 0x88, 0xa0, 0xce, 0x35, 0x2a, 0xe5,
	0x6d, 0x2a, 0x76, 0xd5, 0xae, 0x97, 0x1b, 0x35, 0x6f, 0x7b, 0x56, 0xaf, 0xa3, 0x2b, 0x9b, 0xc5,
	0xd9, 0xe7, 0xa9, 0xd5, 0x35, 0x9c, 0x73, 0x87, 0xf6, 0x74, 0xac, 0x7b, 0xc6, 0x45, 0xe5, 0x5f,
	0xb5, 0x50, 0x2f, 0x37, 0xce, 0x76, 0x99, 0xb4, 0x55, 0xb1, 0xf1, 0xf8, 0x25, 0x1d, 0x8c, 0x90,
	0xfe, 0xb9, 0x01, 0x99, 0x88, 0x4a, 0xa1, 0x6a, 0xd7, 0x8b, 0xdd, 0x35, 0xc5, 0x91, 0xe8, 0xc4,
	0x0c, 0xdb, 0x51, 0xb3, 0x3e, 0xb2, 0x60, 0xf4, 0xa0, 0x26, 0xd5, 0x6d, 0x8b, 0xba, 0xed, 0xc5,
	0xae, 0xb6, 0x4f, 0x9b, 0x60, 0x13, 0x63, 0xbb, 0xb3, 0x73, 0x85, 0xdc, 0x8d, 0x8f, 0x79, 0xcc,
	0xff, 0x3a, 0xe6, 0x8e, 0x0a, 0xb5, 0x1d, 0xbd, 0x7f, 0x1d, 0xb3, 0xf4, 0xf7, 0x76, 0x14, 0xbe,
	0xda, 0xce, 0x8a, 0x6c, 0xb6, 0x66, 0x0b, 0x6c, 0xcf, 0x17, 0xd8, 0xfe, 0x5a, 0x60, 0xfb, 0x6d,
	0x89, 0xad, 0xf9, 0x12, 0x5b, 0x1f, 0x4b, 0x6c, 0x3d, 0x9f, 0x47, 0x4c, 0x0c, 0xe5, 0xc0, 0x0b,
	0x60, 0xec, 0xb7, 0x8d, 0xef, 0x2d, 0x88, 0x1e, 0xc4, 0xbe, 0xf1, 0xf5, 0x27, 0xbe, 0x3e, 0x83,
	0x69, 0x4a, 0xf9, 0xa0, 0xa4, 0xef, 0xe0, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xc6, 0x5f,
	0xbf, 0xa6, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TicksList) > 0 {
		for iNdEx := len(m.TicksList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TicksList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.VirtualPriceTickQueueCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VirtualPriceTickQueueCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.VirtualPriceTickQueueList) > 0 {
		for iNdEx := len(m.VirtualPriceTickQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VirtualPriceTickQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.NodesCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NodesCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodesList) > 0 {
		for iNdEx := len(m.NodesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.NodesList) > 0 {
		for _, e := range m.NodesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NodesCount != 0 {
		n += 1 + sovGenesis(uint64(m.NodesCount))
	}
	if len(m.VirtualPriceTickQueueList) > 0 {
		for _, e := range m.VirtualPriceTickQueueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.VirtualPriceTickQueueCount != 0 {
		n += 1 + sovGenesis(uint64(m.VirtualPriceTickQueueCount))
	}
	if len(m.TicksList) > 0 {
		for _, e := range m.TicksList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodesList = append(m.NodesList, Nodes{})
			if err := m.NodesList[len(m.NodesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodesCount", wireType)
			}
			m.NodesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodesCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualPriceTickQueueList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VirtualPriceTickQueueList = append(m.VirtualPriceTickQueueList, VirtualPriceTickQueue{})
			if err := m.VirtualPriceTickQueueList[len(m.VirtualPriceTickQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VirtualPriceTickQueueCount", wireType)
			}
			m.VirtualPriceTickQueueCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VirtualPriceTickQueueCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TicksList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TicksList = append(m.TicksList, Ticks{})
			if err := m.TicksList[len(m.TicksList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
