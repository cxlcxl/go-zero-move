package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdsRequestLogsModel = (*customAdsRequestLogsModel)(nil)

type (
	// AdsRequestLogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdsRequestLogsModel.
	AdsRequestLogsModel interface {
		adsRequestLogsModel
		FindLogByStatDayApiModule(ctx context.Context, statDay, apiModule string, accountId int64) (*AdsRequestLogs, error)
	}

	customAdsRequestLogsModel struct {
		*defaultAdsRequestLogsModel
	}
)

// NewAdsRequestLogsModel returns a model for the database table.
func NewAdsRequestLogsModel(conn sqlx.SqlConn) AdsRequestLogsModel {
	return &customAdsRequestLogsModel{
		defaultAdsRequestLogsModel: newAdsRequestLogsModel(conn),
	}
}

func (m *defaultAdsRequestLogsModel) FindLogByStatDayApiModule(ctx context.Context, statDay, apiModule string, accountId int64) (*AdsRequestLogs, error) {
	var resp AdsRequestLogs
	query := fmt.Sprintf("select %s from %s where `stat_day` = ? and `api_module` = ? and `account_id` = ? limit 1", adsRequestLogsRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, statDay, apiModule, accountId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
