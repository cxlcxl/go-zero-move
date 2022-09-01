package model

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CampaignsModel = (*customCampaignsModel)(nil)

type (
	// CampaignsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCampaignsModel.
	CampaignsModel interface {
		campaignsModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		CampaignList(ctx context.Context, campaignId, campaignName, campaignType string, offset, limit uint64) ([]*Campaigns, int64, error)
		UpdateByCampaignId(ctx context.Context, campaignId string, values map[string]interface{}) error
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

// export logic
func (m *defaultCampaignsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultCampaignsModel) CampaignList(ctx context.Context, campaignId, campaignName, campaignType string, offset, limit uint64) (campaigns []*Campaigns, total int64, err error) {
	total, err = m.campaignListCount(ctx, campaignId, campaignName, campaignType)
	if total == 0 || err != nil {
		return
	}
	query := squirrel.Select(campaignsRows).From(m.table)
	if len(campaignId) > 0 {
		query = query.Where("campaign_id like ?", "%"+campaignId+"%")
	}
	if len(campaignName) > 0 {
		query = query.Where("campaign_name like ?", "%"+campaignName+"%")
	}
	if len(campaignType) > 0 {
		query = query.Where("campaign_type = ?", campaignType)
	}
	sql, args, err := query.Offset(offset).Limit(limit).OrderBy("updated_at desc").ToSql()
	err = m.conn.QueryRowsCtx(ctx, &campaigns, sql, args...)
	return
}

func (m *defaultCampaignsModel) campaignListCount(ctx context.Context, campaignId, campaignName, campaignType string) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table)
	if len(campaignId) > 0 {
		query = query.Where("campaign_id like ?", "%"+campaignId+"%")
	}
	if len(campaignName) > 0 {
		query = query.Where("campaign_name like ?", "%"+campaignName+"%")
	}
	if len(campaignType) > 0 {
		query = query.Where("campaign_type = ?", campaignType)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.conn.QueryRowCtx(ctx, &total, sql, args...); err != nil {
		return 0, err
	}
	return
}

func (m *defaultCampaignsModel) UpdateByCampaignId(ctx context.Context, campaignId string, values map[string]interface{}) (err error) {
	if len(values) == 0 {
		return errors.New("value can not be empty")
	}
	query := squirrel.Update(m.table).Where("campaign_id = ?", campaignId)
	for k, v := range values {
		query = query.Set(k, v)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = m.conn.ExecCtx(ctx, sql, args...)
	return err
}
