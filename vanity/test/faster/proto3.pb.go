// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3.proto

package vanity

import (
	fmt "fmt"
	proto "github.com/temporalio/gogo-protobuf/proto"
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

type Aproto3 struct {
	B string `protobuf:"bytes,1,opt,name=B,proto3" json:"B,omitempty"`
}

func (m *Aproto3) Reset()         { *m = Aproto3{} }
func (m *Aproto3) String() string { return proto.CompactTextString(m) }
func (*Aproto3) ProtoMessage()    {}
func (*Aproto3) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0}
}
func (m *Aproto3) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Aproto3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Aproto3.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Aproto3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aproto3.Merge(m, src)
}
func (m *Aproto3) XXX_Size() int {
	return m.Size()
}
func (m *Aproto3) XXX_DiscardUnknown() {
	xxx_messageInfo_Aproto3.DiscardUnknown(m)
}

var xxx_messageInfo_Aproto3 proto.InternalMessageInfo

func (m *Aproto3) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func init() {
	proto.RegisterType((*Aproto3)(nil), "vanity.Aproto3")
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor_4fee6d65e34a64b6) }

var fileDescriptor_4fee6d65e34a64b6 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd6, 0x03, 0x53, 0x42, 0x6c, 0x65, 0x89, 0x79, 0x99, 0x25, 0x95, 0x4a, 0xe2, 0x5c,
	0xec, 0x8e, 0x10, 0x09, 0x21, 0x1e, 0x2e, 0x46, 0x27, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x46, 0x27, 0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x83,
	0xa8, 0x07, 0x04, 0x00, 0x00, 0xff, 0xff, 0x73, 0xb1, 0xdb, 0xd2, 0x51, 0x00, 0x00, 0x00,
}

func (m *Aproto3) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Aproto3) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Aproto3) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.B) > 0 {
		i -= len(m.B)
		copy(dAtA[i:], m.B)
		i = encodeVarintProto3(dAtA, i, uint64(len(m.B)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProto3(dAtA []byte, offset int, v uint64) int {
	offset -= sovProto3(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Aproto3) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.B)
	if l > 0 {
		n += 1 + l + sovProto3(uint64(l))
	}
	return n
}

func sovProto3(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProto3(x uint64) (n int) {
	return sovProto3(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Aproto3) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProto3
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
			return fmt.Errorf("proto: Aproto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Aproto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProto3
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
				return ErrInvalidLengthProto3
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProto3
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.B = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProto3(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProto3
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProto3
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
func skipProto3(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProto3
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
					return 0, ErrIntOverflowProto3
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
					return 0, ErrIntOverflowProto3
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
				return 0, ErrInvalidLengthProto3
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProto3
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProto3
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProto3        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProto3          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProto3 = fmt.Errorf("proto: unexpected end of group")
)
