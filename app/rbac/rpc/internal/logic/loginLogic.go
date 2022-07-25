package logic

import (
	"business/app/rbac/rpc/internal/svc"
	"business/app/rbac/rpc/rbac"
	"business/common/utils"
	"business/common/validators"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *rbac.LoginReq) (*rbac.LoginResp, error) {
	if err := validators.New(validators.Email(in.Email), validators.Password(in.Pass)); err != nil {
		return nil, err
	}
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	pass := utils.Password(in.Pass, user.Secret)
	if pass != user.Pass {
		return nil, errors.New("密码错误")
	}

	return &rbac.LoginResp{Id: user.Id, Username: user.Username, Email: user.Email, Mobile: user.Mobile, RoleId: user.RoleId}, nil
}
