// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/ingest/chef.proto

package ingest

import (
	context "context"
	fmt "fmt"
	version "github.com/chef/automate/api/external/common/version"
	request "github.com/chef/automate/api/external/ingest/request"
	response "github.com/chef/automate/api/external/ingest/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("api/external/ingest/chef.proto", fileDescriptor_920646edf89f77d6)
}

var fileDescriptor_920646edf89f77d6 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xd5, 0x1d, 0x38, 0x58, 0xd3, 0x04, 0x9e, 0x36, 0x42, 0x40, 0x43, 0xcb, 0xd4, 0xd1,
	0x15, 0x92, 0x4c, 0xfc, 0x39, 0x10, 0x38, 0xf0, 0x67, 0x12, 0x9a, 0x04, 0x68, 0xea, 0x01, 0x09,
	0x38, 0xa0, 0x2c, 0x7d, 0x9b, 0x5a, 0x4a, 0xec, 0x60, 0x3b, 0x15, 0x15, 0xe2, 0xc2, 0x05, 0x29,
	0x57, 0x3e, 0x4b, 0xbe, 0x05, 0x37, 0x2e, 0x48, 0x5c, 0xb9, 0x22, 0xbe, 0x02, 0x4a, 0xec, 0xac,
	0xcd, 0x08, 0x69, 0x7b, 0x6a, 0x65, 0x3f, 0xcf, 0xfb, 0x3e, 0x3f, 0x3b, 0x7e, 0xd1, 0x8e, 0x9f,
	0x10, 0x17, 0x3e, 0x48, 0xe0, 0xd4, 0x8f, 0x5c, 0x42, 0x43, 0x10, 0xd2, 0x0d, 0xc6, 0x30, 0x72,
	0x12, 0xce, 0x24, 0xc3, 0x46, 0xf9, 0xdf, 0x4f, 0x25, 0x8b, 0x7d, 0x09, 0x8e, 0x9f, 0x10, 0x47,
	0x89, 0xcc, 0x6b, 0x21, 0x63, 0x61, 0x04, 0x6e, 0x51, 0xc0, 0xa7, 0x94, 0x49, 0x5f, 0x12, 0x46,
	0x85, 0xf2, 0x99, 0x8f, 0x02, 0x16, 0x27, 0x8c, 0x02, 0x95, 0xc2, 0xad, 0xdc, 0x76, 0xc8, 0x93,
	0xc0, 0x2d, 0xf7, 0x03, 0x3b, 0x04, 0x6a, 0x27, 0x2c, 0x22, 0xc1, 0xd4, 0x25, 0x7e, 0xdc, 0x50,
	0xe1, 0xa0, 0x96, 0x2c, 0x60, 0x71, 0xcc, 0xa8, 0x3b, 0x01, 0x2e, 0xc8, 0xec, 0x57, 0x4b, 0xf7,
	0x9b, 0x20, 0x38, 0xbc, 0x4f, 0xeb, 0x30, 0x66, 0xaf, 0x4d, 0xe7, 0x07, 0x72, 0x56, 0xb1, 0xdf,
	0xa6, 0x8c, 0xc8, 0x04, 0x28, 0x88, 0x2a, 0xe8, 0x8d, 0x66, 0xad, 0x48, 0x18, 0x15, 0x30, 0xdf,
	0xfe, 0xa0, 0x55, 0x58, 0xeb, 0x7f, 0xb3, 0x55, 0x5a, 0x0f, 0x70, 0xfb, 0x37, 0x42, 0xeb, 0x4f,
	0xc7, 0x30, 0x3a, 0x2e, 0x65, 0xc0, 0xf1, 0xcf, 0x0e, 0xda, 0x38, 0xe1, 0x2c, 0x00, 0x21, 0x8a,
	0xf5, 0x41, 0x4a, 0x71, 0xd7, 0xf9, 0xdf, 0x45, 0x3a, 0x1a, 0xcb, 0x19, 0xa4, 0xd4, 0xbc, 0xdf,
	0x26, 0x53, 0xdd, 0x9d, 0x7a, 0xe5, 0x81, 0x5e, 0xb6, 0xc2, 0x2c, 0x37, 0x76, 0xd1, 0x75, 0xa5,
	0xf6, 0x28, 0x1b, 0x82, 0xf0, 0x3e, 0x02, 0x95, 0x44, 0x4e, 0xdf, 0xa5, 0x29, 0x19, 0x7e, 0xf2,
	0x78, 0x4a, 0x45, 0x96, 0x1b, 0x5b, 0x78, 0xb3, 0x26, 0x0a, 0x38, 0xf8, 0x12, 0x3e, 0x7f, 0xff,
	0xf5, 0x75, 0x6d, 0xcf, 0xda, 0x29, 0xbf, 0xa7, 0xc9, 0x61, 0xc5, 0x0c, 0x93, 0xf2, 0x13, 0x2a,
	0x52, 0xb9, 0x3c, 0xa5, 0x5e, 0xa7, 0x8f, 0x7f, 0x74, 0xd0, 0xa5, 0xb9, 0x0c, 0x8f, 0xcb, 0x73,
	0xc3, 0xbd, 0xc5, 0x80, 0x4a, 0x69, 0x3e, 0x5c, 0x8d, 0x51, 0xb9, 0xce, 0x30, 0xdf, 0x66, 0xb9,
	0x71, 0x11, 0x6d, 0x68, 0x02, 0x75, 0x69, 0x05, 0xd5, 0x65, 0xbc, 0x55, 0x5f, 0x9b, 0xe7, 0xda,
	0xb7, 0x76, 0x5b, 0xb8, 0x94, 0xe1, 0x1c, 0xda, 0x4b, 0x36, 0x84, 0x23, 0x88, 0x40, 0xc2, 0x32,
	0x68, 0x4a, 0xb9, 0x0a, 0xda, 0xac, 0x7e, 0x0d, 0x6d, 0x03, 0xad, 0xcf, 0x5f, 0x4e, 0xc3, 0x75,
	0x0d, 0x4b, 0x53, 0x89, 0xd5, 0xb7, 0xba, 0x2d, 0x58, 0x85, 0x5c, 0x89, 0x0b, 0xb4, 0x2f, 0x6b,
	0xc8, 0xd4, 0xad, 0x5f, 0xa4, 0x91, 0x24, 0x49, 0x04, 0xb3, 0x08, 0x02, 0x3f, 0x58, 0xcc, 0xf8,
	0xaf, 0x6d, 0xa0, 0x76, 0xcc, 0xa3, 0xe5, 0xb1, 0x9b, 0x8a, 0x68, 0xfc, 0xd1, 0xaa, 0xf8, 0xf7,
	0xac, 0xc3, 0x05, 0xf8, 0x76, 0xac, 0xdb, 0xd9, 0xca, 0x26, 0x8a, 0x93, 0xf8, 0xd3, 0x41, 0x9b,
	0x3a, 0xcd, 0x73, 0xfd, 0x90, 0x4f, 0x08, 0x0d, 0x71, 0x7f, 0xf1, 0x11, 0x54, 0x7a, 0xd3, 0x5b,
	0x9e, 0xb8, 0xf2, 0x9c, 0x71, 0x26, 0x59, 0x6e, 0x74, 0xd1, 0x5e, 0xcb, 0x43, 0xad, 0xe6, 0x4b,
	0xfb, 0x63, 0xed, 0x59, 0x7b, 0x2d, 0xf8, 0x55, 0x8d, 0x82, 0xf8, 0x5b, 0x07, 0xa1, 0x67, 0x20,
	0x5f, 0xa9, 0xa1, 0x8d, 0xef, 0x36, 0x84, 0x57, 0xf3, 0xdd, 0xa9, 0xe6, 0xba, 0x96, 0x1e, 0xd3,
	0x11, 0xab, 0x2e, 0xd9, 0x5e, 0xc9, 0x65, 0xbd, 0xce, 0x72, 0xc3, 0x40, 0xdb, 0x62, 0x2a, 0x24,
	0xc4, 0x9e, 0x00, 0x3e, 0x21, 0x01, 0x78, 0x5a, 0x99, 0xe5, 0xc6, 0x55, 0x7c, 0xa5, 0xbe, 0xa7,
	0xcd, 0x5e, 0x08, 0xb2, 0xc4, 0x33, 0xf0, 0xf6, 0x39, 0x3c, 0x6d, 0x7e, 0xe2, 0xbc, 0xb9, 0x15,
	0x12, 0x39, 0x4e, 0x4f, 0x8b, 0xfe, 0xfa, 0x09, 0xeb, 0x54, 0x6e, 0xc3, 0xdc, 0x3e, 0xbd, 0x50,
	0x8e, 0xe9, 0x3b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0xd0, 0x4e, 0x23, 0x6c, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefIngesterClient is the client API for ChefIngester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefIngesterClient interface {
	ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error)
	ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error)
	ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error)
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
}

