// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fx.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	fx.proto

It has these top-level messages:
	FunctionMeta
	UpRequest
	DownRequest
	ListRequest
	UpMsgMeta
	UpResponse
	DownMsgMeta
	DownResponse
	ListItem
	ListResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type FunctionMeta struct {
	Lang    string `protobuf:"bytes,1,opt,name=Lang,json=lang" json:"Lang,omitempty"`
	Path    string `protobuf:"bytes,2,opt,name=Path,json=path" json:"Path,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=Content,json=content" json:"Content,omitempty"`
}

func (m *FunctionMeta) Reset()                    { *m = FunctionMeta{} }
func (m *FunctionMeta) String() string            { return proto.CompactTextString(m) }
func (*FunctionMeta) ProtoMessage()               {}
func (*FunctionMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FunctionMeta) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *FunctionMeta) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FunctionMeta) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpRequest struct {
	Functions []*FunctionMeta `protobuf:"bytes,1,rep,name=Functions,json=functions" json:"Functions,omitempty"`
}

func (m *UpRequest) Reset()                    { *m = UpRequest{} }
func (m *UpRequest) String() string            { return proto.CompactTextString(m) }
func (*UpRequest) ProtoMessage()               {}
func (*UpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpRequest) GetFunctions() []*FunctionMeta {
	if m != nil {
		return m.Functions
	}
	return nil
}

type DownRequest struct {
	ID []string `protobuf:"bytes,1,rep,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *DownRequest) Reset()                    { *m = DownRequest{} }
func (m *DownRequest) String() string            { return proto.CompactTextString(m) }
func (*DownRequest) ProtoMessage()               {}
func (*DownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DownRequest) GetID() []string {
	if m != nil {
		return m.ID
	}
	return nil
}

type ListRequest struct {
	ID []string `protobuf:"bytes,1,rep,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListRequest) GetID() []string {
	if m != nil {
		return m.ID
	}
	return nil
}

type UpMsgMeta struct {
	FunctionID     string `protobuf:"bytes,1,opt,name=FunctionID,json=functionID" json:"FunctionID,omitempty"`
	FunctionSource string `protobuf:"bytes,2,opt,name=FunctionSource,json=functionSource" json:"FunctionSource,omitempty"`
	LocalAddress   string `protobuf:"bytes,3,opt,name=LocalAddress,json=localAddress" json:"LocalAddress,omitempty"`
	RemoteAddress  string `protobuf:"bytes,4,opt,name=RemoteAddress,json=remoteAddress" json:"RemoteAddress,omitempty"`
	Error          string `protobuf:"bytes,5,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *UpMsgMeta) Reset()                    { *m = UpMsgMeta{} }
func (m *UpMsgMeta) String() string            { return proto.CompactTextString(m) }
func (*UpMsgMeta) ProtoMessage()               {}
func (*UpMsgMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpMsgMeta) GetFunctionID() string {
	if m != nil {
		return m.FunctionID
	}
	return ""
}

func (m *UpMsgMeta) GetFunctionSource() string {
	if m != nil {
		return m.FunctionSource
	}
	return ""
}

func (m *UpMsgMeta) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *UpMsgMeta) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *UpMsgMeta) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpResponse struct {
	Instances []*UpMsgMeta `protobuf:"bytes,1,rep,name=Instances,json=instances" json:"Instances,omitempty"`
}

func (m *UpResponse) Reset()                    { *m = UpResponse{} }
func (m *UpResponse) String() string            { return proto.CompactTextString(m) }
func (*UpResponse) ProtoMessage()               {}
func (*UpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpResponse) GetInstances() []*UpMsgMeta {
	if m != nil {
		return m.Instances
	}
	return nil
}

type DownMsgMeta struct {
	ContainerId     string `protobuf:"bytes,1,opt,name=ContainerId,json=containerId" json:"ContainerId,omitempty"`
	ContainerStatus string `protobuf:"bytes,2,opt,name=ContainerStatus,json=containerStatus" json:"ContainerStatus,omitempty"`
	ImageStatus     string `protobuf:"bytes,3,opt,name=ImageStatus,json=imageStatus" json:"ImageStatus,omitempty"`
	Error           string `protobuf:"bytes,4,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *DownMsgMeta) Reset()                    { *m = DownMsgMeta{} }
func (m *DownMsgMeta) String() string            { return proto.CompactTextString(m) }
func (*DownMsgMeta) ProtoMessage()               {}
func (*DownMsgMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DownMsgMeta) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *DownMsgMeta) GetContainerStatus() string {
	if m != nil {
		return m.ContainerStatus
	}
	return ""
}

