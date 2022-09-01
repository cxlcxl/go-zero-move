package vars

import (
	"business/cronJobs/svc"
)

const (
	DateFormat = "2006-01-02"
)

var (
	SvcCtx *svc.ServiceContext
)
