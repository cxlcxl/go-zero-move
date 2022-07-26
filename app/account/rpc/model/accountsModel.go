package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AccountsModel = (*customAccountsModel)(nil)

type (
	// AccountsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAccountsModel.
	AccountsModel interface {
		accountsModel
		AccountList(ctx context.Context, accountType, state int64, accountName string, offset, limit uint64) (ls []*Accounts, total int64, err error)
		RemoteAccounts(ctx context.Context, accountName string, isParent int64) ([]*Accounts, error)
		SetIsAuth(ctx context.Context, accountId int64) error
		GetAccountsByIds(ctx context.Context, ids []int64) ([]*Accounts, error)
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
	sql := squirrel.Select(accountsRows).From(m.table).Where("state = ?", state).OrderBy("updated_at desc")
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

func (m *defaultAccountsModel) SetIsAuth(ctx context.Context, accountId int64) error {
	query := fmt.Sprintf("update %s set `is_auth` = 1 where `id` = ?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, accountId)
	return err
}

func (m *defaultAccountsModel) RemoteAccounts(ctx context.Context, accountName string, isParent int64) (accounts []*Accounts, err error) {
	query := squirrel.Select(accountsRows).From(m.table).Where("state = 1")
	if isParent == 1 {
		query = query.Where("parent_id = 0")
	}
	if accountName != "" {
		query = query.Where("account_name like ?", "%"+accountName+"%")
	} else {
		query = query.OrderBy("updated_at desc").Offset(0).Limit(20)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowsNoCacheCtx(ctx, &accounts, sql, args...); err != nil {
		return nil, err
	}
	return
}

func (m *defaultAccountsModel) GetAccountsByIds(ctx context.Context, ids []int64) (accounts []*Accounts, err error) {
	in, a, err := utils.WhereIn(ids)
	if err != nil {
		return nil, err
	}
	query := squirrel.Select(accountsRows).From(m.table).Where("id in "+in, a...)
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowsNoCacheCtx(ctx, &accounts, sql, args...); err != nil {
		return nil, err
	}
	return
}
