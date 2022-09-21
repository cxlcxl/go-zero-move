package logic

import (
	"business/app/marketing/rpc/marketingcenter"
	"business/common/utils"
	"business/common/vars"
	"context"
	"errors"
	"strings"

	"business/app/marketing/api/internal/svc"
	"business/app/marketing/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TargetingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTargetingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TargetingListLogic {
	return &TargetingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TargetingListLogic) TargetingList(req *types.TargetingListReq) (resp *types.TargetingListResp, err error) {
	info, err := l.svcCtx.MarketingRpcClient.GetCampaignInfo(l.ctx, &marketingcenter.CampaignInfoReq{CampaignId: req.CampaignId})
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	list, err := l.svcCtx.MarketingRpcClient.TargetingList(l.ctx, &marketingcenter.TargetingListReq{AccountId: info.AccountId})
	if err != nil {
		return nil, utils.RpcError(err, "没有本地定向包，需要同步")
	}
	rs, err := l.unFormatRpcData(list.Targetings)
	if err != nil {
		return nil, utils.RpcError(err, "")
	}
	return &types.TargetingListResp{
		BaseResp: types.BaseResp{
			Code: vars.ResponseCodeOk,
			Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
		},
		Data: rs,
	}, nil
}

func (l *TargetingListLogic) unFormatRpcData(src []*marketingcenter.Targeting) (targetings []*types.TargetingListItem, err error) {
	carriers := make([]string, 0)
	appInterests := make([]string, 0)
	mediaAppCategories := make([]string, 0)
	targetings = make([]*types.TargetingListItem, len(src))
	for i, targeting := range src {
		var (
			appCategories = make([]string, 0)
			include       = make([]string, 0)
			exclude       = make([]string, 0)
			audiences     = make([]string, 0)
			notAudiences  = make([]string, 0)
			series        = make([]string, 0)
			language      = make([]string, 0)
		)
		location := ""
		if targeting.IncludeLocation != "" || targeting.ExcludeLocation != "" {
			location = "1"
			include = targetingSplit(targeting.IncludeLocation)
			exclude = targetingSplit(targeting.ExcludeLocation)
		}
		audience := ""
		if targeting.Audiences != "" || targeting.NotAudiences != "" {
			audience = "1"
			audiences = targetingSplit(targeting.Audiences)
			notAudiences = targetingSplit(targeting.NotAudiences)
		}
		carrier := ""
		if targeting.Carriers != "" {
			carrier = "1"
			carriers = append(carriers, targetingSplit(targeting.Carriers)...)
		}
		seriesType := ""
		if targeting.Series != "" {
			seriesType = "1"
			series = targetingSplit(targeting.Series)
		}
		if targeting.AppCategory != "" {
			appCategories = targetingSplit(targeting.AppCategories)
		}
		if targeting.AppInterest != "" {
			appInterestsTmp := targetingSplit(targeting.AppInterests)
			appInterests = append(appInterests, appInterestsTmp...)
		}
		mediaAppCategory := ""
		if targeting.AppCategoryOfMedia != "" {
			mediaAppCategory = "1"
			mediaAppCategories = append(mediaAppCategories, targetingSplit(targeting.AppCategoryOfMedia)...)
		}
		languageCheck := ""
		if targeting.Language != "" {
			languageCheck = "1"
			language = targetingSplit(targeting.Language)
		}
		targetings[i] = &types.TargetingListItem{
			TargetingId: targeting.TargetingId,
			TargetingCreateReq: types.TargetingCreateReq{
				TargetingType:    targeting.TargetingType,
				TargetingName:    targeting.TargetingName,
				Gender:           targeting.Gender,
				Age:              strings.Split(targeting.Age, vars.TargetingDatabaseSeq),
				NetworkType:      strings.Split(targeting.NetworkType, vars.TargetingDatabaseSeq),
				Location:         location,
				LocationType:     targeting.LocationType,
				IncludeLocation:  include,
				ExcludeLocation:  exclude,
				Carrier:          carrier,
				AppCategory:      targeting.AppCategory,
				AppCategories:    appCategories,
				AppInterest:      targeting.AppInterest,
				Audience:         audience,
				Audiences:        audiences,
				NotAudience:      notAudiences,
				SeriesType:       seriesType,
				Series:           series,
				MediaAppCategory: mediaAppCategory,
				LanguageCheck:    languageCheck,
				Language:         language,
				InstalledApps:    targeting.InstalledApps,
			},
		}
	}
	// 运营商
	carrierMap := make(map[string]string)
	carrierParents := make(map[string]string)
	if len(carriers) > 0 {
		dictCarriers, err := l.svcCtx.MarketingRpcClient.DictQuery(l.ctx, &marketingcenter.DictionaryReq{DictKey: []string{"carrier"}, Value: carriers})
		if err != nil {
			return nil, utils.RpcError(err, "")
		}
		countries, err := l.svcCtx.MarketingRpcClient.GetCountries(l.ctx, &marketingcenter.EmptyParamsReq{})
		if err != nil {
			return nil, utils.RpcError(err, "")
		}
		for _, dict := range dictCarriers.Dictionaries {
			carrierMap[dict.Value] = dict.Pid
		}
		for _, country := range countries.Countries {
			carrierParents[country.CCode] = country.CId
		}
	}

	// 媒体类型
	mediaCategoryParents := make(map[string]string)
	if len(mediaAppCategories) > 0 {
		mediaCategories, err := l.svcCtx.MarketingRpcClient.DictQuery(l.ctx, &marketingcenter.DictionaryReq{DictKey: []string{"media_app_category"}, Value: mediaAppCategories})
		if err != nil {
			return nil, utils.RpcError(err, "")
		}
		for _, dictionary := range mediaCategories.Dictionaries {
			mediaCategoryParents[dictionary.Value] = dictionary.Pid
		}
	}

	// App 兴趣
	appInterestsParents := make(map[string]string)
	interestParentMap := make(map[string]string)
	if len(appInterests) > 0 {
		dictAppInterests, err := l.svcCtx.MarketingRpcClient.DictQuery(l.ctx, &marketingcenter.DictionaryReq{DictKey: []string{"app_interest"}, Value: appInterests})
		if err != nil {
			return nil, utils.RpcError(err, "")
		}
		for _, dictionary := range dictAppInterests.Dictionaries {
			appInterestsParents[dictionary.Value] = dictionary.Pid
		}
		for s, s2 := range vars.AppInterest {
			interestParentMap[s2] = s
		}
	}
	// 组合数据
	for i, targeting := range src {
		carriers = targetingSplit(targeting.Carriers)
		mediaAppCategories = targetingSplit(targeting.AppCategoryOfMedia)
		appInterests = targetingSplit(targeting.AppInterests)
		targetingCarriers := make([][]string, 0)
		targetingMediaCategories := make([][]string, 0)
		targetingAppInterests := make([][]string, 0)

		for _, carrier := range carriers {
			if cId, ok := carrierParents[carrierMap[carrier]]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：运营商")
			} else {
				targetingCarriers = append(targetingCarriers, []string{cId, carrier})
			}
		}
		for _, category := range mediaAppCategories {
			if pid, ok := mediaCategoryParents[category]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：媒体类型")
			} else {
				targetingMediaCategories = append(targetingMediaCategories, []string{pid, category})
			}
		}
		for _, appInterest := range appInterests {
			if pid, ok := appInterestsParents[appInterest]; !ok {
				return nil, errors.New("字典与定向包数据匹配异常：App 兴趣")
			} else {
				targetingAppInterests = append(targetingAppInterests, []string{interestParentMap[pid], appInterest})
			}
		}

		targetings[i].Carriers = targetingCarriers
		targetings[i].AppCategoryOfMedia = targetingMediaCategories
		targetings[i].AppInterests = targetingAppInterests
	}
	return
}

func targetingSplit(s string) []string {
	if s == "" {
		return []string{}
	} else {
		return strings.Split(s, vars.TargetingDatabaseSeq)
	}
}
