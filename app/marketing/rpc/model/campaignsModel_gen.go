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
	campaignsFieldNames          = builder.RawFieldNames(&Campaigns{})
	campaignsRows                = strings.Join(campaignsFieldNames, ",")
	campaignsRowsExpectAutoSet   = strings.Join(stringx.Remove(campaignsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	campaignsRowsWithPlaceHolder = strings.Join(stringx.Remove(campaignsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	campaignsModel interface {
		Insert(ctx context.Context, data *Campaigns) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Campaigns, error)
		FindOneByCampaignId(ctx context.Context, campaignId string) (*Campaigns, error)
		Update(ctx context.Context, newData *Campaigns) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCampaignsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Campaigns struct {
		Id                        int64     `db:"id"`
		CampaignId                string    `db:"campaign_id"`                  // 计划 ID
		CampaignName              string    `db:"campaign_name"`                // 计划名称
		AccountId                 int64     `db:"account_id"`                   // 账户 ID
		AdvertiserId              string    `db:"advertiser_id"`                // 广告主账户 ID
		OptStatus                 string    `db:"opt_status"`                   // 操作状态
		CampaignDailyBudgetStatus string    `db:"campaign_daily_budget_status"` // 计划日预算状态
		ProductType               string    `db:"product_type"`                 // 推广产品类型
		ShowStatus                string    `db:"show_status"`                  // 计划状态
		UserBalanceStatus         string    `db:"user_balance_status"`          // 账户余额状态
		FlowResource              string    `db:"flow_resource"`                // 投放网络
		SyncFlowResource          string    `db:"sync_flow_resource"`           // 同时同步投放搜索广告网络
		CampaignType              string    `db:"campaign_type"`                // 计划类型
		TodayDailyBudget          int64     `db:"today_daily_budget"`           // 当日计划日限额
		TomorrowDailyBudget       int64     `db:"tomorrow_daily_budget"`        // 次日计划日限额，不返回表示与当日计划日限额相同
		MarketingGoal             string    `db:"marketing_goal"`               // 营销目标
		IsCallback                int64     `db:"is_callback"`                  // 是否通过查询计划任务回调完整信息
		CreatedAt                 time.Time `db:"created_at"`                   // 添加时间
		UpdatedAt                 time.Time `db:"updated_at"`                   // 最后一次修改时间
	}
)

func newCampaignsModel(conn sqlx.SqlConn) *defaultCampaignsModel {
	return &defaultCampaignsModel{
		conn:  conn,
		table: "`campaigns`",
	}
}

func (m *defaultCampaignsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCampaignsModel) FindOne(ctx context.Context, id int64) (*Campaigns, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", campaignsRows, m.table)
	var resp Campaigns
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

func (m *defaultCampaignsModel) FindOneByCampaignId(ctx context.Context, campaignId string) (*Campaigns, error) {
	var resp Campaigns
	query := fmt.Sprintf("select %s from %s where `campaign_id` = ? limit 1", campaignsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, campaignId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCampaignsModel) Insert(ctx context.Context, data *Campaigns) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, campaignsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CampaignId, data.CampaignName, data.AccountId, data.AdvertiserId, data.OptStatus, data.CampaignDailyBudgetStatus, data.ProductType, data.ShowStatus, data.UserBalanceStatus, data.FlowResource, data.SyncFlowResource, data.CampaignType, data.TodayDailyBudget, data.TomorrowDailyBudget, data.MarketingGoal, data.IsCallback, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultCampaignsModel) Update(ctx context.Context, newData *Campaigns) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, campaignsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.CampaignId, newData.CampaignName, newData.AccountId, newData.AdvertiserId, newData.OptStatus, newData.CampaignDailyBudgetStatus, newData.ProductType, newData.ShowStatus, newData.UserBalanceStatus, newData.FlowResource, newData.SyncFlowResource, newData.CampaignType, newData.TodayDailyBudget, newData.TomorrowDailyBudget, newData.MarketingGoal, newData.IsCallback, newData.CreatedAt, newData.UpdatedAt, newData.Id)
	return err
}

func (m *defaultCampaignsModel) tableName() string {
	return m.table
}
