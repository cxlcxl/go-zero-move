package logic

import (
	"business/app/account/rpc/model"
	"business/common/vars"
	"context"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActClientIdsUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActClientIdsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActClientIdsUpdateLogic {
	return &ActClientIdsUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActClientIdsUpdateLogic) ActClientIdsUpdate(in *account.ActClientIdUpdateReq) (*account.BaseResp, error) {
	err := l.svcCtx.ActClientModel.Update(l.ctx, &model.AccountClientIds{
		Id:        in.AccountClientIds.Id,
		AccountId: in.AccountId,
		ClientId:  in.AccountClientIds.ClientId,
		Secret:    in.AccountClientIds.Secret,
		Comment:   in.AccountClientIds.Comment,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
