package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"business/app/marketing/rpc/internal/svc"
	"business/app/marketing/rpc/marketing"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictQueryLogic {
	return &DictQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictQueryLogic) DictQuery(in *marketing.DictionaryReq) (*marketing.DictionaryResp, error) {
	dictionaries, err := l.svcCtx.DictModel.FindDictionaries(l.ctx, in.DictKey)
	if err != nil {
		return nil, err
	}
	var dict []*marketing.DictionaryResp_Dictionary
	if err = copier.Copy(&dict, dictionaries); err != nil {
		return nil, err
	}
	return &marketing.DictionaryResp{
		Dictionaries: dict,
	}, nil
}
