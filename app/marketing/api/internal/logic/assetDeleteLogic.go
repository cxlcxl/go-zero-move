package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/marketing/api/internal/statements"
	"business/app/marketing/rpc/marketingcenter"
	"business/common/curl"
	"business/common/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetDeleteLogic {
	return &AssetDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetDeleteLogic) AssetDelete(req *types.AssetDeleteReq) (resp *types.BaseResp, err error) {
	if err = l.svcCtx.Validator.StructCtx(l.ctx, req); err != nil {
		return nil, err
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Asset.Delete).Post().JsonData(req)
	if err != nil {
		return nil, err
	}
	token, err := l.svcCtx.AccountRpcClient.GetToken(l.ctx, &accountcenter.GetTokenReq{AccountId: req.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "Token 信息获取失败")
	}
	if time.Unix(token.ExpiredAt, 0).Before(time.Now()) {
		return nil, errors.New("Token 已过期，请先到账户列表页刷新 Token ")
	}
	var response statements.AdsAssetDeleteResponse
	if err = c.Request(&response, curl.Authorization(fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))); err != nil {
		return nil, errors.New("华为接口调用失败：" + err.Error())
	}
	errorMsg := ""
	assetIds := make(map[int64]struct{})
	for _, id := range req.AssetIds {
		assetIds[id] = struct{}{}
	}
	for _, datum := range response.Data {
		// 1005031 是已删除的
		if _, ok := assetIds[datum.AssetId]; ok && datum.ErrorCode != "1005031" {
			delete(assetIds, datum.AssetId)
			errorMsg = datum.Message
		}
	}
	if len(assetIds) == 0 {
		return nil, errors.New("全部删除失败：" + errorMsg)
	}
	removeAssetIds := make([]int64, 0)
	for id := range assetIds {
		removeAssetIds = append(removeAssetIds, id)
	}
	if _, err = l.svcCtx.MarketingRpcClient.DeleteAssets(l.ctx, &marketingcenter.AssetDeleteReq{AssetIds: removeAssetIds}); err != nil {
		return nil, utils.RpcError(err, "")
	}
	return successFul(), nil
}
