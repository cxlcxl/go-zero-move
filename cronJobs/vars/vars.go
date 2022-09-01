package vars

import (
	"business/cronJobs/jobs/country"
	"business/cronJobs/svc"
)

type Job func()

var (
	SvcCtx *svc.ServiceContext

	ScheduleJobs = map[string]Job{
		"Country": country.Country,
	}
)
