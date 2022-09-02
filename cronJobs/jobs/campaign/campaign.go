package campaign

import (
	"business/cronJobs/jobs/campaign/logic"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
	"time"
)

func Campaign() {
	fmt.Println("================= campaign job start ==================")

	ctx := context.Background()
	job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, "Campaign")
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
		if err = logic.NewCampaignQueryLogic(ctx, vars.SvcCtx, d).CampaignQuery(); err != nil {
			log.Fatal(err)
			return
		}
		jobDay = jobDay.AddDate(0, 0, 1)

		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("================= campaign job end ==================")
}
