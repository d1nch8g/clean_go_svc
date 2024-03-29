// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: services/users.proto

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

// UserStorageClient is the client API for UserStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserStorageClient interface {
	// Creates new user
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	// Operation to recieve new user
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UserStorage_ListClient, error)
	// Operation to delete user
	Remove(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	// Operation update information about user
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type userStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewUserStorageClient(cc grpc.ClientConnInterface) UserStorageClient {
	return &userStorageClient{cc}
}

func (c *userStorageClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.v1.UserStorage/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UserStorage_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserStorage_ServiceDesc.Streams[0], "/user.v1.UserStorage/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &userStorageListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserStorage_ListClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userStorageListClient struct {
	grpc.ClientStream
}

func (x *userStorageListClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userStorageClient) Remove(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.v1.UserStorage/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStorageClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.v1.UserStorage/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserStorageServer is the server API for UserStorage service.
// All implementations must embed UnimplementedUserStorageServer
// for forward compatibility
type UserStorageServer interface {
	// Creates new user
	Create(context.Context, *User) (*User, error)
	// Operation to recieve new user
	List(*Empty, UserStorage_ListServer) error
	// Operation to delete user
	Remove(context.Context, *Id) (*Empty, error)
	// Operation update information about user
	Update(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserStorageServer()
}

// UnimplementedUserStorageServer must be embedded to have forward compatible implementations.
type UnimplementedUserStorageServer struct {
}

func (UnimplementedUserStorageServer) Create(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserStorageServer) List(*Empty, UserStorage_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserStorageServer) Remove(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedUserStorageServer) Update(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserStorageServer) mustEmbedUnimplementedUserStorageServer() {}

// UnsafeUserStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserStorageServer will
// result in compilation errors.
type UnsafeUserStorageServer interface {
	mustEmbedUnimplementedUserStorageServer()
}

func RegisterUserStorageServer(s grpc.ServiceRegistrar, srv UserStorageServer) {
	s.RegisterService(&UserStorage_ServiceDesc, srv)
}

func _UserStorage_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserStorage/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorage_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserStorageServer).List(m, &userStorageListServer{stream})
}

type UserStorage_ListServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userStorageListServer struct {
	grpc.ServerStream
}

func (x *userStorageListServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserStorage_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserStorage/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServer).Remove(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserStorage_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserStorageServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.UserStorage/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserStorageServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserStorage_ServiceDesc is the grpc.ServiceDesc for UserStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserStorage",
	HandlerType: (*UserStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserStorage_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _UserStorage_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserStorage_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _UserStorage_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/users.proto",
}
