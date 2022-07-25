package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		FindExistsUsersByRoleId(ctx context.Context, roleId int64) (exists bool, err error)
		UserList(ctx context.Context, username, email string, state int64, offset, limit uint64) (users []*Users, total int64, err error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

// export logic
func (m *defaultUsersModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUsersModel) FindExistsUsersByRoleId(ctx context.Context, roleId int64) (exists bool, err error) {
	var ct int64
	query := fmt.Sprintf("select count(1) as ct from %s where `role_id` = ? limit 1", m.table)
	if err = m.QueryRowNoCacheCtx(ctx, &ct, query, roleId); err != nil {
		return false, err
	}
	return ct > 0, nil
}

func (m *defaultUsersModel) UserList(ctx context.Context, username, email string, state int64, offset, limit uint64) (users []*Users, total int64, err error) {
	total, err = m.userListCount(ctx, username, email, state)
	if total == 0 || err != nil {
		return
	}
	query := squirrel.Select(usersRows).From(m.table).Where("state = ?", state)
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	if len(email) > 0 {
		query = query.Where("email like ?", "%"+email+"%")
	}
	sql, args, err := query.Offset(offset).Limit(limit).ToSql()
	err = m.QueryRowsNoCacheCtx(ctx, &users, sql, args...)
	return
}

func (m *defaultUsersModel) userListCount(ctx context.Context, username, email string, state int64) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table).Where("state = ?", state)
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	if len(email) > 0 {
		query = query.Where("email like ?", "%"+email+"%")
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
