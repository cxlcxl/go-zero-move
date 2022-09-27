package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreativePriceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreativePriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreativePriceLogic {
	return &CreativePriceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreativePriceLogic) CreativePrice(req *types.CreativeSizePriceReq) (resp *types.CreativeSizePriceResp, err error) {
	info, err := l.svcCtx.MarketingRpcClient.GetPositionInfo(l.ctx, &marketingcenter.CreativeSizeInfoReq{CreativeSizeId: req.CreativeSizeId})
	if err != nil {
		return nil, utils.RpcError(err, "请求错误")
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Campaign.CreativeSizePrice).Get().JsonData(statements.CreativeSizePriceRequest{
		AdvertiserId: info.AdvertiserId,
		Filtering:    statements.CreativeSizePriceFilter{CreativeSizeId: req.CreativeSizeId, PriceType: req.PriceType},
	})
	if err != nil {
		return nil, err
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: info.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "Token 信息获取失败")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	var response statements.CreativeSizePriceResponse
	if err = c.Request(&response, curl.Authorization(t)); err != nil {
		return nil, errors.New("华为接口拉取错误：" + response.Message)
	}
	return &types.CreativeSizePriceResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: response.Data.FloorPrice,
	}, nil
}
