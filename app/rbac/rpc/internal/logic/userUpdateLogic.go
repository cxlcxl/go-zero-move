package logic

import (
	"business/app/rbac/rpc/internal/svc"
	model2 "business/app/rbac/rpc/model"
	"business/app/rbac/rpc/rbac"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *rbac.UserUpdateReq) (*rbac.BaseResp, error) {
	email, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err != model2.ErrNotFound {
			return nil, err
		}
	} else if email.Id != in.Id {
		return nil, errors.New("邮箱已存在")
	}
	user, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var pass string
	if in.Pass == "" {
		pass = user.Pass
	} else {
		pass = utils.Password(in.Pass, user.Secret)
	}
	if err = l.svcCtx.UserModel.Update(l.ctx, &model2.Users{
		Id: in.Id, State: in.State, Username: in.Username, Pass: pass, Mobile: in.Mobile, Email: in.Email, UpdatedAt: time.Now(), RoleId: in.RoleId,
	}); err != nil {
		return nil, err
	}
	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
