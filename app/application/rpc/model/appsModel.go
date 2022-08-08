package model

import (
	"business/common/utils"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppsModel = (*customAppsModel)(nil)

type (
	// AppsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppsModel.
	AppsModel interface {
		appsModel
		AppList(ctx context.Context, appId, appName string, appType, channel int64, offset, limit uint64) (apps []*Apps, total int64, err error)
		GetAppsByAppIds(ctx context.Context, appIds []string) ([]*Apps, error)
		GetAppsByIds(ctx context.Context, ids []int64) ([]*Apps, error)
		GetAppsByAppName(ctx context.Context, appName string) ([]*Apps, error)
	}

	customAppsModel struct {
		*defaultAppsModel
	}
)

// NewAppsModel returns a model for the database table.
func NewAppsModel(conn sqlx.SqlConn, c cache.CacheConf) AppsModel {
	return &customAppsModel{
		defaultAppsModel: newAppsModel(conn, c),
	}
}

func (m *defaultAppsModel) AppList(ctx context.Context, appId, appName string, appType, channel int64, offset, limit uint64) (apps []*Apps, total int64, err error) {
	total, err = m.countList(ctx, appId, appName, appType, channel)
	if err != nil {
		return nil, 0, err
	}
	query := squirrel.Select(appsRows).From(m.table).OrderBy("updated_at desc")
	if len(appId) > 0 {
		query = query.Where("app_id like ?", "%"+appId+"%")
	}
	if len(appName) > 0 {
		query = query.Where("app_name like ?", "%"+appName+"%")
	}
	if appType > 0 {
		query = query.Where("app_type = ?", appType)
	}
	if channel > 0 {
		query = query.Where("channel = ?", channel)
	}
	sql, args, err := query.Offset(offset).Limit(limit).ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowsNoCacheCtx(ctx, &apps, sql, args...); err != nil {
		return
	}
	return
}

func (m *defaultAppsModel) countList(ctx context.Context, appId, appName string, appType, channel int64) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table)
	if len(appId) > 0 {
		query = query.Where("app_id like ?", "%"+appId+"%")
	}
	if len(appName) > 0 {
		query = query.Where("app_name like ?", "%"+appName+"%")
	}
	if appType > 0 {
		query = query.Where("app_type = ?", appType)
	}
	if channel > 0 {
		query = query.Where("channel = ?", channel)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return
	}
	if err = m.QueryRowNoCacheCtx(ctx, &total, sql, args...); err != nil {
		return 0, err
	}
	return
}

func (m *defaultAppsModel) GetAppsByAppIds(ctx context.Context, appIds []string) ([]*Apps, error) {
	in, v, _ := utils.WhereIn(appIds)
	query, args, err := squirrel.Select(appsRows).From(m.table).Where("app_id in "+in, v...).ToSql()
	if err != nil {
		return nil, err
	}
	var apps []*Apps
	err = m.QueryRowsNoCacheCtx(ctx, &apps, query, args...)
	if err != nil {
		return nil, err
	}
	return apps, nil
}
func (m *defaultAppsModel) GetAppsByIds(ctx context.Context, ids []int64) ([]*Apps, error) {
	in, v, _ := utils.WhereIn(ids)
	query, args, err := squirrel.Select(appsRows).From(m.table).Where("id in "+in, v...).ToSql()
	if err != nil {
		return nil, err
	}
	var apps []*Apps
	err = m.QueryRowsNoCacheCtx(ctx, &apps, query, args...)
	if err != nil {
		return nil, err
	}
	return apps, nil
}
func (m *defaultAppsModel) GetAppsByAppName(ctx context.Context, appName string) ([]*Apps, error) {
	query, args, err := squirrel.Select(appsRows).From(m.table).Where("app_name like ?", "%"+appName+"%").ToSql()
	if err != nil {
		return nil, err
	}
	var apps []*Apps
	err = m.QueryRowsNoCacheCtx(ctx, &apps, query, args...)
	if err != nil {
		return nil, err
	}
	return apps, nil
}
