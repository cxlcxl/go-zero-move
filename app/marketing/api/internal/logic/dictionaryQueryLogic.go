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
		if _, ok := utils.StringInArray(vars.TargetingDictionaryKeys, key); !ok {
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
				Id:    dictionary.Id,
				Pid:   dictionary.Pid,
				Label: dictionary.Label,
				Value: dictionary.Value,
			})
		} else {
			item := make([]*types.Dictionary, 1)
			item[0] = &types.Dictionary{
				Id:    dictionary.Id,
				Pid:   dictionary.Pid,
				Label: dictionary.Label,
				Value: dictionary.Value,
			}
			rs[dictionary.DictKey] = item
		}
	}
	if _, ok := rs["app_category"]; ok {
		rs["app_category"] = formatAppCategory(rs["app_category"])
	}
	if _, ok := rs["app_interest"]; ok {
		rs["app_interest"] = formatAppInterest(rs["app_interest"])
	}
	if _, ok := rs["carrier"]; ok {
		countries, err := l.svcCtx.MarketingRpcClient.GetCountries(l.ctx, &marketingcenter.EmptyParamsReq{})
		if err != nil {
			return nil, utils.RpcError(err, "国家数据为空")
		}
		rs["carrier"] = formatCarrier(rs["carrier"], countries.Countries)
	}
	if _, ok := rs["not_pre_define_audience"]; ok {
		rs["not_pre_define_audience"] = formatAudience(rs["not_pre_define_audience"])
	}
	if _, ok := rs["pre_define_audience"]; ok {
		rs["pre_define_audience"] = formatAudience(rs["pre_define_audience"])
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
				Id:       src[i].Id,
				Pid:      pid,
				Label:    src[i].Label,
				Value:    src[i].Value,
				Children: nil,
			})
		}
	}
	for name, id := range distPid {
		dist = append(dist, &types.Dictionary{
			Id:       id,
			Pid:      "0",
			Label:    name,
			Value:    id,
			Children: nil,
		})
	}
	return
}

func formatAppCategory(src []*types.Dictionary) (dist []*types.Dictionary) {
	dist = make([]*types.Dictionary, 0)
	for _, dictionary := range src {
		if dictionary.Pid == "0" {
			dist = append(dist, dictionary)
		}
	}
	return
}

func formatCarrier(src []*types.Dictionary, countries []*marketingcenter.CountriesResp_OverseasRegionCountry) (dist []*types.Dictionary) {
	tmp := make(map[string][]*types.Dictionary)
	for _, dictionary := range src {
		if _, ok := tmp[dictionary.Pid]; !ok {
			tmp[dictionary.Pid] = make([]*types.Dictionary, 0)
		}
		tmp[dictionary.Pid] = append(tmp[dictionary.Pid], dictionary)
	}
	for _, country := range countries {
		children, ok := tmp[country.CCode]
		if !ok {
			continue
		}
		dist = append(dist, &types.Dictionary{
			Id:       country.CId,
			Pid:      "0",
			Label:    country.CName,
			Value:    country.CId,
			Children: children,
		})
	}
	return
}

func formatAudience(src []*types.Dictionary) (dist []*types.Dictionary) {
	dist = make([]*types.Dictionary, 0)
	tmp := make(map[string]struct{})
	for _, dictionary := range src {
		if _, ok := tmp[dictionary.Id]; !ok {
			dist = append(dist, dictionary)
		}
		tmp[dictionary.Id] = struct{}{}
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
