package logic

import (
	"business/app/account/rpc/accountcenter"
	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"
	"business/common/utils"
	"business/common/vars"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessTokenLogic {
	return &AccessTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type AT struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}

func (l *AccessTokenLogic) AccessToken(req *types.AccessTokenReq) (resp *types.AccessTokenResp, err error) {
	clientId := "104661969"
	host := `https://login.cloud.huawei.com/oauth2/v2/token`
	data := utils.HttpBuildQuery(map[string]string{
		"grant_type":    "authorization_code",
		"code":          req.AuthorizationCode,
		"client_id":     clientId,
		"client_secret": "ac77ad1f48b543f5730dd85c3cc9f53ffeafbcced09908a07809a935e211e946",
		"redirect_uri":  "http://localhost:19527/#/marketing/callback",
	})
	bs, err := utils.PostHttpRequest(host, data, map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
	if err != nil {
		return nil, err
	}
	var at AT
	err = json.Unmarshal(bs, &at)
	if err != nil {
		return nil, err
	}
	if at.Error != 0 {
		return nil, errors.New("华为接口调用失败：" + at.ErrorDescription)
	}
	_, err = l.svcCtx.AccountRpcClient.SetToken(l.ctx, &accountcenter.TokenInfo{
		ClientId:     104661969,
		AccessToken:  at.AccessToken,
		RefreshToken: at.RefreshToken,
		ExpiredAt:    at.ExpiresIn,
		TokenType:    at.TokenType,
	})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.AccessTokenResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		AccessToken: at.AccessToken,
	}, nil
}
