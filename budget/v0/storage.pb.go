// Code generated by protoc-gen-go. DO NOT EDIT.
// source: budget/v0/storage.proto

package pbudget // import "github.com/n0stack/proto.go/budget/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Storage struct {
	Annotations          map[string]string `protobuf:"bytes,1,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RequestBytes         uint64            `protobuf:"varint,2,opt,name=request_bytes,json=requestBytes" json:"request_bytes,omitempty"`
	LimitBytes           uint64            `protobuf:"varint,3,opt,name=limit_bytes,json=limitBytes" json:"limit_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_1a2b0a94868256f8, []int{0}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (dst *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(dst, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Storage) GetRequestBytes() uint64 {
	if m != nil {
		return m.RequestBytes
	}
	return 0
}

func (m *Storage) GetLimitBytes() uint64 {
	if m != nil {
		return m.LimitBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Storage)(nil), "n0stack.budget.Storage")
	proto.RegisterMapType((map[string]string)(nil), "n0stack.budget.Storage.AnnotationsEntry")
}

func init() { proto.RegisterFile("budget/v0/storage.proto", fileDescriptor_storage_1a2b0a94868256f8) }

var fileDescriptor_storage_1a2b0a94868256f8 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2a, 0x4d, 0x49,
	0x4f, 0x2d, 0xd1, 0x2f, 0x33, 0xd0, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0x33, 0x28, 0x2e, 0x49, 0x4c, 0xce, 0xd6, 0x83, 0x28, 0x50,
	0xba, 0xcd, 0xc8, 0xc5, 0x1e, 0x0c, 0x51, 0x21, 0xe4, 0xc5, 0xc5, 0x9d, 0x98, 0x97, 0x97, 0x5f,
	0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa1, 0x87,
	0xaa, 0x43, 0x0f, 0xaa, 0x5a, 0xcf, 0x11, 0xa1, 0xd4, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0x59,
	0xb3, 0x90, 0x32, 0x17, 0x6f, 0x51, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x49, 0x7c, 0x52, 0x65, 0x49,
	0x6a, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x0f, 0x54, 0xd0, 0x09, 0x24, 0x26, 0x24,
	0xcf, 0xc5, 0x9d, 0x93, 0x99, 0x9b, 0x09, 0x53, 0xc2, 0x0c, 0x56, 0xc2, 0x05, 0x16, 0x02, 0x2b,
	0x90, 0xb2, 0xe3, 0x12, 0x40, 0xb7, 0x46, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05,
	0xdb, 0xc1, 0x19, 0x04, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x3a, 0xe9, 0x47, 0xe9, 0xa6, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x3d, 0xa2, 0x0f, 0x0e, 0x09, 0xbd, 0xf4,
	0x7c, 0x7d, 0x78, 0x20, 0x59, 0x17, 0x40, 0x98, 0x49, 0x6c, 0x60, 0x39, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0xae, 0x5e, 0xec, 0x40, 0x01, 0x00, 0x00,
}
