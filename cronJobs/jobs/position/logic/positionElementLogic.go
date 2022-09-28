package logic

import (
	"business/common/curl"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

type PositionElementLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	requestChan chan *requestElementChan
	workThread  int
	maxThread   int
	doneChan    chan struct{}
	elements    []*model.PositionElements
	elementChan chan *model.PositionElements
}

func NewPositionElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionElementLogic {
	return &PositionElementLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		workThread:  0,
		doneChan:    make(chan struct{}),
		maxThread:   10,
		elements:    make([]*model.PositionElements, 0),
		requestChan: make(chan *requestElementChan),
		elementChan: make(chan *model.PositionElements),
	}
}

func (l *PositionElementLogic) ElementQuery() (err error) {
	tokenMap, accountIds, err := getTokenMap(l.ctx, l.svcCtx)
	if err != nil {
		return err
	}
	positions, err := l.svcCtx.PositionModel.GetPositions(l.ctx, accountIds)
	if err != nil {
		return err
	}
	params := make([]*requestElementChan, 0)
	for _, position := range positions {
		if token, ok := tokenMap[position.AccountId]; ok { // token
			params = append(params, &requestElementChan{
				Params: statements.PositionElementRequest{
					AdvertiserId: position.AdvertiserId,
					Filtering:    statements.PositionElementFilter{CreativeSizeId: position.CreativeSizeId},
				},
				Token: token,
			})
		}
	}
	if len(params) == 0 {
		return
	}
	go func(params []*requestElementChan) {
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
		case elements, ok := <-l.elementChan:
			if ok {
				l.elements = append(l.elements, elements)
			}
		case param, ok := <-l.requestChan:
			if ok {
				l.workThread++
				go l.getElement(param)
			}
		}

		if l.workThread == 0 && finishWorker == len(params) {
			break
		}
	}

	if len(l.elements) > 0 {
		err = l.svcCtx.ElementModel.BatchInsert(l.ctx, l.elements)
	}
	return
}

func (l *PositionElementLogic) getElement(param *requestElementChan) {
	defer func() {
		l.doneChan <- struct{}{}
	}()

	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.PositionElement).Get().JsonData(param.Params)
	if err != nil {
		fmt.Println("参数生成失败：", err)
		return
	}
	var response statements.PositionElementResponse
	if err = c.Request(&response, curl.Authorization(param.Token)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口请求失败：", response.Code, response.Message)
		return
	}
	//_element := make(map[string]struct{})
	for n, list := range response.Data.ElementInfoList {
		//key := fmt.Sprintf("%d-%s", response.Data.CreativeSizeId, list.Subtype)
		//if _, ok := _element[key]; !ok {
		for _, info := range list.ElementInfoList {
			durations := make([]string, len(info.Duration))
			for i, elementDuration := range info.Duration {
				durations[i] = fmt.Sprintf("%d,%d", elementDuration.Min, elementDuration.Max)
			}
			l.elementChan <- &model.PositionElements{
				GroupNumber:     int64(n + 1),
				CreativeSizeId:  strconv.Itoa(response.Data.CreativeSizeId),
				SubType:         list.Subtype,
				ElementId:       strconv.Itoa(info.ElementId),
				ElementName:     info.ElementName,
				ElementTitle:    info.ElementTitle,
				ElementCaption:  info.ElementCaption,
				Width:           info.Width,
				Height:          info.Height,
				MinWidth:        info.MinWidth,
				MinHeight:       info.MinHeight,
				MinLength:       info.MinLength,
				MaxLength:       info.MaxLength,
				FileSizeKbLimit: info.FileSizeKbLimit,
				GifSizeKbLimit:  info.GifSizeKbLimit,
				FileFormat:      info.FileFormat,
				Pattern:         info.Pattern,
				Duration:        strings.Join(durations, "|"),
				MinOccurs:       info.MinOccurs,
				MaxOccurs:       info.MaxOccurs,
			}
		}
		//}
		//_element[key] = struct{}{}
	}

	return
}
