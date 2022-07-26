// Code generated by goctl. DO NOT EDIT!

package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	targetingDictionariesFieldNames = builder.RawFieldNames(&TargetingDictionaries{})
	targetingDictionariesRows       = strings.Join(targetingDictionariesFieldNames, ",")
)

type (
	defaultTargetingDictionariesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TargetingDictionaries struct {
		DictKey    string `db:"dict_key"` // 属于什么字典
		Id         string `db:"id"`       // 字典的元素 ID
		Pid        string `db:"pid"`      // 父节点元素ID
		Label      string `db:"label"`    // 显示的内容
		Value      string `db:"value"`    // 元素的值
		Code       string `db:"code"`
		Seq        string `db:"seq"`
		DataStruct int64  `db:"data_struct"`
	}
)

func newTargetingDictionariesModel(conn sqlx.SqlConn) *defaultTargetingDictionariesModel {
	return &defaultTargetingDictionariesModel{
		conn:  conn,
		table: "`targeting_dictionaries`",
	}
}

func (m *defaultTargetingDictionariesModel) tableName() string {
	return m.table
}
