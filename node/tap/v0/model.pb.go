// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node/tap/v0/model.proto

package ptap // import "github.com/n0stack/proto.go/node/tap/v0"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Spec_NetworkType int32

const (
	Spec_FLAT Spec_NetworkType = 0
	Spec_VLAN Spec_NetworkType = 1
)

var Spec_NetworkType_name = map[int32]string{
	0: "FLAT",
	1: "VLAN",
}
var Spec_NetworkType_value = map[string]int32{
	"FLAT": 0,
	"VLAN": 1,
}

func (x Spec_NetworkType) String() string {
	return proto.EnumName(Spec_NetworkType_name, int32(x))
}
func (Spec_NetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{1, 0}
}

type Tap struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec   *Spec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status *Status `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	// 空の場合は自分であると認識する
	NodeId               string   `protobuf:"bytes,4,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{0}
}
func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (dst *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(dst, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tap) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Tap) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Tap) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type Spec struct {
	Type Spec_NetworkType `protobuf:"varint,1,opt,name=type,enum=n0stack.node.tap.Spec_NetworkType" json:"type,omitempty"`
	// FLATの場合は空を許可
	NetworkId            uint64   `protobuf:"varint,2,opt,name=network_id,json=networkId" json:"network_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{1}
}
func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (dst *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(dst, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetType() Spec_NetworkType {
	if m != nil {
		return m.Type
	}
	return Spec_FLAT
}

func (m *Spec) GetNetworkId() uint64 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type Status struct {
	Tap                  string   `protobuf:"bytes,1,opt,name=tap" json:"tap,omitempty"`
	Bridge               string   `protobuf:"bytes,2,opt,name=bridge" json:"bridge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{2}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (dst *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(dst, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetTap() string {
	if m != nil {
		return m.Tap
	}
	return ""
}

func (m *Status) GetBridge() string {
	if m != nil {
		return m.Bridge
	}
	return ""
}

type GetTapRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTapRequest) Reset()         { *m = GetTapRequest{} }
func (m *GetTapRequest) String() string { return proto.CompactTextString(m) }
func (*GetTapRequest) ProtoMessage()    {}
func (*GetTapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{3}
}
func (m *GetTapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTapRequest.Unmarshal(m, b)
}
func (m *GetTapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTapRequest.Marshal(b, m, deterministic)
}
func (dst *GetTapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTapRequest.Merge(dst, src)
}
func (m *GetTapRequest) XXX_Size() int {
	return xxx_messageInfo_GetTapRequest.Size(m)
}
func (m *GetTapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTapRequest proto.InternalMessageInfo

func (m *GetTapRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetTapRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ApplyTapRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec                 *Spec    `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyTapRequest) Reset()         { *m = ApplyTapRequest{} }
func (m *ApplyTapRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyTapRequest) ProtoMessage()    {}
func (*ApplyTapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{4}
}
func (m *ApplyTapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyTapRequest.Unmarshal(m, b)
}
func (m *ApplyTapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyTapRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyTapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyTapRequest.Merge(dst, src)
}
func (m *ApplyTapRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyTapRequest.Size(m)
}
func (m *ApplyTapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyTapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyTapRequest proto.InternalMessageInfo

func (m *ApplyTapRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplyTapRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ApplyTapRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteTapRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTapRequest) Reset()         { *m = DeleteTapRequest{} }
func (m *DeleteTapRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTapRequest) ProtoMessage()    {}
func (*DeleteTapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_2b6544b1e892f7d3, []int{5}
}
func (m *DeleteTapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTapRequest.Unmarshal(m, b)
}
func (m *DeleteTapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTapRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteTapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTapRequest.Merge(dst, src)
}
func (m *DeleteTapRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTapRequest.Size(m)
}
func (m *DeleteTapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTapRequest proto.InternalMessageInfo

func (m *DeleteTapRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteTapRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Tap)(nil), "n0stack.node.tap.Tap")
	proto.RegisterType((*Spec)(nil), "n0stack.node.tap.Spec")
	proto.RegisterType((*Status)(nil), "n0stack.node.tap.Status")
	proto.RegisterType((*GetTapRequest)(nil), "n0stack.node.tap.GetTapRequest")
	proto.RegisterType((*ApplyTapRequest)(nil), "n0stack.node.tap.ApplyTapRequest")
	proto.RegisterType((*DeleteTapRequest)(nil), "n0stack.node.tap.DeleteTapRequest")
	proto.RegisterEnum("n0stack.node.tap.Spec_NetworkType", Spec_NetworkType_name, Spec_NetworkType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapServiceClient is the client API for TapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapServiceClient interface {
	GetTap(ctx context.Context, in *GetTapRequest, opts ...grpc.CallOption) (*Tap, error)
	ApplyTap(ctx context.Context, in *ApplyTapRequest, opts ...grpc.CallOption) (*Tap, error)
	DeleteTap(ctx context.Context, in *DeleteTapRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tapServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapServiceClient(cc *grpc.ClientConn) TapServiceClient {
	return &tapServiceClient{cc}
}

func (c *tapServiceClient) GetTap(ctx context.Context, in *GetTapRequest, opts ...grpc.CallOption) (*Tap, error) {
	out := new(Tap)
	err := c.cc.Invoke(ctx, "/n0stack.node.tap.TapService/GetTap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tapServiceClient) ApplyTap(ctx context.Context, in *ApplyTapRequest, opts ...grpc.CallOption) (*Tap, error) {
	out := new(Tap)
	err := c.cc.Invoke(ctx, "/n0stack.node.tap.TapService/ApplyTap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tapServiceClient) DeleteTap(ctx context.Context, in *DeleteTapRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.node.tap.TapService/DeleteTap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TapService service

type TapServiceServer interface {
	GetTap(context.Context, *GetTapRequest) (*Tap, error)
	ApplyTap(context.Context, *ApplyTapRequest) (*Tap, error)
	DeleteTap(context.Context, *DeleteTapRequest) (*empty.Empty, error)
}

func RegisterTapServiceServer(s *grpc.Server, srv TapServiceServer) {
	s.RegisterService(&_TapService_serviceDesc, srv)
}

func _TapService_GetTap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapServiceServer).GetTap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.tap.TapService/GetTap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapServiceServer).GetTap(ctx, req.(*GetTapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TapService_ApplyTap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyTapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapServiceServer).ApplyTap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.tap.TapService/ApplyTap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapServiceServer).ApplyTap(ctx, req.(*ApplyTapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TapService_DeleteTap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapServiceServer).DeleteTap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.tap.TapService/DeleteTap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapServiceServer).DeleteTap(ctx, req.(*DeleteTapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.node.tap.TapService",
	HandlerType: (*TapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTap",
			Handler:    _TapService_GetTap_Handler,
		},
		{
			MethodName: "ApplyTap",
			Handler:    _TapService_ApplyTap_Handler,
		},
		{
			MethodName: "DeleteTap",
			Handler:    _TapService_DeleteTap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/tap/v0/model.proto",
}

func init() { proto.RegisterFile("node/tap/v0/model.proto", fileDescriptor_model_2b6544b1e892f7d3) }

var fileDescriptor_model_2b6544b1e892f7d3 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x13, 0x63, 0x9a, 0xa9, 0x28, 0xd6, 0x4a, 0xa4, 0x56, 0x10, 0xa2, 0xdd, 0x53, 0x85,
	0xd0, 0x3a, 0x0a, 0x12, 0x42, 0xea, 0x29, 0x15, 0x14, 0x45, 0xaa, 0x7a, 0x70, 0x2d, 0x0e, 0x5c,
	0xd0, 0xc6, 0x3b, 0x35, 0x56, 0x93, 0xec, 0x10, 0x6f, 0x8a, 0x7c, 0xe3, 0x03, 0xf8, 0x4b, 0x7e,
	0x04, 0xed, 0xda, 0x45, 0x69, 0x9c, 0x22, 0xc1, 0x6d, 0x77, 0xe7, 0xcd, 0x9b, 0xf7, 0xc6, 0x7e,
	0x70, 0xb8, 0xd4, 0x0a, 0x63, 0x23, 0x29, 0xbe, 0x1d, 0xc5, 0x0b, 0xad, 0x70, 0x2e, 0x68, 0xa5,
	0x8d, 0x66, 0xe1, 0x72, 0x54, 0x1a, 0x99, 0xdd, 0x08, 0x0b, 0x10, 0x46, 0xd2, 0xf0, 0x79, 0xae,
	0x75, 0x3e, 0xc7, 0xd8, 0xd5, 0x67, 0xeb, 0xeb, 0x18, 0x17, 0x64, 0xaa, 0x1a, 0xce, 0x7f, 0x7a,
	0xd0, 0x4b, 0x25, 0xb1, 0x03, 0xe8, 0x16, 0x2a, 0xf2, 0x8e, 0xbc, 0x93, 0x7e, 0xd2, 0x2d, 0x14,
	0x7b, 0x05, 0x7e, 0x49, 0x98, 0x45, 0xdd, 0x23, 0xef, 0x64, 0x7f, 0x3c, 0x10, 0xdb, 0xac, 0xe2,
	0x8a, 0x30, 0x4b, 0x1c, 0x86, 0x8d, 0x20, 0x28, 0x8d, 0x34, 0xeb, 0x32, 0xea, 0x39, 0x74, 0xb4,
	0x03, 0xed, 0xea, 0x49, 0x83, 0x63, 0x87, 0xf0, 0xd8, 0x96, 0xbe, 0x14, 0x2a, 0xf2, 0xdd, 0xc8,
	0xc0, 0x5e, 0xa7, 0x8a, 0xff, 0xf0, 0xc0, 0xb7, 0xcc, 0xec, 0x2d, 0xf8, 0xa6, 0x22, 0x74, 0x8a,
	0x0e, 0xc6, 0x7c, 0xf7, 0x7c, 0x71, 0x89, 0xe6, 0xbb, 0x5e, 0xdd, 0xa4, 0x15, 0x61, 0xe2, 0xf0,
	0xec, 0x05, 0xc0, 0xb2, 0x7e, 0xb4, 0xe4, 0x56, 0xbd, 0x9f, 0xf4, 0x9b, 0x97, 0xa9, 0xe2, 0xc7,
	0xb0, 0xbf, 0xd1, 0xc3, 0xf6, 0xc0, 0x3f, 0xbf, 0x98, 0xa4, 0x61, 0xc7, 0x9e, 0x3e, 0x5d, 0x4c,
	0x2e, 0x43, 0x8f, 0x8f, 0x21, 0xa8, 0xd5, 0xb2, 0x10, 0x7a, 0x46, 0x52, 0xb3, 0x14, 0x7b, 0x64,
	0x03, 0x08, 0x66, 0xab, 0x42, 0xe5, 0xe8, 0x98, 0xfb, 0x49, 0x73, 0xe3, 0xef, 0xe0, 0xc9, 0x47,
	0x34, 0xa9, 0xa4, 0x04, 0xbf, 0xad, 0xb1, 0x34, 0xad, 0x75, 0x6e, 0x18, 0x7e, 0x74, 0xcf, 0xf0,
	0x35, 0x3c, 0x9d, 0x10, 0xcd, 0xab, 0xbf, 0xf4, 0xfe, 0xcb, 0xa7, 0x78, 0x70, 0xce, 0x29, 0x84,
	0xef, 0x71, 0x8e, 0x06, 0xff, 0x43, 0xe4, 0xf8, 0x97, 0x07, 0x90, 0x4a, 0xba, 0xc2, 0xd5, 0x6d,
	0x91, 0x21, 0x3b, 0x83, 0xa0, 0x76, 0xcb, 0x5e, 0xb6, 0xc5, 0xdc, 0xdb, 0xc3, 0xf0, 0x59, 0x1b,
	0x90, 0x4a, 0xe2, 0x1d, 0x76, 0x0e, 0x7b, 0x77, 0xbe, 0xd9, 0x71, 0x1b, 0xb4, 0xb5, 0x93, 0x87,
	0x79, 0xa6, 0xd0, 0xff, 0xe3, 0x8b, 0xed, 0xf8, 0x4d, 0xb6, 0x4d, 0x0f, 0x07, 0xa2, 0x8e, 0x83,
	0xb8, 0x8b, 0x83, 0xf8, 0x60, 0xe3, 0xc0, 0x3b, 0x67, 0xe2, 0xf3, 0xeb, 0xbc, 0x30, 0x5f, 0xd7,
	0x33, 0x91, 0xe9, 0x45, 0xdc, 0x30, 0xd5, 0xa9, 0x11, 0xb9, 0x8e, 0x37, 0x02, 0x77, 0x4a, 0x46,
	0xd2, 0x2c, 0x70, 0xa5, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x95, 0xb8, 0xe9, 0x8b,
	0x03, 0x00, 0x00,
}
