// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: hello.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// 定义方法
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// 服务端返回流式数据
	LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfRepliesClient, error)
	// 客户端发送流式数据
	LotofGreetings(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotofGreetingsClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfRepliesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/proto.Greeter/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_LotsOfRepliesClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *greeterLotsOfRepliesClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) LotofGreetings(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotofGreetingsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/proto.Greeter/LotofGreetings", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotofGreetingsClient{stream}
	return x, nil
}

type Greeter_LotofGreetingsClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterLotofGreetingsClient struct {
	grpc.ClientStream
}

func (x *greeterLotofGreetingsClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterLotofGreetingsClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// 定义方法
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	// 服务端返回流式数据
	LotsOfReplies(*HelloRequest, Greeter_LotsOfRepliesServer) error
	// 客户端发送流式数据
	LotofGreetings(Greeter_LotofGreetingsServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) LotsOfReplies(*HelloRequest, Greeter_LotsOfRepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfReplies not implemented")
}
func (UnimplementedGreeterServer) LotofGreetings(Greeter_LotofGreetingsServer) error {
	return status.Errorf(codes.Unimplemented, "method LotofGreetings not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).LotsOfReplies(m, &greeterLotsOfRepliesServer{stream})
}

type Greeter_LotsOfRepliesServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *greeterLotsOfRepliesServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_LotofGreetings_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).LotofGreetings(&greeterLotofGreetingsServer{stream})
}

type Greeter_LotofGreetingsServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterLotofGreetingsServer struct {
	grpc.ServerStream
}

func (x *greeterLotofGreetingsServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterLotofGreetingsServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _Greeter_LotsOfReplies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LotofGreetings",
			Handler:       _Greeter_LotofGreetings_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
