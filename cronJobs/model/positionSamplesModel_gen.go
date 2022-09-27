// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	defaultPositionSamplesModel struct {
		conn  sqlx.Session
		table string
	}

	PositionSamples struct {
		Id                 int64  `db:"id"`
		CreativeSizeId     string `db:"creative_size_id"`     // 版位ID
		CreativeSizeSample string `db:"creative_size_sample"` // 预览图地址
		PreviewTitle       string `db:"preview_title"`        // 预览图标题
	}
)

func newPositionSamplesModel(conn sqlx.Session) *defaultPositionSamplesModel {
	return &defaultPositionSamplesModel{
		conn:  conn,
		table: "`position_samples`",
	}
}

func (m *defaultPositionSamplesModel) tableName() string {
	return m.table
}

func (m *defaultPositionSamplesModel) BatchInsert(ctx context.Context, samples []*PositionSamples, deleteWhere string, args []interface{}) (err error) {
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in %s", m.table, deleteWhere)
	if _, err = m.conn.ExecCtx(ctx, deleteSql, args...); err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (creative_size_id, creative_size_sample, preview_title) values ", m.table)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	for i := 0; i < len(samples); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?)")
		values = append(values, samples[i].CreativeSizeId, samples[i].CreativeSizeSample, samples[i].PreviewTitle)
		// 达到了 500 条数据，或最后一条了
		if chunk == 500 || i == len(samples)-1 {
			// 写入库
			insertSQL := query + strings.Join(valueStatement, ",")
			if _, err = m.conn.ExecCtx(ctx, insertSQL, values...); err != nil {
				return err
			}
			// 重置
			values, valueStatement = make([]interface{}, 0), make([]string, 0)
			chunk = 0
		}
		chunk++
	}
	return nil
}
