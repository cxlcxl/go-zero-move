package logic

import (
	"business/common/utils"
	"context"
	"encoding/json"
	"fmt"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromotionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromotionListLogic {
	return &PromotionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PromotionListLogic) PromotionList(req *types.PromotionListReq) (resp *types.PromotionListResp, err error) {
	requestJson := map[string]interface{}{
		"page":      req.Page,
		"page_size": req.PageSize,
		"filtering": map[string]interface{}{
			"campaign_name": req.CampaignName,
		},
	}
	ds, err := json.Marshal(requestJson)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer DAEAAOda5NC2fNM0IDMNtGtyGHWyYBLYAleF/aTKOeZHwkTvyRqHu/zlJzz/caXcj0Q/b3BGZCPNqqdZGm8R+/1oNnMJlE6KxVLatRbeOWT+MES36uBDtBk=",
	}
	request, err := utils.HttpRequest(l.svcCtx.Config.MarketingApis.Promotion.Query, string(ds), "GET", headers)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(request))
	return
}
