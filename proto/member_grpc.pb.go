// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/member.proto

package member

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MemberService_CreateMember_FullMethodName = "/member.MemberService/CreateMember"
	MemberService_UpdateMember_FullMethodName = "/member.MemberService/UpdateMember"
	MemberService_DeleteMember_FullMethodName = "/member.MemberService/DeleteMember"
)

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Member Service 정의
type MemberServiceClient interface {
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*Response, error)
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*Response, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MemberService_CreateMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MemberService_UpdateMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MemberService_DeleteMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility.
//
// Member Service 정의
type MemberServiceServer interface {
	CreateMember(context.Context, *CreateMemberRequest) (*Response, error)
	UpdateMember(context.Context, *UpdateMemberRequest) (*Response, error)
	DeleteMember(context.Context, *DeleteMemberRequest) (*Response, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMemberServiceServer struct{}

func (UnimplementedMemberServiceServer) CreateMember(context.Context, *CreateMemberRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedMemberServiceServer) UpdateMember(context.Context, *UpdateMemberRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMember not implemented")
}
func (UnimplementedMemberServiceServer) DeleteMember(context.Context, *DeleteMemberRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMember not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}
func (UnimplementedMemberServiceServer) testEmbeddedByValue()                       {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	// If the following call pancis, it indicates UnimplementedMemberServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_CreateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_UpdateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).UpdateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_UpdateMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).UpdateMember(ctx, req.(*UpdateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_DeleteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).DeleteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberService_DeleteMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).DeleteMember(ctx, req.(*DeleteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMember",
			Handler:    _MemberService_CreateMember_Handler,
		},
		{
			MethodName: "UpdateMember",
			Handler:    _MemberService_UpdateMember_Handler,
		},
		{
			MethodName: "DeleteMember",
			Handler:    _MemberService_DeleteMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/member.proto",
}
