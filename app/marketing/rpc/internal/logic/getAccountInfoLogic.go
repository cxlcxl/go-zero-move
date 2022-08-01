package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

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

func (l *GetAccountInfoLogic) GetAccountInfo(in *marketing.AccountInfoReq) (*marketing.AccountInfo, error) {
	act, err := l.svcCtx.AccountModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	var account marketing.AccountInfo
	err = copier.Copy(&account, act)
	if err != nil {
		return nil, err
	}
	account.CreatedAt = act.CreatedAt.Unix()
	account.UpdatedAt = act.UpdatedAt.Unix()
	return &account, nil
}