func (m *DownMsgMeta) GetImageStatus() string {
	if m != nil {
		return m.ImageStatus
	}
	return ""
}

func (m *DownMsgMeta) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DownResponse struct {
	Instances []*DownMsgMeta `protobuf:"bytes,1,rep,name=Instances,json=instances" json:"Instances,omitempty"`
}

func (m *DownResponse) Reset()                    { *m = DownResponse{} }
func (m *DownResponse) String() string            { return proto.CompactTextString(m) }
func (*DownResponse) ProtoMessage()               {}
func (*DownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DownResponse) GetInstances() []*DownMsgMeta {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ListItem struct {
	FunctionID string `protobuf:"bytes,1,opt,name=FunctionID,json=functionID" json:"FunctionID,omitempty"`
	State      string `protobuf:"bytes,2,opt,name=State,json=state" json:"State,omitempty"`
	ServiceURL string `protobuf:"bytes,3,opt,name=ServiceURL,json=serviceURL" json:"ServiceURL,omitempty"`
}

func (m *ListItem) Reset()                    { *m = ListItem{} }
func (m *ListItem) String() string            { return proto.CompactTextString(m) }
func (*ListItem) ProtoMessage()               {}
func (*ListItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListItem) GetFunctionID() string {
	if m != nil {
		return m.FunctionID
	}
	return ""
}

func (m *ListItem) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ListItem) GetServiceURL() string {
	if m != nil {
		return m.ServiceURL
	}
	return ""
}

