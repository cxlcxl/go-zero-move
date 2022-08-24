package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/account/rpc/account"
	"business/app/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParentAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParentAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParentAccountListLogic {
	return &GetParentAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParentAccountListLogic) GetParentAccountList(in *account.ParentListReq) (*account.AccountSearchResp, error) {
	accounts, err := l.svcCtx.AccountModel.RemoteAccounts(l.ctx, in.AccountName, 1)
	if err != nil {
		return nil, err
	}
	var res []*account.AccountInfo
	err = copier.Copy(&res, accounts)
	if err != nil {
		return nil, err
	}
	if in.Id > 0 {
		info, _ := l.svcCtx.AccountModel.FindOne(l.ctx, in.Id)
		res = append(res, &account.AccountInfo{
			Id:           info.Id,
			AccountName:  info.AccountName,
			AccountType:  info.AccountType,
			AdvertiserId: info.AdvertiserId,
			DeveloperId:  info.DeveloperId,
			ClientId:     info.ClientId,
			Secret:       info.Secret,
			State:        info.State,
			IsAuth:       info.IsAuth,
			CreatedAt:    info.CreatedAt.Unix(),
			UpdatedAt:    info.UpdatedAt.Unix(),
			ParentId:     info.ParentId,
		})
	}
	return &account.AccountSearchResp{
		Accounts: res,
	}, nil
}
