package dictionary

import (
	"business/cronJobs/jobs/dictionary/logic"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
	"time"
)

func Dictionary() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Dictionary job start ==================")

	ctx := context.Background()
	job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, vars.ApiModuleDictionary)
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
		if err = logic.NewDictionaryQueryLogic(ctx, vars.SvcCtx, d).DictionaryQuery(); err != nil {
			log.Fatal(err)
			return
		}
		jobDay = jobDay.AddDate(0, 0, 1)
		if err = vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleDictionary, jobDay.Format(vars.DateFormat)); err != nil {
			fmt.Println("数据库调度时间修改失败: ", err)
		}
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("================= Dictionary job end ==================")
}
