// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node/kvm/v0/model.proto

package pkvm // import "github.com/n0stack/proto.go/node/kvm/v0"

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

type Status_RunLevel int32

const (
	Status_SHUTDOWN Status_RunLevel = 0
	Status_RUNNING  Status_RunLevel = 1
	Status_PAUSED   Status_RunLevel = 2
)

var Status_RunLevel_name = map[int32]string{
	0: "SHUTDOWN",
	1: "RUNNING",
	2: "PAUSED",
}
var Status_RunLevel_value = map[string]int32{
	"SHUTDOWN": 0,
	"RUNNING":  1,
	"PAUSED":   2,
}

func (x Status_RunLevel) String() string {
	return proto.EnumName(Status_RunLevel_name, int32(x))
}
func (Status_RunLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{2, 0}
}

type KVM struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec   *Spec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status *Status `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	// 空の場合は自分であると認識する
	NodeId               string   `protobuf:"bytes,4,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVM) Reset()         { *m = KVM{} }
func (m *KVM) String() string { return proto.CompactTextString(m) }
func (*KVM) ProtoMessage()    {}
func (*KVM) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{0}
}
func (m *KVM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVM.Unmarshal(m, b)
}
func (m *KVM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVM.Marshal(b, m, deterministic)
}
func (dst *KVM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVM.Merge(dst, src)
}
func (m *KVM) XXX_Size() int {
	return xxx_messageInfo_KVM.Size(m)
}
func (m *KVM) XXX_DiscardUnknown() {
	xxx_messageInfo_KVM.DiscardUnknown(m)
}

var xxx_messageInfo_KVM proto.InternalMessageInfo

func (m *KVM) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KVM) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *KVM) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *KVM) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type Spec struct {
	// CPU
	Vcpus uint32 `protobuf:"varint,1,opt,name=vcpus" json:"vcpus,omitempty"`
	// Memory
	MemoryBytes uint64 `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes" json:"memory_bytes,omitempty"`
	// Storage
	Qcow2Ids             []string       `protobuf:"bytes,3,rep,name=qcow2_ids,json=qcow2Ids" json:"qcow2_ids,omitempty"`
	Nics                 []*Spec_Netdev `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{1}
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

func (m *Spec) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *Spec) GetMemoryBytes() uint64 {
	if m != nil {
		return m.MemoryBytes
	}
	return 0
}

func (m *Spec) GetQcow2Ids() []string {
	if m != nil {
		return m.Qcow2Ids
	}
	return nil
}

func (m *Spec) GetNics() []*Spec_Netdev {
	if m != nil {
		return m.Nics
	}
	return nil
}

// Network
type Spec_Netdev struct {
	TapId                string   `protobuf:"bytes,1,opt,name=tap_id,json=tapId" json:"tap_id,omitempty"`
	Hwaddr               string   `protobuf:"bytes,2,opt,name=hwaddr" json:"hwaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spec_Netdev) Reset()         { *m = Spec_Netdev{} }
func (m *Spec_Netdev) String() string { return proto.CompactTextString(m) }
func (*Spec_Netdev) ProtoMessage()    {}
func (*Spec_Netdev) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{1, 0}
}
func (m *Spec_Netdev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec_Netdev.Unmarshal(m, b)
}
func (m *Spec_Netdev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec_Netdev.Marshal(b, m, deterministic)
}
func (dst *Spec_Netdev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec_Netdev.Merge(dst, src)
}
func (m *Spec_Netdev) XXX_Size() int {
	return xxx_messageInfo_Spec_Netdev.Size(m)
}
func (m *Spec_Netdev) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec_Netdev.DiscardUnknown(m)
}

var xxx_messageInfo_Spec_Netdev proto.InternalMessageInfo

func (m *Spec_Netdev) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

func (m *Spec_Netdev) GetHwaddr() string {
	if m != nil {
		return m.Hwaddr
	}
	return ""
}

