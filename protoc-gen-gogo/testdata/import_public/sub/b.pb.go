// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/sub/b.proto

package sub

import (
	fmt "fmt"
	proto "github.com/temporalio/gogo-protobuf/proto"
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

type M2 struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc66afda3d7c2232, []int{0}
}

var extRange_M2 = []proto.ExtensionRange{
	{Start: 1, End: 536870911},
}

func (*M2) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_M2
}

func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "goproto.test.import_public.sub.M2")
}

func init() { proto.RegisterFile("import_public/sub/b.proto", fileDescriptor_fc66afda3d7c2232) }

var fileDescriptor_fc66afda3d7c2232 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x2f, 0x2e, 0x4d, 0xd2, 0x4f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcf, 0x07, 0x33, 0xf4, 0x4a, 0x52, 0x8b, 0x4b,
	0xf4, 0x50, 0xd4, 0xe9, 0x15, 0x97, 0x26, 0x29, 0xf1, 0x71, 0x31, 0xf9, 0x1a, 0x69, 0x71, 0x70,
	0x30, 0x0a, 0x34, 0x34, 0x34, 0x34, 0x30, 0x39, 0x05, 0x44, 0xf9, 0xa5, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x97, 0xa4, 0x82, 0x34, 0x24, 0xe6, 0x64, 0xe6, 0xeb, 0xa7,
	0xe7, 0xa7, 0xe7, 0xeb, 0x82, 0x8d, 0x4a, 0x2a, 0x4d, 0xd3, 0x07, 0x33, 0x92, 0x75, 0xd3, 0x53,
	0xf3, 0x74, 0x41, 0x12, 0xfa, 0x20, 0xd3, 0x53, 0x12, 0x4b, 0x12, 0xf5, 0x31, 0x5c, 0x02, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x9a, 0x17, 0xf2, 0x9d, 0x00, 0x00, 0x00,
}
