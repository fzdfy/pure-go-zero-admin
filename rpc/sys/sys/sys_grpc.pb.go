// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc/sys/sys.proto

package sys

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
	Sys_UserLogin_FullMethodName  = "/sys.Sys/UserLogin"
	Sys_UserInfo_FullMethodName   = "/sys.Sys/UserInfo"
	Sys_UserAdd_FullMethodName    = "/sys.Sys/UserAdd"
	Sys_MenuAdd_FullMethodName    = "/sys.Sys/MenuAdd"
	Sys_MenuList_FullMethodName   = "/sys.Sys/MenuList"
	Sys_MenuUpdate_FullMethodName = "/sys.Sys/MenuUpdate"
	Sys_MenuDelete_FullMethodName = "/sys.Sys/MenuDelete"
)

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
	UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
	UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
	MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
	MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
	MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
	MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Sys_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	out := new(InfoResp)
	err := c.cc.Invoke(ctx, Sys_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	out := new(UserAddResp)
	err := c.cc.Invoke(ctx, Sys_UserAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	out := new(MenuAddResp)
	err := c.cc.Invoke(ctx, Sys_MenuAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	out := new(MenuListResp)
	err := c.cc.Invoke(ctx, Sys_MenuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	out := new(MenuUpdateResp)
	err := c.cc.Invoke(ctx, Sys_MenuUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	out := new(MenuDeleteResp)
	err := c.cc.Invoke(ctx, Sys_MenuDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	UserLogin(context.Context, *LoginReq) (*LoginResp, error)
	UserInfo(context.Context, *InfoReq) (*InfoResp, error)
	UserAdd(context.Context, *UserAddReq) (*UserAddResp, error)
	MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error)
	MenuList(context.Context, *MenuListReq) (*MenuListResp, error)
	MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error)
	MenuDelete(context.Context, *MenuDeleteReq) (*MenuDeleteResp, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) UserLogin(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedSysServer) UserInfo(context.Context, *InfoReq) (*InfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSysServer) UserAdd(context.Context, *UserAddReq) (*UserAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedSysServer) MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuAdd not implemented")
}
func (UnimplementedSysServer) MenuList(context.Context, *MenuListReq) (*MenuListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuList not implemented")
}
func (UnimplementedSysServer) MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuUpdate not implemented")
}
func (UnimplementedSysServer) MenuDelete(context.Context, *MenuDeleteReq) (*MenuDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuDelete not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserLogin(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserInfo(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserAdd(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuAdd(ctx, req.(*MenuAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuList(ctx, req.(*MenuListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuUpdate(ctx, req.(*MenuUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuDelete(ctx, req.(*MenuDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sys.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _Sys_UserLogin_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Sys_UserInfo_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Sys_UserAdd_Handler,
		},
		{
			MethodName: "MenuAdd",
			Handler:    _Sys_MenuAdd_Handler,
		},
		{
			MethodName: "MenuList",
			Handler:    _Sys_MenuList_Handler,
		},
		{
			MethodName: "MenuUpdate",
			Handler:    _Sys_MenuUpdate_Handler,
		},
		{
			MethodName: "MenuDelete",
			Handler:    _Sys_MenuDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/sys/sys.proto",
}
