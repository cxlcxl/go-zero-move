package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionBasePricesModel = (*customPositionBasePricesModel)(nil)

type (
	// PositionBasePricesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionBasePricesModel.
	PositionBasePricesModel interface {
		positionBasePricesModel
		FindFloorPrice(ctx context.Context, creativeSizeId, priceType string) (float64, error)
	}

	customPositionBasePricesModel struct {
		*defaultPositionBasePricesModel
	}
)

// NewPositionBasePricesModel returns a model for the database table.
func NewPositionBasePricesModel(conn sqlx.SqlConn) PositionBasePricesModel {
	return &customPositionBasePricesModel{
		defaultPositionBasePricesModel: newPositionBasePricesModel(conn),
	}
}
func (m *defaultPositionBasePricesModel) FindFloorPrice(ctx context.Context, creativeSizeId, priceType string) (f float64, err error) {
	sql := fmt.Sprintf("select base_price from %s where creative_size_id = ? and price_type = ?", m.table)
	err = m.conn.QueryRowCtx(ctx, &f, sql, creativeSizeId, priceType)
	switch err {
	case nil:
		return f, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
