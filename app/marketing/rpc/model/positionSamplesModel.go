package model

import (
	"business/common/utils"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PositionSamplesModel = (*customPositionSamplesModel)(nil)

type (
	// PositionSamplesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPositionSamplesModel.
	PositionSamplesModel interface {
		positionSamplesModel
		SampleList(ctx context.Context, creativeSizeIds []string) (positions []*PositionSamples, err error)
	}

	customPositionSamplesModel struct {
		*defaultPositionSamplesModel
	}
)

// NewPositionSamplesModel returns a model for the database table.
func NewPositionSamplesModel(conn sqlx.SqlConn) PositionSamplesModel {
	return &customPositionSamplesModel{
		defaultPositionSamplesModel: newPositionSamplesModel(conn),
	}
}

func (m *defaultPositionSamplesModel) SampleList(ctx context.Context, creativeSizeIds []string) (samples []*PositionSamples, err error) {
	in, args, _ := utils.WhereIn(creativeSizeIds)
	sql := fmt.Sprintf("select %s from %s where creative_size_id in %s", positionSamplesRows, m.table, in)
	err = m.conn.QueryRowsCtx(ctx, &samples, sql, args...)
	if len(samples) == 0 {
		return nil, sqlc.ErrNotFound
	}
	return
}
