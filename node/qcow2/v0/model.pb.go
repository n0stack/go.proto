// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node/qcow2/v0/model.proto

package pqcow2 // import "github.com/n0stack/proto.go/node/qcow2/v0"

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

type Qcow2 struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec   *Spec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status *Status `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	// 空の場合は自分であると認識する
	NodeId               string   `protobuf:"bytes,4,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Qcow2) Reset()         { *m = Qcow2{} }
func (m *Qcow2) String() string { return proto.CompactTextString(m) }
func (*Qcow2) ProtoMessage()    {}
func (*Qcow2) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{0}
}
func (m *Qcow2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Qcow2.Unmarshal(m, b)
}
func (m *Qcow2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Qcow2.Marshal(b, m, deterministic)
}
func (dst *Qcow2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Qcow2.Merge(dst, src)
}
func (m *Qcow2) XXX_Size() int {
	return xxx_messageInfo_Qcow2.Size(m)
}
func (m *Qcow2) XXX_DiscardUnknown() {
	xxx_messageInfo_Qcow2.DiscardUnknown(m)
}

var xxx_messageInfo_Qcow2 proto.InternalMessageInfo

func (m *Qcow2) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Qcow2) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Qcow2) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Qcow2) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type Spec struct {
	// サイズを指定する
	Bytes                uint64   `protobuf:"varint,1,opt,name=bytes" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{1}
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

func (m *Spec) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type Status struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{2}
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

func (m *Status) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CreateQcow2WithPackerRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec                 *Spec    `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Repository           string   `protobuf:"bytes,6,opt,name=repository" json:"repository,omitempty"`
	Workdir              string   `protobuf:"bytes,7,opt,name=workdir" json:"workdir,omitempty"`
	TemplateFile         string   `protobuf:"bytes,8,opt,name=template_file,json=templateFile" json:"template_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQcow2WithPackerRequest) Reset()         { *m = CreateQcow2WithPackerRequest{} }
func (m *CreateQcow2WithPackerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQcow2WithPackerRequest) ProtoMessage()    {}
func (*CreateQcow2WithPackerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{3}
}
func (m *CreateQcow2WithPackerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQcow2WithPackerRequest.Unmarshal(m, b)
}
func (m *CreateQcow2WithPackerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQcow2WithPackerRequest.Marshal(b, m, deterministic)
}
func (dst *CreateQcow2WithPackerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQcow2WithPackerRequest.Merge(dst, src)
}
func (m *CreateQcow2WithPackerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQcow2WithPackerRequest.Size(m)
}
func (m *CreateQcow2WithPackerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQcow2WithPackerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQcow2WithPackerRequest proto.InternalMessageInfo

func (m *CreateQcow2WithPackerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateQcow2WithPackerRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateQcow2WithPackerRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *CreateQcow2WithPackerRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *CreateQcow2WithPackerRequest) GetWorkdir() string {
	if m != nil {
		return m.Workdir
	}
	return ""
}

func (m *CreateQcow2WithPackerRequest) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

type CreateQcow2WithDownloadRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec   *Spec  `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	NodeId string `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	// URLからイメージファイルをダウンロードすることができる
	// 別のRPCのRequestに含める予定
	SourceUrl            string   `protobuf:"bytes,6,opt,name=source_url,json=sourceUrl" json:"source_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQcow2WithDownloadRequest) Reset()         { *m = CreateQcow2WithDownloadRequest{} }
func (m *CreateQcow2WithDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQcow2WithDownloadRequest) ProtoMessage()    {}
func (*CreateQcow2WithDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{4}
}
func (m *CreateQcow2WithDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQcow2WithDownloadRequest.Unmarshal(m, b)
}
func (m *CreateQcow2WithDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQcow2WithDownloadRequest.Marshal(b, m, deterministic)
}
func (dst *CreateQcow2WithDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQcow2WithDownloadRequest.Merge(dst, src)
}
func (m *CreateQcow2WithDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQcow2WithDownloadRequest.Size(m)
}
func (m *CreateQcow2WithDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQcow2WithDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQcow2WithDownloadRequest proto.InternalMessageInfo

func (m *CreateQcow2WithDownloadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateQcow2WithDownloadRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateQcow2WithDownloadRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *CreateQcow2WithDownloadRequest) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

type GetQcow2Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQcow2Request) Reset()         { *m = GetQcow2Request{} }
func (m *GetQcow2Request) String() string { return proto.CompactTextString(m) }
func (*GetQcow2Request) ProtoMessage()    {}
func (*GetQcow2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{5}
}
func (m *GetQcow2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQcow2Request.Unmarshal(m, b)
}
func (m *GetQcow2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQcow2Request.Marshal(b, m, deterministic)
}
func (dst *GetQcow2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQcow2Request.Merge(dst, src)
}
func (m *GetQcow2Request) XXX_Size() int {
	return xxx_messageInfo_GetQcow2Request.Size(m)
}
func (m *GetQcow2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQcow2Request.DiscardUnknown(m)
}

var xxx_messageInfo_GetQcow2Request proto.InternalMessageInfo

func (m *GetQcow2Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetQcow2Request) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ApplyQcow2Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec                 *Spec    `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyQcow2Request) Reset()         { *m = ApplyQcow2Request{} }
func (m *ApplyQcow2Request) String() string { return proto.CompactTextString(m) }
func (*ApplyQcow2Request) ProtoMessage()    {}
func (*ApplyQcow2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{6}
}
func (m *ApplyQcow2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyQcow2Request.Unmarshal(m, b)
}
func (m *ApplyQcow2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyQcow2Request.Marshal(b, m, deterministic)
}
func (dst *ApplyQcow2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyQcow2Request.Merge(dst, src)
}
func (m *ApplyQcow2Request) XXX_Size() int {
	return xxx_messageInfo_ApplyQcow2Request.Size(m)
}
func (m *ApplyQcow2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyQcow2Request.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyQcow2Request proto.InternalMessageInfo

func (m *ApplyQcow2Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplyQcow2Request) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ApplyQcow2Request) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteQcow2Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteQcow2Request) Reset()         { *m = DeleteQcow2Request{} }
func (m *DeleteQcow2Request) String() string { return proto.CompactTextString(m) }
func (*DeleteQcow2Request) ProtoMessage()    {}
func (*DeleteQcow2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_83409dc2b1a347c2, []int{7}
}
func (m *DeleteQcow2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQcow2Request.Unmarshal(m, b)
}
func (m *DeleteQcow2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQcow2Request.Marshal(b, m, deterministic)
}
func (dst *DeleteQcow2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQcow2Request.Merge(dst, src)
}
func (m *DeleteQcow2Request) XXX_Size() int {
	return xxx_messageInfo_DeleteQcow2Request.Size(m)
}
func (m *DeleteQcow2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQcow2Request.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQcow2Request proto.InternalMessageInfo

func (m *DeleteQcow2Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteQcow2Request) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func init() {
	proto.RegisterType((*Qcow2)(nil), "n0stack.node.qcow2.Qcow2")
	proto.RegisterType((*Spec)(nil), "n0stack.node.qcow2.Spec")
	proto.RegisterType((*Status)(nil), "n0stack.node.qcow2.Status")
	proto.RegisterType((*CreateQcow2WithPackerRequest)(nil), "n0stack.node.qcow2.CreateQcow2WithPackerRequest")
	proto.RegisterType((*CreateQcow2WithDownloadRequest)(nil), "n0stack.node.qcow2.CreateQcow2WithDownloadRequest")
	proto.RegisterType((*GetQcow2Request)(nil), "n0stack.node.qcow2.GetQcow2Request")
	proto.RegisterType((*ApplyQcow2Request)(nil), "n0stack.node.qcow2.ApplyQcow2Request")
	proto.RegisterType((*DeleteQcow2Request)(nil), "n0stack.node.qcow2.DeleteQcow2Request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Qcow2Service service

type Qcow2ServiceClient interface {
	CreateQcow2WithPacker(ctx context.Context, in *CreateQcow2WithPackerRequest, opts ...grpc.CallOption) (*Qcow2, error)
	CreateQcow2WithDownload(ctx context.Context, in *CreateQcow2WithDownloadRequest, opts ...grpc.CallOption) (*Qcow2, error)
	GetQcow2(ctx context.Context, in *GetQcow2Request, opts ...grpc.CallOption) (*Qcow2, error)
	ApplyQcow2(ctx context.Context, in *ApplyQcow2Request, opts ...grpc.CallOption) (*Qcow2, error)
	DeleteQcow2(ctx context.Context, in *DeleteQcow2Request, opts ...grpc.CallOption) (*empty.Empty, error)
}

type qcow2ServiceClient struct {
	cc *grpc.ClientConn
}

func NewQcow2ServiceClient(cc *grpc.ClientConn) Qcow2ServiceClient {
	return &qcow2ServiceClient{cc}
}

func (c *qcow2ServiceClient) CreateQcow2WithPacker(ctx context.Context, in *CreateQcow2WithPackerRequest, opts ...grpc.CallOption) (*Qcow2, error) {
	out := new(Qcow2)
	err := grpc.Invoke(ctx, "/n0stack.node.qcow2.Qcow2Service/CreateQcow2WithPacker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qcow2ServiceClient) CreateQcow2WithDownload(ctx context.Context, in *CreateQcow2WithDownloadRequest, opts ...grpc.CallOption) (*Qcow2, error) {
	out := new(Qcow2)
	err := grpc.Invoke(ctx, "/n0stack.node.qcow2.Qcow2Service/CreateQcow2WithDownload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qcow2ServiceClient) GetQcow2(ctx context.Context, in *GetQcow2Request, opts ...grpc.CallOption) (*Qcow2, error) {
	out := new(Qcow2)
	err := grpc.Invoke(ctx, "/n0stack.node.qcow2.Qcow2Service/GetQcow2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qcow2ServiceClient) ApplyQcow2(ctx context.Context, in *ApplyQcow2Request, opts ...grpc.CallOption) (*Qcow2, error) {
	out := new(Qcow2)
	err := grpc.Invoke(ctx, "/n0stack.node.qcow2.Qcow2Service/ApplyQcow2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qcow2ServiceClient) DeleteQcow2(ctx context.Context, in *DeleteQcow2Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/n0stack.node.qcow2.Qcow2Service/DeleteQcow2", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Qcow2Service service

type Qcow2ServiceServer interface {
	CreateQcow2WithPacker(context.Context, *CreateQcow2WithPackerRequest) (*Qcow2, error)
	CreateQcow2WithDownload(context.Context, *CreateQcow2WithDownloadRequest) (*Qcow2, error)
	GetQcow2(context.Context, *GetQcow2Request) (*Qcow2, error)
	ApplyQcow2(context.Context, *ApplyQcow2Request) (*Qcow2, error)
	DeleteQcow2(context.Context, *DeleteQcow2Request) (*empty.Empty, error)
}

func RegisterQcow2ServiceServer(s *grpc.Server, srv Qcow2ServiceServer) {
	s.RegisterService(&_Qcow2Service_serviceDesc, srv)
}

func _Qcow2Service_CreateQcow2WithPacker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQcow2WithPackerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Qcow2ServiceServer).CreateQcow2WithPacker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.qcow2.Qcow2Service/CreateQcow2WithPacker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Qcow2ServiceServer).CreateQcow2WithPacker(ctx, req.(*CreateQcow2WithPackerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qcow2Service_CreateQcow2WithDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQcow2WithDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Qcow2ServiceServer).CreateQcow2WithDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.qcow2.Qcow2Service/CreateQcow2WithDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Qcow2ServiceServer).CreateQcow2WithDownload(ctx, req.(*CreateQcow2WithDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qcow2Service_GetQcow2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQcow2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Qcow2ServiceServer).GetQcow2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.qcow2.Qcow2Service/GetQcow2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Qcow2ServiceServer).GetQcow2(ctx, req.(*GetQcow2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qcow2Service_ApplyQcow2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyQcow2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Qcow2ServiceServer).ApplyQcow2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.qcow2.Qcow2Service/ApplyQcow2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Qcow2ServiceServer).ApplyQcow2(ctx, req.(*ApplyQcow2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qcow2Service_DeleteQcow2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQcow2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Qcow2ServiceServer).DeleteQcow2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.qcow2.Qcow2Service/DeleteQcow2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Qcow2ServiceServer).DeleteQcow2(ctx, req.(*DeleteQcow2Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Qcow2Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.node.qcow2.Qcow2Service",
	HandlerType: (*Qcow2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQcow2WithPacker",
			Handler:    _Qcow2Service_CreateQcow2WithPacker_Handler,
		},
		{
			MethodName: "CreateQcow2WithDownload",
			Handler:    _Qcow2Service_CreateQcow2WithDownload_Handler,
		},
		{
			MethodName: "GetQcow2",
			Handler:    _Qcow2Service_GetQcow2_Handler,
		},
		{
			MethodName: "ApplyQcow2",
			Handler:    _Qcow2Service_ApplyQcow2_Handler,
		},
		{
			MethodName: "DeleteQcow2",
			Handler:    _Qcow2Service_DeleteQcow2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/qcow2/v0/model.proto",
}

func init() { proto.RegisterFile("node/qcow2/v0/model.proto", fileDescriptor_model_83409dc2b1a347c2) }

var fileDescriptor_model_83409dc2b1a347c2 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xfa, 0xb9, 0xdd, 0x8d, 0xaf, 0x2b, 0x60, 0x59, 0x19, 0xd3, 0xe4, 0x09, 0xb4, 0x07,
	0xe4, 0x54, 0xe1, 0x0d, 0xc4, 0x03, 0x30, 0x40, 0x48, 0x48, 0x8c, 0x4e, 0x08, 0x89, 0x97, 0x2a,
	0x4d, 0x6e, 0x3b, 0xaf, 0x6e, 0x9d, 0x39, 0xce, 0xaa, 0xfe, 0x0d, 0x1e, 0xf8, 0x63, 0xf0, 0x83,
	0x50, 0x9c, 0x56, 0x94, 0x36, 0x8d, 0x86, 0x04, 0x6f, 0xb6, 0xef, 0xb9, 0xc7, 0xe7, 0xf8, 0x58,
	0x17, 0xf6, 0xc6, 0x2a, 0x22, 0xef, 0x32, 0x54, 0x13, 0xdf, 0xbb, 0x6a, 0x7b, 0x23, 0x15, 0x91,
	0xe4, 0xb1, 0x56, 0x46, 0x21, 0x8e, 0xdb, 0x89, 0x09, 0xc2, 0x21, 0xcf, 0x20, 0xdc, 0x42, 0x5a,
	0x0f, 0x06, 0x4a, 0x0d, 0x24, 0x79, 0x16, 0xd1, 0x4b, 0xfb, 0x1e, 0x8d, 0x62, 0x33, 0xcd, 0x1b,
	0xd8, 0x37, 0x07, 0xea, 0x9f, 0x32, 0x18, 0xde, 0x84, 0x8a, 0x88, 0x5c, 0xe7, 0xd0, 0x39, 0xde,
	0xea, 0x54, 0x44, 0x84, 0x4f, 0xa0, 0x96, 0xc4, 0x14, 0xba, 0x95, 0x43, 0xe7, 0x78, 0xdb, 0x77,
	0xf9, 0x2a, 0x33, 0x3f, 0x8b, 0x29, 0xec, 0x58, 0x14, 0xfa, 0xd0, 0x48, 0x4c, 0x60, 0xd2, 0xc4,
	0xad, 0x5a, 0x7c, 0xab, 0x10, 0x6f, 0x11, 0x9d, 0x19, 0x12, 0x77, 0xa1, 0x99, 0x15, 0xbb, 0x22,
	0x72, 0x6b, 0xf6, 0xda, 0x46, 0xb6, 0x7d, 0x1f, 0xb1, 0x7d, 0xa8, 0x65, 0xd4, 0x78, 0x17, 0xea,
	0xbd, 0xa9, 0xa1, 0xc4, 0xaa, 0xaa, 0x75, 0xf2, 0x0d, 0x6b, 0x41, 0x23, 0x27, 0xc2, 0xdb, 0x50,
	0x4d, 0xb5, 0x9c, 0x69, 0xce, 0x96, 0xec, 0xa7, 0x03, 0xfb, 0xaf, 0x35, 0x05, 0x86, 0xac, 0xa9,
	0x2f, 0xc2, 0x9c, 0x9f, 0x06, 0xe1, 0x90, 0x74, 0x87, 0x2e, 0x53, 0x4a, 0xcc, 0x5a, 0x97, 0xd5,
	0x6b, 0xb9, 0x5c, 0x50, 0x5c, 0x5f, 0x54, 0x8c, 0x07, 0x00, 0x9a, 0x62, 0x95, 0x08, 0xa3, 0xf4,
	0xd4, 0x6d, 0xd8, 0xda, 0xc2, 0x09, 0xba, 0xd0, 0x9c, 0x28, 0x3d, 0x8c, 0x84, 0x76, 0x9b, 0xb6,
	0x38, 0xdf, 0xe2, 0x11, 0xdc, 0x30, 0x34, 0x8a, 0x65, 0x60, 0xa8, 0xdb, 0x17, 0x92, 0xdc, 0x4d,
	0x5b, 0xdf, 0x99, 0x1f, 0xbe, 0x15, 0x92, 0xd8, 0x77, 0x07, 0x0e, 0x96, 0x6c, 0x9d, 0xa8, 0xc9,
	0x58, 0xaa, 0x20, 0xfa, 0xcf, 0xc6, 0x1e, 0x02, 0x24, 0x2a, 0xd5, 0x21, 0x75, 0xb3, 0x97, 0xce,
	0x8d, 0x6d, 0xe5, 0x27, 0x9f, 0xb5, 0x64, 0xcf, 0xe0, 0xd6, 0x3b, 0x32, 0x56, 0xd4, 0x3a, 0x21,
	0xeb, 0xa8, 0xd9, 0x05, 0xdc, 0x79, 0x19, 0xc7, 0x72, 0x5a, 0xda, 0xfd, 0x6f, 0x6c, 0xb0, 0x17,
	0x80, 0x27, 0x24, 0x69, 0xf6, 0x7e, 0x7f, 0x2b, 0xd5, 0xff, 0x51, 0x85, 0x1d, 0xdb, 0x79, 0x46,
	0xfa, 0x4a, 0x84, 0x84, 0x7d, 0xb8, 0x57, 0xf8, 0xcd, 0xb0, 0x5d, 0xa4, 0xb0, 0xec, 0x47, 0xb6,
	0xf6, 0x8a, 0x3a, 0x2c, 0x96, 0x6d, 0xe0, 0x05, 0xec, 0xae, 0xc9, 0x1d, 0xfd, 0x6b, 0xdc, 0xb4,
	0xf4, 0x49, 0xca, 0xef, 0xfa, 0x00, 0x9b, 0xf3, 0x2c, 0xf1, 0xa8, 0x08, 0xb8, 0x94, 0x74, 0x39,
	0xdb, 0x29, 0xc0, 0xef, 0x74, 0xf1, 0x51, 0x11, 0x74, 0x25, 0xfd, 0x72, 0xc6, 0x8f, 0xb0, 0xbd,
	0x90, 0x21, 0x3e, 0x2e, 0xc2, 0xae, 0x86, 0xdc, 0xba, 0xcf, 0xf3, 0xf9, 0xc7, 0xe7, 0xf3, 0x8f,
	0xbf, 0xc9, 0xe6, 0x1f, 0xdb, 0x78, 0xe5, 0x7f, 0x6d, 0x0f, 0x84, 0x39, 0x4f, 0x7b, 0x3c, 0x54,
	0x23, 0x6f, 0xc6, 0x96, 0x8f, 0x49, 0x3e, 0x50, 0xde, 0x1f, 0x53, 0xf6, 0x79, 0x6c, 0x57, 0xbd,
	0x86, 0x2d, 0x3f, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xde, 0x8e, 0xe3, 0x06, 0x84, 0x05, 0x00,
	0x00,
}
