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

type ActUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActUpdateLogic {
	return &ActUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActUpdateLogic) ActUpdate(in *account.ActUpdateReq) (*account.BaseResp, error) {
	err := l.svcCtx.ActModel.Update(l.ctx, &model.Accounts{
		Id:          in.Id,
		AccountId:   in.AccountId,
		AccountType: in.AccountType,
		Email:       in.Email,
		AccountName: in.AccountName,
		UpdatedAt:   time.Now(),
		State:       in.State,
	})
	if err != nil {
		return nil, err
	}
	return &account.BaseResp{
		Code: vars.ResponseCodeOk,
	}, nil
}
