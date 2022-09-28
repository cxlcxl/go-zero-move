package model

import (
	"business/common/utils"
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
		GetPositions(ctx context.Context, accountIds []int64) (positions []*Positions, err error)
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
	in, args, _ := utils.WhereIn(creativeSizeIds)
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

func (m *defaultPositionsModel) GetPositions(ctx context.Context, accountIds []int64) (positions []*Positions, err error) {
	if len(accountIds) == 0 {
		return nil, err
	}
	fields := []string{
		"id",
		"account_id",
		"advertiser_id",
		"creative_size_id",
		"'' as creative_size_name_dsp",
		"'' as creative_size_description",
		"'' as category",
		"'' as support_product_type",
		"'' as support_objective_type",
		"'' as is_support_time_period",
		"'' as is_support_multiple_creatives",
		"support_price_type",
		"last_pull_time",
	}
	in, args, _ := utils.WhereIn(accountIds)
	sql := fmt.Sprintf("select %s from %s where account_id in %s", strings.Join(fields, ","), m.table, in)
	err = m.conn.QueryRowsCtx(ctx, &positions, sql, args...)
	return
}
