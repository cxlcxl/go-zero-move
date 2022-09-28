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

var (
	moduleName string
	day        string
	pauseRule  int
)

func init() {
	flag.IntVar(&pauseRule, "pause", 0, "* 调度停止规则：0 只调度一天；99 一直调度到当天，[1-98] 调度到当天之前的天数；default：0")
	flag.StringVar(&moduleName, "module", "", "* 调度的模块；default：''")
	flag.StringVar(&day, "date", time.Now().Format("2006-01-02"), "要调度的日期")
}

func main() {
	flag.Parse()

	var c config.Config
	if err := config.Unmarshal("etc/jobs."+vars.Env+".yaml", &c); err != nil {
		log.Fatal(err)
		return
	}
	vars.SvcCtx = svc.NewServiceContext(c)

	if len(moduleName) > 0 {
		job, ok := scripts.ManualScheduleJobs[moduleName]
		if !ok {
			log.Fatal("调度模块不存在：", moduleName)
			return
		}
		fmt.Println(fmt.Sprintf("手动调度参数，模块：%s，日期：%s，规则：%d", moduleName, day, pauseRule))
		job(day, int64(pauseRule))
	} else {
		(Schedule{}).TaskScheduling()
	}
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
