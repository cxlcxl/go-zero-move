package main

import (
	"business/cronJobs/config"
	"business/cronJobs/svc"
	"business/cronJobs/vars"
	"business/cronJobs/vars/scripts"
	"context"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// Schedule 任务调度
type Schedule struct{}

func main() {
	flag.Parse()

	var c config.Config
	if err := config.Unmarshal("etc/jobs."+vars.Env+".yaml", &c); err != nil {
		log.Fatal(err)
		return
	}
	vars.SvcCtx = svc.NewServiceContext(c)
	(Schedule{}).TaskScheduling()
}

func (s Schedule) TaskScheduling() {
	c := cron.New()
	// cron 表达式与 linux 的 crontab 一致：*  *  *  *  *
	//                           分别对应：分 时 日 月 周
	_jobs, err := vars.SvcCtx.JobModel.GetJobs(context.Background())
	if err != nil {
		log.Fatal("调度启动失败", err)
		return
	}
	now := time.Now()
	for _, job := range _jobs {
		if job.PauseRule == -1 {
			continue
		}
		pauseDay := now.AddDate(0, 0, -1*int(job.PauseRule))
		if job.StatDay.After(pauseDay) {
			continue
		}

		if fn, ok := scripts.ScheduleJobs[job.ApiModule]; ok {
			if _, err = c.AddFunc(job.JobSchedule, fn); err != nil {
				fmt.Println(job.JobSchedule, " 添加调度失败：", err)
			}
		}
	}

	c.Start()
	defer c.Stop()
	select {} // 阻塞进程
}
