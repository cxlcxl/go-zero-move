package targeting

import (
	"business/cronJobs/jobs/targeting/logic"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
	"time"
)

func Targeting() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Targeting job start ==================")

	ctx := context.Background()
	if err := logic.NewTargetingLogic(ctx, vars.SvcCtx).TargetingQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleTargeting, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= Targeting job end ==================")
	fmt.Println()
	fmt.Println()
}
