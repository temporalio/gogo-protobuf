// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3/proto3.proto

package proto3

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

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "Sweet",
	1: "Sour",
	2: "Umami",
	3: "Gopherlicious",
}

var Request_Flavour_value = map[string]int32{
	"Sweet":         0,
	"Sour":          1,
	"Umami":         2,
	"Gopherlicious": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}

func (Request_Flavour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0, 0}
}

type Request struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  []int64         `protobuf:"varint,2,rep,packed,name=key,proto3" json:"key,omitempty"`
	Taste                Request_Flavour `protobuf:"varint,3,opt,name=taste,proto3,enum=proto3.Request_Flavour" json:"taste,omitempty"`
	Book                 *Book           `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
	Unpacked             []int64         `protobuf:"varint,5,rep,name=unpacked,proto3" json:"unpacked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
}

func init() { proto.RegisterFile("proto3/proto3.proto", fileDescriptor_ab04eb4084a521db) }

var fileDescriptor_ab04eb4084a521db = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x99, 0xfe, 0x08, 0x5c, 0xd1, 0xd4, 0xd1, 0xc4, 0xba, 0x31, 0x0d, 0xab, 0x6e, 0xda,
	0x26, 0xb8, 0x70, 0xe3, 0x46, 0xb0, 0x2a, 0x89, 0x04, 0x33, 0x48, 0x4c, 0xdc, 0x98, 0x29, 0x8c,
	0x95, 0x14, 0xb8, 0x58, 0x6e, 0x25, 0xbe, 0xac, 0xcf, 0x62, 0xa6, 0x2d, 0xae, 0xee, 0x6f, 0xbe,
	0x93, 0x73, 0xe0, 0x74, 0x93, 0x23, 0xe1, 0x55, 0x54, 0x95, 0xb0, 0x2c, 0xfc, 0xa0, 0x9a, 0xba,
	0xbf, 0x0c, 0x9a, 0x42, 0x7d, 0x15, 0x6a, 0x4b, 0x9c, 0x83, 0xb5, 0x96, 0x2b, 0xe5, 0x32, 0x8f,
	0xf9, 0x6d, 0x51, 0xf6, 0xdc, 0x01, 0x33, 0x53, 0x3f, 0xae, 0xe1, 0x99, 0xbe, 0x29, 0x74, 0xcb,
	0x03, 0xb0, 0x49, 0x6e, 0x49, 0xb9, 0xa6, 0xc7, 0xfc, 0xe3, 0xde, 0x79, 0x58, 0x73, 0x6b, 0x4a,
	0x78, 0xbf, 0x94, 0xdf, 0x58, 0xe4, 0xa2, 0xfa, 0xe2, 0x1e, 0x58, 0x09, 0x62, 0xe6, 0x5a, 0x1e,
	0xf3, 0x0f, 0x7b, 0x9d, 0xfd, 0x77, 0x1f, 0x31, 0x13, 0xe5, 0x85, 0x5f, 0x42, 0xab, 0x58, 0x6f,
	0xe4, 0x2c, 0x53, 0x73, 0xd7, 0xd6, 0x3a, 0x7d, 0xc3, 0x69, 0x88, 0xff, 0x5d, 0xf7, 0x06, 0x9a,
	0x35, 0x93, 0xb7, 0xc1, 0x9e, 0xbc, 0xc6, 0xf1, 0x8b, 0xd3, 0xe0, 0x2d, 0xb0, 0x26, 0xe3, 0xa9,
	0x70, 0x98, 0x5e, 0x4e, 0x47, 0xb7, 0xa3, 0xa1, 0x63, 0xf0, 0x13, 0x38, 0x7a, 0x18, 0x3f, 0x3f,
	0xc6, 0xe2, 0x69, 0x38, 0x18, 0x8e, 0xa7, 0x13, 0xc7, 0xec, 0x5e, 0x83, 0xa5, 0xb5, 0xf8, 0x19,
	0xd8, 0xb4, 0xa0, 0xe5, 0xde, 0x5d, 0x35, 0xf0, 0x0b, 0x68, 0xe5, 0x72, 0xf7, 0x3e, 0x97, 0x24,
	0x5d, 0xc3, 0x63, 0x7e, 0x47, 0x34, 0x73, 0xb9, 0xbb, 0x93, 0x24, 0xfb, 0xf1, 0xdb, 0x20, 0x5d,
	0xd0, 0x67, 0x91, 0x84, 0x33, 0x5c, 0x45, 0xa4, 0x56, 0x1b, 0xcc, 0xe5, 0x72, 0x81, 0x51, 0x8a,
	0x29, 0x06, 0xa5, 0x8d, 0xa4, 0xf8, 0xa8, 0x52, 0x9d, 0x05, 0xa9, 0x5a, 0x07, 0xfa, 0x10, 0x91,
	0xda, 0x92, 0xc6, 0xd5, 0x71, 0x27, 0x75, 0xd0, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x0b,
	0x09, 0x4b, 0x86, 0x01, 0x00, 0x00,
}
