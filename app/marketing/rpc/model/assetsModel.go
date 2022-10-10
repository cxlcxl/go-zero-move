package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ AssetsModel = (*customAssetsModel)(nil)

type (
	// AssetsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetsModel.
	AssetsModel interface {
		assetsModel
		DeleteByAssetIds(ctx context.Context, assetIds []int64) error
		BatchInsert(ctx context.Context, assets []*Assets) error
		AssetList(ctx context.Context, appId, assetType string, w, h int64, offset, limit uint64) (campaigns []*Assets, total int64, err error)
		AssetBind(ctx context.Context, appId string, assetIds []int64) (err error)
		AssetElementList(ctx context.Context, appId, assetType string, width, height, fileSize int64, assetName string) ([]*Assets, error)
	}

	customAssetsModel struct {
		*defaultAssetsModel
	}
)

// NewAssetsModel returns a model for the database table.
func NewAssetsModel(conn sqlx.SqlConn) AssetsModel {
	return &customAssetsModel{
		defaultAssetsModel: newAssetsModel(conn),
	}
}

func (m *defaultAssetsModel) BatchInsert(ctx context.Context, assets []*Assets) error {
	query := fmt.Sprintf("insert ignore into %s (%s) values ", m.table, assetsRowsExpectAutoSet)
	values := make([]interface{}, 0)
	valueStatement := make([]string, 0)
	for i := 0; i < len(assets); i++ {
		valueStatement = append(valueStatement, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		values = append(values,
			assets[i].AccountId, assets[i].AdvertiserId, assets[i].AppId, assets[i].AssetId, assets[i].AssetName,
			assets[i].AssetType, assets[i].FileUrl, assets[i].Width, assets[i].Height, assets[i].VideoPlayDuration,
			assets[i].FileSize, assets[i].FileFormat, assets[i].FileHashSha256,
		)
	}
	// 写入库
	insertSQL := query + strings.Join(valueStatement, ",")
	if _, err := m.conn.ExecCtx(ctx, insertSQL, values...); err != nil {
		return err
	}
	return nil
}

func (m *defaultAssetsModel) AssetElementList(ctx context.Context, appId, assetType string, w, h, fileSize int64, assetName string) (assets []*Assets, err error) {
	query := squirrel.Select(assetsRows).From(m.table).Where("app_id in (?, '')", appId).Where("asset_type = ?", assetType).
		Where("width = ?", w).Where("height = ?", h).Where("file_size <= ?", fileSize)
	if len(assetName) > 0 {
		query = query.Where("asset_name like ?", "%"+assetName+"%")
	}
	sql, args, err := query.OrderBy("width asc,height asc").ToSql()
	if err != nil {
		return nil, err
	}
	err = m.conn.QueryRowsCtx(ctx, &assets, sql, args...)
	if len(assets) == 0 {
		return nil, sqlc.ErrNotFound
	}
	return
}

func (m *defaultAssetsModel) AssetList(ctx context.Context, appId, assetType string, w, h int64, offset, limit uint64) (assets []*Assets, total int64, err error) {
	total, err = m.assetListCount(ctx, appId, assetType, w, h)
	if total == 0 || err != nil {
		return
	}
	query := squirrel.Select(assetsRows).From(m.table).Where("app_id = ?", appId)
	if len(assetType) > 0 {
		query = query.Where("asset_type = ?", assetType)
	}
	if w > 0 {
		query = query.Where("width = ?", w)
	}
	if h > 0 {
		query = query.Where("height = ?", h)
	}
	sql, args, err := query.Offset(offset).Limit(limit).OrderBy("width asc,height asc,asset_type asc").ToSql()
	err = m.conn.QueryRowsCtx(ctx, &assets, sql, args...)
	return
}

func (m *defaultAssetsModel) assetListCount(ctx context.Context, appId, assetType string, w, h int64) (total int64, err error) {
	query := squirrel.Select("COUNT(id) as ct").From(m.table).Where("app_id = ?", appId)
	if len(assetType) > 0 {
		query = query.Where("asset_type = ?", assetType)
	}
	if w > 0 {
		query = query.Where("width = ?", w)
	}
	if h > 0 {
		query = query.Where("height = ?", h)
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

func (m *defaultAssetsModel) DeleteByAssetIds(ctx context.Context, assetIds []int64) (err error) {
	in, args, _ := utils.WhereIn(assetIds)
	sql := fmt.Sprintf("delete from %s where asset_id in %s", m.table, in)
	_, err = m.conn.ExecCtx(ctx, sql, args...)
	return
}

func (m *defaultAssetsModel) AssetBind(ctx context.Context, appId string, assetIds []int64) (err error) {
	in, args, _ := utils.WhereIn(assetIds)
	sql := fmt.Sprintf("update %s set app_id = ? where asset_id in %s", m.table, in)
	_args := []interface{}{appId}
	_args = append(_args, args...)
	_, err = m.conn.ExecCtx(ctx, sql, _args...)
	return
}
