// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvm/v0/model.proto

/*
Package pkvm is a generated protocol buffer package.

It is generated from these files:
	kvm/v0/model.proto

It has these top-level messages:
	KVM
	Spec
	Status
	GetRequest
	ApplyRequest
	DeleteRequest
	ActionRequest
*/
package pkvm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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
func (Status_RunLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type KVM struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec   *Spec   `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
	Status *Status `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *KVM) Reset()                    { *m = KVM{} }
func (m *KVM) String() string            { return proto.CompactTextString(m) }
func (*KVM) ProtoMessage()               {}
func (*KVM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

type Spec struct {
	// CPU
	Vcpus uint32 `protobuf:"varint,1,opt,name=vcpus" json:"vcpus,omitempty"`
	// Memory
	MemoryBytes uint64 `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes" json:"memory_bytes,omitempty"`
	// Storage
	Volumes []string    `protobuf:"bytes,3,rep,name=volumes" json:"volumes,omitempty"`
	Nics    []*Spec_NIC `protobuf:"bytes,4,rep,name=nics" json:"nics,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

func (m *Spec) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Spec) GetNics() []*Spec_NIC {
	if m != nil {
		return m.Nics
	}
	return nil
}

// Network
type Spec_NIC struct {
	Tap    string `protobuf:"bytes,1,opt,name=tap" json:"tap,omitempty"`
	Hwaddr string `protobuf:"bytes,2,opt,name=hwaddr" json:"hwaddr,omitempty"`
}

func (m *Spec_NIC) Reset()                    { *m = Spec_NIC{} }
func (m *Spec_NIC) String() string            { return proto.CompactTextString(m) }
func (*Spec_NIC) ProtoMessage()               {}
func (*Spec_NIC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Spec_NIC) GetTap() string {
	if m != nil {
		return m.Tap
	}
	return ""
}

func (m *Spec_NIC) GetHwaddr() string {
	if m != nil {
		return m.Hwaddr
	}
	return ""
}

type Status struct {
	Node     string          `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	RunLevel Status_RunLevel `protobuf:"varint,2,opt,name=runLevel,enum=n0stack.kvm.Status_RunLevel" json:"runLevel,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Status) GetRunLevel() Status_RunLevel {
	if m != nil {
		return m.RunLevel
	}
	return Status_SHUTDOWN
}

type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ApplyRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec *Spec  `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *ApplyRequest) Reset()                    { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string            { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()               {}
func (*ApplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ApplyRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplyRequest) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ActionRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ActionRequest) Reset()                    { *m = ActionRequest{} }
func (m *ActionRequest) String() string            { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()               {}
func (*ActionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ActionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*KVM)(nil), "n0stack.kvm.KVM")
	proto.RegisterType((*Spec)(nil), "n0stack.kvm.Spec")
	proto.RegisterType((*Spec_NIC)(nil), "n0stack.kvm.Spec.NIC")
	proto.RegisterType((*Status)(nil), "n0stack.kvm.Status")
	proto.RegisterType((*GetRequest)(nil), "n0stack.kvm.GetRequest")
	proto.RegisterType((*ApplyRequest)(nil), "n0stack.kvm.ApplyRequest")
	proto.RegisterType((*DeleteRequest)(nil), "n0stack.kvm.DeleteRequest")
	proto.RegisterType((*ActionRequest)(nil), "n0stack.kvm.ActionRequest")
	proto.RegisterEnum("n0stack.kvm.Status_RunLevel", Status_RunLevel_name, Status_RunLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KVMService service

type KVMServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*KVM, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*KVM, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// VM actions
	Boot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
	Reboot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
	HardReboot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
	Shutdown(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
	HardShutdown(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
	Save(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error)
}

type kVMServiceClient struct {
	cc *grpc.ClientConn
}

func NewKVMServiceClient(cc *grpc.ClientConn) KVMServiceClient {
	return &kVMServiceClient{cc}
}

func (c *kVMServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Boot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Boot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Reboot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Reboot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardReboot(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/HardReboot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Shutdown(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) HardShutdown(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/HardShutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVMServiceClient) Save(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*KVM, error) {
	out := new(KVM)
	err := grpc.Invoke(ctx, "/n0stack.kvm.KVMService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KVMService service

type KVMServiceServer interface {
	Get(context.Context, *GetRequest) (*KVM, error)
	Apply(context.Context, *ApplyRequest) (*KVM, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
	// VM actions
	Boot(context.Context, *ActionRequest) (*KVM, error)
	Reboot(context.Context, *ActionRequest) (*KVM, error)
	HardReboot(context.Context, *ActionRequest) (*KVM, error)
	Shutdown(context.Context, *ActionRequest) (*KVM, error)
	HardShutdown(context.Context, *ActionRequest) (*KVM, error)
	Save(context.Context, *ActionRequest) (*KVM, error)
}

func RegisterKVMServiceServer(s *grpc.Server, srv KVMServiceServer) {
	s.RegisterService(&_KVMService_serviceDesc, srv)
}

func _KVMService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Boot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Boot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Boot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Boot(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Reboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Reboot(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardReboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardReboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/HardReboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardReboot(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Shutdown(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_HardShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).HardShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/HardShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).HardShutdown(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVMService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVMServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n0stack.kvm.KVMService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVMServiceServer).Save(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KVMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "n0stack.kvm.KVMService",
	HandlerType: (*KVMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVMService_Get_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _KVMService_Apply_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVMService_Delete_Handler,
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
	Metadata: "kvm/v0/model.proto",
}

func init() { proto.RegisterFile("kvm/v0/model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x6f, 0x6f, 0x93, 0x5e,
	0x14, 0x1e, 0x85, 0xb1, 0xee, 0xb4, 0x5b, 0xfa, 0xbb, 0x3f, 0x9d, 0x58, 0x97, 0xac, 0x92, 0x18,
	0xbb, 0x98, 0xc0, 0xac, 0xc9, 0xb2, 0xa8, 0x59, 0xb2, 0x7f, 0xe9, 0x9a, 0x5a, 0x34, 0x17, 0x5b,
	0x13, 0xdf, 0x18, 0x0a, 0xd7, 0x96, 0x14, 0xb8, 0x08, 0x17, 0x96, 0x7e, 0x04, 0xbf, 0x83, 0x5f,
	0xc3, 0xef, 0x67, 0xb8, 0x50, 0x1d, 0xb3, 0x4d, 0x5c, 0xdf, 0xdd, 0x73, 0x9e, 0xe7, 0xe1, 0x39,
	0x0f, 0x37, 0xe7, 0x02, 0x9a, 0xa5, 0xbe, 0x9e, 0x1e, 0xe9, 0x3e, 0x75, 0x88, 0xa7, 0x85, 0x11,
	0x65, 0x14, 0xd5, 0x82, 0xa3, 0x98, 0x59, 0xf6, 0x4c, 0x9b, 0xa5, 0x7e, 0xf3, 0xc9, 0x84, 0xd2,
	0x89, 0x47, 0x74, 0x0e, 0x8d, 0x93, 0xaf, 0x3a, 0xf1, 0x43, 0x36, 0xcf, 0x99, 0xaa, 0x0b, 0x62,
	0x7f, 0x34, 0x40, 0xbb, 0x50, 0x71, 0x1d, 0x45, 0x68, 0x09, 0xed, 0x6d, 0x5c, 0x71, 0x1d, 0xf4,
	0x0c, 0xa4, 0x38, 0x24, 0xb6, 0x22, 0xb6, 0x84, 0x76, 0xad, 0xf3, 0x9f, 0x76, 0xeb, 0x7b, 0x9a,
	0x19, 0x12, 0x1b, 0x73, 0x18, 0xbd, 0x00, 0x39, 0x66, 0x16, 0x4b, 0x62, 0x45, 0xe2, 0xc4, 0xff,
	0xcb, 0x44, 0x0e, 0xe1, 0x82, 0xa2, 0xfe, 0x14, 0x40, 0xca, 0xb4, 0xe8, 0x01, 0x6c, 0xa6, 0x76,
	0x98, 0xc4, 0xdc, 0x6f, 0x07, 0xe7, 0x05, 0x7a, 0x0a, 0x75, 0x9f, 0xf8, 0x34, 0x9a, 0x7f, 0x19,
	0xcf, 0x19, 0x89, 0x95, 0x4a, 0x4b, 0x68, 0x4b, 0xb8, 0x96, 0xf7, 0xce, 0xb3, 0x16, 0x52, 0x60,
	0x2b, 0xa5, 0x5e, 0xe2, 0x93, 0x58, 0x11, 0x5b, 0x62, 0x7b, 0x1b, 0x2f, 0x4a, 0x74, 0x08, 0x52,
	0xe0, 0xda, 0xd9, 0x18, 0x62, 0xbb, 0xd6, 0x79, 0xf8, 0xd7, 0xbc, 0x9a, 0xd1, 0xbb, 0xc0, 0x9c,
	0xd2, 0xd4, 0x41, 0x34, 0x7a, 0x17, 0xa8, 0x01, 0x22, 0xb3, 0xc2, 0x22, 0x72, 0x76, 0x44, 0x7b,
	0x20, 0x4f, 0x6f, 0x2c, 0xc7, 0x89, 0xb8, 0xf5, 0x36, 0x2e, 0x2a, 0xf5, 0xbb, 0x00, 0x72, 0x1e,
	0x05, 0x21, 0x90, 0x02, 0xea, 0x90, 0x42, 0xc5, 0xcf, 0xe8, 0x04, 0xaa, 0x51, 0x12, 0xbc, 0x23,
	0x29, 0xf1, 0xb8, 0x70, 0xb7, 0xb3, 0xbf, 0xe4, 0x2f, 0x68, 0xb8, 0xe0, 0xe0, 0xdf, 0x6c, 0xf5,
	0x25, 0x54, 0x17, 0x5d, 0x54, 0x87, 0xaa, 0x79, 0x3d, 0xfc, 0x78, 0xf9, 0xfe, 0x93, 0xd1, 0xd8,
	0x40, 0x35, 0xd8, 0xc2, 0x43, 0xc3, 0xe8, 0x19, 0xdd, 0x86, 0x80, 0x00, 0xe4, 0x0f, 0x67, 0x43,
	0xf3, 0xea, 0xb2, 0x51, 0x51, 0xf7, 0x01, 0xba, 0x84, 0x61, 0xf2, 0x2d, 0x21, 0x31, 0xbb, 0x7b,
	0x6b, 0xea, 0x15, 0xd4, 0xcf, 0xc2, 0xd0, 0x9b, 0xaf, 0xc0, 0xff, 0xf1, 0x56, 0xd5, 0x03, 0xd8,
	0xb9, 0x24, 0x1e, 0x61, 0x64, 0x95, 0xcf, 0x01, 0xec, 0x9c, 0xd9, 0xcc, 0xa5, 0xc1, 0x0a, 0x42,
	0xe7, 0x87, 0x04, 0xd0, 0x1f, 0x0d, 0x4c, 0x12, 0xa5, 0xae, 0x4d, 0x50, 0x07, 0xc4, 0x2e, 0x61,
	0xe8, 0x51, 0xc9, 0xf0, 0x4f, 0x8e, 0x66, 0xa3, 0x04, 0xf4, 0x47, 0x03, 0x75, 0x03, 0x1d, 0xc3,
	0x26, 0xcf, 0x82, 0x1e, 0x97, 0xc0, 0xdb, 0xf9, 0x96, 0xea, 0x4e, 0x41, 0xce, 0x87, 0x47, 0xcd,
	0x12, 0x5a, 0x4a, 0xd4, 0xdc, 0xd3, 0xf2, 0xa5, 0xd0, 0x16, 0x4b, 0xa1, 0x5d, 0x65, 0x4b, 0xc1,
	0x7d, 0xa5, 0x73, 0x4a, 0xd9, 0x1d, 0x75, 0x29, 0xee, 0x52, 0xdf, 0x13, 0x90, 0x31, 0x19, 0xaf,
	0xa3, 0x7c, 0x0b, 0x70, 0x6d, 0x45, 0xce, 0x9a, 0xea, 0xd7, 0x50, 0x35, 0xa7, 0x09, 0x73, 0xe8,
	0x4d, 0x70, 0x6f, 0xed, 0x29, 0xd4, 0x33, 0xe7, 0xb5, 0xf5, 0xc7, 0x20, 0x99, 0x56, 0x4a, 0xee,
	0xab, 0x3b, 0x3f, 0xfc, 0xfc, 0x7c, 0xe2, 0xb2, 0x69, 0x32, 0xd6, 0x6c, 0xea, 0xeb, 0x05, 0x9e,
	0xbf, 0x4f, 0xda, 0x84, 0xea, 0xf9, 0x83, 0xf6, 0x26, 0x9c, 0xa5, 0xfe, 0x58, 0xe6, 0xdd, 0x57,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x77, 0xac, 0x84, 0xe6, 0x04, 0x00, 0x00,
}
