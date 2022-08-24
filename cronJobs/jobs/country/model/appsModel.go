package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AppsModel = (*customAppsModel)(nil)

type (
	// AppsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAppsModel.
	AppsModel interface {
		appsModel
		GetApps(ctx context.Context) ([]*AdsAppMap, error)
	}

	customAppsModel struct {
		*defaultAppsModel
	}

	AdsAppMap struct {
		AppId   string `db:"app_id"`
		AppName string `db:"app_name"`
		PkgName string `db:"pkg_name"`
	}
)

// NewAppsModel returns a model for the database table.
func NewAppsModel(conn sqlx.SqlConn) AppsModel {
	return &customAppsModel{
		defaultAppsModel: newAppsModel(conn),
	}
}

func (m *defaultAppsModel) GetApps(ctx context.Context) (apps []*AdsAppMap, err error) {
	query := fmt.Sprintf("select app_id,app_name,pkg_name from %s where pkg_name != ''", m.table)
	err = m.conn.QueryRowsCtx(ctx, &apps, query)
	return
}
