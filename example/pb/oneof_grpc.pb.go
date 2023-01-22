// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: oneof.proto

package pb

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

// OneofClient is the client API for Oneof service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OneofClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	FetchNested(ctx context.Context, in *FetchNestedRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type oneofClient struct {
	cc grpc.ClientConnInterface
}

func NewOneofClient(cc grpc.ClientConnInterface) OneofClient {
	return &oneofClient{cc}
}

func (c *oneofClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/example.Oneof/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oneofClient) FetchNested(ctx context.Context, in *FetchNestedRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/example.Oneof/FetchNested", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OneofServer is the server API for Oneof service.
// All implementations must embed UnimplementedOneofServer
// for forward compatibility
type OneofServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	FetchNested(context.Context, *FetchNestedRequest) (*FetchResponse, error)
	mustEmbedUnimplementedOneofServer()
}

// UnimplementedOneofServer must be embedded to have forward compatible implementations.
type UnimplementedOneofServer struct {
}

func (UnimplementedOneofServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedOneofServer) FetchNested(context.Context, *FetchNestedRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchNested not implemented")
}
func (UnimplementedOneofServer) mustEmbedUnimplementedOneofServer() {}

// UnsafeOneofServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OneofServer will
// result in compilation errors.
type UnsafeOneofServer interface {
	mustEmbedUnimplementedOneofServer()
}

func RegisterOneofServer(s grpc.ServiceRegistrar, srv OneofServer) {
	s.RegisterService(&Oneof_ServiceDesc, srv)
}

func _Oneof_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OneofServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Oneof/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OneofServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oneof_FetchNested_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchNestedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OneofServer).FetchNested(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Oneof/FetchNested",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OneofServer).FetchNested(ctx, req.(*FetchNestedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oneof_ServiceDesc is the grpc.ServiceDesc for Oneof service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oneof_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Oneof",
	HandlerType: (*OneofServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Oneof_Fetch_Handler,
		},
		{
			MethodName: "FetchNested",
			Handler:    _Oneof_FetchNested_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oneof.proto",
}
