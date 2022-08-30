package model

import (
	"context"
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
		CampaignList(ctx context.Context, username, email string, state int64, offset, limit uint64) ([]*Campaigns, int64, error)
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

func (m *defaultCampaignsModel) CampaignList(ctx context.Context, username, email string, state int64, offset, limit uint64) (campaigns []*Campaigns, total int64, err error) {
	total, err = m.campaignListCount(ctx, username, email, state)
	if total == 0 || err != nil {
		return
	}
	query := squirrel.Select(campaignsRows).From(m.table).Where("state = ?", state)
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	if len(email) > 0 {
		query = query.Where("email like ?", "%"+email+"%")
	}
	sql, args, err := query.Offset(offset).Limit(limit).ToSql()
	err = m.conn.QueryRowsCtx(ctx, &campaigns, sql, args...)
	return
}

func (m *defaultCampaignsModel) campaignListCount(ctx context.Context, username, email string, state int64) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table).Where("state = ?", state)
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	if len(email) > 0 {
		query = query.Where("email like ?", "%"+email+"%")
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.conn.QueryRowsCtx(ctx, &total, sql, args...); err != nil {
		return 0, err
	}
	return
}
