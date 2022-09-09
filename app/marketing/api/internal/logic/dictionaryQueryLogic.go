package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/vars"
	"context"
	"errors"
	"strings"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictionaryQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictionaryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryQueryLogic {
	return &DictionaryQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictionaryQueryLogic) DictionaryQuery(req *types.DictReq) (resp *types.DictResp, err error) {
	keys := strings.Split(req.DictKey, ",")
	for _, key := range keys {
		if !stringInArray(key, vars.TargetingDictionaryKeys) {
			return nil, errors.New("参数有误")
		}
	}
	dictionaries, err := l.svcCtx.MarketingRpcClient.DictQuery(l.ctx, &marketingcenter.DictionaryReq{DictKey: keys})
	if err != nil {
		return nil, err
	}
	rs := make(map[string][]*types.Dictionary)
	for _, dictionary := range dictionaries.Dictionaries {
		if _, ok := rs[dictionary.DictKey]; ok {
			rs[dictionary.DictKey] = append(rs[dictionary.DictKey], &types.Dictionary{
				DictKey: dictionary.DictKey,
				Id:      dictionary.Id,
				Pid:     dictionary.Pid,
				Label:   dictionary.Label,
				Value:   dictionary.Value,
				Code:    dictionary.Code,
				Seq:     dictionary.Seq,
			})
		} else {
			item := make([]*types.Dictionary, 1)
			item[0] = &types.Dictionary{
				DictKey: dictionary.DictKey,
				Id:      dictionary.Id,
				Pid:     dictionary.Pid,
				Label:   dictionary.Label,
				Value:   dictionary.Value,
				Code:    dictionary.Code,
				Seq:     dictionary.Seq,
			}
			rs[dictionary.DictKey] = item
		}
	}
	return &types.DictResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs,
	}, nil
}

func stringInArray(k string, s []string) bool {
	for _, _s := range s {
		if _s == k {
			return true
		}
	}
	return false
}
