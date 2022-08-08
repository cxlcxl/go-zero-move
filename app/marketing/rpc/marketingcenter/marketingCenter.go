// Code generated by goctl. DO NOT EDIT!
// Source: marketing.proto

package marketingcenter

import (
	"context"

	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResp           = marketing.BaseResp
	PromotionCreateReq = marketing.PromotionCreateReq
	AccountCreateReq   = marketing.AccountCreateReq
	AccountUpdateReq   = marketing.AccountUpdateReq
	AccountInfoReq     = marketing.AccountInfoReq
	AccountInfo        = marketing.AccountInfo
	GetTokenReq        = marketing.GetTokenReq
	AccountListReq     = marketing.AccountListReq
	AccountListResp    = marketing.AccountListResp
	TokenInfo          = marketing.TokenInfo
	GetByAccountIdsReq = marketing.GetByAccountIdsReq
	AccountSearchResp  = marketing.AccountSearchResp
	AccountSearchReq   = marketing.AccountSearchReq

	MarketingCenter interface {
		//  账户板块 RPC 服务
		AccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*BaseResp, error)
		AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetAccountInfo(ctx context.Context, in *AccountInfoReq, opts ...grpc.CallOption) (*AccountInfo, error)
		GetAccountByClientId(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*AccountInfo, error)
		AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error)
		AccountSearch(ctx context.Context, in *AccountSearchReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
		GetAccountsByAccountIds(ctx context.Context, in *GetByAccountIdsReq, opts ...grpc.CallOption) (*AccountSearchResp, error)
		GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error)
		SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
		PromotionCreate(ctx context.Context, in *PromotionCreateReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultMarketingCenter struct {
		cli zrpc.Client
	}
)

func NewMarketingCenter(cli zrpc.Client) MarketingCenter {
	return &defaultMarketingCenter{
		cli: cli,
	}
}

//  账户板块 RPC 服务
func (m *defaultMarketingCenter) AccountCreate(ctx context.Context, in *AccountCreateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.AccountCreate(ctx, in, opts...)
}

func (m *defaultMarketingCenter) AccountUpdate(ctx context.Context, in *AccountUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.AccountUpdate(ctx, in, opts...)
}

func (m *defaultMarketingCenter) GetAccountInfo(ctx context.Context, in *AccountInfoReq, opts ...grpc.CallOption) (*AccountInfo, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.GetAccountInfo(ctx, in, opts...)
}

func (m *defaultMarketingCenter) GetAccountByClientId(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*AccountInfo, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.GetAccountByClientId(ctx, in, opts...)
}

func (m *defaultMarketingCenter) AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.AccountList(ctx, in, opts...)
}

func (m *defaultMarketingCenter) AccountSearch(ctx context.Context, in *AccountSearchReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.AccountSearch(ctx, in, opts...)
}

func (m *defaultMarketingCenter) GetAccountsByAccountIds(ctx context.Context, in *GetByAccountIdsReq, opts ...grpc.CallOption) (*AccountSearchResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.GetAccountsByAccountIds(ctx, in, opts...)
}

func (m *defaultMarketingCenter) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.GetToken(ctx, in, opts...)
}

func (m *defaultMarketingCenter) SetToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.SetToken(ctx, in, opts...)
}

func (m *defaultMarketingCenter) PromotionCreate(ctx context.Context, in *PromotionCreateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := marketing.NewMarketingCenterClient(m.cli.Conn())
	return client.PromotionCreate(ctx, in, opts...)
}
