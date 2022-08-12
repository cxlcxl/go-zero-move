package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountInfoLogic {
	return &GetAccountInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountInfoLogic) GetAccountInfo(in *account.AccountInfoReq) (*account.AccountInfo, error) {
	act, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var actInfo account.AccountInfo
	err = copier.Copy(&actInfo, act)
	if err != nil {
		return nil, err
	}
	actInfo.CreatedAt = act.CreatedAt.Unix()
	actInfo.UpdatedAt = act.UpdatedAt.Unix()
	return &actInfo, nil
}
