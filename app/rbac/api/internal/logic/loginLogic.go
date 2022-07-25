package logic

import (
	"business/app/rbac/api/internal/svc"
	"business/app/rbac/api/internal/types"
	"business/app/rbac/rpc/rbac"
	"business/app/rbac/rpc/rbaccenter"
	"business/common/utils"
	"business/common/validators"
	"business/common/vars"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Login) (resp *types.LoginResp, err error) {
	if err = validators.New(validators.Email(req.Email), validators.Password(req.Pass)); err != nil {
		return
	}
	res, err := l.svcCtx.RbacClient.Login(l.ctx, &rbaccenter.LoginReq{Email: req.Email, Pass: req.Pass})
	if err != nil {
		return nil, utils.RpcError(err, "没有此邮箱信息")
	}
	token, e := l.jwtToken(res)
	if e != nil {
		return nil, e
	}
	return &types.LoginResp{
		BaseResp: types.BaseResp{Code: vars.ResponseCodeOk, Msg: vars.ResponseMsg[vars.ResponseCodeOk]},
		Data: types.LoginTokenResp{
			Id: res.Id, Username: res.Username, Email: res.Email, Token: token,
			Mobile: res.Mobile,
			RoleId: res.RoleId,
		},
	}, nil
}

func (l *LoginLogic) jwtToken(u *rbac.LoginResp) (token string, err error) {
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"id": u.Id, "username": u.Username, "email": u.Email, "mobile": u.Mobile, "role_id": u.RoleId,
		"exp": now + l.svcCtx.Config.Auth.Expire,
		"iat": now,
	}
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = claims

	token, err = t.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	return
}