type Status struct {
	RunLevel Status_RunLevel `protobuf:"varint,1,opt,name=run_level,json=runLevel,enum=n0stack.node.kvm.Status_RunLevel" json:"run_level,omitempty"`
	// TCP port of websocket vnc which is opened by qemu
	VncPort              uint32   `protobuf:"varint,2,opt,name=vnc_port,json=vncPort" json:"vnc_port,omitempty"`
	VncWebsocketPort     uint32   `protobuf:"varint,3,opt,name=vnc_websocket_port,json=vncWebsocketPort" json:"vnc_websocket_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{2}
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

func (m *Status) GetRunLevel() Status_RunLevel {
	if m != nil {
		return m.RunLevel
	}
	return Status_SHUTDOWN
}

func (m *Status) GetVncPort() uint32 {
	if m != nil {
		return m.VncPort
	}
	return 0
}

func (m *Status) GetVncWebsocketPort() uint32 {
	if m != nil {
		return m.VncWebsocketPort
	}
	return 0
}

type GetKVMRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKVMRequest) Reset()         { *m = GetKVMRequest{} }
func (m *GetKVMRequest) String() string { return proto.CompactTextString(m) }
func (*GetKVMRequest) ProtoMessage()    {}
func (*GetKVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{3}
}
func (m *GetKVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKVMRequest.Unmarshal(m, b)
}
func (m *GetKVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKVMRequest.Marshal(b, m, deterministic)
}
func (dst *GetKVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKVMRequest.Merge(dst, src)
}
func (m *GetKVMRequest) XXX_Size() int {
	return xxx_messageInfo_GetKVMRequest.Size(m)
}
func (m *GetKVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKVMRequest proto.InternalMessageInfo

func (m *GetKVMRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetKVMRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ApplyKVMRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec                 *Spec    `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyKVMRequest) Reset()         { *m = ApplyKVMRequest{} }
func (m *ApplyKVMRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyKVMRequest) ProtoMessage()    {}
func (*ApplyKVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{4}
}
func (m *ApplyKVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyKVMRequest.Unmarshal(m, b)
}
func (m *ApplyKVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyKVMRequest.Marshal(b, m, deterministic)
}
func (dst *ApplyKVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyKVMRequest.Merge(dst, src)
}
func (m *ApplyKVMRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyKVMRequest.Size(m)
}
func (m *ApplyKVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyKVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyKVMRequest proto.InternalMessageInfo

func (m *ApplyKVMRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplyKVMRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ApplyKVMRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type DeleteKVMRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKVMRequest) Reset()         { *m = DeleteKVMRequest{} }
func (m *DeleteKVMRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteKVMRequest) ProtoMessage()    {}
func (*DeleteKVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{5}
}
func (m *DeleteKVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKVMRequest.Unmarshal(m, b)
}
func (m *DeleteKVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKVMRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteKVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKVMRequest.Merge(dst, src)
}
func (m *DeleteKVMRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteKVMRequest.Size(m)
}
func (m *DeleteKVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKVMRequest proto.InternalMessageInfo

func (m *DeleteKVMRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteKVMRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ActionKVMRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionKVMRequest) Reset()         { *m = ActionKVMRequest{} }
func (m *ActionKVMRequest) String() string { return proto.CompactTextString(m) }
func (*ActionKVMRequest) ProtoMessage()    {}
func (*ActionKVMRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_da0852159d3d4774, []int{6}
}
func (m *ActionKVMRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionKVMRequest.Unmarshal(m, b)
}
func (m *ActionKVMRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionKVMRequest.Marshal(b, m, deterministic)
}
func (dst *ActionKVMRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionKVMRequest.Merge(dst, src)
}
func (m *ActionKVMRequest) XXX_Size() int {
	return xxx_messageInfo_ActionKVMRequest.Size(m)
}
func (m *ActionKVMRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionKVMRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionKVMRequest proto.InternalMessageInfo

func (m *ActionKVMRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ActionKVMRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func init() {
	proto.RegisterType((*KVM)(nil), "n0stack.node.kvm.KVM")
	proto.RegisterType((*Spec)(nil), "n0stack.node.kvm.Spec")
	proto.RegisterType((*Spec_Netdev)(nil), "n0stack.node.kvm.Spec.Netdev")
	proto.RegisterType((*Status)(nil), "n0stack.node.kvm.Status")
	proto.RegisterType((*GetKVMRequest)(nil), "n0stack.node.kvm.GetKVMRequest")
	proto.RegisterType((*ApplyKVMRequest)(nil), "n0stack.node.kvm.ApplyKVMRequest")
	proto.RegisterType((*DeleteKVMRequest)(nil), "n0stack.node.kvm.DeleteKVMRequest")
	proto.RegisterType((*ActionKVMRequest)(nil), "n0stack.node.kvm.ActionKVMRequest")
	proto.RegisterEnum("n0stack.node.kvm.Status_RunLevel", Status_RunLevel_name, Status_RunLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVMServiceClient is the client API for KVMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVMServiceClient interface {
	GetKVM(ctx context.Context, in *GetKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	ApplyKVM(ctx context.Context, in *ApplyKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	DeleteKVM(ctx context.Context, in *DeleteKVMRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// VM actions
	Boot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	Reboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	HardReboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	Shutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	HardShutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
	Save(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error)
}

type kVMServiceClient struct {
	cc *grpc.ClientConn
}

func NewKVMServiceClient(cc *grpc.ClientConn) KVMServiceClient {
	return &kVMServiceClient{cc}
}

func (c *kVMServiceClient) GetKVM(ctx context.Context, in *GetKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/GetKVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) ApplyKVM(ctx context.Context, in *ApplyKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/ApplyKVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) DeleteKVM(ctx context.Context, in *DeleteKVMRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/DeleteKVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Boot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/Boot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Reboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/Reboot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardReboot(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/HardReboot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Shutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardShutdown(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/HardShutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Save(ctx context.Context, in *ActionKVMRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := c.cc.Invoke(ctx, "/n0stack.node.kvm.KVMService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KVMService service

type KVMServiceServer interface {
	GetKVM(context.Context, *GetKVMRequest) (*KVM, error)
	ApplyKVM(context.Context, *ApplyKVMRequest) (*KVM, error)
	DeleteKVM(context.Context, *DeleteKVMRequest) (*empty.Empty, error)
	// VM actions
	Boot(context.Context, *ActionKVMRequest) (*KVM, error)
	Reboot(context.Context, *ActionKVMRequest) (*KVM, error)
	HardReboot(context.Context, *ActionKVMRequest) (*KVM, error)
	Shutdown(context.Context, *ActionKVMRequest) (*KVM, error)
	HardShutdown(context.Context, *ActionKVMRequest) (*KVM, error)
	Save(context.Context, *ActionKVMRequest) (*KVM, error)
}

func RegisterKVMServiceServer(s *grpc.Server, srv KVMServiceServer) {
	s.RegisterService(&_KVMService_serviceDesc, srv)
}

func _KVMService_GetKVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).GetKVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/GetKVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).GetKVM(ctx, req.(*GetKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_ApplyKVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).ApplyKVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/ApplyKVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).ApplyKVM(ctx, req.(*ApplyKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_DeleteKVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).DeleteKVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/DeleteKVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).DeleteKVM(ctx, req.(*DeleteKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Boot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Boot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/Boot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Boot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/Reboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Reboot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardReboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardReboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/HardReboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardReboot(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Shutdown(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/HardShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardShutdown(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionKVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.node.kvm.KVMService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Save(ctx, req.(*ActionKVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.node.kvm.KVMService",
	HandlerType: (*KVMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKVM",
			Handler:    _KVMService_GetKVM_Handler,
		},
		{
			MethodName: "ApplyKVM",
			Handler:    _KVMService_ApplyKVM_Handler,
		},
		{
			MethodName: "DeleteKVM",
			Handler:    _KVMService_DeleteKVM_Handler,
		},
		{
			MethodName: "Boot",
			Handler:    _KVMService_Boot_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _KVMService_Reboot_Handler,
		},
		{
			MethodName: "HardReboot",
			Handler:    _KVMService_HardReboot_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _KVMService_Shutdown_Handler,
		},
		{
			MethodName: "HardShutdown",
			Handler:    _KVMService_HardShutdown_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _KVMService_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node/kvm/v0/model.proto",
}

func init() { proto.RegisterFile("node/kvm/v0/model.proto", fileDescriptor_model_da0852159d3d4774) }

var fileDescriptor_model_da0852159d3d4774 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0xc6, 0x69, 0x9b, 0x86, 0xf4, 0x14, 0x58, 0x64, 0x0d, 0xe8, 0x8a, 0xa6, 0x41, 0xae, 0xaa,
	0x09, 0x25, 0xd0, 0x5d, 0x6c, 0x12, 0xd2, 0x24, 0x0a, 0x0c, 0xaa, 0xae, 0x1d, 0x72, 0x47, 0x91,
	0x76, 0x53, 0xa5, 0xb6, 0x29, 0x51, 0x9b, 0x38, 0x24, 0x4e, 0xaa, 0xbe, 0xc3, 0x6e, 0xf7, 0x46,
	0xbb, 0xd8, 0x63, 0x4d, 0x71, 0x5a, 0x04, 0xfd, 0x33, 0x4d, 0xc0, 0x5d, 0x8e, 0xcf, 0xcf, 0x9f,
	0x8f, 0x3f, 0x1f, 0x9d, 0xc0, 0xb6, 0xc7, 0x29, 0xb3, 0x06, 0xb1, 0x6b, 0xc5, 0x07, 0x96, 0xcb,
	0x29, 0x1b, 0x9a, 0x7e, 0xc0, 0x05, 0x47, 0xba, 0x77, 0x10, 0x0a, 0x9b, 0x0c, 0xcc, 0x04, 0x30,
	0x07, 0xb1, 0x5b, 0xde, 0xe9, 0x73, 0xde, 0x1f, 0x32, 0x4b, 0xe6, 0x7b, 0xd1, 0x8d, 0xc5, 0x5c,
	0x5f, 0x8c, 0x53, 0xdc, 0xf8, 0x99, 0x81, 0x5c, 0xa3, 0xd3, 0x44, 0x1b, 0x90, 0x75, 0x68, 0x29,
	0xb3, 0x9b, 0xa9, 0x14, 0x70, 0xd6, 0xa1, 0xe8, 0x3d, 0x28, 0xa1, 0xcf, 0x48, 0x29, 0xbb, 0x9b,
	0xa9, 0x14, 0xab, 0x5b, 0xe6, 0xac, 0xaa, 0xd9, 0xf6, 0x19, 0xc1, 0x92, 0x41, 0x07, 0xa0, 0x86,
	0xc2, 0x16, 0x51, 0x58, 0xca, 0x49, 0xba, 0xb4, 0x80, 0x96, 0x79, 0x3c, 0xe1, 0xd0, 0x36, 0xac,
	0x26, 0xa9, 0xae, 0x43, 0x4b, 0x8a, 0x3c, 0x52, 0x4d, 0xc2, 0x3a, 0x35, 0xfe, 0x64, 0x40, 0x49,
	0x94, 0xd1, 0x6b, 0xc8, 0xc7, 0xc4, 0x8f, 0x42, 0x59, 0xd2, 0x3a, 0x4e, 0x03, 0xb4, 0x07, 0x6b,
	0x2e, 0x73, 0x79, 0x30, 0xee, 0xf6, 0xc6, 0x82, 0x85, 0xb2, 0x3a, 0x05, 0x17, 0xd3, 0xb5, 0x5a,
	0xb2, 0x84, 0x76, 0xa0, 0x70, 0x47, 0xf8, 0xa8, 0xda, 0x75, 0x68, 0x52, 0x4f, 0xae, 0x52, 0xc0,
	0x9a, 0x5c, 0xa8, 0xd3, 0x10, 0x1d, 0x82, 0xe2, 0x39, 0x24, 0x2c, 0x29, 0xbb, 0xb9, 0x4a, 0xb1,
	0xfa, 0x76, 0xf1, 0xad, 0xcc, 0x16, 0x13, 0x94, 0xc5, 0x58, 0xa2, 0xe5, 0x8f, 0xa0, 0xa6, 0x31,
	0xda, 0x04, 0x55, 0xd8, 0x7e, 0xf7, 0xde, 0xa6, 0xbc, 0xb0, 0xfd, 0x3a, 0x45, 0x5b, 0xa0, 0xde,
	0x8e, 0x6c, 0x4a, 0x03, 0x59, 0x4d, 0x01, 0x4f, 0x22, 0xe3, 0x77, 0x06, 0xd4, 0xf4, 0xda, 0xe8,
	0x33, 0x14, 0x82, 0xc8, 0xeb, 0x0e, 0x59, 0xcc, 0x86, 0x72, 0xf3, 0x46, 0x75, 0x6f, 0x99, 0x47,
	0x26, 0x8e, 0xbc, 0xaf, 0x09, 0x88, 0xb5, 0x60, 0xf2, 0x85, 0xde, 0x80, 0x16, 0x7b, 0xa4, 0xeb,
	0xf3, 0x40, 0xc8, 0x43, 0xd6, 0xf1, 0x6a, 0xec, 0x91, 0x4b, 0x1e, 0x08, 0xb4, 0x0f, 0x28, 0x49,
	0x8d, 0x58, 0x2f, 0xe4, 0x64, 0xc0, 0x44, 0x0a, 0xe5, 0x24, 0xa4, 0xc7, 0x1e, 0xb9, 0x9e, 0x26,
	0x12, 0xda, 0x38, 0x04, 0x6d, 0x2a, 0x8f, 0xd6, 0x40, 0x6b, 0x5f, 0x5c, 0x7d, 0x3f, 0xfd, 0x76,
	0xdd, 0xd2, 0x57, 0x50, 0x11, 0x56, 0xf1, 0x55, 0xab, 0x55, 0x6f, 0x9d, 0xeb, 0x19, 0x04, 0xa0,
	0x5e, 0x1e, 0x5f, 0xb5, 0xcf, 0x4e, 0xf5, 0xac, 0xf1, 0x09, 0xd6, 0xcf, 0x99, 0x68, 0x74, 0x9a,
	0x98, 0xdd, 0x45, 0x2c, 0x14, 0x73, 0x9d, 0xf2, 0xe0, 0x2d, 0xf3, 0x8f, 0xde, 0xf2, 0x06, 0x5e,
	0x1d, 0xfb, 0xfe, 0x70, 0xfc, 0x8f, 0xbd, 0xd3, 0x2e, 0xcb, 0xfd, 0x47, 0x97, 0x2d, 0x3d, 0xe7,
	0x08, 0xf4, 0x53, 0x36, 0x64, 0x82, 0x3d, 0xa5, 0xc8, 0x23, 0xd0, 0x8f, 0x89, 0x70, 0xb8, 0xf7,
	0x84, 0xcd, 0xd5, 0x5f, 0x79, 0x80, 0x46, 0xa7, 0xd9, 0x66, 0x41, 0xec, 0x10, 0x86, 0x6a, 0xa0,
	0xa6, 0x56, 0xa1, 0x77, 0xf3, 0x37, 0x79, 0x64, 0x62, 0x79, 0x73, 0x1e, 0x68, 0x74, 0x9a, 0xc6,
	0x0a, 0xfa, 0x02, 0xda, 0xd4, 0x34, 0xb4, 0xa0, 0x47, 0x66, 0x0c, 0x5d, 0xae, 0x53, 0x87, 0xc2,
	0xbd, 0x29, 0xc8, 0x98, 0xa7, 0x66, 0x1d, 0x2b, 0x6f, 0x99, 0xe9, 0x98, 0x30, 0xa7, 0x63, 0xc2,
	0x3c, 0x4b, 0xc6, 0x84, 0xb1, 0x82, 0x4e, 0x40, 0xa9, 0x71, 0x2e, 0x16, 0xa9, 0xcc, 0x5a, 0xb7,
	0xbc, 0x9e, 0x33, 0x50, 0x31, 0xeb, 0x3d, 0x5b, 0xa6, 0x0e, 0x70, 0x61, 0x07, 0xf4, 0x25, 0xa4,
	0xce, 0x41, 0x6b, 0xdf, 0x46, 0x82, 0xf2, 0x91, 0xf7, 0x3c, 0xa1, 0x06, 0xac, 0x25, 0x35, 0xbd,
	0x8c, 0xd8, 0x09, 0x28, 0x6d, 0x3b, 0x66, 0xcf, 0x12, 0xa9, 0x99, 0x3f, 0xf6, 0xfb, 0x8e, 0xb8,
	0x8d, 0x7a, 0x26, 0xe1, 0xae, 0x35, 0x81, 0xd2, 0xf9, 0x6f, 0xf6, 0xb9, 0xf5, 0xe0, 0xd7, 0x71,
	0xe4, 0x0f, 0x62, 0xb7, 0xa7, 0xca, 0xd4, 0x87, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x0a,
	0xfc, 0x4d, 0x55, 0x06, 0x00, 0x00,
}
