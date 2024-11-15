// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: nf6.proto

package nf6

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

const (
	Nf6Insecure_GetCaCert_FullMethodName = "/nf6.Nf6Insecure/GetCaCert"
	Nf6Insecure_Ping_FullMethodName      = "/nf6.Nf6Insecure/Ping"
	Nf6Insecure_Register_FullMethodName  = "/nf6.Nf6Insecure/Register"
)

// Nf6InsecureClient is the client API for Nf6Insecure service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Nf6InsecureClient interface {
	GetCaCert(ctx context.Context, in *GetCaCertRequest, opts ...grpc.CallOption) (*GetCaCertReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
}

type nf6InsecureClient struct {
	cc grpc.ClientConnInterface
}

func NewNf6InsecureClient(cc grpc.ClientConnInterface) Nf6InsecureClient {
	return &nf6InsecureClient{cc}
}

func (c *nf6InsecureClient) GetCaCert(ctx context.Context, in *GetCaCertRequest, opts ...grpc.CallOption) (*GetCaCertReply, error) {
	out := new(GetCaCertReply)
	err := c.cc.Invoke(ctx, Nf6Insecure_GetCaCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nf6InsecureClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, Nf6Insecure_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nf6InsecureClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, Nf6Insecure_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Nf6InsecureServer is the server API for Nf6Insecure service.
// All implementations must embed UnimplementedNf6InsecureServer
// for forward compatibility
type Nf6InsecureServer interface {
	GetCaCert(context.Context, *GetCaCertRequest) (*GetCaCertReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	mustEmbedUnimplementedNf6InsecureServer()
}

// UnimplementedNf6InsecureServer must be embedded to have forward compatible implementations.
type UnimplementedNf6InsecureServer struct {
}

func (UnimplementedNf6InsecureServer) GetCaCert(context.Context, *GetCaCertRequest) (*GetCaCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaCert not implemented")
}
func (UnimplementedNf6InsecureServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedNf6InsecureServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedNf6InsecureServer) mustEmbedUnimplementedNf6InsecureServer() {}

// UnsafeNf6InsecureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Nf6InsecureServer will
// result in compilation errors.
type UnsafeNf6InsecureServer interface {
	mustEmbedUnimplementedNf6InsecureServer()
}

func RegisterNf6InsecureServer(s grpc.ServiceRegistrar, srv Nf6InsecureServer) {
	s.RegisterService(&Nf6Insecure_ServiceDesc, srv)
}

func _Nf6Insecure_GetCaCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6InsecureServer).GetCaCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Insecure_GetCaCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6InsecureServer).GetCaCert(ctx, req.(*GetCaCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nf6Insecure_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6InsecureServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Insecure_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6InsecureServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nf6Insecure_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6InsecureServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Insecure_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6InsecureServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nf6Insecure_ServiceDesc is the grpc.ServiceDesc for Nf6Insecure service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nf6Insecure_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nf6.Nf6Insecure",
	HandlerType: (*Nf6InsecureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaCert",
			Handler:    _Nf6Insecure_GetCaCert_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Nf6Insecure_Ping_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Nf6Insecure_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nf6.proto",
}

const (
	Nf6Secure_CreateRepo_FullMethodName = "/nf6.Nf6Secure/CreateRepo"
	Nf6Secure_ListRepos_FullMethodName  = "/nf6.Nf6Secure/ListRepos"
	Nf6Secure_RenameRepo_FullMethodName = "/nf6.Nf6Secure/RenameRepo"
	Nf6Secure_WhoAmI_FullMethodName     = "/nf6.Nf6Secure/WhoAmI"
)

// Nf6SecureClient is the client API for Nf6Secure service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Nf6SecureClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoReply, error)
	ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposReply, error)
	RenameRepo(ctx context.Context, in *RenameRepoRequest, opts ...grpc.CallOption) (*RenameRepoReply, error)
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error)
}

type nf6SecureClient struct {
	cc grpc.ClientConnInterface
}

func NewNf6SecureClient(cc grpc.ClientConnInterface) Nf6SecureClient {
	return &nf6SecureClient{cc}
}

func (c *nf6SecureClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoReply, error) {
	out := new(CreateRepoReply)
	err := c.cc.Invoke(ctx, Nf6Secure_CreateRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nf6SecureClient) ListRepos(ctx context.Context, in *ListReposRequest, opts ...grpc.CallOption) (*ListReposReply, error) {
	out := new(ListReposReply)
	err := c.cc.Invoke(ctx, Nf6Secure_ListRepos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nf6SecureClient) RenameRepo(ctx context.Context, in *RenameRepoRequest, opts ...grpc.CallOption) (*RenameRepoReply, error) {
	out := new(RenameRepoReply)
	err := c.cc.Invoke(ctx, Nf6Secure_RenameRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nf6SecureClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIReply, error) {
	out := new(WhoAmIReply)
	err := c.cc.Invoke(ctx, Nf6Secure_WhoAmI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Nf6SecureServer is the server API for Nf6Secure service.
// All implementations must embed UnimplementedNf6SecureServer
// for forward compatibility
type Nf6SecureServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoReply, error)
	ListRepos(context.Context, *ListReposRequest) (*ListReposReply, error)
	RenameRepo(context.Context, *RenameRepoRequest) (*RenameRepoReply, error)
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error)
	mustEmbedUnimplementedNf6SecureServer()
}

// UnimplementedNf6SecureServer must be embedded to have forward compatible implementations.
type UnimplementedNf6SecureServer struct {
}

func (UnimplementedNf6SecureServer) CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepo not implemented")
}
func (UnimplementedNf6SecureServer) ListRepos(context.Context, *ListReposRequest) (*ListReposReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepos not implemented")
}
func (UnimplementedNf6SecureServer) RenameRepo(context.Context, *RenameRepoRequest) (*RenameRepoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameRepo not implemented")
}
func (UnimplementedNf6SecureServer) WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedNf6SecureServer) mustEmbedUnimplementedNf6SecureServer() {}

// UnsafeNf6SecureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Nf6SecureServer will
// result in compilation errors.
type UnsafeNf6SecureServer interface {
	mustEmbedUnimplementedNf6SecureServer()
}

func RegisterNf6SecureServer(s grpc.ServiceRegistrar, srv Nf6SecureServer) {
	s.RegisterService(&Nf6Secure_ServiceDesc, srv)
}

func _Nf6Secure_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6SecureServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Secure_CreateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6SecureServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nf6Secure_ListRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6SecureServer).ListRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Secure_ListRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6SecureServer).ListRepos(ctx, req.(*ListReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nf6Secure_RenameRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6SecureServer).RenameRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Secure_RenameRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6SecureServer).RenameRepo(ctx, req.(*RenameRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nf6Secure_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Nf6SecureServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nf6Secure_WhoAmI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Nf6SecureServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nf6Secure_ServiceDesc is the grpc.ServiceDesc for Nf6Secure service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nf6Secure_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nf6.Nf6Secure",
	HandlerType: (*Nf6SecureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _Nf6Secure_CreateRepo_Handler,
		},
		{
			MethodName: "ListRepos",
			Handler:    _Nf6Secure_ListRepos_Handler,
		},
		{
			MethodName: "RenameRepo",
			Handler:    _Nf6Secure_RenameRepo_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _Nf6Secure_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nf6.proto",
}
