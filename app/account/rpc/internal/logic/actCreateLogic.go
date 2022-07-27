package logic

import (
	"business/app/account/rpc/model"
	"business/common/vars"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActCreateLogic {
	return &ActCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActCreateLogic) ActCreate(in *account.ActCreateReq) (*account.BaseResp, error) {
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
	err := l.svcCtx.ActModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		if _, err := l.svcCtx.ActModel.Insert(l.ctx, session, &model.Accounts{
			AccountId:   in.AccountId,
			AccountType: in.AccountType,
			Email:       in.Email,
			AccountName: in.AccountName,
			CreatedAt:   now,
			UpdatedAt:   now,
			State:       1,
		}); err != nil {
			return err
		}
		if err := l.svcCtx.ActClientModel.BatchInsert(l.ctx, session, accountClientIds); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
