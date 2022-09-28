package logic

import (
	"business/common/curl"
	"business/common/vars"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type PositionPriceLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	maxThread   int
	workThread  int
	doneChan    chan struct{}
	requestChan chan *requestPriceChan
	priceChan   chan *model.PositionBasePrices
	basePrices  []*model.PositionBasePrices
}

func NewPositionPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionPriceLogic {
	return &PositionPriceLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		maxThread:   10,
		workThread:  0,
		doneChan:    make(chan struct{}),
		priceChan:   make(chan *model.PositionBasePrices),
		basePrices:  make([]*model.PositionBasePrices, 0),
		requestChan: make(chan *requestPriceChan),
	}
}

func (l *PositionPriceLogic) PriceQuery() (err error) {
	tokenMap, accountIds, err := getTokenMap(l.ctx, l.svcCtx)
	if err != nil {
		return err
	}
	positions, err := l.svcCtx.PositionModel.GetPositions(l.ctx, accountIds)
	if err != nil {
		return err
	}

	tmpPriceTypes := make(map[string]string)
	for s, s2 := range vars.Pricing {
		tmpPriceTypes[s2] = s
	}
	params := make([]*requestPriceChan, 0)
	for _, position := range positions {
		if token, ok := tokenMap[position.AccountId]; ok { // token
			priceTypes := strings.Split(position.SupportPriceType, ",")
			for _, priceType := range priceTypes {
				if t, _ok := tmpPriceTypes[priceType]; _ok {
					params = append(params, &requestPriceChan{
						Params: statements.CreativeSizePriceRequest{
							AdvertiserId: position.AdvertiserId,
							Filtering:    statements.CreativeSizePriceFilter{CreativeSizeId: position.CreativeSizeId, PriceType: t},
						},
						Token: token,
					})
				}
			}
		}
	}

	if len(params) == 0 {
		return
	}

	go func(params []*requestPriceChan) {
		i := 0
		for {
			if l.workThread < l.maxThread {
				l.requestChan <- params[i]
				i++
			}
			if i >= len(params) {
				break
			}
		}
	}(params)

	finishWorker := 0
	for {
		select {
		case <-l.doneChan:
			l.workThread--
			finishWorker++
		case basePrice, ok := <-l.priceChan:
			if ok {
				l.basePrices = append(l.basePrices, basePrice)
			}
		case param, ok := <-l.requestChan:
			if ok {
				l.workThread++
				go l.getBasePrice(param)
			}
		}

		if l.workThread == 0 && finishWorker == len(params) {
			break
		}
	}

	if len(l.basePrices) > 0 {
		err = l.svcCtx.PriceModel.BatchInsert(l.ctx, l.basePrices)
	}
	return
}

func (l *PositionPriceLogic) getBasePrice(param *requestPriceChan) {
	defer func() {
		l.doneChan <- struct{}{}
	}()
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.PositionPrice).Get().JsonData(param.Params)
	if err != nil {
		fmt.Println("参数生成失败：", err)
		return
	}
	var response statements.CreativeSizePriceResponse
	if err = c.Request(&response, curl.Authorization(param.Token)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败：", response.Code, response.Message)
		return
	}
	l.priceChan <- &model.PositionBasePrices{
		CreativeSizeId: param.Params.Filtering.CreativeSizeId,
		PriceType:      param.Params.Filtering.PriceType,
		BasePrice:      response.Data.FloorPrice,
	}
	return
}
