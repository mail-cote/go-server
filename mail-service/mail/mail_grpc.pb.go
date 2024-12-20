// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: mail-service/mail/mail.proto

package mail

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
	Mail_FetchQuizFromBucket_FullMethodName = "/mail.Mail/FetchQuizFromBucket"
	Mail_SendMail_FullMethodName            = "/mail.Mail/SendMail"
)

// MailClient is the client API for Mail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailClient interface {
	// 1. mysql 연결 후, 유저 정보 가져오기-> 이건 어짜피 다른 모듈에서 해야하는거니까 넘겨도 됨.
	// 2. 버킷에서 랜덤 파일 가져오기
	// 3. smtp로 메일 전송하기
	FetchQuizFromBucket(ctx context.Context, in *FetchQuizFromBucketRequest, opts ...grpc.CallOption) (*FetchQuizFromBucketResponse, error)
	SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error)
}

type mailClient struct {
	cc grpc.ClientConnInterface
}

func NewMailClient(cc grpc.ClientConnInterface) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) FetchQuizFromBucket(ctx context.Context, in *FetchQuizFromBucketRequest, opts ...grpc.CallOption) (*FetchQuizFromBucketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchQuizFromBucketResponse)
	err := c.cc.Invoke(ctx, Mail_FetchQuizFromBucket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailClient) SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMailResponse)
	err := c.cc.Invoke(ctx, Mail_SendMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServer is the server API for Mail service.
// All implementations must embed UnimplementedMailServer
// for forward compatibility.
type MailServer interface {
	// 1. mysql 연결 후, 유저 정보 가져오기-> 이건 어짜피 다른 모듈에서 해야하는거니까 넘겨도 됨.
	// 2. 버킷에서 랜덤 파일 가져오기
	// 3. smtp로 메일 전송하기
	FetchQuizFromBucket(context.Context, *FetchQuizFromBucketRequest) (*FetchQuizFromBucketResponse, error)
	SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error)
	mustEmbedUnimplementedMailServer()
}

// UnimplementedMailServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMailServer struct{}

func (UnimplementedMailServer) FetchQuizFromBucket(context.Context, *FetchQuizFromBucketRequest) (*FetchQuizFromBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchQuizFromBucket not implemented")
}
func (UnimplementedMailServer) SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMailServer) mustEmbedUnimplementedMailServer() {}
func (UnimplementedMailServer) testEmbeddedByValue()              {}

// UnsafeMailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailServer will
// result in compilation errors.
type UnsafeMailServer interface {
	mustEmbedUnimplementedMailServer()
}

func RegisterMailServer(s grpc.ServiceRegistrar, srv MailServer) {
	// If the following call pancis, it indicates UnimplementedMailServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mail_ServiceDesc, srv)
}

func _Mail_FetchQuizFromBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchQuizFromBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).FetchQuizFromBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mail_FetchQuizFromBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).FetchQuizFromBucket(ctx, req.(*FetchQuizFromBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mail_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mail_SendMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).SendMail(ctx, req.(*SendMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mail_ServiceDesc is the grpc.ServiceDesc for Mail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mail.Mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchQuizFromBucket",
			Handler:    _Mail_FetchQuizFromBucket_Handler,
		},
		{
			MethodName: "SendMail",
			Handler:    _Mail_SendMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail-service/mail/mail.proto",
}
