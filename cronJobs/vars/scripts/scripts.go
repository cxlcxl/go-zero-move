package scripts

import (
	"business/cronJobs/jobs/campaign"
	"business/cronJobs/jobs/country"
	"business/cronJobs/jobs/dictionary"
	"business/cronJobs/jobs/position"
	"business/cronJobs/jobs/refreshToken"
	"business/cronJobs/jobs/targeting"
	"business/cronJobs/vars"
)

type Job func()

var (
	ScheduleJobs = map[string]Job{
		vars.ApiModuleCountry:      country.Country,
		vars.ApiModuleCampaign:     campaign.Campaign,
		vars.ApiModuleDictionary:   dictionary.Dictionary,
		vars.ApiModuleRefreshToken: refreshToken.RefreshToken,
		vars.ApiModuleTargeting:    targeting.Targeting,
		vars.ApiModulePosition:     position.Position,
	}
)
