// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_b_1/m2.proto

package beta

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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1becddceeb586f2, []int{0}
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
	proto.RegisterType((*M2)(nil), "test.b.part2.M2")
}

func init() { proto.RegisterFile("imports/test_b_1/m2.proto", fileDescriptor_a1becddceeb586f2) }

var fileDescriptor_a1becddceeb586f2 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8a, 0x37, 0xd4, 0xcf, 0x35, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x09, 0xe9, 0x25, 0xe9, 0x15, 0x24, 0x16, 0x95,
	0x18, 0x29, 0xb1, 0x70, 0x31, 0xf9, 0x1a, 0x39, 0x85, 0x44, 0x05, 0xa5, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x97, 0xa4, 0x82, 0x34, 0x27, 0xe6, 0x64, 0xe6, 0xeb, 0xa7,
	0xe7, 0xa7, 0xe7, 0xeb, 0x82, 0xf5, 0x25, 0x95, 0xa6, 0xe9, 0x83, 0x19, 0xc9, 0xba, 0xe9, 0xa9,
	0x79, 0xba, 0x20, 0x09, 0xb0, 0xe1, 0x29, 0x89, 0x25, 0x89, 0xfa, 0xe8, 0xb6, 0x59, 0x27, 0xa5,
	0x96, 0x24, 0x26, 0xb1, 0x81, 0xd5, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xd8, 0xf5,
	0x1b, 0x8d, 0x00, 0x00, 0x00,
}
