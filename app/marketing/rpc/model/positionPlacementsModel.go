package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionPlacementsModel = (*customPositionPlacementsModel)(nil)

type (
	// PositionPlacementsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionPlacementsModel.
	PositionPlacementsModel interface {
		positionPlacementsModel
		PlacementList(ctx context.Context, creativeSizeIds []string) (placements []*PositionPlacements, err error)
		PlacementsByCreativeSizeId(ctx context.Context, creativeSizeId string) (placements []*PositionPlacement, err error)
	}

	customPositionPlacementsModel struct {
		*defaultPositionPlacementsModel
	}

	PositionPlacement struct {
		CreativeSize        string `db:"creative_size"`          // 尺寸
		CreativeSizeSubType string `db:"creative_size_sub_type"` // 版位子形式
	}
)

// NewPositionPlacementsModel returns a model for the database table.
func NewPositionPlacementsModel(conn sqlx.SqlConn) PositionPlacementsModel {
	return &customPositionPlacementsModel{
		defaultPositionPlacementsModel: newPositionPlacementsModel(conn),
	}
}

func (m *defaultPositionPlacementsModel) PlacementList(ctx context.Context, creativeSizeIds []string) (placements []*PositionPlacements, err error) {
	in, args, _ := utils.WhereIn(creativeSizeIds)
	sql := fmt.Sprintf("select %s from %s where creative_size_id in %s", positionPlacementsRows, m.table, in)
	err = m.conn.QueryRowsCtx(ctx, &placements, sql, args...)
	if len(placements) == 0 {
		return nil, sqlc.ErrNotFound
	}
	return
}

func (m *defaultPositionPlacementsModel) PlacementsByCreativeSizeId(ctx context.Context, creativeSizeId string) (placements []*PositionPlacement, err error) {
	sql, args, err := squirrel.Select("creative_size,creative_size_sub_type").
		From(m.table).Where("creative_size_id = ?", creativeSizeId).
		GroupBy("creative_size_sub_type,creative_size").
		OrderBy("creative_size_sub_type asc,creative_size asc").ToSql()
	if err != nil {
		return nil, err
	}
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
