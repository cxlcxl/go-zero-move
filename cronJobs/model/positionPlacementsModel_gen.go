// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type (
	defaultPositionPlacementsModel struct {
		conn  sqlx.Session
		table string
	}

	PositionPlacements struct {
		Id                         int64  `db:"id"`
		CreativeSizeId             string `db:"creative_size_id"`              // 版位ID
		PlacementSizeId            string `db:"placement_size_id"`             // 规格ID
		CreativeSize               string `db:"creative_size"`                 // 尺寸
		CreativeSizeSubType        string `db:"creative_size_sub_type"`        // 版位子形式
		IsSupportMultipleCreatives string `db:"is_support_multiple_creatives"` // 是否支持多创意
	}
)

func newPositionPlacementsModel(conn sqlx.Session) *defaultPositionPlacementsModel {
	return &defaultPositionPlacementsModel{
		conn:  conn,
		table: "`position_placements`",
	}
}

func (m *defaultPositionPlacementsModel) tableName() string {
	return m.table
}

func (m *defaultPositionPlacementsModel) BatchInsert(ctx context.Context, placements []*PositionPlacements, deleteWhere string, args []interface{}) (err error) {
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in %s", m.table, deleteWhere)
	if _, err = m.conn.ExecCtx(ctx, deleteSql, args...); err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (creative_size_id, placement_size_id, creative_size, creative_size_sub_type, is_support_multiple_creatives) values ", m.table)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	for i := 0; i < len(placements); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?)")
		values = append(values,
			placements[i].CreativeSizeId,
			placements[i].PlacementSizeId,
			placements[i].CreativeSize,
			placements[i].CreativeSizeSubType,
			placements[i].IsSupportMultipleCreatives,
		)
		// 达到了 500 条数据，或最后一条了
		if chunk == 500 || i == len(placements)-1 {
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
