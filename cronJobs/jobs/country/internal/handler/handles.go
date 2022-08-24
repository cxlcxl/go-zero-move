package handler

import (
	"business/cronJobs/jobs/country/internal/logic"
	"business/cronJobs/jobs/country/internal/svc"
	"context"
	"log"
)

func StartHandlers(svcCtx *svc.ServiceContext) {
	ctx := context.Background()
	// 获取 token 列表
	if err := logic.NewCountryQueryLogic(ctx, svcCtx).CountryQuery(); err != nil {
		log.Fatal(err)
		return
	}
}
