package country

import (
	"business/cronJobs/jobs/country/logic"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
)

func Country() {
	fmt.Println("country job start...")

	ctx := context.Background()
	job, err := vars.SvcCtx.JobModel.FindOneByApiModule(ctx, "Country")
	if err != nil {
		log.Fatal("调度模块查询错误：", err)
		return
	}
	// 获取 token 列表
	if err = logic.NewCountryQueryLogic(ctx, vars.SvcCtx, job).CountryQuery(); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("country job end...")
}
