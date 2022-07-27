// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package accountcenter

import (
	"context"

	"business/app/account/rpc/account"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActClientId          = account.ActClientId
	ActClientIdCreateReq = account.ActClientIdCreateReq
	ActClientIdUpdateReq = account.ActClientIdUpdateReq
	ActCreateReq         = account.ActCreateReq
	ActInfo              = account.ActInfo
	ActInfoReq           = account.ActInfoReq
	ActInfoResp          = account.ActInfoResp
	ActListReq           = account.ActListReq
	ActListResp          = account.ActListResp
	ActUpdateReq         = account.ActUpdateReq
	BaseResp             = account.BaseResp
	GetTokenReq          = account.GetTokenReq
	TokenInfo            = account.TokenInfo

	AccountCenter interface {
		ActList(ctx context.Context, in *ActListReq, opts ...grpc.CallOption) (*ActListResp, error)
		ActCreate(ctx context.Context, in *ActCreateReq, opts ...grpc.CallOption) (*BaseResp, error)
		ActUpdate(ctx context.Context, in *ActUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		ActInfo(ctx context.Context, in *ActInfoReq, opts ...grpc.CallOption) (*ActInfoResp, error)
		ActClientIdsCreate(ctx context.Context, in *ActClientIdCreateReq, opts ...grpc.CallOption) (*BaseResp, error)
		ActClientIdsUpdate(ctx context.Context, in *ActClientIdUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error)
		SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultAccountCenter struct {
		cli zrpc.Client
	}
)

func NewAccountCenter(cli zrpc.Client) AccountCenter {
	return &defaultAccountCenter{
		cli: cli,
	}
}

func (m *defaultAccountCenter) ActList(ctx context.Context, in *ActListReq, opts ...grpc.CallOption) (*ActListResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActList(ctx, in, opts...)
}

func (m *defaultAccountCenter) ActCreate(ctx context.Context, in *ActCreateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActCreate(ctx, in, opts...)
}

func (m *defaultAccountCenter) ActUpdate(ctx context.Context, in *ActUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActUpdate(ctx, in, opts...)
}

func (m *defaultAccountCenter) ActInfo(ctx context.Context, in *ActInfoReq, opts ...grpc.CallOption) (*ActInfoResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActInfo(ctx, in, opts...)
}

func (m *defaultAccountCenter) ActClientIdsCreate(ctx context.Context, in *ActClientIdCreateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActClientIdsCreate(ctx, in, opts...)
}

func (m *defaultAccountCenter) ActClientIdsUpdate(ctx context.Context, in *ActClientIdUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.ActClientIdsUpdate(ctx, in, opts...)
}

func (m *defaultAccountCenter) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.GetToken(ctx, in, opts...)
}

func (m *defaultAccountCenter) SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := account.NewAccountCenterClient(m.cli.Conn())
	return client.SetToken(ctx, in, opts...)
}
