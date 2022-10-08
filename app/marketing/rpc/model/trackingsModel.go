package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ TrackingsModel = (*customTrackingsModel)(nil)

type (
	// TrackingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTrackingsModel.
	TrackingsModel interface {
		trackingsModel
		TrackingList(ctx context.Context, appId string) (trackings []*Trackings, err error)
		BatchInsert(ctx context.Context, trackings []*Trackings, disabled []int64) error
	}

	customTrackingsModel struct {
		*defaultTrackingsModel
	}
)

// NewTrackingsModel returns a model for the database table.
func NewTrackingsModel(conn sqlx.SqlConn) TrackingsModel {
	return &customTrackingsModel{
		defaultTrackingsModel: newTrackingsModel(conn),
	}
}

func (m *defaultTrackingsModel) TrackingList(ctx context.Context, appId string) (trackings []*Trackings, err error) {
	sqlString := "select id,account_id,advertiser_id,app_id,tracking_id,effect_type,effect_name,'' as click_tracking_url,'' as imp_tracking_url from %s where app_id = ?"
	sql := fmt.Sprintf(sqlString, m.table)
	err = m.conn.QueryRowsCtx(ctx, &trackings, sql, appId)
	if len(trackings) == 0 {
		return nil, sqlc.ErrNotFound
	}
	return
}

func (m *defaultTrackingsModel) BatchInsert(ctx context.Context, trackings []*Trackings, disabled []int64) error {
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.table, trackingsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	for i := 0; i < len(trackings); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?)")
		values = append(values,
			trackings[i].AccountId, trackings[i].AdvertiserId, trackings[i].AppId, trackings[i].TrackingId,
			trackings[i].EffectType, trackings[i].EffectName, trackings[i].ClickTrackingUrl, trackings[i].ImpTrackingUrl,
		)
	}
	// 写入库
	insertSQL := query + strings.Join(valueStatement, ",")
	deleteSql := ""
	args := []interface{}{trackings[0].AppId}
	if len(disabled) > 0 {
		in, _args, _ := utils.WhereIn(disabled)
		deleteSql = fmt.Sprintf("delete from %s where app_id = ? and tracking_id in "+in, m.table)
		args = append(args, _args...)
	}
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if len(deleteSql) > 0 {
			if _, err := session.ExecCtx(ctx, deleteSql, args...); err != nil {
				return err
			}
		}
		if _, err := session.ExecCtx(ctx, insertSQL, values...); err != nil {
			return err
		}
		return nil
	})
}
