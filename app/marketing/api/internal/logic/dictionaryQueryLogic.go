package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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
		if !utils.StringInArray(vars.TargetingDictionaryKeys, key) {
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
	if _, ok := rs["app_interest"]; ok {
		rs["app_interest"] = formatAppInterest(rs["app_interest"])
	}
	return &types.DictResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs,
	}, nil
}

func formatAppInterest(src []*types.Dictionary) (dist []*types.Dictionary) {
	dist = make([]*types.Dictionary, 0)
	distPid := make(map[string]string)
	for i := range src {
		if src[i].Pid != "" {
			pidId := fmt.Sprintf("%d%d", time.Now().Unix(), len(distPid)+1)
			pid := pidId
			if existsPid, ok := distPid[src[i].Pid]; !ok {
				distPid[src[i].Pid] = pidId
			} else {
				pid = existsPid
			}
			dist = append(dist, &types.Dictionary{
				DictKey:  src[i].DictKey,
				Id:       src[i].Id,
				Pid:      pid,
				Label:    src[i].Label,
				Value:    src[i].Value,
				Code:     src[i].Code,
				Seq:      src[i].Seq,
				Children: nil,
			})
		}
	}
	for name, id := range distPid {
		dist = append(dist, &types.Dictionary{
			DictKey:  "app_interest",
			Id:       id,
			Pid:      "0",
			Label:    name,
			Value:    id,
			Code:     "",
			Seq:      "",
			Children: nil,
		})
	}
	return
}

// 数组转树
func dictionaryToTree(origin []*types.Dictionary, pid string) []*types.Dictionary {
	d := make([]*types.Dictionary, 0)
	for _, v := range origin {
		if v.Pid == pid {
			v.Children = dictionaryToTree(origin, v.Id)
			d = append(d, v)
		}
	}
	return d
}
