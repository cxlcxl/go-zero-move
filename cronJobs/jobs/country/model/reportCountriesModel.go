package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ ReportCountriesModel = (*customReportCountriesModel)(nil)

type (
	// ReportCountriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReportCountriesModel.
	ReportCountriesModel interface {
		reportCountriesModel
		BatchInsertAndLog(ctx context.Context, data []*ReportCountries, adsLog *AdsRequestLogs, hasLog bool) error
	}

	customReportCountriesModel struct {
		*defaultReportCountriesModel
	}
)

// NewReportCountriesModel returns a model for the database table.
func NewReportCountriesModel(conn sqlx.SqlConn) ReportCountriesModel {
	return &customReportCountriesModel{
		defaultReportCountriesModel: newReportCountriesModel(conn),
	}
}

func (m *defaultReportCountriesModel) BatchInsertAndLog(ctx context.Context, data []*ReportCountries, adsLog *AdsRequestLogs, hasLog bool) (err error) {
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, reportCountriesRowsExpectAutoSet)
	values, valueStatement := make([]interface{}, 0), make([]string, 0)
	chunk := 1

	err = m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		for i := 0; i < len(data); i++ {
			valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			values = append(values, data[i].StatDate, data[i].StatHour, data[i].AdvertiserId, data[i].AccountId,
				data[i].CampaignId, data[i].CampaignName, data[i].AdgroupId, data[i].AdgroupName, data[i].CreativeId,
				data[i].CreativeName, data[i].Country, data[i].AppId, data[i].AppName, data[i].ShowCount,
				data[i].ClickCount, data[i].Cpc, data[i].Cpm, data[i].Cost, data[i].InstallCount, data[i].Cpi, data[i].DownloadCount,
				data[i].Cpd, data[i].BrowseCount, data[i].BrowseCost, data[i].AppCustomCount, data[i].AppCustomCost, data[i].WebCustomCount,
				data[i].WebCustomCost, data[i].PlayCount, data[i].PlayOverCount, data[i].ThreeDayRetainCount, data[i].ThreeDayRetainCost,
				data[i].ShareCount, data[i].ShareCost, data[i].SevenDayRetainCount, data[i].SevenDayRetainCost, data[i].ActiveCountNormalized,
				data[i].Cpa, data[i].RetainCountNormalized, data[i].RetainCostNormalized, data[i].PayCountNormalized, data[i].PayCostNormalized,
				data[i].CreatedAt, data[i].UpdatedAt,
			)
			// 达到了 500 条数据，或最后一条了
			if chunk == 500 || i == len(data)-1 {
				// 写入库
				insertSQL := query + strings.Join(valueStatement, ",")
				if _, err = session.ExecCtx(ctx, insertSQL, values...); err != nil {
					return err
				}
				// 重置
				values, valueStatement = make([]interface{}, 0), make([]string, 0)
				chunk = 0
			}
			chunk++
		}

		adsModel := newAdsRequestLogsModel(nil)
		if hasLog {
			query = fmt.Sprintf("update %s set %s where `id` = ?", adsModel.table, adsRequestLogsRowsWithPlaceHolder)
			_, err = session.ExecCtx(ctx, query, adsLog.StatDay, adsLog.ApiModule, adsLog.AccountId, adsLog.RequestJsonBody, adsLog.CurrentRequestPage, adsLog.NextRequestPage, adsLog.IsCompleted, adsLog.TotalPage, adsLog.TotalNum, adsLog.PageSize, adsLog.LastRequestTime, adsLog.Id)
		} else {
			query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", adsModel.table, adsRequestLogsRowsExpectAutoSet)
			_, err = session.ExecCtx(ctx, query, adsLog.StatDay, adsLog.ApiModule, adsLog.AccountId, adsLog.RequestJsonBody, adsLog.CurrentRequestPage, adsLog.NextRequestPage, adsLog.IsCompleted, adsLog.TotalPage, adsLog.TotalNum, adsLog.PageSize, adsLog.LastRequestTime)
		}
		return err
	})

	return err
}
