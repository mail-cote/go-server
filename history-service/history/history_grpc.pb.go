// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: history.proto

package history

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
	History_SaveHistory_FullMethodName   = "/history.v1.History/saveHistory"
	History_GetAllHistory_FullMethodName = "/history.v1.History/getAllHistory"
)

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryClient interface {
	// 1. history 저장하기
	// 2. history 조회하기
	SaveHistory(ctx context.Context, in *SaveHistoryRequest, opts ...grpc.CallOption) (*SaveHistoryResponse, error)
	GetAllHistory(ctx context.Context, in *GetAllHistoryRequest, opts ...grpc.CallOption) (*GetAllHistoryResponse, error)
}

type historyClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryClient(cc grpc.ClientConnInterface) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) SaveHistory(ctx context.Context, in *SaveHistoryRequest, opts ...grpc.CallOption) (*SaveHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveHistoryResponse)
	err := c.cc.Invoke(ctx, History_SaveHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyClient) GetAllHistory(ctx context.Context, in *GetAllHistoryRequest, opts ...grpc.CallOption) (*GetAllHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllHistoryResponse)
	err := c.cc.Invoke(ctx, History_GetAllHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
// All implementations must embed UnimplementedHistoryServer
// for forward compatibility.
type HistoryServer interface {
	// 1. history 저장하기
	// 2. history 조회하기
	SaveHistory(context.Context, *SaveHistoryRequest) (*SaveHistoryResponse, error)
	GetAllHistory(context.Context, *GetAllHistoryRequest) (*GetAllHistoryResponse, error)
	mustEmbedUnimplementedHistoryServer()
}

// UnimplementedHistoryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHistoryServer struct{}

func (UnimplementedHistoryServer) SaveHistory(context.Context, *SaveHistoryRequest) (*SaveHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveHistory not implemented")
}
func (UnimplementedHistoryServer) GetAllHistory(context.Context, *GetAllHistoryRequest) (*GetAllHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHistory not implemented")
}
func (UnimplementedHistoryServer) mustEmbedUnimplementedHistoryServer() {}
func (UnimplementedHistoryServer) testEmbeddedByValue()                 {}

// UnsafeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServer will
// result in compilation errors.
type UnsafeHistoryServer interface {
	mustEmbedUnimplementedHistoryServer()
}

func RegisterHistoryServer(s grpc.ServiceRegistrar, srv HistoryServer) {
	// If the following call pancis, it indicates UnimplementedHistoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&History_ServiceDesc, srv)
}

func _History_SaveHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).SaveHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: History_SaveHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).SaveHistory(ctx, req.(*SaveHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _History_GetAllHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).GetAllHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: History_GetAllHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).GetAllHistory(ctx, req.(*GetAllHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// History_ServiceDesc is the grpc.ServiceDesc for History service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var History_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "history.v1.History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveHistory",
			Handler:    _History_SaveHistory_Handler,
		},
		{
			MethodName: "getAllHistory",
			Handler:    _History_GetAllHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}
