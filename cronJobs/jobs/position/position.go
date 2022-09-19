package position

import (
	"business/cronJobs/jobs/position/logic"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
	"time"
)

func Position() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= Position job start ==================")

	ctx := context.Background()
	if err := logic.NewPositionLogic(ctx, vars.SvcCtx).PositionQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModulePosition, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= Position job end ==================")
	fmt.Println()
	fmt.Println()
}
