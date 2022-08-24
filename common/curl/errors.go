package curl

import "errors"

var (
	NotSetMethod = errors.New("未设置请求 Method")
)
