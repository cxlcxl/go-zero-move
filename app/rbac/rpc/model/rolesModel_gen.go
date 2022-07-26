// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rolesFieldNames          = builder.RawFieldNames(&Roles{})
	rolesRows                = strings.Join(rolesFieldNames, ",")
	rolesRowsExpectAutoSet   = strings.Join(stringx.Remove(rolesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	rolesRowsWithPlaceHolder = strings.Join(stringx.Remove(rolesFieldNames, "`id`", "`create_time`", "`update_time`", "`created_at`", "`update_at`"), "=?,") + "=?"

	cacheRolesIdPrefix = "cache:roles:id:"
)

type (
	rolesModel interface {
		Insert(ctx context.Context, data *Roles) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Roles, error)
		Update(ctx context.Context, newData *Roles) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolesModel struct {
		sqlc.CachedConn
		table string
	}

	Roles struct {
		Id        int64     `db:"id"`
		RoleName  string    `db:"role_name"`
		State     int64     `db:"state"`      // 1正常0停用
		Sys       int64     `db:"sys"`        // 角色所属系统
		CreatedAt time.Time `db:"created_at"` // 添加时间
		UpdatedAt time.Time `db:"updated_at"` // 最后一次修改时间
	}
)

func newRolesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRolesModel {
	return &defaultRolesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`roles`",
	}
}

func (m *defaultRolesModel) Delete(ctx context.Context, id int64) error {
	rolesIdKey := fmt.Sprintf("%s%v", cacheRolesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, rolesIdKey)
	return err
}

func (m *defaultRolesModel) FindOne(ctx context.Context, id int64) (*Roles, error) {
	rolesIdKey := fmt.Sprintf("%s%v", cacheRolesIdPrefix, id)
	var resp Roles
	err := m.QueryRowCtx(ctx, &resp, rolesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRolesModel) Insert(ctx context.Context, data *Roles) (sql.Result, error) {
	rolesIdKey := fmt.Sprintf("%s%v", cacheRolesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, rolesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RoleName, data.State, data.Sys, data.CreatedAt, data.UpdatedAt)
	}, rolesIdKey)
	return ret, err
}

func (m *defaultRolesModel) Update(ctx context.Context, data *Roles) error {
	rolesIdKey := fmt.Sprintf("%s%v", cacheRolesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rolesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RoleName, data.State, data.Sys, data.UpdatedAt, data.Id)
	}, rolesIdKey)
	return err
}

func (m *defaultRolesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRolesIdPrefix, primary)
}

func (m *defaultRolesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRolesModel) tableName() string {
	return m.table
}
