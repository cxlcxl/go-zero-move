package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ CampaignsModel = (*customCampaignsModel)(nil)

type (
	// CampaignsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCampaignsModel.
	CampaignsModel interface {
		campaignsModel
		BatchInsert(ctx context.Context, campaigns []*Campaigns) error
	}

	customCampaignsModel struct {
		*defaultCampaignsModel
	}
)

// NewCampaignsModel returns a model for the database table.
func NewCampaignsModel(conn sqlx.SqlConn) CampaignsModel {
	return &customCampaignsModel{
		defaultCampaignsModel: newCampaignsModel(conn),
	}
}
func (m *defaultCampaignsModel) BatchInsert(ctx context.Context, campaigns []*Campaigns) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, campaignsRowsExpectAutoSet)
	updateFields := []string{
		"opt_status", "campaign_daily_budget_status", "show_status", "user_balance_status", "flow_resource",
		"today_daily_budget", "tomorrow_daily_budget", "marketing_goal", "is_callback",
	}
	fieldSql := make([]string, 0)
	for _, field := range updateFields {
		fieldSql = append(fieldSql, fmt.Sprintf("%s=values(%s)", field, field))
	}
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	chunk := 1
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		for i := 0; i < len(campaigns); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			data := campaigns[i]
			values = append(values,
				data.CampaignId, "", data.CampaignName, data.AccountId, data.AdvertiserId, data.OptStatus,
				data.CampaignDailyBudgetStatus, data.ProductType, data.ShowStatus, data.UserBalanceStatus, data.FlowResource,
				data.SyncFlowResource, data.CampaignType, data.TodayDailyBudget, data.TomorrowDailyBudget,
				data.MarketingGoal, 1, data.CreatedAt, data.UpdatedAt,
			)
			// 达到了 300 条数据，或最后一条了
			if chunk == 300 || i == len(campaigns)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",") + " on duplicate key update " + strings.Join(fieldSql, ",")
				if _, err = session.ExecCtx(ctx, insertSQL, values...); err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}
		return nil
	})
}
