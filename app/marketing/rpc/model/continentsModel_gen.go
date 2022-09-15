// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	continentsFieldNames          = builder.RawFieldNames(&Continents{})
	continentsRows                = strings.Join(continentsFieldNames, ",")
	continentsRowsExpectAutoSet   = strings.Join(stringx.Remove(continentsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	continentsRowsWithPlaceHolder = strings.Join(stringx.Remove(continentsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	continentsModel interface {
		Insert(ctx context.Context, data *Continents) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Continents, error)
		Update(ctx context.Context, newData *Continents) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContinentsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Continents struct {
		Id    int64  `db:"id"`
		CName string `db:"c_name"`
	}
)

func newContinentsModel(conn sqlx.SqlConn) *defaultContinentsModel {
	return &defaultContinentsModel{
		conn:  conn,
		table: "`continents`",
	}
}

func (m *defaultContinentsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultContinentsModel) FindOne(ctx context.Context, id int64) (*Continents, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", continentsRows, m.table)
	var resp Continents
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContinentsModel) Insert(ctx context.Context, data *Continents) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, continentsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CName)
	return ret, err
}

func (m *defaultContinentsModel) Update(ctx context.Context, data *Continents) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, continentsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CName, data.Id)
	return err
}

func (m *defaultContinentsModel) tableName() string {
	return m.table
}
