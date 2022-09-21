// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	positionsFieldNames          = builder.RawFieldNames(&Positions{})
	positionsRows                = strings.Join(positionsFieldNames, ",")
	positionsRowsExpectAutoSet   = strings.Join(stringx.Remove(positionsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	positionsRowsWithPlaceHolder = strings.Join(stringx.Remove(positionsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	positionsModel interface {
		Insert(ctx context.Context, data *Positions) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Positions, error)
		FindOneByCreativeSizeId(ctx context.Context, creativeSizeId int64) (*Positions, error)
		Update(ctx context.Context, newData *Positions) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPositionsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Positions struct {
		Id                         int64     `db:"id"`
		AccountId                  int64     `db:"account_id"`
		AdvertiserId               string    `db:"advertiser_id"`             // 广告主账户ID
		CreativeSizeId             int64     `db:"creative_size_id"`          // 版位ID
		CreativeSizeNameDsp        string    `db:"creative_size_name_dsp"`    // 版位名称
		CreativeSizeDescription    string    `db:"creative_size_description"` // 版位描述
		Category                   string    `db:"category"`                  // 版位所属分类
		SupportProductType         string    `db:"support_product_type"`      // 支持的推广产品
		SupportObjectiveType       string    `db:"support_objective_type"`
		IsSupportTimePeriod        string    `db:"is_support_time_period"`        // 是否支持选择投放时段
		IsSupportMultipleCreatives string    `db:"is_support_multiple_creatives"` // 是否支持多创意
		SupportPriceType           string    `db:"support_price_type"`            // 付费方式
		LastPullTime               time.Time `db:"last_pull_time"`                // 最后拉取时间
	}
)

func newPositionsModel(conn sqlx.SqlConn) *defaultPositionsModel {
	return &defaultPositionsModel{
		conn:  conn,
		table: "`positions`",
	}
}

func (m *defaultPositionsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPositionsModel) FindOne(ctx context.Context, id int64) (*Positions, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", positionsRows, m.table)
	var resp Positions
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

func (m *defaultPositionsModel) FindOneByCreativeSizeId(ctx context.Context, creativeSizeId int64) (*Positions, error) {
	var resp Positions
	query := fmt.Sprintf("select %s from %s where `creative_size_id` = ? limit 1", positionsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, creativeSizeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPositionsModel) Insert(ctx context.Context, data *Positions) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, positionsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AccountId, data.AdvertiserId, data.CreativeSizeId, data.CreativeSizeNameDsp, data.CreativeSizeDescription, data.Category, data.SupportProductType, data.SupportObjectiveType, data.IsSupportTimePeriod, data.IsSupportMultipleCreatives, data.SupportPriceType, data.LastPullTime)
	return ret, err
}

func (m *defaultPositionsModel) Update(ctx context.Context, newData *Positions) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, positionsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.AccountId, newData.AdvertiserId, newData.CreativeSizeId, newData.CreativeSizeNameDsp, newData.CreativeSizeDescription, newData.Category, newData.SupportProductType, newData.SupportObjectiveType, newData.IsSupportTimePeriod, newData.IsSupportMultipleCreatives, newData.SupportPriceType, newData.LastPullTime, newData.Id)
	return err
}

func (m *defaultPositionsModel) tableName() string {
	return m.table
}
