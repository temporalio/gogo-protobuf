// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/b.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/temporalio/gogo-protobuf/proto"
	sub "github.com/temporalio/gogo-protobuf/protoc-gen-gogo/testdata/import_public/sub"
	math "math"
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

type Local struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.test.import_public.sub.E" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Local) Reset()         { *m = Local{} }
func (m *Local) String() string { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()    {}
func (*Local) Descriptor() ([]byte, []int) {
	return fileDescriptor_84995586b3d09710, []int{0}
}
func (m *Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Local.Unmarshal(m, b)
}
func (m *Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Local.Marshal(b, m, deterministic)
}
func (m *Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Local.Merge(m, src)
}
func (m *Local) XXX_Size() int {
	return xxx_messageInfo_Local.Size(m)
}
func (m *Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Local proto.InternalMessageInfo

func (m *Local) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Local) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func init() {
	proto.RegisterType((*Local)(nil), "goproto.test.import_public.Local")
}

func init() { proto.RegisterFile("import_public/b.proto", fileDescriptor_84995586b3d09710) }

var fileDescriptor_84995586b3d09710 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcc, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0xf2, 0xc2, 0xbb, 0x54, 0x70, 0x28, 0x08, 0xb5, 0x53, 0x75, 0xea, 0xd2, 0x3b,
	0xf0, 0x23, 0x08, 0x2e, 0x52, 0x97, 0x8e, 0x2e, 0x92, 0x8b, 0x31, 0x06, 0x1a, 0xaf, 0xb4, 0x97,
	0xef, 0x2f, 0xa9, 0x53, 0x07, 0x71, 0x3b, 0xb8, 0xe7, 0xf7, 0xcf, 0x36, 0x3e, 0x0c, 0x3c, 0xca,
	0x6d, 0x88, 0xd4, 0x7b, 0x83, 0x04, 0xc3, 0xc8, 0xc2, 0x79, 0xe9, 0x78, 0x3e, 0x40, 0xec, 0x24,
	0xb0, 0xd8, 0x94, 0xdb, 0x25, 0x99, 0x22, 0xa1, 0xfe, 0xb0, 0xbd, 0xcf, 0xfe, 0x5b, 0x36, 0xba,
	0xcf, 0x31, 0x53, 0xa1, 0x50, 0x95, 0xaa, 0x57, 0x87, 0x1d, 0x7c, 0x6f, 0xc1, 0x14, 0x09, 0x2e,
	0x9d, 0x0a, 0x09, 0xd8, 0xe2, 0xaf, 0x52, 0xf5, 0xfa, 0x37, 0x38, 0x75, 0xca, 0x1e, 0xdb, 0xeb,
	0xd9, 0x79, 0x79, 0x46, 0x02, 0xc3, 0x01, 0xc5, 0xa6, 0x95, 0xee, 0x3d, 0xa3, 0x63, 0xc7, 0xcd,
	0xec, 0x29, 0x3e, 0x70, 0x3e, 0x4c, 0xe3, 0xec, 0xab, 0x49, 0x0f, 0x4c, 0xc9, 0xbb, 0x16, 0x8d,
	0x8b, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x99, 0xeb, 0xce, 0x3a, 0x07, 0x01, 0x00, 0x00,
}
