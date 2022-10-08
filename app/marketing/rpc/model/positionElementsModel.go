package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionElementsModel = (*customPositionElementsModel)(nil)

type (
	// PositionElementsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionElementsModel.
	PositionElementsModel interface {
		positionElementsModel
		PositionElement(ctx context.Context, creativeSizeId, subType string, width, height uint64) (elements []*PositionElements, err error)
	}

	customPositionElementsModel struct {
		*defaultPositionElementsModel
	}
)

// NewPositionElementsModel returns a model for the database table.
func NewPositionElementsModel(conn sqlx.SqlConn) PositionElementsModel {
	return &customPositionElementsModel{
		defaultPositionElementsModel: newPositionElementsModel(conn),
	}
}

func (m *defaultPositionElementsModel) PositionElement(ctx context.Context, creativeSizeId, subType string, width, height uint64) (elements []*PositionElements, err error) {
	var groupNumber int64
	gnSql := fmt.Sprintf("select group_number from %s where creative_size_id = ? and sub_type = ? and width = ? and height = ?", m.table)
	if err = m.conn.QueryRowCtx(ctx, &groupNumber, gnSql, creativeSizeId, subType, width, height); err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("select %s from %s where creative_size_id = ? and sub_type = ? and group_number = ?", positionElementsRows, m.table)
	if err = m.conn.QueryRowsCtx(ctx, &elements, sql, creativeSizeId, subType, groupNumber); err != nil {
		return nil, err
	}
	return
}
