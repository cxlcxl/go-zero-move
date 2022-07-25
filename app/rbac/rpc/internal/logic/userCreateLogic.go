package logic

import (
	"business/app/rbac/rpc/internal/svc"
	model2 "business/app/rbac/rpc/model"
	"business/app/rbac/rpc/rbac"
	"business/common/utils"
	"business/common/validators"
	"business/common/vars"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *rbac.UserCreateReq) (*rbac.BaseResp, error) {
	if err := validators.New(
		validators.Email(in.Email),
		validators.Password(in.Pass),
		validators.Empty(in.Username),
	); err != nil {
		return nil, err
	}
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model2.ErrNotFound {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("邮箱已存在")
	}

	secret := utils.GeneratePassSecret()
	_, err = l.svcCtx.UserModel.Insert(l.ctx, nil, &model2.Users{
		Email:     in.Email,
		Username:  in.Username,
		Mobile:    in.Mobile,
		State:     1,
		RoleId:    in.RoleId,
		Secret:    secret,
		Pass:      utils.Password(in.Pass, secret),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &rbac.BaseResp{Code: vars.ResponseCodeOk}, nil
}