type chefIngesterClient struct {
	cc grpc.ClientConnInterface
}

func NewChefIngesterClient(cc grpc.ClientConnInterface) ChefIngesterClient {
	return &chefIngesterClient{cc}
}

func (c *chefIngesterClient) ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error) {
	out := new(response.ProcessChefRunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessChefRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error) {
	out := new(response.ProcessChefActionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessChefAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error) {
	out := new(response.ProcessNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessNodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error) {
	out := new(response.ProcessMultipleNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessMultipleNodeDeletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error) {
	out := new(response.ProcessLivenessResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessLivenessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefIngesterServer is the server API for ChefIngester service.
type ChefIngesterServer interface {
	ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error)
	ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error)
	ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error)
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
}

// UnimplementedChefIngesterServer can be embedded to have forward compatible implementations.
type UnimplementedChefIngesterServer struct {
}

func (*UnimplementedChefIngesterServer) ProcessChefRun(ctx context.Context, req *request.Run) (*response.ProcessChefRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefRun not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessChefAction(ctx context.Context, req *request.Action) (*response.ProcessChefActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefAction not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessNodeDelete(ctx context.Context, req *request.Delete) (*response.ProcessNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessNodeDelete not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessMultipleNodeDeletes(ctx context.Context, req *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMultipleNodeDeletes not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessLivenessPing(ctx context.Context, req *request.Liveness) (*response.ProcessLivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLivenessPing not implemented")
}
func (*UnimplementedChefIngesterServer) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterChefIngesterServer(s *grpc.Server, srv ChefIngesterServer) {
	s.RegisterService(&_ChefIngester_serviceDesc, srv)
}

func _ChefIngester_ProcessChefRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Run)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessChefRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, req.(*request.Run))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessChefAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessChefAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, req.(*request.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessNodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Delete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessNodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, req.(*request.Delete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessMultipleNodeDeletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.MultipleNodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessMultipleNodeDeletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, req.(*request.MultipleNodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessLivenessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Liveness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessLivenessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, req.(*request.Liveness))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefIngester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.ingest.ChefIngester",
	HandlerType: (*ChefIngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessChefRun",
			Handler:    _ChefIngester_ProcessChefRun_Handler,
		},
		{
			MethodName: "ProcessChefAction",
			Handler:    _ChefIngester_ProcessChefAction_Handler,
		},
		{
			MethodName: "ProcessNodeDelete",
			Handler:    _ChefIngester_ProcessNodeDelete_Handler,
		},
		{
			MethodName: "ProcessMultipleNodeDeletes",
			Handler:    _ChefIngester_ProcessMultipleNodeDeletes_Handler,
		},
		{
			MethodName: "ProcessLivenessPing",
			Handler:    _ChefIngester_ProcessLivenessPing_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ChefIngester_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/external/ingest/chef.proto",
}
