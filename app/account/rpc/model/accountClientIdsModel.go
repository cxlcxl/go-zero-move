package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ AccountClientIdsModel = (*customAccountClientIdsModel)(nil)

type (
	// AccountClientIdsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountClientIdsModel.
	AccountClientIdsModel interface {
		accountClientIdsModel
		BatchInsert(ctx context.Context, session sqlx.Session, data []*AccountClientIds) error
		GetClientIdsByAccountIds(ctx context.Context, accountIds []string) ([]*AccountClientIds, error)
	}

	customAccountClientIdsModel struct {
		*defaultAccountClientIdsModel
	}
)

// NewAccountClientIdsModel returns a model for the database table.
func NewAccountClientIdsModel(conn sqlx.SqlConn, c cache.CacheConf) AccountClientIdsModel {
	return &customAccountClientIdsModel{
		defaultAccountClientIdsModel: newAccountClientIdsModel(conn, c),
	}
}

func (m *defaultAccountClientIdsModel) BatchInsert(ctx context.Context, session sqlx.Session, data []*AccountClientIds) (err error) {
	args := make([]interface{}, 0)
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, accountClientIdsRowsExpectAutoSet)
	valueColumns := make([]string, 0)
	for _, datum := range data {
		valueColumns = append(valueColumns, "(?, ?, ?, ?, ?, ?)")
		args = append(args, datum.AccountId, datum.ClientId, datum.Secret, datum.Comment, datum.CreatedAt, datum.UpdatedAt)
	}
	if session != nil {
		_, err = session.ExecCtx(ctx, query+strings.Join(valueColumns, ","), args...)
	} else {
		_, err = m.ExecNoCacheCtx(ctx, query+strings.Join(valueColumns, ","), args...)
	}
	return
}

func (m *defaultAccountClientIdsModel) GetClientIdsByAccountIds(ctx context.Context, accountIds []string) (clientIds []*AccountClientIds, err error) {
	query := squirrel.Select(accountClientIdsRows).From(m.table)
	if len(accountIds) > 0 {
		in, v := whereIn(accountIds)
		query = query.Where("account_id in "+in, v...)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCacheCtx(ctx, &clientIds, sql, args...)
	return
}

func whereIn(actIds []string) (w string, args []interface{}) {
	query := make([]string, 0)
	for i := range actIds {
		query = append(query, "?")
		args = append(args, actIds[i])
	}

	return "(" + strings.Join(query, ",") + ")", args
}
