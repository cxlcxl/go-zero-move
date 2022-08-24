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
	reportCountriesFieldNames          = builder.RawFieldNames(&ReportCountries{})
	reportCountriesRows                = strings.Join(reportCountriesFieldNames, ",")
	reportCountriesRowsExpectAutoSet   = strings.Join(stringx.Remove(reportCountriesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	reportCountriesRowsWithPlaceHolder = strings.Join(stringx.Remove(reportCountriesFieldNames, "`id`", "`create_time`", "`update_time`", "`created_at`", "`update_at`"), "=?,") + "=?"
)

type (
	reportCountriesModel interface {
		Insert(ctx context.Context, data *ReportCountries) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ReportCountries, error)
		Update(ctx context.Context, newData *ReportCountries) error
		Delete(ctx context.Context, id int64) error
	}

	defaultReportCountriesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ReportCountries struct {
		Id                    int64     `db:"id"`
		StatDate              time.Time `db:"stat_date"`               // 日期
		StatHour              string    `db:"stat_hour"`               // 时间
		AdvertiserId          string    `db:"advertiser_id"`           // 广告主账户ID
		AccountId             int64     `db:"account_id"`              // 账户 ID
		CampaignId            string    `db:"campaign_id"`             // 计划ID
		CampaignName          string    `db:"campaign_name"`           // 计划名称
		AdgroupId             string    `db:"adgroup_id"`              // 任务 ID
		AdgroupName           string    `db:"adgroup_name"`            // 任务名称
		CreativeId            string    `db:"creative_id"`             // 创意 ID
		CreativeName          string    `db:"creative_name"`           // 创意名称
		Country               string    `db:"country"`                 // 国家地域编码
		AppId                 string    `db:"app_id"`                  // 应用 ID
		AppName               string    `db:"app_name"`                // 应用名
		ShowCount             int64     `db:"show_count"`              // 曝光量
		ClickCount            int64     `db:"click_count"`             // 点击量
		Cpc                   float64   `db:"cpc"`                     // 点击均价
		Cpm                   float64   `db:"cpm"`                     // 千次展示均价 thousand_show_cost
		Cost                  float64   `db:"cost"`                    // 花费
		InstallCount          int64     `db:"install_count"`           // 安装量
		Cpi                   float64   `db:"cpi"`                     // 安装成本 install_cost
		DownloadCount         int64     `db:"download_count"`          // 下载量
		Cpd                   float64   `db:"cpd"`                     // 下载成本 download_cost
		BrowseCount           int64     `db:"browse_count"`            // 浏览量
		BrowseCost            float64   `db:"browse_cost"`             // 浏览成本
		AppCustomCount        int64     `db:"app_custom_count"`        // 应用自定义转化量
		AppCustomCost         float64   `db:"app_custom_cost"`         // 应用自定义转化成本
		WebCustomCount        int64     `db:"web_custom_count"`        // 网页自定义转化量
		WebCustomCost         float64   `db:"web_custom_cost"`         // 网页自定义转化成本
		PlayCount             int64     `db:"play_count"`              // 视频播放次数
		PlayOverCount         int64     `db:"play_over_count"`         // 视频播放完成次数
		ThreeDayRetainCount   int64     `db:"three_day_retain_count"`  // 3日留存
		ThreeDayRetainCost    float64   `db:"three_day_retain_cost"`   // 3日留存成本
		ShareCount            int64     `db:"share_count"`             // 分享
		ShareCost             float64   `db:"share_cost"`              // 分享成本
		SevenDayRetainCount   int64     `db:"seven_day_retain_count"`  // 7日留存
		SevenDayRetainCost    float64   `db:"seven_day_retain_cost"`   // 7日留存成本
		ActiveCountNormalized int64     `db:"active_count_normalized"` // 归一化激活
		Cpa                   float64   `db:"cpa"`                     // 归一化激活成本 active_cost_normalized
		RetainCountNormalized int64     `db:"retain_count_normalized"` // 归一化次留
		RetainCostNormalized  float64   `db:"retain_cost_normalized"`  // 归一化次留成本
		PayCountNormalized    int64     `db:"pay_count_normalized"`    // 归一化付费
		PayCostNormalized     float64   `db:"pay_cost_normalized"`     // 归一化付费成本
		CreatedAt             time.Time `db:"created_at"`              // 添加时间
		UpdatedAt             time.Time `db:"updated_at"`              // 最后一次修改时间
	}
)

func newReportCountriesModel(conn sqlx.SqlConn) *defaultReportCountriesModel {
	return &defaultReportCountriesModel{
		conn:  conn,
		table: "`report_countries`",
	}
}

func (m *defaultReportCountriesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultReportCountriesModel) FindOne(ctx context.Context, id int64) (*ReportCountries, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", reportCountriesRows, m.table)
	var resp ReportCountries
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

func (m *defaultReportCountriesModel) Insert(ctx context.Context, data *ReportCountries) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, reportCountriesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.StatDate, data.StatHour, data.AdvertiserId, data.AccountId, data.CampaignId, data.CampaignName, data.AdgroupId, data.AdgroupName, data.CreativeId, data.CreativeName, data.Country, data.AppId, data.AppName, data.ShowCount, data.ClickCount, data.Cpc, data.Cpm, data.Cost, data.InstallCount, data.Cpi, data.DownloadCount, data.Cpd, data.BrowseCount, data.BrowseCost, data.AppCustomCount, data.AppCustomCost, data.WebCustomCount, data.WebCustomCost, data.PlayCount, data.PlayOverCount, data.ThreeDayRetainCount, data.ThreeDayRetainCost, data.ShareCount, data.ShareCost, data.SevenDayRetainCount, data.SevenDayRetainCost, data.ActiveCountNormalized, data.Cpa, data.RetainCountNormalized, data.RetainCostNormalized, data.PayCountNormalized, data.PayCostNormalized, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultReportCountriesModel) Update(ctx context.Context, data *ReportCountries) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, reportCountriesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.StatDate, data.StatHour, data.AdvertiserId, data.AccountId, data.CampaignId, data.CampaignName, data.AdgroupId, data.AdgroupName, data.CreativeId, data.CreativeName, data.Country, data.AppId, data.AppName, data.ShowCount, data.ClickCount, data.Cpc, data.Cpm, data.Cost, data.InstallCount, data.Cpi, data.DownloadCount, data.Cpd, data.BrowseCount, data.BrowseCost, data.AppCustomCount, data.AppCustomCost, data.WebCustomCount, data.WebCustomCost, data.PlayCount, data.PlayOverCount, data.ThreeDayRetainCount, data.ThreeDayRetainCost, data.ShareCount, data.ShareCost, data.SevenDayRetainCount, data.SevenDayRetainCost, data.ActiveCountNormalized, data.Cpa, data.RetainCountNormalized, data.RetainCostNormalized, data.PayCountNormalized, data.PayCostNormalized, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultReportCountriesModel) tableName() string {
	return m.table
}
