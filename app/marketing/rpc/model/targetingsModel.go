package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TargetingsModel = (*customTargetingsModel)(nil)

type (
	// TargetingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTargetingsModel.
	TargetingsModel interface {
		targetingsModel
		GetTargetings(ctx context.Context) (targetings []*Targetings, err error)
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
func (m *defaultTargetingsModel) GetTargetings(ctx context.Context) (targetings []*Targetings, err error) {
	query := fmt.Sprintf("select * from %s order by id desc", m.table)
	err = m.conn.QueryRowsCtx(ctx, &targetings, query)
	switch err {
	case nil:
		return targetings, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
