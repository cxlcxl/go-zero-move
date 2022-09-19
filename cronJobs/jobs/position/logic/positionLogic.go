package logic

import (
	"business/common/curl"
	"business/common/vars"
	"business/cronJobs/jobs"
	"business/cronJobs/model"
	"business/cronJobs/svc"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type PositionLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	tokenChan chan *jobs.QueryParam
}

func NewPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PositionLogic {
	return &PositionLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		tokenChan: make(chan *jobs.QueryParam),
	}
}

func (l *PositionLogic) PositionQuery() (err error) {
	go jobs.GetTokens(l.ctx, l.svcCtx, l.tokenChan)

	for token := range l.tokenChan {
		account, err := l.svcCtx.AccountModel.FindOne(l.ctx, token.AccountId) // 只有几条数据
		if err != nil {
			fmt.Println("定向包调度失败：", err)
		}

		for category := range vars.CreativeCategory {
			l.query(token, account.AdvertiserId, category)
		}
	}

	return
}

func (l *PositionLogic) query(param *jobs.QueryParam, advertiserId, category string) {
	data := statements.PositionRequest{
		AdvertiserId: advertiserId,
		Filtering:    statements.PositionFiltering{Category: category},
	}
	c, err := curl.New(l.svcCtx.Config.MarketingApis.Tools.Position).Get().JsonData(data)
	if err != nil {
		fmt.Println("请求生成失败：", err)
		return
	}
	var response statements.PositionResponse
	if err = c.Request(&response, curl.Authorization(param.AccessToken)); err != nil {
		fmt.Println("接口请求失败：", err)
		return
	}
	if response.Code != "200" {
		fmt.Println("接口返回错误：", response.Message)
		return
	}

	positions, samples, placements := formatPosition(response.Data.CreativeSizeInfoList, param.AccountId, advertiserId, category)
	if err = l.svcCtx.PositionModel.BatchInsert(l.ctx, positions, samples, placements); err != nil {
		fmt.Println("写入数据错误：", err)
	}

	return
}

func formatPosition(src []*statements.CreativeSizeInfo, accountId int64, advertiserId, category string) ([]*model.Positions, []*model.PositionSamples, []*model.PositionPlacements) {
	var (
		positions          = make([]*model.Positions, len(src))
		positionSamples    = make([]*model.PositionSamples, 0)
		positionPlacements = make([]*model.PositionPlacements, 0)
	)
	for i, info := range src {
		positions[i] = &model.Positions{
			AccountId:                  accountId,
			AdvertiserId:               advertiserId,
			Category:                   category,
			CreativeSizeId:             info.CreativeSizeId,
			CreativeSizeNameDsp:        info.CreativeSizeBaseInfo.CreativeSizeNameDsp,
			CreativeSizeDescription:    info.CreativeSizeBaseInfo.CreativeSizeDescription,
			SupportProductType:         info.CreativeSizeOperationInfo.SupportProductType,
			SupportObjectiveType:       info.CreativeSizeOperationInfo.SupportObjectiveType,
			IsSupportTimePeriod:        info.CreativeSizeOperationInfo.IsSupportTimePeriod,
			IsSupportMultipleCreatives: info.CreativeSizeOperationInfo.IsSupportMultipleCreatives,
			SupportPriceType:           info.CreativeSizePriceInfo.SupportPriceType,
			LastPullTime:               time.Now(),
		}
		for _, sample := range info.CreativeSizeBaseInfo.CreativeSizeSampleList {
			positionSamples = append(positionSamples, &model.PositionSamples{
				CreativeSizeId:     info.CreativeSizeId,
				CreativeSizeSample: sample.CreativeSizeSample,
				PreviewTitle:       sample.PreviewTitle,
			})
		}
		for _, placement := range info.CreativeSizeBaseInfo.CreativeSizePlacementList {
			positionPlacements = append(positionPlacements, &model.PositionPlacements{
				CreativeSizeId:             info.CreativeSizeId,
				PlacementSizeId:            placement.PlacementSizeId,
				CreativeSize:               placement.CreativeSize,
				CreativeSizeSubType:        placement.CreativeSizeSubType,
				IsSupportMultipleCreatives: placement.IsSupportMultipleCreatives,
			})
		}
	}
	return positions, positionSamples, positionPlacements
}
