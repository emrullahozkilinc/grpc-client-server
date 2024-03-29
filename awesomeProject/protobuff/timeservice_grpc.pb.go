// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: timeservice.proto

package protobuff

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

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*TimeUpdate, error)
	Stream(ctx context.Context, in *TimeStreamRequest, opts ...grpc.CallOption) (TimeService_StreamClient, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*TimeUpdate, error) {
	out := new(TimeUpdate)
	err := c.cc.Invoke(ctx, "/TimeService/Now", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeServiceClient) Stream(ctx context.Context, in *TimeStreamRequest, opts ...grpc.CallOption) (TimeService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TimeService_ServiceDesc.Streams[0], "/TimeService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeService_StreamClient interface {
	Recv() (*TimeUpdate, error)
	grpc.ClientStream
}

type timeServiceStreamClient struct {
	grpc.ClientStream
}

func (x *timeServiceStreamClient) Recv() (*TimeUpdate, error) {
	m := new(TimeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility
type TimeServiceServer interface {
	Now(context.Context, *NowRequest) (*TimeUpdate, error)
	Stream(*TimeStreamRequest, TimeService_StreamServer) error
}

// UnimplementedTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeServiceServer struct {
}

func (UnimplementedTimeServiceServer) Now(context.Context, *NowRequest) (*TimeUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}
func (UnimplementedTimeServiceServer) Stream(*TimeStreamRequest, TimeService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	s.RegisterService(&TimeService_ServiceDesc, srv)
}

func _TimeService_Now_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).Now(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimeService/Now",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).Now(ctx, req.(*NowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TimeStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeServiceServer).Stream(m, &timeServiceStreamServer{stream})
}

type TimeService_StreamServer interface {
	Send(*TimeUpdate) error
	grpc.ServerStream
}

type timeServiceStreamServer struct {
	grpc.ServerStream
}

func (x *timeServiceStreamServer) Send(m *TimeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// TimeService_ServiceDesc is the grpc.ServiceDesc for TimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Now",
			Handler:    _TimeService_Now_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _TimeService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timeservice.proto",
}
