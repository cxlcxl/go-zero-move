package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TargetingsModel = (*customTargetingsModel)(nil)

type (
	// TargetingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTargetingsModel.
	TargetingsModel interface {
		targetingsModel
		GetTargetings(ctx context.Context, accountId int64) (targetings []*Targetings, err error)
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
func (m *defaultTargetingsModel) GetTargetings(ctx context.Context, accountId int64) (targetings []*Targetings, err error) {
	query := squirrel.Select("*").From(m.table).OrderBy("id desc")
	if accountId > 0 {
		query = query.Where("account_id = ?", accountId)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.conn.QueryRowsCtx(ctx, &targetings, sql, args...)
	switch err {
	case nil:
		return targetings, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
