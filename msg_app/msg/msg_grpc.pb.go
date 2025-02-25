// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: msg/msg.proto

package msg

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
	Messenger_GetTopics_FullMethodName = "/msg.Messenger/GetTopics"
	Messenger_GetMsgs_FullMethodName   = "/msg.Messenger/GetMsgs"
	Messenger_GetTopic_FullMethodName  = "/msg.Messenger/GetTopic"
	Messenger_GetMsg_FullMethodName    = "/msg.Messenger/GetMsg"
	Messenger_AddMsg_FullMethodName    = "/msg.Messenger/AddMsg"
	Messenger_AddTopic_FullMethodName  = "/msg.Messenger/AddTopic"
	Messenger_DelTopic_FullMethodName  = "/msg.Messenger/DelTopic"
	Messenger_DelMsg_FullMethodName    = "/msg.Messenger/DelMsg"
)

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Interface exported by the server.
type MessengerClient interface {
	GetTopics(ctx context.Context, in *Filter, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Topic], error)
	GetMsgs(ctx context.Context, in *Topic, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Msg], error)
	GetTopic(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Topic, error)
	GetMsg(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Msg, error)
	AddMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error)
	AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Topic, error)
	DelTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Topic, error)
	DelMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) GetTopics(ctx context.Context, in *Filter, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Topic], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], Messenger_GetTopics_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Filter, Topic]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messenger_GetTopicsClient = grpc.ServerStreamingClient[Topic]

func (c *messengerClient) GetMsgs(ctx context.Context, in *Topic, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Msg], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[1], Messenger_GetMsgs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Topic, Msg]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messenger_GetMsgsClient = grpc.ServerStreamingClient[Msg]

func (c *messengerClient) GetTopic(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Topic, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Topic)
	err := c.cc.Invoke(ctx, Messenger_GetTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) GetMsg(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, Messenger_GetMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) AddMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, Messenger_AddMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Topic, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Topic)
	err := c.cc.Invoke(ctx, Messenger_AddTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) DelTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Topic, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Topic)
	err := c.cc.Invoke(ctx, Messenger_DelTopic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) DelMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, Messenger_DelMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility.
//
// Interface exported by the server.
type MessengerServer interface {
	GetTopics(*Filter, grpc.ServerStreamingServer[Topic]) error
	GetMsgs(*Topic, grpc.ServerStreamingServer[Msg]) error
	GetTopic(context.Context, *ID) (*Topic, error)
	GetMsg(context.Context, *ID) (*Msg, error)
	AddMsg(context.Context, *Msg) (*Msg, error)
	AddTopic(context.Context, *Topic) (*Topic, error)
	DelTopic(context.Context, *Topic) (*Topic, error)
	DelMsg(context.Context, *Msg) (*Msg, error)
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessengerServer struct{}

func (UnimplementedMessengerServer) GetTopics(*Filter, grpc.ServerStreamingServer[Topic]) error {
	return status.Errorf(codes.Unimplemented, "method GetTopics not implemented")
}
func (UnimplementedMessengerServer) GetMsgs(*Topic, grpc.ServerStreamingServer[Msg]) error {
	return status.Errorf(codes.Unimplemented, "method GetMsgs not implemented")
}
func (UnimplementedMessengerServer) GetTopic(context.Context, *ID) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (UnimplementedMessengerServer) GetMsg(context.Context, *ID) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsg not implemented")
}
func (UnimplementedMessengerServer) AddMsg(context.Context, *Msg) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMsg not implemented")
}
func (UnimplementedMessengerServer) AddTopic(context.Context, *Topic) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTopic not implemented")
}
func (UnimplementedMessengerServer) DelTopic(context.Context, *Topic) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTopic not implemented")
}
func (UnimplementedMessengerServer) DelMsg(context.Context, *Msg) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMsg not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}
func (UnimplementedMessengerServer) testEmbeddedByValue()                   {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	// If the following call pancis, it indicates UnimplementedMessengerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_GetTopics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).GetTopics(m, &grpc.GenericServerStream[Filter, Topic]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messenger_GetTopicsServer = grpc.ServerStreamingServer[Topic]

func _Messenger_GetMsgs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).GetMsgs(m, &grpc.GenericServerStream[Topic, Msg]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messenger_GetMsgsServer = grpc.ServerStreamingServer[Msg]

func _Messenger_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_GetTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).GetTopic(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_GetMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).GetMsg(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_AddMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).AddMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_AddMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).AddMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_AddTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).AddTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_DelTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).DelTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_DelTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).DelTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_DelMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).DelMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messenger_DelMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).DelMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopic",
			Handler:    _Messenger_GetTopic_Handler,
		},
		{
			MethodName: "GetMsg",
			Handler:    _Messenger_GetMsg_Handler,
		},
		{
			MethodName: "AddMsg",
			Handler:    _Messenger_AddMsg_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _Messenger_AddTopic_Handler,
		},
		{
			MethodName: "DelTopic",
			Handler:    _Messenger_DelTopic_Handler,
		},
		{
			MethodName: "DelMsg",
			Handler:    _Messenger_DelMsg_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTopics",
			Handler:       _Messenger_GetTopics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMsgs",
			Handler:       _Messenger_GetMsgs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "msg/msg.proto",
}
