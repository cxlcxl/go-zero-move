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

type ActClientIdsCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActClientIdsCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActClientIdsCreateLogic {
	return &ActClientIdsCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActClientIdsCreateLogic) ActClientIdsCreate(in *account.ActClientIdCreateReq) (*account.BaseResp, error) {
	now := time.Now()
	accountClientIds := make([]*model.AccountClientIds, 0)
	for _, clientIds := range in.AccountClientIds {
		clientIdModel := &model.AccountClientIds{
			AccountId: in.AccountId,
			ClientId:  clientIds.ClientId,
			Secret:    clientIds.Secret,
			Comment:   clientIds.Comment,
			CreatedAt: now,
			UpdatedAt: now,
		}
		accountClientIds = append(accountClientIds, clientIdModel)
	}
	err := l.svcCtx.ActClientModel.BatchInsert(l.ctx, nil, accountClientIds)
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
