package logic

import (
	"business/app/marketing/api/internal/types"
	"business/common/utils"
	"business/common/vars"
	"strconv"
	"time"
)

// 创建一个批次码
func generateBatchCode(salt string) string {
	//node, _ := snowflake.NewNode(0)
	//node.Generate().Int64()
	var rs []string
	if salt != "" {
		rs = []string{salt, "-", time.Now().Format("060102150405"), strconv.FormatInt(time.Now().UnixNano(), 10)[10:]}
	} else {
		rs = []string{time.Now().Format("060102150405"), strconv.FormatInt(time.Now().UnixNano(), 10)[10:]}
	}
	return utils.BufferConcat(rs)
}

func successFul() (resp *types.BaseResp) {
	return &types.BaseResp{
		Code: vars.ResponseCodeOk,
		Msg:  vars.ResponseMsg[vars.ResponseCodeOk],
	}
}
