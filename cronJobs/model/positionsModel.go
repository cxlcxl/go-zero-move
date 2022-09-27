package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ PositionsModel = (*customPositionsModel)(nil)

type (
	// PositionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionsModel.
	PositionsModel interface {
		BatchInsert(ctx context.Context, positions []*Positions, samples []*PositionSamples, placements []*PositionPlacements) (err error)
	}

	customPositionsModel struct {
		*defaultPositionsModel
	}
)

// NewPositionsModel returns a model for the database table.
func NewPositionsModel(conn sqlx.SqlConn) PositionsModel {
	return &customPositionsModel{
		defaultPositionsModel: newPositionsModel(conn),
	}
}

func (m *defaultPositionsModel) BatchInsert(ctx context.Context, positions []*Positions, samples []*PositionSamples, placements []*PositionPlacements) (err error) {
	creativeSizeIds := make([]string, 0)
	for _, position := range positions {
		creativeSizeIds = append(creativeSizeIds, position.CreativeSizeId)
	}
	in, args := whereIn(creativeSizeIds)
	deleteSql := fmt.Sprintf("delete from %s where creative_size_id in %s", m.table, in)
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, positionsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = session.ExecCtx(ctx, deleteSql, args...); err != nil {
			return err
		}
		for i := 0; i < len(positions); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				positions[i].AccountId, positions[i].AdvertiserId, positions[i].CreativeSizeId, positions[i].CreativeSizeNameDsp,
				positions[i].CreativeSizeDescription, positions[i].Category, positions[i].SupportProductType,
				positions[i].SupportObjectiveType, positions[i].IsSupportTimePeriod, positions[i].IsSupportMultipleCreatives,
				positions[i].SupportPriceType, positions[i].LastPullTime,
			)
			// 达到了 300 条数据，或最后一条了
			if chunk == 300 || i == len(positions)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",")
				if _, err = session.ExecCtx(ctx, insertSQL, values...); err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		if err = newPositionSamplesModel(session).BatchInsert(ctx, samples, in, args); err != nil {
			return err
		}
		if err = newPositionPlacementsModel(session).BatchInsert(ctx, placements, in, args); err != nil {
			return err
		}
		return nil
	})
}

// where in 条件组合
func whereIn(s []string) (w string, args []interface{}) {
	query := make([]string, 0)
	for i := range s {
		query = append(query, "?")
		args = append(args, s[i])
	}

	return "(" + strings.Join(query, ",") + ")", args
}
