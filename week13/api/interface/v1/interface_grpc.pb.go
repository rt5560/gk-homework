// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// InterfaceClient is the client API for Interface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterfaceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
}

type interfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterfaceClient(cc grpc.ClientConnInterface) InterfaceClient {
	return &interfaceClient{cc}
}

func (c *interfaceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.interface.v1.Interface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfaceServer is the server API for Interface service.
// All implementations must embed UnimplementedInterfaceServer
// for forward compatibility
type InterfaceServer interface {
	Login(context.Context, *LoginReq) (*LoginReply, error)
	mustEmbedUnimplementedInterfaceServer()
}

// UnimplementedInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedInterfaceServer struct {
}

func (UnimplementedInterfaceServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedInterfaceServer) mustEmbedUnimplementedInterfaceServer() {}

// UnsafeInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterfaceServer will
// result in compilation errors.
type UnsafeInterfaceServer interface {
	mustEmbedUnimplementedInterfaceServer()
}

func RegisterInterfaceServer(s grpc.ServiceRegistrar, srv InterfaceServer) {
	s.RegisterService(&Interface_ServiceDesc, srv)
}

func _Interface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.interface.v1.Interface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Interface_ServiceDesc is the grpc.ServiceDesc for Interface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.interface.v1.Interface",
	HandlerType: (*InterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Interface_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface/v1/interface.proto",
}