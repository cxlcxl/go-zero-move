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
	fmt.Println()
	fmt.Println()
	fmt.Println("================= campaign job start ==================")

	defer func() {
		fmt.Println("================= campaign job end ==================")
		fmt.Println()
		fmt.Println()
	}()
	ctx := context.Background()
	job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, vars.ApiModuleCampaign)
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
		if err = vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleCampaign, jobDay.Format(vars.DateFormat)); err != nil {
			fmt.Println("数据库调度时间修改失败: ", err)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func CampaignManual(day time.Time, pauseRule int64) {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= campaign job start ==================")
	defer func() {
		fmt.Println("================= campaign job end ==================")
		fmt.Println()
		fmt.Println()
	}()
	ctx := context.Background()
	now := time.Now()
	if pauseRule == 0 {
		if err := campaignDoSchedule(ctx, day.Format(vars.DateFormat), day); err != nil {
			fmt.Println("调度失败：", err)
		}
	} else if pauseRule == 99 {
		job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, vars.ApiModuleCampaign)
		if err != nil {
			log.Fatal("调度模块查询错误：", err)
			return
		}
		if job.PauseRule == -1 {
			fmt.Println("原规则此模块已停止调度：", err)
			return
		}
		pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
		for {
			if day.After(pauseDay) {
				break
			}
			d := day.Format(vars.DateFormat)
			if err = logic.NewCampaignQueryLogic(ctx, vars.SvcCtx, d).CampaignQuery(); err != nil {
				log.Fatal(err)
				return
			}
			day = day.AddDate(0, 0, 1)
			fmt.Println(d, day.Format(vars.DateFormat))
			if err = vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleCampaign, day.Format(vars.DateFormat)); err != nil {
				fmt.Println("数据库调度时间修改失败: ", err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	} else if pauseRule >= 1 && pauseRule <= 31 {
		pauseDay := now.AddDate(0, 0, -1*int(pauseRule))
		for {
			if day.After(pauseDay) {
				break
			}
			d := day.Format(vars.DateFormat)
			if err := logic.NewCampaignQueryLogic(ctx, vars.SvcCtx, d).CampaignQuery(); err != nil {
				log.Fatal(err)
				return
			}
			day = day.AddDate(0, 0, 1)
			fmt.Println(d, day.Format(vars.DateFormat))
			if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleCampaign, day.Format(vars.DateFormat)); err != nil {
				fmt.Println("数据库调度时间修改失败: ", err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	} else {
		log.Fatal("规则有误，可使用 -h 查看")
	}
}

func campaignDoSchedule(ctx context.Context, d string, jobDay time.Time) (err error) {
	if err = logic.NewCampaignQueryLogic(ctx, vars.SvcCtx, d).CampaignQuery(); err != nil {
		log.Fatal(err)
		return
	}
	jobDay = jobDay.AddDate(0, 0, 1)
	if err = vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleCampaign, jobDay.Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}
	return err
}
