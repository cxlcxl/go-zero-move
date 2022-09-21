package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionPlacementsModel = (*customPositionPlacementsModel)(nil)

type (
	// PositionPlacementsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionPlacementsModel.
	PositionPlacementsModel interface {
		positionPlacementsModel
		PlacementList(ctx context.Context, creativeSizeIds []int64) (placements []*PositionPlacements, err error)
	}

	customPositionPlacementsModel struct {
		*defaultPositionPlacementsModel
	}
)

// NewPositionPlacementsModel returns a model for the database table.
func NewPositionPlacementsModel(conn sqlx.SqlConn) PositionPlacementsModel {
	return &customPositionPlacementsModel{
		defaultPositionPlacementsModel: newPositionPlacementsModel(conn),
	}
}

func (m *defaultPositionPlacementsModel) PlacementList(ctx context.Context, creativeSizeIds []int64) (placements []*PositionPlacements, err error) {
	in, args, _ := utils.WhereIn(creativeSizeIds)
	sql := fmt.Sprintf("select %s from %s where creative_size_id in %s", positionPlacementsRows, m.table, in)
	err = m.conn.QueryRowsCtx(ctx, &placements, sql, args...)
	switch err {
	case nil:
		return placements, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
