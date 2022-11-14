// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: main.proto

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

// SubClient is the client API for Sub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubClient interface {
	Func(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Coffee, error)
}

type subClient struct {
	cc grpc.ClientConnInterface
}

func NewSubClient(cc grpc.ClientConnInterface) SubClient {
	return &subClient{cc}
}

func (c *subClient) Func(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Coffee, error) {
	out := new(Coffee)
	err := c.cc.Invoke(ctx, "/proto.Sub/Func", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubServer is the server API for Sub service.
// All implementations must embed UnimplementedSubServer
// for forward compatibility
type SubServer interface {
	Func(context.Context, *Input) (*Coffee, error)
	mustEmbedUnimplementedSubServer()
}

// UnimplementedSubServer must be embedded to have forward compatible implementations.
type UnimplementedSubServer struct {
}

func (UnimplementedSubServer) Func(context.Context, *Input) (*Coffee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Func not implemented")
}
func (UnimplementedSubServer) mustEmbedUnimplementedSubServer() {}

// UnsafeSubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubServer will
// result in compilation errors.
type UnsafeSubServer interface {
	mustEmbedUnimplementedSubServer()
}

func RegisterSubServer(s grpc.ServiceRegistrar, srv SubServer) {
	s.RegisterService(&Sub_ServiceDesc, srv)
}

func _Sub_Func_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubServer).Func(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sub/Func",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubServer).Func(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Sub_ServiceDesc is the grpc.ServiceDesc for Sub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sub",
	HandlerType: (*SubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Func",
			Handler:    _Sub_Func_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}