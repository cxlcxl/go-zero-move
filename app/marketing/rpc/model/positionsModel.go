package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionsModel = (*customPositionsModel)(nil)

type (
	// PositionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionsModel.
	PositionsModel interface {
		positionsModel
		PositionList(ctx context.Context, category, productType string, accountId int64) ([]*Positions, error)
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

func (m *defaultPositionsModel) PositionList(ctx context.Context, category, productType string, accountId int64) (positions []*Positions, err error) {
	sql := fmt.Sprintf("select %s from %s where category = ? and account_id = ? and FIND_IN_SET(?, support_product_type)", positionsRows, m.table)
	err = m.conn.QueryRowsCtx(ctx, &positions, sql, category, accountId, productType)
	if len(positions) == 0 {
		return nil, sqlc.ErrNotFound
	}
	return
}
