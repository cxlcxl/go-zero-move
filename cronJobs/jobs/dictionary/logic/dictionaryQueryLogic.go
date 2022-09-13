package logic

import (
	"business/common/curl"
	"business/common/utils"
	"business/common/vars"
	"business/cronJobs/jobs"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type DictionaryQueryLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *jobs.QueryParam
	statDay   string
	pageSize  int64
}

func NewDictionaryQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictionaryQueryLogic {
	return &DictionaryQueryLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
	}
}

func (l *DictionaryQueryLogic) DictionaryQuery() (err error) {
	go jobs.GetTokens(l.ctx, l.svcCtx, l.tokenChan)

	for token := range l.tokenChan {
		// 字典数据同步一次成功即可
		if err = l.query(token); err == nil {
			break
		} else {
			continue
		}
	}
	return
}

func (l *DictionaryQueryLogic) query(param *jobs.QueryParam) (err error) {
	data := statements.DictionaryRequest{Language: "zh_CN", TargetingList: vars.TargetingDictionaryKeys}
	var response statements.DictionaryResponse
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.Dictionary).Get().JsonData(data)
	if err != nil {
		return err
	}
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		return err
	}
	if response.Code != "200" {
		return errors.New("接口返回错误：" + response.Message)
	}
	var dict []*model.TargetingDictionaries
	dict = append(dict, dataCopy("age", response.Data.LinearTargetingMap.Age)...)
	dict = append(dict, dataCopy("device_price", response.Data.LinearTargetingMap.DevicePrice)...)
	dict = append(dict, dataCopy("not_pre_define_audience", response.Data.LinearTargetingMap.NotPreDefineAudience)...)
	dict = append(dict, dataCopy("gender", response.Data.LinearTargetingMap.Gender)...)
	dict = append(dict, dataCopy("pre_define_audience", response.Data.LinearTargetingMap.PreDefineAudience)...)
	dict = append(dict, dataCopy("media_app_category", response.Data.LinearTargetingMap.MediaAppCategory)...)
	dict = append(dict, dataCopy("network_type", response.Data.LinearTargetingMap.NetworkType)...)
	dict = append(dict, dataCopy("app_interest", response.Data.LinearTargetingMap.AppInterest)...)
	dict = append(dict, dataCopy("language", response.Data.LinearTargetingMap.Language)...)
	dict = append(dict, dataCopy("carrier", response.Data.TreeTargetingMap.Carrier)...)
	dict = append(dict, dataCopy("app_category", response.Data.TreeTargetingMap.AppCategory)...)
	dict = append(dict, dataCopy("series_type", response.Data.TreeTargetingMap.SeriesType)...)
	dict = append(dict, dataCopy("location_category", response.Data.TreeTargetingMap.LocationCategory)...)

	return l.svcCtx.DictModel.BatchInsert(l.ctx, dict)
}

func dataCopy(key string, items []*statements.DictionaryItem) (dict []*model.TargetingDictionaries) {
	dict = make([]*model.TargetingDictionaries, 0)
	var dataStruct int64 = 0
	for ds, keys := range vars.TargetingDictionaryStruct {
		if utils.StringInArray(keys, key) {
			dataStruct = ds
			break
		}
	}
	if dataStruct == 0 {
		log.Println("调度异常，字典 Key 未设置：", key)
		return
	}
	for _, item := range items {
		dict = append(dict, &model.TargetingDictionaries{
			DictKey:    key,
			Id:         item.Id,
			Pid:        item.Pid,
			Label:      item.Label,
			Value:      item.Value,
			Code:       item.Code,
			Seq:        item.Seq,
			DataStruct: dataStruct,
		})
	}
	return
}
