package vars

const (
	OptStatusEnable  = "OPERATION_STATUS_ENABLE"
	OptStatusDisable = "OPERATION_STATUS_DISABLE"
	OptStatusDelete  = "OPERATION_STATUS_DELETE"
)

var (
	// OptStatus 计划/任务/创意操作状态
	OptStatus = map[string]string{
		OptStatusEnable:  "投放中",
		OptStatusDisable: "已暂停",
		OptStatusDelete:  "已删除",
	}
)
