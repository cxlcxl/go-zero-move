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
	tokensFieldNames          = builder.RawFieldNames(&Tokens{})
	tokensRows                = strings.Join(tokensFieldNames, ",")
	tokensRowsExpectAutoSet   = strings.Join(stringx.Remove(tokensFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tokensRowsWithPlaceHolder = strings.Join(stringx.Remove(tokensFieldNames, "`id`", "`create_time`", "`update_time`", "`created_at`", "`update_at`"), "=?,") + "=?"

	cacheTokensIdPrefix       = "cache:tokens:id:"
	cacheTokensClientIdPrefix = "cache:tokens:clientId:"
)

type (
	tokensModel interface {
		Insert(ctx context.Context, data *Tokens) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Tokens, error)
		FindOneByClientId(ctx context.Context, clientId int64) (*Tokens, error)
		Update(ctx context.Context, newData *Tokens) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTokensModel struct {
		sqlc.CachedConn
		table string
	}

	Tokens struct {
		Id           int64     `db:"id"`
		ClientId     int64     `db:"client_id"` // 客户端ID
		AccessToken  string    `db:"access_token"`
		RefreshToken string    `db:"refresh_token"`
		ExpiredAt    time.Time `db:"expired_at"` // access_token 过期时间
		CreatedAt    time.Time `db:"created_at"` // 添加时间
		UpdatedAt    time.Time `db:"updated_at"` // 最后一次修改时间
		TokenType    string    `db:"token_type"`
	}
)

func newTokensModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTokensModel {
	return &defaultTokensModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tokens`",
	}
}

func (m *defaultTokensModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	tokensClientIdKey := fmt.Sprintf("%s%v", cacheTokensClientIdPrefix, data.ClientId)
	tokensIdKey := fmt.Sprintf("%s%v", cacheTokensIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tokensClientIdKey, tokensIdKey)
	return err
}

func (m *defaultTokensModel) FindOne(ctx context.Context, id int64) (*Tokens, error) {
	tokensIdKey := fmt.Sprintf("%s%v", cacheTokensIdPrefix, id)
	var resp Tokens
	err := m.QueryRowCtx(ctx, &resp, tokensIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tokensRows, m.table)
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

func (m *defaultTokensModel) FindOneByClientId(ctx context.Context, clientId int64) (*Tokens, error) {
	tokensClientIdKey := fmt.Sprintf("%s%v", cacheTokensClientIdPrefix, clientId)
	var resp Tokens
	err := m.QueryRowIndexCtx(ctx, &resp, tokensClientIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `client_id` = ? limit 1", tokensRows, m.table)
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

func (m *defaultTokensModel) Insert(ctx context.Context, data *Tokens) (sql.Result, error) {
	tokensClientIdKey := fmt.Sprintf("%s%v", cacheTokensClientIdPrefix, data.ClientId)
	tokensIdKey := fmt.Sprintf("%s%v", cacheTokensIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tokensRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ClientId, data.AccessToken, data.RefreshToken, data.ExpiredAt, data.CreatedAt, data.UpdatedAt, data.TokenType)
	}, tokensClientIdKey, tokensIdKey)
	return ret, err
}

func (m *defaultTokensModel) Update(ctx context.Context, newData *Tokens) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	tokensClientIdKey := fmt.Sprintf("%s%v", cacheTokensClientIdPrefix, data.ClientId)
	tokensIdKey := fmt.Sprintf("%s%v", cacheTokensIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tokensRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ClientId, newData.AccessToken, newData.RefreshToken, newData.ExpiredAt, newData.UpdatedAt, newData.TokenType, newData.Id)
	}, tokensClientIdKey, tokensIdKey)
	return err
}

func (m *defaultTokensModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTokensIdPrefix, primary)
}

func (m *defaultTokensModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tokensRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTokensModel) tableName() string {
	return m.table
}
