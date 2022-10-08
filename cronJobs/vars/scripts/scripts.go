package scripts

import (
	"business/cronJobs/jobs/campaign"
	"business/cronJobs/jobs/country"
	"business/cronJobs/jobs/dictionary"
	"business/cronJobs/jobs/position"
	"business/cronJobs/jobs/refreshToken"
	"business/cronJobs/jobs/targeting"
	"business/cronJobs/vars"
	"time"
)

type Job func()
type ManualJob func(day time.Time, pauseRule int64)

var (
	ScheduleJobs = map[string]Job{
		vars.ApiModuleCountry:         country.Country,
		vars.ApiModuleCampaign:        campaign.Campaign,
		vars.ApiModuleDictionary:      dictionary.Dictionary,
		vars.ApiModuleRefreshToken:    refreshToken.RefreshToken,
		vars.ApiModuleTargeting:       targeting.Targeting,
		vars.ApiModulePosition:        position.Position,
		vars.ApiModulePositionPrice:   position.Price,
		vars.ApiModulePositionElement: position.Element,
	}

	ManualScheduleJobs = map[string]ManualJob{
		vars.ApiModuleCampaign:        campaign.CampaignManual,
		vars.ApiModuleRefreshToken:    refreshToken.RefreshTokenManual,
		vars.ApiModulePositionPrice:   position.PriceManual,
		vars.ApiModulePositionElement: position.ElementManual,
	}
)
