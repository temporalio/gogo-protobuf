// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gogovanity.proto

package vanity

import (
	fmt "fmt"
	_ "github.com/temporalio/gogo-protobuf/gogoproto"
	proto "github.com/temporalio/gogo-protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type B struct {
	String_ *string `protobuf:"bytes,1,opt,name=String" json:"String,omitempty"`
	Int64   int64   `protobuf:"varint,2,opt,name=Int64" json:"Int64"`
	Int32   *int32  `protobuf:"varint,3,opt,name=Int32,def=1234" json:"Int32,omitempty"`
}

func (m *B) Reset()      { *m = B{} }
func (*B) ProtoMessage() {}
func (*B) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c9b51615339d8e, []int{0}
}
func (m *B) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_B.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_B.Merge(m, src)
}
func (m *B) XXX_Size() int {
	return m.Size()
}
func (m *B) XXX_DiscardUnknown() {
	xxx_messageInfo_B.DiscardUnknown(m)
}

var xxx_messageInfo_B proto.InternalMessageInfo

const Default_B_Int32 int32 = 1234

func (m *B) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func (m *B) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *B) GetInt32() int32 {
	if m != nil && m.Int32 != nil {
		return *m.Int32
	}
	return Default_B_Int32
}

func init() {
	proto.RegisterType((*B)(nil), "vanity.B")
}

func init() { proto.RegisterFile("gogovanity.proto", fileDescriptor_f8c9b51615339d8e) }

var fileDescriptor_f8c9b51615339d8e = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcf, 0x4f, 0xcf,
	0x2f, 0x4b, 0xcc, 0xcb, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0,
	0xa4, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0x8a, 0xf4,
	0xc1, 0xd2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xb4, 0x29, 0x45, 0x72, 0x31,
	0x3a, 0x09, 0xc9, 0x70, 0xb1, 0x05, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x3a, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x18, 0x04, 0x15, 0x13, 0x92, 0xe2, 0x62, 0xf5, 0xcc,
	0x2b, 0x31, 0x33, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x06, 0x4b, 0x32, 0x04, 0x41, 0x84, 0xa0,
	0x72, 0xc6, 0x46, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x56, 0x2c, 0x86, 0x46, 0xc6, 0x26, 0x41,
	0x10, 0x21, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50, 0x8e,
	0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xab, 0xeb, 0xeb, 0x07, 0xda, 0x00, 0x00, 0x00,
}

func (this *B) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*B)
	if !ok {
		that2, ok := that.(B)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.String_ != nil && that1.String_ != nil {
		if *this.String_ != *that1.String_ {
			return false
		}
	} else if this.String_ != nil {
		return false
	} else if that1.String_ != nil {
		return false
	}
	if this.Int64 != that1.Int64 {
		return false
	}
	if this.Int32 != nil && that1.Int32 != nil {
		if *this.Int32 != *that1.Int32 {
			return false
		}
	} else if this.Int32 != nil {
		return false
	} else if that1.Int32 != nil {
		return false
	}
	return true
}
func (this *B) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&vanity.B{")
	if this.String_ != nil {
		s = append(s, "String_: "+valueToGoStringGogovanity(this.String_, "string")+",\n")
	}
	s = append(s, "Int64: "+fmt.Sprintf("%#v", this.Int64)+",\n")
	if this.Int32 != nil {
		s = append(s, "Int32: "+valueToGoStringGogovanity(this.Int32, "int32")+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGogovanity(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *B) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Int32 != nil {
		i = encodeVarintGogovanity(dAtA, i, uint64(*m.Int32))
		i--
		dAtA[i] = 0x18
	}
	i = encodeVarintGogovanity(dAtA, i, uint64(m.Int64))
	i--
	dAtA[i] = 0x10
	if m.String_ != nil {
		i -= len(*m.String_)
		copy(dAtA[i:], *m.String_)
		i = encodeVarintGogovanity(dAtA, i, uint64(len(*m.String_)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGogovanity(dAtA []byte, offset int, v uint64) int {
	offset -= sovGogovanity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovGogovanity(uint64(l))
	}
	n += 1 + sovGogovanity(uint64(m.Int64))
	if m.Int32 != nil {
		n += 1 + sovGogovanity(uint64(*m.Int32))
	}
	return n
}

func sovGogovanity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGogovanity(x uint64) (n int) {
	return sovGogovanity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *B) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&B{`,
		`String_:` + valueToStringGogovanity(this.String_) + `,`,
		`Int64:` + fmt.Sprintf("%v", this.Int64) + `,`,
		`Int32:` + valueToStringGogovanity(this.Int32) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGogovanity(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogovanity
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
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
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
				return ErrInvalidLengthGogovanity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGogovanity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int64 |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogovanity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Int32 = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGogovanity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogovanity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGogovanity
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
func skipGogovanity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGogovanity
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
					return 0, ErrIntOverflowGogovanity
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
					return 0, ErrIntOverflowGogovanity
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
				return 0, ErrInvalidLengthGogovanity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGogovanity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGogovanity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGogovanity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGogovanity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGogovanity = fmt.Errorf("proto: unexpected end of group")
)
