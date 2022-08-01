package logic

import (
	"business/common/utils"
	"business/common/vars"
	"context"
	"time"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizeCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizeCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizeCodeLogic {
	return &AuthorizeCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizeCodeLogic) AuthorizeCode(req *types.AuthCodeUrlReq) (resp *types.AuthCodeUrlResp, err error) {
	state := utils.MD5(time.Now().String())
	baseUrl := `https://login.cloud.huawei.com/oauth2/v2/authorize?`
	params := utils.HttpBuildQuery(map[string]string{
		"response_type": "code",
		"access_type":   "offline",
		"client_id":     "104661969",
		"state":         state,
		"redirect_uri":  "http://localhost:19527/#/marketing/callback",
		"scope":         "https://ads.cloud.huawei.com/promotion https://ads.cloud.huawei.com/report",
	})
	return &types.AuthCodeUrlResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		AuthCodeUrl: baseUrl + params,
	}, nil
}
