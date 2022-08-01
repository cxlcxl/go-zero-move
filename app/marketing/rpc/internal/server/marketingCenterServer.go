// Code generated by goctl. DO NOT EDIT!
// Source: marketing.proto

package server

import (
	"context"

	"business/app/marketing/rpc/internal/logic"
	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"
)

type MarketingCenterServer struct {
	svcCtx *svc.ServiceContext
	marketing.UnimplementedMarketingCenterServer
}

func NewMarketingCenterServer(svcCtx *svc.ServiceContext) *MarketingCenterServer {
	return &MarketingCenterServer{
		svcCtx: svcCtx,
	}
}

//  账户板块 RPC 服务
func (s *MarketingCenterServer) AccountCreate(ctx context.Context, in *marketing.AccountCreateReq) (*marketing.BaseResp, error) {
	l := logic.NewAccountCreateLogic(ctx, s.svcCtx)
	return l.AccountCreate(in)
}

func (s *MarketingCenterServer) AccountUpdate(ctx context.Context, in *marketing.AccountUpdateReq) (*marketing.BaseResp, error) {
	l := logic.NewAccountUpdateLogic(ctx, s.svcCtx)
	return l.AccountUpdate(in)
}

func (s *MarketingCenterServer) GetAccountInfo(ctx context.Context, in *marketing.AccountInfoReq) (*marketing.AccountInfo, error) {
	l := logic.NewGetAccountInfoLogic(ctx, s.svcCtx)
	return l.GetAccountInfo(in)
}

func (s *MarketingCenterServer) AccountList(ctx context.Context, in *marketing.AccountListReq) (*marketing.AccountListResp, error) {
	l := logic.NewAccountListLogic(ctx, s.svcCtx)
	return l.AccountList(in)
}

func (s *MarketingCenterServer) GetToken(ctx context.Context, in *marketing.GetTokenReq) (*marketing.TokenInfo, error) {
	l := logic.NewGetTokenLogic(ctx, s.svcCtx)
	return l.GetToken(in)
}

func (s *MarketingCenterServer) SetToken(ctx context.Context, in *marketing.TokenInfo) (*marketing.BaseResp, error) {
	l := logic.NewSetTokenLogic(ctx, s.svcCtx)
	return l.SetToken(in)
}

func (s *MarketingCenterServer) PromotionCreate(ctx context.Context, in *marketing.PromotionCreateReq) (*marketing.BaseResp, error) {
	l := logic.NewPromotionCreateLogic(ctx, s.svcCtx)
	return l.PromotionCreate(in)
}