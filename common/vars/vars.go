package vars

const (
	MaxPageSize    uint64 = 100
	SysCachePrefix string = "business:cache:"

	ApiModuleCountry  = "Country"
	ApiModuleCampaign = "Campaign"
)

const (
	ResponseCodeOk = iota
	ResponseCodeError
	ResponseCodeOvertime
	ResponseCodeDatabaseErr
	ResponseCodeValidFailed
	ResponseCodeUnauthorized
	ResponseCodeEmptyToken
	ResponseCodeTokenErr
	ResponseCodeTokenExpire
)

const (
	_ = iota
	AccountTypeAds
	AccountTypeReport
)

const (
	_ = iota
	AppChannelGallery
	AppChannelGooglePlay
	AppChannelAppStore
	AppChannelOther
)

const (
	CommonStateVoid = iota
	CommonStateValid
)

const (
	JobPauseRuleStop  = -1
	JobPauseRuleToday = iota
	JobPauseRuleADayAgo
	JobPauseRuleTwoDayAgo
	JobPauseRuleThreeDayAgo
	JobPauseRuleFourDayAgo
	JobPauseRuleFiveDayAgo
	JobPauseRuleAWeekAgo
)

var (
	ResponseMsg = map[int]string{
		ResponseCodeOk:           "OK",
		ResponseCodeError:        "请求失败",
		ResponseCodeOvertime:     "请求超时",
		ResponseCodeDatabaseErr:  "数据库查询失败",
		ResponseCodeValidFailed:  "数据验证失败",
		ResponseCodeUnauthorized: "Unauthorized:权限不足",
		ResponseCodeEmptyToken:   "缺少 TOKEN",
		ResponseCodeTokenErr:     "TOKEN 错误",
		ResponseCodeTokenExpire:  "TOKEN 过期",
	}
	// AccountType 账号类型
	AccountType = map[int]string{
		AccountTypeAds:    "ADS",
		AccountTypeReport: "变现",
	}
	// CommonState 通用数据库状态字段
	CommonState = map[int]string{
		CommonStateVoid:  "停用",
		CommonStateValid: "正常",
	}
	// AppChannel 系统平台(渠道)
	AppChannel = map[int]string{
		AppChannelGallery:    "AppGallery",
		AppChannelGooglePlay: "GooglePlay",
		AppChannelAppStore:   "AppStore",
		AppChannelOther:      "Other",
	}
	// CronModules 调度模块 - 与 cronJobs 模块保持一致
	CronModules = map[string]string{
		ApiModuleCountry:  "国家/地区数据",
		ApiModuleCampaign: "计划数据",
	}
	JobPauseRule = map[int]string{
		JobPauseRuleStop:        "已停止",
		JobPauseRuleToday:       "到当天",
		JobPauseRuleADayAgo:     "一天前",
		JobPauseRuleTwoDayAgo:   "两天前",
		JobPauseRuleThreeDayAgo: "三天前",
		JobPauseRuleFourDayAgo:  "四天前",
		JobPauseRuleFiveDayAgo:  "五天前",
		JobPauseRuleAWeekAgo:    "一周前",
	}
)
