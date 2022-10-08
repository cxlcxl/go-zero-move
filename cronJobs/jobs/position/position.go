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

func PriceManual(_ time.Time, _ int64) {
	Price()
}
func Price() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= PositionPrice job start ==================")

	ctx := context.Background()
	if err := logic.NewPositionPriceLogic(ctx, vars.SvcCtx).PriceQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModulePositionPrice, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= PositionPrice job end ==================")
	fmt.Println()
	fmt.Println()
}

func ElementManual(_ time.Time, _ int64) {
	Element()
}

// Element 版位元素
func Element() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= PositionElement job start ==================")

	ctx := context.Background()
	if err := logic.NewPositionElementLogic(ctx, vars.SvcCtx).ElementQuery(); err != nil {
		log.Fatal(err)
		return
	}

	if err := vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModulePositionElement, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= PositionElement job end ==================")
	fmt.Println()
	fmt.Println()
}
