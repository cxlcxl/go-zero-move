// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/account.proto

package account

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

// AccountCenterClient is the client API for AccountCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountCenterClient interface {
	AccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*BaseResp, error)
	AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
	GetAccountInfo(ctx context.Context, in *AccountInfoReq, opts ...grpc.CallOption) (*AccountInfo, error)
	GetAccountByClientId(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*AccountInfo, error)
	AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error)
	AccountSearch(ctx context.Context, in *AccountSearchReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
	GetAccountsByAccountIds(ctx context.Context, in *GetByAccountIdsReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
	GetDefaultAccountList(ctx context.Context, in *DefaultListReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
	GetParentAccountList(ctx context.Context, in *ParentListReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error)
	RefreshToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error)
	SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
}

type accountCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountCenterClient(cc grpc.ClientConnInterface) AccountCenterClient {
	return &accountCenterClient{cc}
}

func (c *accountCenterClient) AccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/AccountCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/AccountUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetAccountInfo(ctx context.Context, in *AccountInfoReq, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetAccountByClientId(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetAccountByClientId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error) {
	out := new(AccountListResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/AccountList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) AccountSearch(ctx context.Context, in *AccountSearchReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	out := new(AccountSearchResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/AccountSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetAccountsByAccountIds(ctx context.Context, in *GetByAccountIdsReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	out := new(AccountSearchResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetAccountsByAccountIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetDefaultAccountList(ctx context.Context, in *DefaultListReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	out := new(AccountSearchResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetDefaultAccountList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetParentAccountList(ctx context.Context, in *ParentListReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	out := new(AccountSearchResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetParentAccountList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	out := new(TokenInfo)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) RefreshToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	out := new(TokenInfo)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountCenterClient) SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/account.AccountCenter/SetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountCenterServer is the server API for AccountCenter service.
// All implementations must embed UnimplementedAccountCenterServer
// for forward compatibility
type AccountCenterServer interface {
	AccountCreate(context.Context, *AccountCreateReq) (*BaseResp, error)
	AccountUpdate(context.Context, *AccountUpdateReq) (*BaseResp, error)
	GetAccountInfo(context.Context, *AccountInfoReq) (*AccountInfo, error)
	GetAccountByClientId(context.Context, *GetTokenReq) (*AccountInfo, error)
	AccountList(context.Context, *AccountListReq) (*AccountListResp, error)
	AccountSearch(context.Context, *AccountSearchReq) (*AccountSearchResp, error)
	GetAccountsByAccountIds(context.Context, *GetByAccountIdsReq) (*AccountSearchResp, error)
	GetDefaultAccountList(context.Context, *DefaultListReq) (*AccountSearchResp, error)
	GetParentAccountList(context.Context, *ParentListReq) (*AccountSearchResp, error)
	GetToken(context.Context, *GetTokenReq) (*TokenInfo, error)
	RefreshToken(context.Context, *GetTokenReq) (*TokenInfo, error)
	SetToken(context.Context, *TokenInfo) (*BaseResp, error)
	mustEmbedUnimplementedAccountCenterServer()
}

// UnimplementedAccountCenterServer must be embedded to have forward compatible implementations.
type UnimplementedAccountCenterServer struct {
}

func (UnimplementedAccountCenterServer) AccountCreate(context.Context, *AccountCreateReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountCreate not implemented")
}
func (UnimplementedAccountCenterServer) AccountUpdate(context.Context, *AccountUpdateReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountUpdate not implemented")
}
func (UnimplementedAccountCenterServer) GetAccountInfo(context.Context, *AccountInfoReq) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (UnimplementedAccountCenterServer) GetAccountByClientId(context.Context, *GetTokenReq) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByClientId not implemented")
}
func (UnimplementedAccountCenterServer) AccountList(context.Context, *AccountListReq) (*AccountListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountList not implemented")
}
func (UnimplementedAccountCenterServer) AccountSearch(context.Context, *AccountSearchReq) (*AccountSearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountSearch not implemented")
}
func (UnimplementedAccountCenterServer) GetAccountsByAccountIds(context.Context, *GetByAccountIdsReq) (*AccountSearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsByAccountIds not implemented")
}
func (UnimplementedAccountCenterServer) GetDefaultAccountList(context.Context, *DefaultListReq) (*AccountSearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultAccountList not implemented")
}
func (UnimplementedAccountCenterServer) GetParentAccountList(context.Context, *ParentListReq) (*AccountSearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentAccountList not implemented")
}
func (UnimplementedAccountCenterServer) GetToken(context.Context, *GetTokenReq) (*TokenInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedAccountCenterServer) RefreshToken(context.Context, *GetTokenReq) (*TokenInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAccountCenterServer) SetToken(context.Context, *TokenInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToken not implemented")
}
func (UnimplementedAccountCenterServer) mustEmbedUnimplementedAccountCenterServer() {}

// UnsafeAccountCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountCenterServer will
// result in compilation errors.
type UnsafeAccountCenterServer interface {
	mustEmbedUnimplementedAccountCenterServer()
}

func RegisterAccountCenterServer(s grpc.ServiceRegistrar, srv AccountCenterServer) {
	s.RegisterService(&AccountCenter_ServiceDesc, srv)
}

func _AccountCenter_AccountCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).AccountCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/AccountCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).AccountCreate(ctx, req.(*AccountCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_AccountUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).AccountUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/AccountUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).AccountUpdate(ctx, req.(*AccountUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetAccountInfo(ctx, req.(*AccountInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetAccountByClientId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetAccountByClientId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetAccountByClientId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetAccountByClientId(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_AccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).AccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/AccountList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).AccountList(ctx, req.(*AccountListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_AccountSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).AccountSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/AccountSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).AccountSearch(ctx, req.(*AccountSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetAccountsByAccountIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccountIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetAccountsByAccountIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetAccountsByAccountIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetAccountsByAccountIds(ctx, req.(*GetByAccountIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetDefaultAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetDefaultAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetDefaultAccountList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetDefaultAccountList(ctx, req.(*DefaultListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetParentAccountList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetParentAccountList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetParentAccountList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetParentAccountList(ctx, req.(*ParentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).RefreshToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountCenter_SetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountCenterServer).SetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountCenter/SetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountCenterServer).SetToken(ctx, req.(*TokenInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountCenter_ServiceDesc is the grpc.ServiceDesc for AccountCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountCenter",
	HandlerType: (*AccountCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountCreate",
			Handler:    _AccountCenter_AccountCreate_Handler,
		},
		{
			MethodName: "AccountUpdate",
			Handler:    _AccountCenter_AccountUpdate_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _AccountCenter_GetAccountInfo_Handler,
		},
		{
			MethodName: "GetAccountByClientId",
			Handler:    _AccountCenter_GetAccountByClientId_Handler,
		},
		{
			MethodName: "AccountList",
			Handler:    _AccountCenter_AccountList_Handler,
		},
		{
			MethodName: "AccountSearch",
			Handler:    _AccountCenter_AccountSearch_Handler,
		},
		{
			MethodName: "GetAccountsByAccountIds",
			Handler:    _AccountCenter_GetAccountsByAccountIds_Handler,
		},
		{
			MethodName: "GetDefaultAccountList",
			Handler:    _AccountCenter_GetDefaultAccountList_Handler,
		},
		{
			MethodName: "GetParentAccountList",
			Handler:    _AccountCenter_GetParentAccountList_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _AccountCenter_GetToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AccountCenter_RefreshToken_Handler,
		},
		{
			MethodName: "SetToken",
			Handler:    _AccountCenter_SetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/account.proto",
}
