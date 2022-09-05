package country

import (
	"business/cronJobs/jobs/country/logic"
	"business/cronJobs/vars"
	"business/cronJobs/vars/statements"
	"context"
	"fmt"
	"log"
	"time"
)

func Country() {
	fmt.Println("================= country job start ==================")

	ctx := context.Background()
	job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, statements.ApiModuleCountry)
	if err != nil {
		log.Fatal("调度模块查询错误：", err)
		return
	}
	jobDay := job.StatDay
	now := time.Now()
	for {
		pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
		if jobDay.After(pauseDay) {
			break
		}
		d := jobDay.Format(vars.DateFormat)
		fmt.Println("schedule day: ", d)
		if err = logic.NewCountryQueryLogic(ctx, vars.SvcCtx, d).CountryQuery(); err != nil {
			log.Fatal(err)
			return
		}
		jobDay = jobDay.AddDate(0, 0, 1)

		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("================= country job end ==================")
}
