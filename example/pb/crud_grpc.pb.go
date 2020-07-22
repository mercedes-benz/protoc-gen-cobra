// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	Create(ctx context.Context, in *CreateCRUD, opts ...grpc.CallOption) (*CRUDObject, error)
	Get(ctx context.Context, in *GetCRUD, opts ...grpc.CallOption) (*CRUDObject, error)
	Update(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*CRUDObject, error)
	Delete(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*Empty, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) Create(ctx context.Context, in *CreateCRUD, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Get(ctx context.Context, in *GetCRUD, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Update(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*CRUDObject, error) {
	out := new(CRUDObject)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Delete(ctx context.Context, in *CRUDObject, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.CRUD/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	Create(context.Context, *CreateCRUD) (*CRUDObject, error)
	Get(context.Context, *GetCRUD) (*CRUDObject, error)
	Update(context.Context, *CRUDObject) (*CRUDObject, error)
	Delete(context.Context, *CRUDObject) (*Empty, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (*UnimplementedCRUDServer) Create(context.Context, *CreateCRUD) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCRUDServer) Get(context.Context, *GetCRUD) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCRUDServer) Update(context.Context, *CRUDObject) (*CRUDObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCRUDServer) Delete(context.Context, *CRUDObject) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

func RegisterCRUDServer(s *grpc.Server, srv CRUDServer) {
	s.RegisterService(&_CRUD_serviceDesc, srv)
}

func _CRUD_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCRUD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Create(ctx, req.(*CreateCRUD))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCRUD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Get(ctx, req.(*GetCRUD))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRUDObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Update(ctx, req.(*CRUDObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CRUDObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CRUD/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Delete(ctx, req.(*CRUDObject))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRUD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CRUD_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CRUD_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CRUD_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CRUD_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}