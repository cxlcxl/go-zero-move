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

var (
	moduleName string
	day        string
	pauseRule  int
	_cron      *cron.Cron
)

type schedule struct {
	versions map[string]int64
	entries  map[string]cron.EntryID
}

func init() {
	flag.IntVar(&pauseRule, "pause", 0, "* 调度停止规则：0 只调度一天；99 使用配置的规则，[1-31] 调度到当天之前的天数；default：0")
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
		parse, err := time.Parse("2006-01-02", day)
		if err != nil {
			log.Fatal(fmt.Sprintf("日期格式错误：%s, 正确格式: %s", day, "2022-10-01"))
			return
		}
		fmt.Println(fmt.Sprintf("手动调度参数，模块：%s，日期：%s，规则：%d", moduleName, parse.Format("2006-01-02"), pauseRule))
		job(parse, int64(pauseRule))
	} else {
		listener := make(chan string)
		go startListener(listener)

		s := &schedule{
			versions: make(map[string]int64),
			entries:  make(map[string]cron.EntryID),
		}
		_cron = cron.New()
		firstStart := true
		for v := range listener {
			fmt.Println("刷新调度：", v)

			s.taskScheduling()

			if firstStart {
				fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>> CRON START <<<<<<<<<<<<<<<<<<<<<<<<<<")
				_cron.Start()
			}
			firstStart = false
		}

		defer _cron.Stop()
		select {} // 阻塞进程
	}
}

func (s *schedule) taskScheduling() {
	// cron 表达式与 linux 的 crontab 一致：*  *  *  *  *
	//                           分别对应：分 时 日 月 周
	_jobs, err := vars.SvcCtx.JobModel.GetJobs(context.Background())
	if err != nil {
		fmt.Println("调度启动失败", err)
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

		version, ok := s.versions[job.ApiModule]
		if !ok {
			if fn, _ok := scripts.ScheduleJobs[job.ApiModule]; _ok {
				fmt.Println("添加：", job.ApiModule, job.Version)
				s.versions[job.ApiModule] = job.Version
				entryId, _ := _cron.AddFunc(job.JobSchedule, fn)
				s.entries[job.ApiModule] = entryId
			}
		} else {
			if version == job.Version {
				continue
			}
			s.versions[job.ApiModule] = job.Version

			if fn, _ok := scripts.ScheduleJobs[job.ApiModule]; _ok {
				if entryId, exists := s.entries[job.ApiModule]; exists {
					_cron.Remove(entryId)
				}
				entryId, _ := _cron.AddFunc(job.JobSchedule, fn)
				s.entries[job.ApiModule] = entryId
				fmt.Println("替换：", job.ApiModule, job.Version)
			}
		}
	}
}

func startListener(tc chan string) {
	for {
		tc <- time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(time.Second * 60)
	}
}
