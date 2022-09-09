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
	if err := logic.NewDictionaryQueryLogic(ctx, vars.SvcCtx).DictionaryQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleDictionary, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}
	time.Sleep(time.Millisecond * 500)

	fmt.Println("================= Dictionary job end ==================")
}
