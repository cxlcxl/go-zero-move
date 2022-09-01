package jobs

import (
	"business/cronJobs/jobs/country"
)

type Job func()

var (
	ScheduleJobs = map[string]Job{
		"Country": country.Country,
	}
)
