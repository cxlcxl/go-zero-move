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
	positionSamplesFieldNames          = builder.RawFieldNames(&PositionSamples{})
	positionSamplesRows                = strings.Join(positionSamplesFieldNames, ",")
	positionSamplesRowsExpectAutoSet   = strings.Join(stringx.Remove(positionSamplesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	positionSamplesRowsWithPlaceHolder = strings.Join(stringx.Remove(positionSamplesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	positionSamplesModel interface {
		Insert(ctx context.Context, data *PositionSamples) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PositionSamples, error)
		Update(ctx context.Context, newData *PositionSamples) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPositionSamplesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PositionSamples struct {
		Id                 int64  `db:"id"`
		CreativeSizeId     int64  `db:"creative_size_id"`     // 版位ID
		CreativeSizeSample string `db:"creative_size_sample"` // 预览图地址
		PreviewTitle       string `db:"preview_title"`        // 预览图标题
	}
)

func newPositionSamplesModel(conn sqlx.SqlConn) *defaultPositionSamplesModel {
	return &defaultPositionSamplesModel{
		conn:  conn,
		table: "`position_samples`",
	}
}

func (m *defaultPositionSamplesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPositionSamplesModel) FindOne(ctx context.Context, id int64) (*PositionSamples, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", positionSamplesRows, m.table)
	var resp PositionSamples
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

func (m *defaultPositionSamplesModel) Insert(ctx context.Context, data *PositionSamples) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, positionSamplesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreativeSizeId, data.CreativeSizeSample, data.PreviewTitle)
	return ret, err
}

func (m *defaultPositionSamplesModel) Update(ctx context.Context, data *PositionSamples) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, positionSamplesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CreativeSizeId, data.CreativeSizeSample, data.PreviewTitle, data.Id)
	return err
}

func (m *defaultPositionSamplesModel) tableName() string {
	return m.table
}
