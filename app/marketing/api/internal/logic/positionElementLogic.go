package logic

import (
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PositionElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPositionElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionElementLogic {
	return &PositionElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PositionElementLogic) PositionElement(req *types.PositionElementReq) (resp *types.PositionElementResp, err error) {
	//info, err := l.svcCtx.MarketingRpcClient.GetPositionInfo(l.ctx, &marketingcenter.CreativeSizeInfoReq{CreativeSizeId: req.CreativeSizeId})
	//if err != nil {
	//	return nil, utils.RpcError(err, "请求错误")
	//}
	//d := statements.PositionElementRequest{
	//	AdvertiserId: info.AdvertiserId,
	//	Filtering: statements.PositionElementFilter{
	//		CreativeSizeId: req.CreativeSizeId,
	//	},
	//}
	//c, err := curl.New(l.svcCtx.Config.MarketingApis.Campaign.PositionDetail).Get().JsonData(d)
	//if err != nil {
	//	return nil, err
	//}
	//token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: info.AccountId})
	//if err != nil {
	//	return nil, utils.RpcError(err, "Token 信息获取失败")
	//}
	//if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
	//	return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	//}
	//t := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	//var response statements.PositionElementResponse
	//if err = c.Request(&response, curl.Authorization(t)); err != nil {
	//	return nil, errors.New("华为接口调用失败：" + err.Error())
	//}
	//if response.Code != "200" {
	//	return nil, errors.New("华为接口调用失败：" + response.Message)
	//}
	//subTypeElement := make(map[string]map[string]interface{})
	//for _, list := range response.Data.ElementInfoList {
	//	if _, ok := subTypeElement[list.Subtype]; ok {
	//		// subTypeElement[list.Subtype]
	//	} else {
	//		subTypeElement[list.Subtype] = make(map[string]interface{})
	//	}
	//}
	//return &types.PositionElementResp{
	//	BaseResp: types.BaseResp{
	//		Code: vars.ResponseCodeOk,
	//		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	//	},
	//	Data: nil,
	//}, nil
	return
}
