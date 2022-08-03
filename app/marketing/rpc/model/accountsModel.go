package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AccountsModel = (*customAccountsModel)(nil)

type (
	// AccountsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountsModel.
	AccountsModel interface {
		accountsModel
		AccountList(ctx context.Context, accountType, state int64, accountName string, offset, limit uint64) (ls []*Accounts, total int64, err error)
		FindOneByClientId(ctx context.Context, clientId int64) (*Accounts, error)
		SetIsAuth(ctx context.Context, clientId int64) error
	}

	customAccountsModel struct {
		*defaultAccountsModel
	}
)

// NewAccountsModel returns a model for the database table.
func NewAccountsModel(conn sqlx.SqlConn, c cache.CacheConf) AccountsModel {
	return &customAccountsModel{
		defaultAccountsModel: newAccountsModel(conn, c),
	}
}

func (m *defaultAccountsModel) AccountList(ctx context.Context, accountType, state int64, accountName string, offset, limit uint64) (ls []*Accounts, total int64, err error) {
	total, err = m.actListCount(ctx, accountType, state, accountName)
	if total == 0 || err != nil {
		return
	}
	sql := squirrel.Select(accountsRows).From(m.table).Where("state = ?", state)
	if len(accountName) > 0 {
		sql = sql.Where("account_name like ?", "%"+accountName+"%")
	}
	if accountType > 0 {
		sql = sql.Where("account_type = ?", accountType)
	}
	query, args, err := sql.Offset(offset).Limit(limit).ToSql()
	if err != nil {
		return
	}
	err = m.QueryRowsNoCacheCtx(ctx, &ls, query, args...)
	return
}

func (m *defaultAccountsModel) actListCount(ctx context.Context, accountType, state int64, accountName string) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table).Where("state = ?", state)
	if len(accountName) > 0 {
		query = query.Where("account_name like ?", "%"+accountName+"%")
	}
	if accountType > 0 {
		query = query.Where("account_type = ?", accountType)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, sql, args...); err != nil {
		return 0, err
	}
	return
}

func (m *defaultAccountsModel) FindOneByClientId(ctx context.Context, clientId int64) (*Accounts, error) {
	accountsClientIdKey := fmt.Sprintf("%s%v", cacheAccountsClientIdPrefix, clientId)
	var resp Accounts
	err := m.QueryRowIndexCtx(ctx, &resp, accountsClientIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `client_id` = ? limit 1", accountsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, clientId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAccountsModel) SetIsAuth(ctx context.Context, clientId int64) error {
	query := fmt.Sprintf("update %s set `is_auth` = 1 where `client_id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, clientId)
	return err
}
