package refreshToken

import (
	"business/cronJobs/jobs"
	"business/cronJobs/vars"
	"context"
	"fmt"
	"log"
	"time"
)

func RefreshToken() {
	fmt.Println()
	fmt.Println()
	fmt.Println("================= RefreshToken job start ==================")

	ctx := context.Background()
	list, err := vars.SvcCtx.TokenModel.GetAccessTokenList(ctx)
	if err != nil {
		fmt.Println("刷新 Token 任务失败，查询 Token 数据失败：", err)
	}
	for _, tokens := range list {
		_, err = jobs.Refresh(ctx, vars.SvcCtx, tokens)
		if err != nil {
			log.Println("Token 刷新失败，账户 ID：", tokens.AccountId, err)
			continue
		}
	}

	if err = vars.SvcCtx.JobModel.UpdateJobDayByModule(ctx, vars.ApiModuleRefreshToken, time.Now().Format(vars.DateFormat)); err != nil {
		fmt.Println("数据库调度时间修改失败: ", err)
	}

	fmt.Println("================= RefreshToken job end ==================")
	fmt.Println()
	fmt.Println()
}

func RefreshTokenManual(_ string, _ int64) {
	RefreshToken()
}
