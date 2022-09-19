package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ TargetingsModel = (*customTargetingsModel)(nil)

type (
	// TargetingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTargetingsModel.
	TargetingsModel interface {
		targetingsModel
		BatchInsert(ctx context.Context, targetings []*Targetings) (err error)
	}

	customTargetingsModel struct {
		*defaultTargetingsModel
	}
)

// NewTargetingsModel returns a model for the database table.
func NewTargetingsModel(conn sqlx.SqlConn) TargetingsModel {
	return &customTargetingsModel{
		defaultTargetingsModel: newTargetingsModel(conn),
	}
}

func (m *defaultTargetingsModel) BatchInsert(ctx context.Context, targetings []*Targetings) (err error) {
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.table, targetingsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	targetingIds := make([]int64, 0)
	for _, targeting := range targetings {
		targetingIds = append(targetingIds, targeting.TargetingId)
	}
	in, args := whereIntIn(targetingIds)
	err = m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		deleteQuery := fmt.Sprintf("delete from %s where targeting_id in "+in, m.table)
		if _, err = session.ExecCtx(ctx, deleteQuery, args...); err != nil {
			return err
		}

		for i := 0; i < len(targetings); i++ {
			data := targetings[i]
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values,
				data.AccountId, data.AdvertiserId, data.TargetingId, data.TargetingName, data.TargetingType,
				data.LocationType, data.IncludeLocation, data.ExcludeLocation, data.Carriers, data.Language,
				data.Age, data.Gender, data.AppCategory, data.AppCategories, data.InstalledApps, data.AppInterest,
				data.AppInterests, data.Series, data.NetworkType, data.NotAudiences, data.Audiences,
				data.AppCategoryOfMedia, data.CreatedAt, data.UpdatedAt,
			)
			// 达到了 500 条数据，或最后一条了
			if chunk == 500 || i == len(targetings)-1 {
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
		return nil
	})
	return err
}

func whereIntIn(ids []int64) (w string, args []interface{}) {
	query := make([]string, 0)
	for i := range ids {
		query = append(query, "?")
		args = append(args, ids[i])
	}

	return "(" + strings.Join(query, ",") + ")", args
}