type ListResponse struct {
	Instances []*ListItem `protobuf:"bytes,1,rep,name=Instances,json=instances" json:"Instances,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetInstances() []*ListItem {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
	proto.RegisterType((*FunctionMeta)(nil), "api.FunctionMeta")
	proto.RegisterType((*UpRequest)(nil), "api.UpRequest")
	proto.RegisterType((*DownRequest)(nil), "api.DownRequest")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*UpMsgMeta)(nil), "api.UpMsgMeta")
	proto.RegisterType((*UpResponse)(nil), "api.UpResponse")
	proto.RegisterType((*DownMsgMeta)(nil), "api.DownMsgMeta")
	proto.RegisterType((*DownResponse)(nil), "api.DownResponse")
	proto.RegisterType((*ListItem)(nil), "api.ListItem")
	proto.RegisterType((*ListResponse)(nil), "api.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FxService service

type FxServiceClient interface {
	Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error)
	Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type fxServiceClient struct {
	cc *grpc.ClientConn
}

func NewFxServiceClient(cc *grpc.ClientConn) FxServiceClient {
	return &fxServiceClient{cc}
}

func (c *fxServiceClient) Up(ctx context.Context, in *UpRequest, opts ...grpc.CallOption) (*UpResponse, error) {
	out := new(UpResponse)
	err := grpc.Invoke(ctx, "/api.FxService/Up", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxServiceClient) Down(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (*DownResponse, error) {
	out := new(DownResponse)
	err := grpc.Invoke(ctx, "/api.FxService/Down", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fxServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/api.FxService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FxService service

type FxServiceServer interface {
	Up(context.Context, *UpRequest) (*UpResponse, error)
	Down(context.Context, *DownRequest) (*DownResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterFxServiceServer(s *grpc.Server, srv FxServiceServer) {
	s.RegisterService(&_FxService_serviceDesc, srv)
}

func _FxService_Up_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxServiceServer).Up(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FxService/Up",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxServiceServer).Up(ctx, req.(*UpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxService_Down_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxServiceServer).Down(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FxService/Down",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxServiceServer).Down(ctx, req.(*DownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FxService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FxService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.FxService",
	HandlerType: (*FxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Up",
			Handler:    _FxService_Up_Handler,
		},
		{
			MethodName: "Down",
			Handler:    _FxService_Down_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FxService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fx.proto",
}

func init() { proto.RegisterFile("fx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x66, 0xfa, 0xb3, 0xdb, 0x39, 0xd3, 0x1f, 0x1b, 0xf7, 0x62, 0x58, 0x54, 0x4a, 0x10, 0x29,
	0xab, 0xb4, 0xb8, 0xde, 0x48, 0x15, 0x51, 0xac, 0x0b, 0x85, 0x2e, 0x2c, 0x53, 0x7a, 0x6f, 0x9c,
	0xa6, 0xb3, 0x81, 0x36, 0x19, 0x27, 0xe9, 0xee, 0x5e, 0xfb, 0x06, 0xe2, 0xa3, 0xf8, 0x14, 0x5e,
	0xfb, 0x0a, 0x3e, 0x88, 0x24, 0x93, 0x4c, 0x53, 0x16, 0xf1, 0xf2, 0x7c, 0xe7, 0x3b, 0xe7, 0x7c,
	0xdf, 0x39, 0x99, 0x81, 0xd6, 0xfa, 0x6e, 0x94, 0x17, 0x42, 0x09, 0x54, 0x27, 0x39, 0x3b, 0x7d,
	0x94, 0x09, 0x91, 0x6d, 0xe8, 0x98, 0xe4, 0x6c, 0x4c, 0x38, 0x17, 0x8a, 0x28, 0x26, 0xb8, 0x2c,
	0x29, 0xf8, 0x0a, 0xda, 0x17, 0x3b, 0x9e, 0x6a, 0xe8, 0x92, 0x2a, 0x82, 0x10, 0x34, 0xe6, 0x84,
	0x67, 0x71, 0x30, 0x08, 0x86, 0x61, 0xd2, 0xd8, 0x10, 0x9e, 0x69, 0xec, 0x8a, 0xa8, 0xeb, 0xb8,
	0x56, 0x62, 0x39, 0x51, 0xd7, 0x28, 0x86, 0xe3, 0x8f, 0x82, 0x2b, 0xca, 0x55, 0x5c, 0x37, 0xf0,
	0x71, 0x5a, 0x86, 0xf8, 0x2d, 0x84, 0xcb, 0x3c, 0xa1, 0x5f, 0x77, 0x54, 0x2a, 0x34, 0x86, 0xd0,
	0xb5, 0x97, 0x71, 0x30, 0xa8, 0x0f, 0xa3, 0xf3, 0xfe, 0x88, 0xe4, 0x6c, 0xe4, 0x0f, 0x4d, 0xc2,
	0xb5, 0xe3, 0xe0, 0xc7, 0x10, 0x4d, 0xc5, 0x2d, 0x77, 0xf5, 0x5d, 0xa8, 0xcd, 0xa6, 0xa6, 0x30,
	0x4c, 0x6a, 0x6c, 0xaa, 0xd3, 0x73, 0x26, 0xd5, 0xbf, 0xd2, 0x3f, 0x03, 0x3d, 0xfc, 0x52, 0x66,
	0xc6, 0xcb, 0x13, 0x00, 0x37, 0xc6, 0xb0, 0xb4, 0x4c, 0x58, 0x57, 0x08, 0x7a, 0x06, 0x5d, 0x97,
	0x5f, 0x88, 0x5d, 0x91, 0x52, 0xeb, 0xb0, 0xbb, 0x3e, 0x40, 0x11, 0x86, 0xf6, 0x5c, 0xa4, 0x64,
	0xf3, 0x61, 0xb5, 0x2a, 0xa8, 0x94, 0xd6, 0x70, 0x7b, 0xe3, 0x61, 0xe8, 0x29, 0x74, 0x12, 0xba,
	0x15, 0x8a, 0x3a, 0x52, 0xc3, 0x90, 0x3a, 0x85, 0x0f, 0xa2, 0x13, 0x68, 0x7e, 0x2a, 0x0a, 0x51,
	0xc4, 0x4d, 0x93, 0x6d, 0x52, 0x1d, 0xe0, 0x09, 0x80, 0xde, 0x98, 0xcc, 0x05, 0x97, 0x14, 0xbd,
	0x80, 0x70, 0xc6, 0xa5, 0x22, 0x3c, 0xa5, 0x6e, 0x65, 0x5d, 0xb3, 0xb2, 0xca, 0x58, 0x12, 0x32,
	0x47, 0xc0, 0xdf, 0x83, 0x72, 0x61, 0xce, 0xf3, 0x00, 0x22, 0x7d, 0x17, 0xc2, 0x38, 0x2d, 0x66,
	0x2b, 0x6b, 0x3a, 0x4a, 0xf7, 0x10, 0x1a, 0x42, 0xaf, 0x62, 0x2c, 0x14, 0x51, 0x3b, 0x69, 0x6d,
	0xf7, 0xd2, 0x43, 0x58, 0xf7, 0x9a, 0x6d, 0x49, 0x46, 0x2d, 0xab, 0xb4, 0x1d, 0xb1, 0x3d, 0xb4,
	0xf7, 0xd3, 0xf0, 0xfd, 0xbc, 0x83, 0x76, 0x79, 0x43, 0xeb, 0x68, 0x74, 0xdf, 0xd1, 0x03, 0xe3,
	0xc8, 0x13, 0xee, 0x7b, 0xfa, 0x0c, 0x2d, 0x7d, 0xe4, 0x99, 0xa2, 0xdb, 0xff, 0xde, 0xf0, 0x04,
	0x9a, 0x5a, 0x8b, 0x3b, 0x5d, 0x53, 0xea, 0x40, 0x57, 0x2d, 0x68, 0x71, 0xc3, 0x52, 0xba, 0x4c,
	0xe6, 0x56, 0x38, 0xc8, 0x0a, 0xc1, 0x6f, 0xa0, 0x5d, 0x3e, 0x23, 0xab, 0xf0, 0xf9, 0x7d, 0x85,
	0x1d, 0xa3, 0xd0, 0xe9, 0xf0, 0xe4, 0x9d, 0xff, 0x0a, 0x20, 0xbc, 0xb8, 0xb3, 0xfd, 0xd1, 0x6b,
	0xa8, 0x2d, 0x73, 0xe4, 0x2e, 0x64, 0x1f, 0xe6, 0x69, 0xaf, 0x8a, 0xcb, 0x09, 0xb8, 0xff, 0xed,
	0xf7, 0x9f, 0x1f, 0xb5, 0x08, 0x1f, 0x8d, 0x6f, 0x5e, 0x8e, 0x77, 0xf9, 0x24, 0x38, 0x43, 0xef,
	0xa1, 0xa1, 0x17, 0x80, 0xf6, 0xbb, 0x70, 0xd5, 0x7d, 0x0f, 0xb1, 0xf5, 0x0f, 0x4d, 0x7d, 0x07,
	0xb7, 0x74, 0xfd, 0x4a, 0xdc, 0x72, 0xdb, 0x41, 0x0b, 0xb4, 0x1d, 0xbc, 0x0f, 0xc3, 0x76, 0xf0,
	0x3d, 0x1e, 0x76, 0xd8, 0x30, 0xa9, 0x26, 0xc1, 0xd9, 0x97, 0x23, 0xf3, 0x17, 0x78, 0xf5, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x79, 0x9e, 0x25, 0x34, 0x04, 0x00, 0x00,
}
